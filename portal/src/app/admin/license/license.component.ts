import { Component, OnInit } from '@angular/core';
import {RequestService} from "../../request.service";
import {ActiveComponent} from "../active/active.component";
import {NzModalService} from "ng-zorro-antd/modal";

@Component({
  selector: 'app-license',
  templateUrl: './license.component.html',
  styleUrls: ['./license.component.scss']
})
export class LicenseComponent implements OnInit {

  license!: any

  constructor(private rs: RequestService, private ms: NzModalService) {
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
    this.rs.get("/license").subscribe(res=>{
      if (!res.data) {
        this.active()
        return
      }
      this.license = res.data;
    })
  }

  ngOnInit(): void {
  }

}
