import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { RequestService } from '../request.service';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { AppService } from '../app.service';
import { UserService } from '../user.service';
import { PasswordComponent } from '../user/password/password.component';
declare var window: any;

@Component({
    selector: 'app-desktop',
    templateUrl: './desktop.component.html',
    styleUrls: ['./desktop.component.scss'],
})
export class DesktopComponent {
    title: any;
    show: any;
    entries: any = [];
    items: any[] = [];
    userInfo: any;
    oem: any = {
        title: '物联大师',
        logo: '/assets/logo.png',
        company: '无锡真格智能科技有限公司',
        copyright: '©2016-2023'
    }
    constructor(
        private router: Router,
        private rs: RequestService,
        private ms: NzModalService,
        private us: UserService,
        protected _as: AppService
    ) {
        this.userInfo = us && us.user;

    }
    handlePassword() {
        const modal: NzModalRef = this.ms.create({
            nzTitle: '修改密码',
            nzCentered: true,
            nzMaskClosable: false,
            nzContent: PasswordComponent,
            nzFooter: [
                {
                    label: '取消',
                    onClick: () => {
                        modal.destroy();
                    },
                },
                {
                    label: '保存',
                    type: 'primary',
                    onClick: (componentInstance) => {
                        componentInstance!.submit().then(
                            () => {
                                modal.destroy();
                            },
                            () => { }
                        );
                    },
                },
            ],
        });
    }

    hide(mes: any) {
        this.items.filter((item: any, index: any) => {
            if (item.title === mes) {
                item.show = false;
                item.tab = true;
            }
        });
    }
    setIndex(mes: any) {
        this.items.filter((item: any, index: any) => {
            item.index = 0;
            if (item.title === mes) {
                item.index = 9999;
            }
        });
    }
    close(mes: any) {
        this.items.filter((item: any, index: any) => {
            if (item.title === mes) {
                this.items.splice(index, 1);
            }
        });
    }

    showTab(mes: any) {
        this.items.filter((item: any, index: any) => {
            if (item.title === mes) {
                item.show = true;
                item.tab = false;
            }
        });
        this.setIndex(mes)

    }

    open(app: any) {
        if (window.innerWidth < 800) {
            this.router.navigate([app.entries[0].path]);
            return;
        }

        if (
            !this.items.some((item: any) => {
                return item.title === app.name;
            })
        )
            this.items.push({
                show: true,
                entries: app.entries,
                title: app.name,
                index: 0,
            });
        this.showTab(app.name)
        this.setIndex(app.name);

    }

    logout() {
        this.rs
            .get('logout')
            .subscribe((res) => { })
            .add(() => this.router.navigateByUrl('/login'));
    }
}
