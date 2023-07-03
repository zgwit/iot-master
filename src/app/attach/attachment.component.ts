import { Component, ViewContainerRef } from '@angular/core';
import { ParseTableQuery } from 'src/app/base/table';
import { RequestService } from 'src/app/request.service';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { NzTableQueryParams } from 'ng-zorro-antd/table';
import { NzMessageService } from 'ng-zorro-antd/message';
import { UploadComponent } from './upload/upload.component';
import { RenameComponent } from './rename/rename.component';
@Component({
  selector: 'app-attachment',
  templateUrl: './attachment.component.html',
  styleUrls: ['./attachment.component.scss']
})
export class AttachmentComponent {
  loading = false;
  inputValue = '';
  datum: any[] = []
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {}
  fromIframe = true;
  constructor(
    private rs: RequestService,
    private modal: NzModalService,
    private msg: NzMessageService,
    private viewContainerRef: ViewContainerRef,
  ) {
    this.load();
  }

  load() {
    this.loading = true
    this.rs.get(`attach/list/${this.inputValue || ''}`).subscribe(res => {
      const { data, total } = res;
      this.datum = data || [];
      this.total = total || 0;
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
  search() {
    this.query.keyword = {};
    this.query.skip = 0;
    this.datum = [];
    this.load();
  }
  handleRedictTo(url: string) {
    this.inputValue = this.inputValue ? `${this.inputValue}/${url}` : url;
    this.search();
  }
  handleUpload() {
    this.modal.create({
      nzTitle: '上传文件',
      nzContent: UploadComponent,
      nzViewContainerRef: this.viewContainerRef,
      nzComponentParams: {
        inputValue: this.inputValue,
      },
      nzFooter: null,
      nzOnCancel: ({ isSuccess }) => {
        if (isSuccess) {
          this.load();
        }
      }
    });
  }
  handleRename(currentName: string) {
    const modal: NzModalRef = this.modal.create({
      nzTitle: '重命名',
      nzContent: RenameComponent,
      nzComponentParams: {
        currentName
      },
      nzViewContainerRef: this.viewContainerRef,
      nzFooter: [
        {
          label: '取消',
          onClick: () => modal.destroy()
        },
        {
          label: '保存',
          type: 'primary',
          onClick: () => {
            const comp = modal.getContentComponent();
            this.rs.post(`attach/rename/${this.inputValue ? this.inputValue + '/' : ''}${currentName}`, { name: comp.name }).subscribe(res => {
              this.msg.success("保存成功");
              modal.destroy();
              this.load();
            })
          }
        },
      ]
    });
  }
  handleSrc(name: string) {
    const link = this.inputValue ? `${this.inputValue}/` : '';
    return `/attach/${link}${name}`;
  }
  handleOpenLink(name: string) {
    window.open(this.handleSrc(name))
  }
  getMatched(mime: string) {
    const reg = mime.match(/image|video|word|powerpoint|excel|text|zip|pdf|html|flash|exe|xml|psd/g)
    return reg ? reg[0] : 'unknown';
  }
  handleCopy(name: string) {
    const url = this.handleSrc(name);
    // navigator && navigator.clipboard && navigator.clipboard.writeText(url).then(res => {
    //   this.msg.success('复制成功');
    // }).catch(err => { })
    this.msg.success('复制成功');
    const inputEle = document.createElement('input');
    inputEle.value = url;
    inputEle.setAttribute('readonly', '');
    document.body.appendChild(inputEle);
    inputEle.select();
    document.execCommand('copy');
    document.body.removeChild(inputEle);
  }

  handlePre() {
    const arr = this.inputValue.split('/');
    arr.pop();
    this.inputValue = arr.join('/');
    this.search();
  }
  handleDownLoad(name: string) {
    const a = document.createElement('a');
    const event = new MouseEvent('click');
    a.download = name;
    a.href = this.handleSrc(name);
    a.dispatchEvent(event)
  }
  handleSelect(name: string) {
    const url = this.handleSrc(name);
    console.log("🚀 ~ file: attachment.component.ts:145 ~ AttachmentComponent ~ handleSelect ~ url:", url)
    window.top?.postMessage(url);
  }
}