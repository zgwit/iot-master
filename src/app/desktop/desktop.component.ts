import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { RequestService } from '../request.service';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { WindowComponent } from '../window/window.component';
import { OemService } from '../oem.service';
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
    
    constructor(
        private router: Router,
        private rs: RequestService,
        private ms: NzModalService,
        private us: UserService,
        protected os: OemService,
        protected _as: AppService
    ) {
        this.userInfo = us && us.user;
    }
    handlePassword() {
        const modal: NzModalRef = this.ms.create({
            nzTitle: '修改密码',
            nzCentered: true,
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
                            () => {}
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
                item.tab=true
            }
        });
        //this.show=false
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
            }
        });
    }
zindex(mes:any){ 
    this.items.filter((item: any, index: any) => {  item.index = 0;
    if (item.title === mes) {
        item.index = 9999;
    }
    console.log(this.items)
});}
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
                index:0
            });

        //   this.items.filter((item: any, index: any) => {console.log(1)
        //     if (item.title === app.name)    return
        //     console.log(index)
        //     console.log(this.items.length)
        //     if(index+1===this.items.length)
        //     this.items.push({
        //       show: true,
        //       entries: app.entries,
        //       title: app.name,
        //   });
        // });

        // this.show=true
        // this.entries=app.entries
        // this.title=app.name
        // this.ms.create({
        //   nzTitle: app.name,
        //   nzFooter: null,
        //   //nzMask: false,
        //   nzMaskClosable: false,
        //   nzWidth: "90%",
        //   //nzStyle: {height: "90%"},
        //   nzBodyStyle: { padding: "0", overflow: "hidden" },
        //   nzContent: WindowComponent,
        //   nzComponentParams: {
        //     entries: app.entries || []
        //   }
        // })
    }

    logout() {
        this.rs
            .get('logout')
            .subscribe((res) => {})
            .add(() => this.router.navigateByUrl('/login'));
    }
}
