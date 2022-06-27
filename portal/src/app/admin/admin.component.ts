import {Component, OnDestroy, OnInit} from '@angular/core';
import {SideMenu} from './side.menu';
import {RequestService} from '../request.service';
import {UserService} from "../user.service";
import {Router} from '@angular/router';
import {NzModalService} from "ng-zorro-antd/modal";
import {PasswordComponent} from "./password/password.component";
import {InfoService} from "../info.service";
import {ActiveComponent} from "./active/active.component";
import * as dayjs from "dayjs";


@Component({
  selector: 'app-main',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.scss']
})
export class AdminComponent implements OnInit, OnDestroy {

  isCollapsed = false;

  menus: Array<any> = [];

  tabs: Array<any> = [{url: 'welcome'}]

  constructor(private rs: RequestService,
              public us: UserService,
              public is: InfoService,
              private route: Router,
              private ms: NzModalService) {
    this.initMenu();
  }

  checkLicenseInterval!: any;
  checkLicenseTimeout!: any;

  ngOnInit(): void {
    this.checkLicenseInterval = setInterval(()=> this.checkLicense(), 1000 * 60 * 60) //每小时执行一次
    this.checkLicenseTimeout = setTimeout(()=> this.checkLicense(), 1000 * 10) //10秒检查一次
  }

  ngOnDestroy() {
    clearInterval(this.checkLicenseInterval)
    clearTimeout(this.checkLicenseTimeout)
  }

  active() {
    this.ms.create({
      nzTitle: "在线激活",
      nzContent: ActiveComponent,
      nzMaskClosable: false,
    })
  }

  checkLicense() {
    this.rs.get("license").subscribe(res=>{
      if (!res.data) {
        this.ms.error({
          nzContent: "产品未激活",
          nzOkText: "在线激活",
          nzOnOk: instance => {
            this.active()
          }
        })
        return
      }

      let lic = res.data;
      if (dayjs(lic.expireAt).isBefore(dayjs())) {
        this.ms.error({
          nzContent: "授权码已经失效",
          nzOkText: "在线激活",
          nzOnOk: instance => {
            this.active()
          }
        })
      }
    })
  }

  noop(): void {
  }


  initMenu(): void {
    this.menus = SideMenu;
  }

  closeTab($event: any) {
    this.tabs.splice($event.index, 1);
  }

  password() {
    this.ms.create({
      nzTitle:"修改密码",
      nzContent: PasswordComponent,
      nzFooter: ""
    })
  }

  exit() {
    this.rs.get("logout").subscribe(console.log)

    //localStorage.removeItem("token");
    return this.route.navigate(["/login"]);
  }
}
