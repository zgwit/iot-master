import { Component, Optional } from '@angular/core';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { Router } from '@angular/router';
import { RequestService } from '../../request.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzTableQueryParams } from 'ng-zorro-antd/table';
import { ParseTableQuery } from '../../base/table';
import { ProductEditComponentComponent } from '../product-edit-component/product-edit-component.component';
import {
    isIncludeAdmin,
    tableHeight,
    onAllChecked,
    onItemChecked,
    batchdel,
    refreshCheckedStatus,
} from '../../../public';
@Component({
    selector: 'app-products',
    templateUrl: './products.component.html',
    styleUrls: ['./products.component.scss'],
})
export class ProductsComponent {
    isVisible = false;
    chooseData: any = {};
    href!: string;
    loading = false;
    datum: any[] = [];
    total = 1;
    pageSize = 20;
    pageIndex = 1;
    query: any = {};
    showAddBtn: Boolean = true;
    columnKeyNameArr: any = ['name', 'desc'];
    uploading: Boolean = false;
    checked = false;
    indeterminate = false;
    setOfCheckedId = new Set<number>();
    delResData: any = [];

    constructor(
        private modal: NzModalService,
        @Optional() protected ref: NzModalRef,
        private router: Router,
        private rs: RequestService,
        private msg: NzMessageService
    ) {
        this.pageSizeChange(this.pageSize);
        //this.load();
    }

    reload() {
        this.datum = [];
        this.load();
    }
    informate(mes: any) {
        this.isVisible = true;

        this.rs
            .get(`product/${mes.id}`, this.query)
            .subscribe((res) => {
                this.chooseData = res.data;
            })
            .add(() => {
                this.loading = false;
            });
    }
    load() {
        this.loading = true;
        this.rs
            .post('product/search', this.query)
            .subscribe((res) => {
                this.datum = res.data || [];
                this.total = res.total;

                //  this.setOfCheckedId.clear();
                //  refreshCheckedStatus(this);
            })
            .add(() => {
                this.loading = false;
            });
    }

    create() {
        let path = '/product/create';
        if (location.pathname.startsWith('/admin')) path = '/admin' + path;
        this.router.navigateByUrl(path);
    }

    delete(id: number, size?: number) {
        this.rs.get(`product/${id}/delete`).subscribe((res) => {
            if (!size) {
                this.msg.success('删除成功');
                this.datum = this.datum.filter((d) => d.id !== id);
                this.isVisible = false;
                this.load();
                return;
            }
            this.delResData.push(res);
            if (size === this.delResData.length) {
                this.msg.success('删除成功');
                this.isVisible = false;
                this.load();
                this.delResData = [];
            }
        });
    }
    close() {
        this.isVisible = false;
    }
    onQuery($event: NzTableQueryParams) {
        ParseTableQuery($event, this.query);
        this.load();
    }
    pageIndexChange(pageIndex: number) {
        this.query.skip = pageIndex - 1;
    }
    pageSizeChange(pageSize: number) {
        this.query.limit = pageSize;
    }
    pageIndexCardChange(pageIndex: number) {
        this.query.skip = (pageIndex - 1) * this.pageSize;
        this.load();
    }
    pageSizeCardChange(pageSize: number) {
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

    edit(id: any) {
        const path = `${isIncludeAdmin()}/product/edit/${id}`;
        this.router.navigateByUrl(path);
    }

    handleNew() {
        const path = `${isIncludeAdmin()}/product/create`;
        this.router.navigateByUrl(path);
    }

    handleEdit(id?: string) {
        this.isVisible = false;
        const nzTitle = id ? '编辑产品' : '创建产品';
        const modal: NzModalRef = this.modal.create({
            nzTitle,
            nzStyle: { top: '20px' },
            nzWidth: '80%',
            nzContent: ProductEditComponentComponent,
            nzComponentParams: { id },
            nzMaskClosable: false,
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
                                this.load();
                            },
                            () => { }
                        );
                    },
                },
            ],
        });
    }

    select(obj: object) {
        this.ref && this.ref.close(obj);
    }
    cancel() {
        this.msg.info('取消操作');
    }
    handleExport() {
        this.href = `/api/product/export`;
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
