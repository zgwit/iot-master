 import { Component, OnInit } from '@angular/core';
  import { FormBuilder, Validators, FormGroup } from "@angular/forms";
  import { ActivatedRoute, Router } from "@angular/router";
  import { RequestService } from "src/app/request.service";
  import { NzMessageService } from "ng-zorro-antd/message";
  import { isIncludeAdmin } from "src/public"; 
  @Component({
    selector: 'app-device-type-edit',
    templateUrl: './device-type-edit.component.html',
    styleUrls: ['./device-type-edit.component.scss']
  })
  export class DeviceTypeEditComponent  implements OnInit {
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
        this.rs.get(`device/type/${this.id}`).subscribe(res => {
          //let data = res.data;
          this.build(res.data)
        })
  
      }
  
      this.build()
    }
  
    build(obj?: any) {
      obj = obj || {}
      this.group = this.fb.group({ 
        name: [obj.name || '' ,[]], 
        desc: [obj.desc || '', []],  
      })
    }
  
    submit() {
  
      if (this.group.valid) {
  
        let url = this.id ? `device/type/${this.id}` : `device/type/create`;  
        this.rs.post(url, this.group.value).subscribe(res => { 
          const path = `${isIncludeAdmin()}/device/type`;
          this.router.navigateByUrl(path)
          this.msg.success("保存成功")
        })
  
        return;
      }
      else {
        Object.values(this.group.controls).forEach(control => {
          if (control.invalid) {
            control.markAsDirty();
            control.updateValueAndValidity({ onlySelf: true });
          }
        });
  
      }
    }
    handleCancel() {
      const path = `${isIncludeAdmin()}/device/type`;
      this.router.navigateByUrl(path);
    }
  
  }
  
