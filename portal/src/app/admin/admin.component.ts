import {Component,  OnInit} from '@angular/core';
import {SideMenu} from './side.menu';
import {RequestService} from '../request.service';
import {UserService} from "../user.service";
import {Router} from '@angular/router';
import {NzModalService} from "ng-zorro-antd/modal";
import {PasswordComponent} from "./password/password.component";


@Component({
  selector: 'app-main',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.scss']
})
export class AdminComponent implements OnInit {

  isCollapsed = false;

  menus: Array<any> = [];

  tabs: Array<any> = [{url: 'welcome'}]

  version: any = {
    version: "1.0.0"
  }

  constructor(private rs: RequestService, public us: UserService, private route: Router, private ms: NzModalService) {
    this.initMenu();
  }

  ngOnInit(): void {
    this.rs.get("system/version").subscribe(res=>{
      this.version = res.data
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
