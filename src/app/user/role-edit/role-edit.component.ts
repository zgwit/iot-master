import { Component, OnInit } from '@angular/core';
import { DatePipe } from '@angular/common';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { RequestService } from '../../request.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { isIncludeAdmin } from "../../../public";

@Component({
  selector: 'app-role-edit',
  templateUrl: './role-edit.component.html',
  styleUrls: ['./role-edit.component.scss'],
  providers: [DatePipe]
})
export class RoleEditComponent implements OnInit {
  group!: FormGroup;
  id: any = 0;

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private route: ActivatedRoute,
    private rs: RequestService,
    private msg: NzMessageService,
    private datePipe: DatePipe
  ) { }

  listOfOption: Array<{ label: string; value: string }> = [];
  listOfSelectedValue = [];

  ngOnInit(): void {
    this.build();
    if (this.route.snapshot.paramMap.has('id')) {
      this.id = this.route.snapshot.paramMap.get('id');
      this.rs.get(`role/${this.id}`).subscribe((res) => {
        this.build(res.data);
      });
    }
    this.getRoleList();
  }

  build(obj?: any) {
    obj = obj || {};
    const { name, id, privileges } = obj || {};
    this.group = this.fb.group({
      name: [name || '', [Validators.required]],
      id: [id || '', [Validators.required]],
      privileges: [privileges || [], [Validators.required]]
    });
  }
  getRoleList() {
    this.rs
      .get('privileges')
      .subscribe((res) => {
        const { data } = res;
        const listData = [];
        for (const key in data) {
          listData.push({
            label: data[key],
            value: key
          })
        }
        this.listOfOption = listData;
      })

     
  }
  submit() {
    if (this.group.valid) {
      let url = this.id ? `role/${this.id}` : `role/create`
      this.rs.post(url, this.group.value).subscribe(res => {
        const path = `${isIncludeAdmin()}/user/role`;
        this.router.navigateByUrl(path);
        this.msg.success("保存成功");
      })
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
    const path = `${isIncludeAdmin()}/user/role`;
    this.router.navigateByUrl(path);
  }
}
