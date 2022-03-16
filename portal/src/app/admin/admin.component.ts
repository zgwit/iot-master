import {Component,  OnInit} from '@angular/core';
import {SideMenu} from './side.menu';
import {RequestService} from '../request.service';
import {UserService} from "../user.service";
import {Router} from '@angular/router';


@Component({
  selector: 'app-main',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.scss']
})
export class AdminComponent implements OnInit {

  isCollapsed = false;

  menus: Array<any> = [];

  tabs: Array<any> = [{url: 'welcome'}]

  constructor(private rs: RequestService, public us: UserService, private route: Router) {
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
}
