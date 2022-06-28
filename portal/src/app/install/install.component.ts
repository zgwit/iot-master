import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {Router} from "@angular/router";
import {RequestService} from "../request.service";
import {InfoService} from "../info.service";

function randomString(length: number) {
  const str = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
  let result = '';
  for (let i = length; i > 0; --i)
    result += str[Math.floor(Math.random() * str.length)];
  return result;
}

@Component({
  selector: 'app-install',
  templateUrl: './install.component.html',
  styleUrls: ['./install.component.scss']
})
export class InstallComponent implements OnInit {
  current = 0;

  baseForm!: FormGroup;
  databaseForm!: FormGroup;
  historyForm!: FormGroup;


  constructor(private fb: FormBuilder,
              private route: Router,
              private rs: RequestService,
              private is: InfoService,
              ) {
    this.buildForms()
  }

  ngOnInit(): void {
  }

  buildForms(): void {
    let node = 'iot-master-' + randomString(5)
    this.baseForm = this.fb.group({
      node: [node, [Validators.required]],
      data: ['data', [Validators.required]],
      port: ['8080', [Validators.required]],
      default_password: ['123456', [Validators.required]],
    })
    this.databaseForm = this.fb.group({
      type: ['sqlite', [Validators.required]],
      url: ['sqlite.db', [Validators.required]],
    })
    this.historyForm = this.fb.group({
      type: ['embed', [Validators.required]],
      options: [{"data_path": "history"}, [Validators.required]],
    })

  }

  submitBase() {
    this.rs.post("install/base", this.baseForm.value).subscribe(res=>{
      this.current++
    })
  }

  submitDatabase() {
    this.rs.post("install/database", this.databaseForm.value).subscribe(res=>{
      this.current++
    })
  }

  submitHistory() {
    this.rs.post("install/history", this.historyForm.value).subscribe(res=>{
      this.current++
    })
  }

  install() {
    this.rs.get("install/system").subscribe(res=>{
      this.is.info.installed = true //强制赋值已经安装，不太优雅~~
      this.route.navigate(["/admin"])
    })
  }

}
