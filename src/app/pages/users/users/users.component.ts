import {Component, Optional} from '@angular/core';
import {NzModalRef,} from 'ng-zorro-antd/modal';
import {CommonModule} from '@angular/common';
import {
    ParamSearch,
    RequestService,
    SmartTableButton,
    SmartTableColumn,
    SmartTableComponent,
    SmartTableOperator
} from "iot-master-smart";

@Component({
    selector: 'app-users',
    standalone: true,
    imports: [
        CommonModule, SmartTableComponent,
    ],
    templateUrl: './users.component.html',
    styleUrl: './users.component.scss',
})
export class UsersComponent {
    datum: any[] = [];
    total = 0;
    loading = false;

    buttons: SmartTableButton[] = [
        {icon: 'plus', label: '创建', link: () => `/admin/user/create`},
    ];

    columns: SmartTableColumn[] = [
        {key: 'id', sortable: true, label: 'ID', keyword: true,},
        {key: 'name', sortable: true, label: '姓名', keyword: true},
        {key: 'cellphone', sortable: true, label: '手机', keyword: true,},
        {key: 'email', sortable: true, label: '邮箱', keyword: true,},
        {key: 'disabled', sortable: true, label: '状态', keyword: true,},
        {key: 'created', sortable: true, label: '创建时间', date: true},
    ];

    columnsSelect: SmartTableColumn[] = [
        {key: 'id', label: 'ID', keyword: true},
        {key: 'name', label: '名称', keyword: true},
    ];

    operators: SmartTableOperator[] = [
        {icon: 'edit', title: '编辑', link: (data) => `/admin/user/${data.id}/edit`,},
        {
            icon: 'delete', title: '删除', confirm: '确认删除？', action: (data) => {
                this.rs.get(`user/${data.id}/delete`).subscribe((res) => this.refresh())
            },
        },
    ];

    operatorsSelect: SmartTableOperator[] = [
        {label: '选择', action: (data) => this.ref.close(data)},
    ];

    constructor(private rs: RequestService, @Optional() protected ref: NzModalRef) {
    }


    query!: ParamSearch

    refresh() {
        this.search(this.query)
    }

    search(query: ParamSearch) {
        //console.log('onQuery', query)
        this.query = query
        this.loading = true
        this.rs.post('user/search', query).subscribe((res) => {
            this.datum = res.data;
            this.total = res.total;
        }).add(() => this.loading = false);
    }

}
