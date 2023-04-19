import { Component, Input, OnInit } from '@angular/core';
import { FormBuilder, Validators } from "@angular/forms";
import { ActivatedRoute, Router } from "@angular/router";
import { RequestService } from "../../request.service";
import { NzModalRef, NzModalService } from "ng-zorro-antd/modal";
import { NzMessageService } from "ng-zorro-antd/message";

@Component({
  selector: 'app-device-group-edit',
  templateUrl: './device-group-edit.component.html',
  styleUrls: ['./device-group-edit.component.scss']
})
export class DeviceGroupEditComponent implements OnInit {
  @Input() id = 0;
  areaID: any[] = [];
  group: any = {};

  constructor(private fb: FormBuilder,
    private rs: RequestService,
    private ms: NzModalService,
    private msg: NzMessageService,
    protected ref: NzModalRef) {

  }


  ngOnInit(): void {
    //console.log('init', this.id)
    if (this.id) {
      this.rs.get(`device/group/${this.id}`).subscribe(res => {
        //let data = res.data;
        this.build(res.data)
      })
    }
    this.build()

    this.rs
    .post('device/area/search', {})
    .subscribe((res) => {  
      const data: any[] = [];

      res.data.filter((item: { id: string; name: string }) =>
        data.push({ label: item.id + ' / ' + item.name, value: item.id })
      );
      this.areaID = data;
    })
    .add(() => {});
  }

  build(obj?: any) {
    obj = obj || {}
    this.group = this.fb.group({
      name: [obj.name || '', [Validators.required]],
      desc: [obj.desc || '', []], 
      area_id: [obj.area_id || '', []],
    })
  }

  submit() {
    let url = this.id ? `device/group/${this.id}` : `device/group/create`;
    this.rs.post(url, this.group.value).subscribe(res => {
      this.ref.close(res.data);
      this.msg.success("保存成功");
    })
  }
}
