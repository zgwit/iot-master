import {Component, Input, OnInit} from '@angular/core';
import {NzModalRef} from "ng-zorro-antd/modal";
import {HttpClient} from "@angular/common/http";

@Component({
  selector: 'hmi-attachment',
  templateUrl: './attachment.component.html',
  styleUrls: ['./attachment.component.scss']
})
export class AttachmentComponent implements OnInit {

  @Input() url = "/hmi/"
  @Input() path = ""

  files: Array<any> = [];
  select: string = "";

  constructor(private mr: NzModalRef, private http: HttpClient) {

  }

  ngOnInit(): void {
    this.load()
  }

  load() {
    this.http.get(this.url + this.path).subscribe((res: any) => {
      console.log('attachment', res)
      this.files = res.data;
    })
  }

  cancel() {
    this.mr.close()
  }

  ok() {
    this.mr.close(this.url + this.path + '/' + this.select);
  }

  onItemClick(data: any) {
    if (data.folder) {
      if (this.path != "")
        this.path += "/"
      this.path += data.name
      this.load()
      return
    }

    this.select = data.name
  }

  onItemDoubleClick(data: any) {
    this.mr.close(this.url + this.path + '/' + data.name);
  }

  onItemChecked(name: string, $event: boolean) {
    if ($event)
      this.select = name
  }

  upper() {
    let st = this.path.split('/')
    st.pop()
    this.path = st.join('/')
    this.load()
  }

  remove(data: any, i: number) {
    this.http.delete(this.url + this.path + '/' + data.name).subscribe(res => {
      this.files.splice(i, 1)
    })
  }

  rename(data: any) {
    let filename = prompt("请输入新名称", data.name)
    if (filename)
      this.http.patch(this.url + this.path + '/' + data.name, {filename}).subscribe(res => {
        data.name = filename
      })
  }
}
