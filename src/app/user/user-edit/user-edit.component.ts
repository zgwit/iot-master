import { Component, Input, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RequestService } from '../../request.service';
import { NzMessageService } from 'ng-zorro-antd/message';

@Component({
  selector: 'app-users-edit',
  templateUrl: './user-edit.component.html',
  styleUrls: ['./user-edit.component.scss'],
})
export class UserEditComponent implements OnInit {
  group!: FormGroup;
  // id: any = 0;
  listOfOption: any[] = [];
  constructor(
    private fb: FormBuilder,
    private rs: RequestService,
    private msg: NzMessageService
  ) { }
  @Input() id: string = '';
  ngOnInit(): void {
    if (this.id) {
      this.rs.get(`user/${this.id}`).subscribe((res) => {
        //let data = res.data;
        this.build(res.data);
      });
    }
    this.build();
    this.getRoleList();

  }

  build(obj?: any) {
    obj = obj || {};
    this.group = this.fb.group({
      username: [obj.username || '', [Validators.required]],
      name: [obj.name || '', []],
      email: [obj.email || '', []],
      roles: [obj.roles || [], []],
      disabled: [obj.disabled || false, []],
    });
  }

  submit() {  // 从父组件调用
    return new Promise((resolve, reject) => {
      if (this.group.valid) {
        let url = this.id ? `user/${this.id}` : `user/create`
        this.rs.post(url, this.group.value).subscribe(res => {
          this.msg.success("保存成功");
          resolve(true);
        })
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
  getRoleList() {
    this.rs.get('role/list').subscribe((res) => {
      const data: any[] = []
      res.data.filter((item: { id: any; name: any }) => {
        data.push({
          value: item.name,
          label: item.name,
        });
        this.listOfOption = data
      })
        ;
    });
  }
}
