import {Component, OnInit, ViewChild} from '@angular/core';
import {NzButtonComponent} from 'ng-zorro-antd/button';
import {ActivatedRoute, Router, RouterLink} from '@angular/router';
import {RequestService, SmartEditorComponent, SmartField} from 'iot-master-smart';
import {NzMessageService} from 'ng-zorro-antd/message';
import {CommonModule} from '@angular/common';
import {NzCardComponent} from "ng-zorro-antd/card";
import {GetParentRouteParam, GetParentRouteUrl} from "../../../app.routes";

@Component({
    selector: 'app-project-edit',
    standalone: true,
    imports: [
        CommonModule,
        NzButtonComponent,
        RouterLink,
        NzCardComponent,
        SmartEditorComponent,
    ],
    templateUrl: './project-edit.component.html',
    styleUrl: './project-edit.component.scss',
})
export class ProjectEditComponent implements OnInit {
    base = '/admin'
    project_id!: any;

    id: any = '';

    @ViewChild('form') form!: SmartEditorComponent

    fields: SmartField[] = [
        {key: "id", label: "ID", type: "text", min: 2, max: 30, placeholder: "选填"},
        {key: "name", label: "名称", type: "text", required: true, default: '新项目'},
        {key: "keywords", label: "关键字", type: "tags", default: []},
        {key: "url", label: "链接", type: "text"},
        {key: "description", label: "说明", type: "textarea"},
    ]

    data: any = {}


    constructor(private router: Router,
                private msg: NzMessageService,
                private rs: RequestService,
                private route: ActivatedRoute
    ) {
    }

    ngOnInit(): void {
        this.base = GetParentRouteUrl(this.route)
        this.project_id ||= GetParentRouteParam(this.route, "project")

        if (this.route.snapshot.paramMap.has('id')) {
            this.id = this.route.snapshot.paramMap.get('id');
        } else {
            this.id = this.project_id
        }
	
        if (this.id) {
            this.load()
        }
    }

    load() {
        this.rs.get(`project/` + this.id).subscribe((res) => {
            this.data = res.data
        });
    }

    onSubmit() {
        if (!this.form.valid) {
            this.msg.error('请检查数据')
            return
        }

        let url = `project/${this.id || 'create'}`
        this.rs.post(url, this.form.value).subscribe((res) => {
            if (this.project_id)
                this.router.navigateByUrl(`${this.base}/detail`);
            else
                this.router.navigateByUrl(`${this.base}/project/` + res.data.id);
            this.msg.success('保存成功');
        });
    }
}
