import { Component } from '@angular/core';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { Router } from '@angular/router';
import { RequestService } from '../../request.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzTableQueryParams } from 'ng-zorro-antd/table';
import { ParseTableQuery } from '../../base/table';
import { PluginEditComponent } from "../plugin-edit/plugin-edit.component"
import {
    onAllChecked,
    onItemChecked,
    batchdel,
    refreshCheckedStatus,
} from '../../../public';

@Component({
    selector: 'app-plugins',
    templateUrl: './plugins.component.html',
    styleUrls: ['./plugins.component.scss'],
})
export class PluginsComponent {
    href!: string;
    loading = true;
    datum: any[] = [];
    total = 1;
    pageSize = 20;
    pageIndex = 1;
    query: any = {};
    checked = false;
    indeterminate = false;
    setOfCheckedId = new Set<number>();
    delResData: any = [];

    constructor(
        private modal: NzModalService,
        private router: Router,
        private rs: RequestService,
        private msg: NzMessageService
    ) {
        //this.load();
    }

    reload() {
        this.datum = [];
        this.load();
    }

    load() {
        this.loading = true;
        this.rs
            .post('plugin/search', this.query)
            .subscribe((res) => {
                this.datum = res.data || [];
                this.total = res.total;
                this.setOfCheckedId.clear();
                refreshCheckedStatus(this);
            })
            .add(() => {
                this.loading = false;
            });
    }

    enable(id: any) {
        this.rs.get(`plugin/${id}/enable`).subscribe((res) => {
            this.reload();
        });
    }

    disable(id: any) {
        this.rs.get(`plugin/${id}/disable`).subscribe((res) => {
            this.reload();
        });
    }

    plugin(num: number, id: any) {
        switch (num) {
            case 0: {
                this.rs.get(`plugin/${id}/start`).subscribe((res) => {
                    this.load();
                });
            }
                break;
            case 1: {
                this.rs.get(`plugin/${id}/stop`).subscribe((res) => {
                    this.load();
                });
            }
                break;
            case 2: {
                this.rs.get(`plugin/${id}/restart`).subscribe((res) => {
                    this.load();
                });
            }
                break;
            default:
                break;
        }
    }

    create() {
        let path = '/plugin/create';
        if (location.pathname.startsWith('/admin')) path = '/admin' + path;
        this.router.navigateByUrl(path);
    }

    handleExport() {
        this.href = `/api/plugin/export`;
    }

    delete(id: number, size?: number) {
        this.rs.get(`plugin/${id}/delete`).subscribe((res) => {
            if (!size) {
                this.msg.success('删除成功');
                this.datum = this.datum.filter((d) => d.id !== id);
            } else if (size) {
                this.delResData.push(res);
                if (size === this.delResData.length) {
                    this.msg.success('删除成功');
                    this.load();
                }
            }
        });
    }

    onQuery($event: NzTableQueryParams) {
        ParseTableQuery($event, this.query);
        this.load();
    }

    pageIndexChange(pageIndex: number) {
        console.log('pageIndex:', pageIndex);
    }

    pageSizeChange(pageSize: number) {
        this.query.limit = pageSize;
        this.load();
    }

    search($event: string) {
        this.query.keyword = {
            name: $event,
        };
        this.query.skip = 0;
        this.load();
    }

    cancel() {
        this.msg.info('取消操作');
    }

    handleEdit(id?: string) {
        const nzTitle = id ? "编辑插件" : "创建插件";
        const modal: NzModalRef = this.modal.create({
            nzTitle,
            nzStyle: { top: '20px' },
            nzContent: PluginEditComponent,
            nzComponentParams: { id },
            nzFooter: [
                {
                    label: '取消',
                    onClick: () => {
                        modal.destroy();
                    }
                },
                {
                    label: '保存',
                    type: 'primary',
                    onClick: componentInstance => {
                        componentInstance!.submit().then(() => {
                            modal.destroy();
                            this.load();
                        }, () => { })
                    }
                }
            ]
        });
    }

    handleBatchDel() {
        batchdel(this);
    }

    handleAllChecked(id: any) {
        onAllChecked(id, this);
    }

    handleItemChecked(id: number, checked: boolean) {
        onItemChecked(id, checked, this);
    }
}
