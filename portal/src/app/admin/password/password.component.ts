import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {RequestService} from "../../request.service";
import {AbstractControl, FormBuilder, FormControl, FormGroup, ValidationErrors, Validators} from "@angular/forms";
import {NzMessageService} from "ng-zorro-antd/message";
import {Md5} from "ts-md5/dist/md5";
import {NzModalRef} from "ng-zorro-antd/modal";

@Component({
  selector: 'app-password',
  templateUrl: './password.component.html',
  styleUrls: ['./password.component.scss']
})
export class PasswordComponent implements OnInit {

  basicForm: FormGroup = new FormGroup({
    old: new FormControl("", [Validators.required]),
    new: new FormControl("", [Validators.required, Validators.minLength(6), Validators.maxLength(20)]),
    new2: new FormControl("", [Validators.required, (control: AbstractControl): ValidationErrors | null => {
      if (!this.basicForm) return null;
      const ret = this.basicForm.value.new !== control.value;
      return ret ? {diff: {value: control.value}} : null;
    }]),
  });

  constructor(private fb: FormBuilder,
              private rs: RequestService,
              private ms: NzMessageService,
              private mr: NzModalRef) {

  }

  ngOnInit(): void {
  }

  submit(): void {
    //const val = this.basicForm.value;
    if(!this.basicForm.valid) return;

    this.rs.post('password', {
      old: Md5.hashStr(this.basicForm.value.old),
      new: Md5.hashStr(this.basicForm.value.new),
    }).subscribe(res=>{
      this.ms.success("修改成功");
      this.mr.close()
    })
  }
}
