import {Component,  OnInit} from '@angular/core';
import {SideMenu} from './side.menu';
import {RequestService} from '../request.service';
import {UserService} from "../user.service";
import {Router} from '@angular/router';
import {NzModalService} from "ng-zorro-antd/modal";
import {PasswordComponent} from "./password/password.component";
import {InfoService} from "../info.service";


@Component({
  selector: 'app-main',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.scss']
})
export class AdminComponent implements OnInit {

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

  ngOnInit(): void {
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
