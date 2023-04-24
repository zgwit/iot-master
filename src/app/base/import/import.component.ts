import { Component, EventEmitter, Input, Output } from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd/message';
@Component({
  selector: 'app-import',
  templateUrl: './import.component.html',
  styleUrls: ['./import.component.scss'],
})
export class ImportComponent {
  @Input() url!: string;
  @Output() onLoad = new EventEmitter<string>();
  uploading: Boolean = false;
  constructor(private msg: NzMessageService
  ) { }
  load() {
    this.onLoad.emit();
  }
  handleChange(info: any): void {
    if(info.type === 'error'){
      this.msg.error(`上传失败`);
      return;
    }
    if (info.file && info.file.response) {
      const res = info.file.response;
      if (!res.error) {
        this.msg.success(`成功导入${res.data}条数据!`);
        this.load();
      } else {
        this.msg.error(`${res.error}`);
      }
    }
  }
}
