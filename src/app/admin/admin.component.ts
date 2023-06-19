import {Component} from '@angular/core';
import {AppService} from "../app.service"; 
import {Router} from "@angular/router";
import {UserService} from "../user.service";
import {RequestService} from "../request.service";

@Component({
    selector: 'app-admin',
    templateUrl: './admin.component.html',
    styleUrls: ['./admin.component.scss']
})
export class AdminComponent {
    userInfo: any;

    constructor(
        protected _as: AppService, 
        private router: Router,
        private us: UserService,
        private rs: RequestService
    ) {
        this.userInfo = us && us.user;
        const menuList = _as.apps;

        for (let index = 0; index < menuList.length; index++) {
            const item = menuList[index];
            const {entries} = item;
            for (let i = 0; i < entries.length; i++) {
                const it = entries[i];
                if (`/admin${it.path}` === location.pathname) {
                    item.open = true;
                }
            }
        }
    }

    logout() {
        this.rs.get("logout").subscribe(res => {
        }).add(() => this.router.navigateByUrl("/login"))
    }
}
