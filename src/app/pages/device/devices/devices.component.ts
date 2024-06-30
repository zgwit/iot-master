import {Component, Inject, Input, OnInit, Optional,} from '@angular/core';
import {ActivatedRoute,} from '@angular/router';
import {CommonModule} from '@angular/common';
import {
    ParamSearch,
    RequestService,
    SmartTableButton,
    SmartTableColumn,
    SmartTableComponent,
    SmartTableOperator
} from 'iot-master-smart';
import {NZ_MODAL_DATA, NzModalRef} from 'ng-zorro-antd/modal';
import {GetParentRouteParam, GetParentRouteUrl} from "../../../app.routes";

@Component({
    selector: 'app-devices',
    standalone: true,
    imports: [CommonModule, SmartTableComponent],
    templateUrl: './devices.component.html',
    styleUrl: './devices.component.scss',
})
export class DevicesComponent implements OnInit {
    base = "/admin"

    //从Modal中传参过来
    //readonly data: any = inject(NZ_MODAL_DATA, {optional:true});
    @Input() project_id: any = '';
    @Input() product_id: any = '';
    @Input() tunnel_id: any = '';


    datum: any[] = [];
    total = 0;
    loading = false;


    buttons: SmartTableButton[] = [
        {icon: 'plus', label: '创建', link: () => `${this.base}/device/create`},
    ];

    columns: SmartTableColumn[] = [
        {
            key: 'id', sortable: true, label: 'ID', keyword: true,
            link: (data) => `${this.base}/device/${data.id}`,
        },
        {key: 'name', sortable: true, label: '名称', keyword: true},
        {
            key: 'product', sortable: true, label: '产品', keyword: true,
            link: (data) => `${this.base}/product/${data.product_id}`,
        },
        {
            key: 'project', sortable: true, label: '项目', keyword: true,
            link: (data) => `${this.base}/project/${data.project_id}`,
        },
        {key: 'online', sortable: true, label: '上线时间', date: true},
    ];

    columnsSelect: SmartTableColumn[] = [
        {key: 'id', label: 'ID', keyword: true},
        {key: 'name', label: '名称', keyword: true},
    ];

    operators: SmartTableOperator[] = [
        {icon: 'edit', title: '编辑', link: (data) => `${this.base}/device/${data.id}/edit`,},
        {
            icon: 'delete', title: '删除', confirm: '确认删除？', action: (data) => {
                this.rs.get(`device/${data.id}/delete`).subscribe((res) => this.refresh())
            },
        },
    ];

    operatorsSelect: SmartTableOperator[] = [
        {label: '选择', action: (data) => this.ref.close(data)},
    ];


    constructor(
        private route: ActivatedRoute,
        private rs: RequestService,
        @Optional() protected ref: NzModalRef,
        @Optional() @Inject(NZ_MODAL_DATA) protected data: any
    ) {
        this.project_id = data?.project_id;
    }

    ngOnInit(): void {
        this.base = GetParentRouteUrl(this.route)
        this.project_id ||= GetParentRouteParam(this.route, "project")
    }

    query!: ParamSearch

    refresh() {
        this.search(this.query)
    }

    search(query: ParamSearch) {
        //console.log('onQuery', query)        this.query = query
        if (this.product_id) query.filter['product_id'] = this.product_id;
        if (this.project_id) query.filter['project_id'] = this.project_id;
        if (this.tunnel_id) query.filter['tunnel_id'] = this.tunnel_id;

        this.loading = true;
        this.rs.post('device/search', query).subscribe((res) => {
            this.datum = res.data;
            this.total = res.total;
        }).add(() => this.loading = false);
    }
}
