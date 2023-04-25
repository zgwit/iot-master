import { Component, Input, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RequestService } from '../../request.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzModalService } from 'ng-zorro-antd/modal';
import { DevicesComponent } from '../devices/devices.component';
@Component({
  selector: 'app-products-edit',
  templateUrl: './device-edit.component.html',
  styleUrls: ['./device-edit.component.scss'],
})
export class DeviceEditComponent implements OnInit {
  group!: FormGroup;
  // id: any = 0;
  @Input() id = '';
  constructor(
    private fb: FormBuilder,
    private rs: RequestService,
    private ms: NzModalService,
    private msg: NzMessageService
  ) { }

  ngOnInit(): void {
    if (this.id) {
      this.rs.get(`device/${this.id}`).subscribe((res) => {
        this.build(res.data);
      });
    }
    this.build();
  }

  build(obj?: any) {
    obj = obj || {};
    this.group = this.fb.group({
      id: [obj.id || '', []],
      product_id: [obj.product_id || '', []],
      gateway_id: [obj.gateway_id || '', []],
      name: [obj.name || '', [Validators.required]],
      desc: [obj.desc || '', []],
      // disabled: [obj.disabled || false, []],
    });
  }

  submit() {
    return new Promise((resolve, reject) => {
      if (this.group.valid) {
        let url = this.id ? `device/${this.id}` : `device/create`;
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

  chooseGateway() {
    this.ms
      .create({
        nzTitle: '选择网关',
        nzContent: DevicesComponent,
        nzComponentParams: {
          chooseGateway: true,
          showAddBtn: false,
        },
        nzFooter: null,
      })
      .afterClose.subscribe((res) => {
        if (res) {
          this.group.patchValue({ gateway_id: res });
        }
      });
  }
}
