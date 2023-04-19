import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { RequestService } from '../../request.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzTableQueryParams } from 'ng-zorro-antd/table';
import { NzModalService } from "ng-zorro-antd/modal";
import { ParseTableQuery } from '../../base/table';
import { tableHeight, onAllChecked, onItemChecked, batchdel, refreshCheckedStatus, readCsv } from "../../../public";

@Component({
  selector: 'app-role',
  templateUrl: './role.component.html',
  styleUrls: ['./role.component.scss'],
})
export class RoleComponent {
  href!:string
  loading = true;
  uploading: Boolean = false;
  datum: any[] = [];
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {};
  isVisible = false;
  roleObj = {};
  checked = false;
  indeterminate = false;
  setOfCheckedId = new Set<number>();
  delResData: any = [];
  constructor(
    private router: Router,
    private rs: RequestService,
    private modal: NzModalService,
    private msg: NzMessageService,
  ) {
    this.load();
  }

  reload() {
    this.datum = [];
    this.load();
  }

  load() {
    this.loading = true;
    this.rs
      .post('role/search', this.query)
      .subscribe((res) => {
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
    let path = '/user/create';
    if (location.pathname.startsWith('/admin')) path = '/admin' + path;
    this.router.navigateByUrl(path);
  }

  delete(id: number, size?: number) {
    this.rs.get(`role/${id}/delete`).subscribe((res) => {
      if (!size  ) {
        this.msg.success("删除成功");
        this.datum = this.datum.filter(d => d.id !== id);
      } else if (size) {
        this.delResData.push(res);
        if (size === this.delResData.length) {
          this.msg.success("删除成功");
          this.load();
        }
      }
    });
  }
  handleExport(){
    this.href = `/api/role/export`;  
  }
   
  onQuery($event: NzTableQueryParams) {
    ParseTableQuery($event, this.query);
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
      name: $event,
    };
    this.query.skip = 0;
    this.load();
  }

  edit(id: any) {
    let path = '/user/privillege/' + id;
    if (location.pathname.startsWith('/admin')) path = '/admin' + path;
    this.router.navigateByUrl(path);
  }
  add() {
    this.router.navigateByUrl('/admin/user/role/create');
  }
  cancel() {
    this.msg.info('点击取消');
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
