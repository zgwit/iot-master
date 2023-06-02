import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { RequestService } from '../../request.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzTableQueryParams } from 'ng-zorro-antd/table';
import { NzModalService } from 'ng-zorro-antd/modal';
import { ParseTableQuery } from '../../base/table';
import {
    tableHeight,
    onAllChecked,
    onItemChecked,
    batchdel,
    refreshCheckedStatus,
} from '../../../public';

@Component({
    selector: 'app-validators',
    templateUrl: './validators.component.html',
    styleUrls: ['./validators.component.scss'],
})
export class ValidatorsComponent {
    loading = true;
    datum: any[] = [];
    total = 1;
    pageSize = 20;
    uploading: Boolean = false;
    pageIndex = 1;
    query: any = {};
    url = '';
    href!: string;
    filterLevel = [
        { text: '1', value: 1 },
        { text: '2', value: 2 },
        { text: '3', value: 3 },
    ];
    checked = false;
    indeterminate = false;
    setOfCheckedId = new Set<number>();
    delResData: any = [];

    constructor(
        private router: Router,
        private rs: RequestService,
        private modal: NzModalService,
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
            .post('validator/search', this.query)
            .subscribe((res) => {
                const { data, total } = res;
                this.datum = data || [];
                this.total = total || 0;
                this.setOfCheckedId.clear();
                refreshCheckedStatus(this);
            })
            .add(() => {
                this.loading = false;
            });
    }

    edit(id: any) {
        const path = `/validator/edit/${id}`;
        this.router.navigateByUrl(path);
    }
    handleNew() {
        const path = `/validator/create`;
        this.router.navigateByUrl(path);
    }

    delete(id: number, size?: number) {
        this.rs.get(`validator/${id}/delete`).subscribe((res) => {
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
        this.query.skip = pageIndex - 1;
    }

    pageSizeChange(pageSize: number) {
        this.query.limit = pageSize;
    }

    search($event: string) {
        console.log()
        this.query.keyword = {
            title: $event,
            //   Message: $event,
        };

        this.query.skip = 0;
        this.load();
    }

    handleImport(e: any) {
        const file: File = e.target.files[0];
        const formData = new FormData();
        formData.append('file', file);
        this.rs
            .post(`validator/import`, formData)
            .subscribe((res) => {
                console.log(res);
            });
    }



    read(data: any) {

    }

    disable(mess: number, id: any) {
        if (mess)
            this.rs
                .get(`validator/${id}/disable`)
                .subscribe((res) => {
                    this.reload();
                });
        else
            this.rs
                .get(`validator/${id}/enable`)
                .subscribe((res) => {
                    this.reload();
                });
    }
    cancel() {
        this.msg.info('取消操作');
    }

    getTableHeight() {
        return tableHeight(this);
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
