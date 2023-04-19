import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from "@angular/forms";
import { ActivatedRoute, Router } from "@angular/router";
import { RequestService } from "../../request.service";
import { NzMessageService } from "ng-zorro-antd/message";
import { isIncludeAdmin } from "../../../public";

@Component({
  selector: 'app-brokers-edit',
  templateUrl: './broker-edit.component.html',
  styleUrls: ['./broker-edit.component.scss']
})
export class BrokerEditComponent implements OnInit {
  group!: FormGroup;
  id: any = 0

  constructor(private fb: FormBuilder,
    private router: Router,
    private route: ActivatedRoute,
    private rs: RequestService,
    private msg: NzMessageService) {
  }


  ngOnInit(): void {
    if (this.route.snapshot.paramMap.has("id")) {
      this.id = this.route.snapshot.paramMap.get("id");
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
      id:[obj.id || '', [ ]],
      name: [obj.name || '', [Validators.required]],
      desc: [obj.desc || '', []],
      port: [obj.port || 1883, []],
      cert: [obj.cert || '', []],
      key: [obj.key || '', []],
    })
  }

  submit() {
    if (this.group.valid) {
        
      let url = this.id ? `broker/${this.id}` : `broker/create`
      this.rs.post(url, this.group.value).subscribe(res => {
        const path = `${isIncludeAdmin()}/broker/list`;
        this.router.navigateByUrl(path)
        this.msg.success("保存成功")
      })
    } else {
      Object.values(this.group.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  handleCancel() {
    const path = `${isIncludeAdmin()}/broker/list`;
    this.router.navigateByUrl(path)
  }
}
