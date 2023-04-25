import { Component } from '@angular/core';
import { FormBuilder, Validators, FormGroup } from "@angular/forms";
import { ActivatedRoute, Router } from "@angular/router";
import { RequestService } from "../../request.service";
import { NzMessageService } from "ng-zorro-antd/message";
import { Md5 } from "ts-md5";

@Component({
  selector: 'app-password',
  templateUrl: './password.component.html',
  styleUrls: ['./password.component.scss']
})
export class PasswordComponent {
  group!: FormGroup;
  id: any = 0

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private rs: RequestService,
    private msg: NzMessageService) {
    this.build()
  }


  build() {
    this.group = this.fb.group({
      old: ['', [Validators.required]],
      new: ['', [Validators.required]],
      repeat: ['', [Validators.required]],
    })
  }

  submit() {
    let body = {
      old: Md5.hashStr(this.group.value.old),
      new: Md5.hashStr(this.group.value.new),
    }
    return new Promise((resolve, reject) => {
      if (this.group.valid) {
        this.rs.post("password", body).subscribe(res => {
          //清空session
          this.rs.get("logout").subscribe(res => { })
          this.router.navigateByUrl("/login");
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
}
