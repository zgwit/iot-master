import { Component } from '@angular/core';
import { Router } from "@angular/router";
import { RequestService } from "../../request.service";
import { NzMessageService } from "ng-zorro-antd/message";
import { NzTableQueryParams } from "ng-zorro-antd/table";
import { NzModalService } from "ng-zorro-antd/modal";
import { ParseTableQuery } from "../../base/table";
import { tableHeight, onAllChecked, onItemChecked, batchdel, refreshCheckedStatus } from "../../../public";
@Component({
  selector: 'app-alarms',
  templateUrl: './alarms.component.html',
  styleUrls: ['./alarms.component.scss']
})
export class AlarmsComponent {


  loading = true
  datum: any[] = []
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {}
  filterRead = [
    { text: 'true', value: 1 },
    { text: 'false', value: 0 }
  ]
  filterLevel = [
    { text: '1', value: 1 },
    { text: '2', value: 2 },
    { text: '3', value: 3 },
  ]
  checked = false;
  indeterminate = false;
  setOfCheckedId = new Set<number>();
  delResData: any = [];
  constructor(private router: Router,
    private rs: RequestService,
    private modal: NzModalService,
    private msg: NzMessageService
  ) {
    //this.load();
  }

  reload() {
    this.datum = [];
    this.load()
  }

  load() {
    this.loading = true
    this.rs.post("alarm/search", this.query).subscribe(res => {
      const { data, total } = res;
      this.datum = data || [];
      this.total = total || 0;
      this.setOfCheckedId.clear();
      refreshCheckedStatus(this);
    }).add(() => {
      this.loading = false;
    })
  }

  delete(id: number, size?: number) {
    this.rs.get(`alarm/${id}/delete`).subscribe(res => {
      if (!size ) {
        this.msg.success("删除成功");
        this.datum = this.datum.filter(d => d.id !== id);
      } else if (size) {
        this.delResData.push(res);
        if (size === this.delResData.length) {
          this.msg.success("删除成功");
          this.load();
        }
      }
    })
  }

  onQuery($event: NzTableQueryParams) {
    ParseTableQuery($event, this.query)
    this.load();
  }
  pageIndexChange(pageIndex: number) {
    console.log("pageIndex:", pageIndex)
  }
  pageSizeChange(pageSize: number) {
    this.query.limit = pageSize;
    this.load();
  }

  search($event: string) {
    this.query.keyword = {
      title: $event,
      Message: $event,
    };
    this.query.skip = 0;
    this.load();
  }


  read(data: any) {
    this.rs.get(`alarm/${data.id}/read`).subscribe(res => {
      data.read = true;
      //this.msg.success("删除成功")
    })
  }

  getTableHeight() {
    return tableHeight(this);
  }
  handleBatchDel() {
    batchdel(this);
  }
  handleAllChecked(id: any) {
    onAllChecked(id, this);
  }
  handleItemChecked(id: number, checked: boolean) {
    onItemChecked(id, checked, this);
  }
}
