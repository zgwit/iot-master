import { Component, Input, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from "@angular/forms";
import { RequestService } from "../../request.service";
import { NzMessageService } from "ng-zorro-antd/message";

@Component({
  selector: 'app-brokers-edit',
  templateUrl: './broker-edit.component.html',
  styleUrls: ['./broker-edit.component.scss']
})
export class BrokerEditComponent implements OnInit {
  group!: FormGroup;
  // id: any = 0
  @Input() id: string = '';

  constructor(private fb: FormBuilder,
    private rs: RequestService,
    private msg: NzMessageService) {
  }

  ngOnInit(): void {
    if (this.id) {
      this.rs.get(`broker/${this.id}`).subscribe(res => {
        //let data = res.data;
        this.build(res.data)
      })
    }
    this.build()
  }

  build(obj?: any) {
    obj = obj || {}
    this.group = this.fb.group({
      id: [obj.id || '', []],
      name: [obj.name || '', [Validators.required]],
      desc: [obj.desc || '', []],
      port: [obj.port || 1883, []],
      cert: [obj.cert || '', []],
      key: [obj.key || '', []],
    })
  }

  submit() {
    return new Promise((resolve, reject) => {
      if (this.group.valid) {
        let url = this.id ? `broker/${this.id}` : `broker/create`
        this.rs.post(url, this.group.value).subscribe(res => {
          this.msg.success("保存成功");
          resolve(true);
        })
      } else {
        Object.values(this.group.controls).forEach(control => {
          if (control.invalid) {
            control.markAsDirty();
            control.updateValueAndValidity({ onlySelf: true });
            reject();
          }
        });
      }

    })
  }
}
