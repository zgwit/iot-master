import { Component, ViewContainerRef } from '@angular/core';
import { Router } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ParseTableQuery } from 'src/app/base/table';
import { RequestService } from 'src/app/request.service';
import { batchdel, onAllChecked, onItemChecked, refreshCheckedStatus, tableHeight } from 'src/public';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { NzTableQueryParams } from 'ng-zorro-antd/table';
import { NzMessageService } from 'ng-zorro-antd/message';
import { UploadComponent } from './upload/upload.component';
@Component({
  selector: 'app-attachment',
  templateUrl: './attachment.component.html',
  styleUrls: ['./attachment.component.scss']
})
export class AttachmentComponent {
  loading = true;
  inputValue = '';
  datum: any[] = []
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {}
  checked = false;
  indeterminate = false;
  setOfCheckedId = new Set<number>();
  delResData: any = [];
  constructor(
    private rs: RequestService,
    private modal: NzModalService,
    private msg: NzMessageService,
    private viewContainerRef: ViewContainerRef,
  ) { }
  handleUpload() {
    this.modal.create({
      nzTitle: '上传文件',
      nzContent: UploadComponent,
      nzViewContainerRef: this.viewContainerRef,
      nzFooter: null
    });
  }

  reload() {
    this.datum = [];
    this.load()
  }

  load() {
    this.loading = true
    this.rs.get(`attach/list/${this.inputValue || ''}`).subscribe(res => {
      const { data, total } = res;
      this.datum = data || [];
      this.total = total || 0;
      this.setOfCheckedId.clear();
      refreshCheckedStatus(this);
    }).add(() => {
      this.loading = false;
    })
  }

  delete(name: number, size?: number) {
    this.rs.get(`attach/remove/${this.inputValue}/${name}`).subscribe(res => {
      this.msg.success("删除成功");
      this.load();
    })
  }
  cancel() { }
  onQuery($event: NzTableQueryParams) {
    ParseTableQuery($event, this.query)
    this.search();
  }
  pageIndexChange(pageIndex: number) {
    console.log("pageIndex:", pageIndex)
    this.query.skip = pageIndex - 1;
  }
  pageSizeChange(pageSize: number) {
    this.query.limit = pageSize;
  }
  search() {
    this.query.keyword = {};
    this.query.skip = 0;
    this.load();
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
  handleRedictTo(url: string) {
    this.inputValue = this.inputValue ? `${this.inputValue}/${url}` : url;
    this.search();
  }
  handleRename(name: string) {

  }
}
