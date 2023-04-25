import { Component, Input, OnInit } from '@angular/core';
import { FormBuilder, Validators, FormGroup } from '@angular/forms';
import { RequestService } from '../../request.service';
import { NzMessageService } from 'ng-zorro-antd/message';

@Component({
  selector: 'app-gateway-edit',
  templateUrl: './gateway-edit.component.html',
  styleUrls: ['./gateway-edit.component.scss'],
})
export class GatewayEditComponent implements OnInit {
  group!: FormGroup;
  // id: any = 0;
  @Input() id = '';
  constructor(
    private fb: FormBuilder,
    private rs: RequestService,
    private msg: NzMessageService
  ) { }

  ngOnInit(): void {
    if (this.id) {
      this.rs.get(`gateway/${this.id}`).subscribe((res) => {
        //let data = res.data;
        this.build(res.data);
      });
    }

    this.build();
  }

  build(obj?: any) {
    obj = obj || {};
    this.group = this.fb.group({
      id: [obj.id || '', []],
      name: [obj.name || '', [Validators.required]],
      username: [obj.username || '', []],
      password: [obj.password || '', []],
      // disabled: [obj.disabled || false, []],
      desc: [obj.desc || '', []],
    });
  }

  submit() {//从父组件调用
    return new Promise((resolve, reject) => {
      if (this.group.valid) {
        let url = this.id ? `gateway/${this.id}` : `gateway/create`;
        this.rs.post(url, this.group.value).subscribe((res) => {
          this.msg.success('保存成功');
          resolve(true);
        });
      } else {
        Object.values(this.group.controls).forEach((control) => {
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
