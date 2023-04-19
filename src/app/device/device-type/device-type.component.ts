 import { Component } from '@angular/core';
import { NzModalService } from 'ng-zorro-antd/modal';
import { Router } from '@angular/router';
import { RequestService } from 'src/app/request.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzTableQueryParams } from 'ng-zorro-antd/table';
import { ParseTableQuery } from 'src/app/base/table';
import {
  isIncludeAdmin,
  tableHeight,
  onAllChecked,
  onItemChecked,
  batchdel,
  refreshCheckedStatus,
  readCsv,
} from 'src/public';
@Component({
  selector: 'app-device-type',
  templateUrl: './device-type.component.html',
  styleUrls: ['./device-type.component.scss']
})
export class DeviceTypeComponent {
  loading = true;
  uploading: Boolean = false;
  datum: any[] = [];
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {};
  checked = false;
  indeterminate = false;
  setOfCheckedId = new Set<number>();
  delResData: any = [];
  href!: string;

  // handleCancel(){}
  // submit(){}
  // listOfOption = ['Apples', 'Nails', 'Bananas', 'Helicopters'];
  // listOfSelectedValue: string[] = [];

  // isNotSelected(value: string): boolean {
  //   return this.listOfSelectedValue.indexOf(value) === -1;
  // }


  constructor(
    private modal: NzModalService,
    private router: Router,
    private rs: RequestService,
    private msg: NzMessageService
  ) {
    //this.load();
  }

  reload() {
    this.datum = [];
    this.load();
  }

  load() {
    this.loading = true;
    this.rs
      .post('device/type/search', this.query)
      .subscribe((res) => {
     //   console.log(res);
        this.datum = res.data||[];
        this.total = res.total;
        this.setOfCheckedId.clear();
        refreshCheckedStatus(this);
      })
      .add(() => {
        this.loading = false;
      });
  }
   
  create() {
    let path = '/device/type/create';
    if (location.pathname.startsWith('/admin')) path = '/admin' + path;
    this.router.navigateByUrl(path);
  }

  delete(id: number, size?: number) {
    this.rs.get(`device/type/${id}/delete`).subscribe((res) => {
      if (!size ) {
        this.msg.success('删除成功');
        this.datum = this.datum.filter((d) => d.id !== id);
      } else if (size) {
        this.delResData.push(res);
        if (size === this.delResData.length) {
          this.msg.success('删除成功');
          this.load();
        }
      }
    });
  }
  handleExport(){ 
    this.href = `/api/device/type/export`;
  }
  
  onQuery($event: NzTableQueryParams) {
    ParseTableQuery($event, this.query);
    this.load();
  }
  pageIndexChange(pageIndex: number) {
    console.log('pageIndex:', pageIndex);
  }
  pageSizeChange(pageSize: number) {
    this.query.limit = pageSize;
    this.load();
  }
  search($event: string) {
    this.query.keyword = {
      name: $event,
    };
    this.query.skip = 0;
    this.load();
  }
  use(id: any) {
    this.rs.get(`device/type/${id}/enable`).subscribe((res) => {
      this.load();
    });
  }
  forbid(id: any) {
    this.rs.get(`device/type/${id}/disable`).subscribe((res) => {
      this.load();
    });
  }
  edit(id: any) {
    const path = `${isIncludeAdmin()}/device/type/edit/${id}`;
    this.router.navigateByUrl(path);
  }
  cancel() {
    this.msg.info('取消操作');
  }

  handleNew() {
    const path = `${isIncludeAdmin()}/device/type/create`;
    this.router.navigateByUrl(path);
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
