import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators, FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { RequestService } from '../../request.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { isIncludeAdmin } from '../../../public';

@Component({
  selector: 'app-gateway-edit',
  templateUrl: './gateway-edit.component.html',
  styleUrls: ['./gateway-edit.component.scss'],
})
export class GatewayEditComponent implements OnInit {
  group!: FormGroup;
  id: any = 0;

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private route: ActivatedRoute,
    private rs: RequestService,
    private msg: NzMessageService
  ) {}

  ngOnInit(): void {
    if (this.route.snapshot.paramMap.has('id')) {
      this.id = this.route.snapshot.paramMap.get('id');
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

  submit() {
    if (this.group.valid) {
      let url = this.id ? `gateway/${this.id}` : `gateway/create`;
      this.rs.post(url, this.group.value).subscribe((res) => {
      
        const path = `${isIncludeAdmin()}/gateway/list`;
        this.router.navigateByUrl(path);
        this.msg.success('保存成功');
      });

      return;
    } else {
      Object.values(this.group.controls).forEach((control) => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }
  handleCancel() {
    const path = `${isIncludeAdmin()}/gateway/list`;
    this.router.navigateByUrl(path);
  }
}
