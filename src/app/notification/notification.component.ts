import { Component } from '@angular/core';
import { Router } from "@angular/router";
import { RequestService } from '../request.service';
import { NzMessageService } from "ng-zorro-antd/message";
import { NzTableQueryParams } from "ng-zorro-antd/table";
import { NzModalService } from "ng-zorro-antd/modal";
import { ParseTableQuery } from "../base/table";
import {  onAllChecked, onItemChecked, batchdel, refreshCheckedStatus } from 'src/public';
@Component({
  selector: 'app-notification',
  templateUrl: './notification.component.html',
  styleUrls: ['./notification.component.scss']
})
export class NotificationComponent {

  loading = true
  datum: any[] = []
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {}
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
    this.rs.post("notification/search", this.query).subscribe(res => {
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
    this.rs.get(`notification/${id}/delete`).subscribe(res => {
      if (!size) {
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
    this.query.skip = pageIndex - 1;
  }

  pageSizeChange(pageSize: number) {
    this.query.limit = pageSize;
  }

  search($event: string) {
    this.query.keyword = {
      id: $event
    };
    this.query.skip = 0;
    this.load();
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
