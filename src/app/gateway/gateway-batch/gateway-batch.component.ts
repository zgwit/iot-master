import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators, FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { RequestService } from '../../request.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { isIncludeAdmin } from '../../../public';

@Component({
  selector: 'app-gateway-batch',
  templateUrl: './gateway-batch.component.html',
  styleUrls: ['./gateway-batch.component.scss'],
})
export class GatewayBatchComponent implements OnInit {
  group!: FormGroup;
  id: any = 0;
  datum: any[] = [];
  nzTitle = '保存';
  nzLoading = false;
  constructor(
    private fb: FormBuilder,
    private router: Router,
    private route: ActivatedRoute,
    private rs: RequestService,
    private msg: NzMessageService
  ) { }

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
      amount: [0, []],
    });
  }
  handleExport() {
    const listColumns = [
      'ID',
      '名称',
      '用户名',
      '密码',
      '描述',
      '启用',
      '日期',
    ];
    const data: any[][] = [];
    data.push(listColumns);
    this.datum.forEach((item) => {
      const arr = [];
      arr.push(item.id);
      arr.push(item.name);
      arr.push(item.username);
      arr.push(item.password);
      arr.push(item.desc);
      arr.push(item.disabled);
      arr.push(String(item.created));
      data.push(arr);
    });
    let csvContent = 'data:text/csv;charset=utf-8,';
    data.forEach((row) => {
      csvContent += row.join(',') + '\n';
    });
    let encodedUri = encodeURI(csvContent);
    window.open(encodedUri);
  }
  submit() {
    if (this.group.valid) {
      let url = `gateway/create`;
      this.nzTitle = "创建中..."
      this.nzLoading = true;
      const resData: any[] = []
      for (let i = 0; i < this.group.value.amount; i++) {
        this.rs.post(url, this.group.value).subscribe((res) => {
          resData.push(res.data)
          if (resData.length === this.group.value.amount) {
            this.msg.success('创建成功');
            this.nzTitle = "保存"
            this.nzLoading = false
            this.datum = resData
          }
        });
      }

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
}
