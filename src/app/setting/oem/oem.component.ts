import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {ActivatedRoute, Router} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzMessageService} from "ng-zorro-antd/message";
import { NzUploadChangeParam,NzUploadFile } from 'ng-zorro-antd/upload';
@Component({
  selector: 'app-oem',
  templateUrl: './oem.component.html',
  styleUrls: ['./oem.component.scss']
})
export class OemComponent implements OnInit {
  group!: FormGroup;

  constructor(private fb: FormBuilder,
              private router: Router,
              private route: ActivatedRoute,
              private rs: RequestService,
              private msg: NzMessageService) {
  }


  ngOnInit(): void {
    // this.rs.get(`config`).subscribe(res => {
    //   //let data = res.data;
    //   this.build(res.data)
    // })

    this.build()
  }

  build(obj?: any) {
    obj = obj || {}
    this.group = this.fb.group({ 
      Title: ['', [Validators.required]],
      Logo: ['', [Validators.required]],
      Company: ['', [Validators.required]] ,
      Copyright: ['', [Validators.required]]
    })
  }
  fileList: NzUploadFile[] = [ 
  ];
  submit() {
    
    if (this.group.valid) {
 
      this.rs.post(`config`, this.group.value).subscribe(res => {
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
  handleChange(info: NzUploadChangeParam): void {
    let fileList = [...info.fileList];
 
    fileList = fileList.slice(-1);

    this.group.patchValue({Logo: fileList[0].originFileObj?.name})
   
    // fileList = fileList.map(file => {
    //   if (file.response) {
    //     // Component will show file.url as link
    //     file.url = file.response.url; 
    //   }
    //   return file;
    // });

    // this.fileList = fileList;
    switch (info.file.status) {
      case 'uploading': 
      this.msg.info( '加载中')
      break;
      case 'done': 
      this.msg.success('上传成功')
        break;
      case 'error':
        this.msg.error('上传失败'); 
        break;
        default:
          break
    }
  }
}
