import { Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from "@angular/forms";
import { NzMessageService } from "ng-zorro-antd/message";
import { RequestService } from "../../request.service";

@Component({
  selector: 'app-batch',
  templateUrl: './batch.component.html',
  styleUrls: ['./batch.component.scss']
})
export class BatchComponent implements OnInit {
  group!: FormGroup;
  isSpinning = false;
  btnTitle = '提交';
  datum: any[] = []
  constructor(
    private fb: FormBuilder,
    private msg: NzMessageService,
    private rs: RequestService,
  ) { }
  ngOnInit(): void {
    this.build();
  }
  build(obj?: any) {
    obj = obj || {}
    this.group = this.fb.group({
      product_id: [obj.product_id || '', []],
      group_id: [obj.group_id || '', []],
      name: [obj.name || '', [Validators.required]],
      desc: [obj.desc || '', []],
      amount: [obj.amount || 1, []],
    })
  }
  handleSubmit() {
    if (this.group.valid) {
      const sendData = Object.assign({}, this.group.value);
      const amount = this.group.value.amount;
      const resData: any = [];
      this.isSpinning = true;
      this.btnTitle = '创建中...';
      for (let index = 0; index < amount; index++) {
        this.rs.post(`device/create`, sendData).subscribe(res => {
          resData.push(res.data);
          if (resData.length === amount) {
            this.isSpinning = false;
            this.btnTitle = '提交';
            this.msg.success("创建成功!");
            this.datum = resData;
          }
        })
      }
    }
  }
  handleExport() {
    const listColumns = ['ID', '产品ID', '分组ID', '名称', '说明', '日期'];
    const data: any[][] = [];
    data.push(listColumns);
    this.datum.forEach(item => {
      const arr = [];
      arr.push(item.id);
      arr.push(item.product_id);
      arr.push(item.group_id);
      arr.push(item.name);
      arr.push(item.desc);
      arr.push(item.created);
      data.push(arr);
    });
    let csvContent = 'data:text/csv;charset=utf-8,';
    data.forEach(row => { csvContent += row.join(',') + '\n'; });
    let encodedUri = encodeURI(csvContent);
    window.open(encodedUri);
  }
}
