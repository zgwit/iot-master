import { Component, OnInit } from '@angular/core';
import {RequestService} from "../../request.service";
import {ActiveComponent} from "../active/active.component";
import {NzModalService} from "ng-zorro-antd/modal";
import {ChooseService} from "../choose.service";

@Component({
  selector: 'app-license',
  templateUrl: './license.component.html',
  styleUrls: ['./license.component.scss']
})
export class LicenseComponent implements OnInit {

  license!: any

  constructor(private rs: RequestService, private ms: NzModalService, private cs: ChooseService) {
    this.load()
  }

  active() {
    this.ms.create({
      nzTitle: "在线激活",
      nzContent: ActiveComponent,
      nzMaskClosable: false,
    })
  }

  load() {
    this.rs.get("license").subscribe(res=>{
      if (!res.data) {
        this.active()
        return
      }
      this.license = res.data;
    })
  }

  input() {
    this.cs.prompt({message: "输入激活码"}).subscribe(res=>{
      this.rs.post("license", {license: res}).subscribe(res=>{
      })
    })

  }

  ngOnInit(): void {
  }

}
