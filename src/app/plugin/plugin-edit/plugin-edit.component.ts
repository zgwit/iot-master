import { Component, Input, OnInit } from '@angular/core';
import { FormBuilder, Validators, FormGroup } from "@angular/forms";
import { RequestService } from "../../request.service";
import { NzMessageService } from "ng-zorro-antd/message";

@Component({
    selector: 'app-plugin-edit',
    templateUrl: './plugin-edit.component.html',
    styleUrls: ['./plugin-edit.component.scss']
})
export class PluginEditComponent implements OnInit {
    group!: FormGroup;
    // id: any = 0
    @Input() id = '';
    constructor(private fb: FormBuilder,
        private rs: RequestService,
        private msg: NzMessageService) {
    }

    ngOnInit(): void {
        if (this.id) {
            this.rs.get(`plugin/${this.id}`).subscribe(res => {
                //let data = res.data;
                this.build(res.data)
            })

        }

        this.build()
    }

    build(obj?: any) {
        obj = obj || {}
        this.group = this.fb.group({
            id: [obj.id || '', []],
            name: [obj.name || '', [Validators.required]],
            version: [obj.version || '', []],
            username: [obj.username || '', []],
            password: [obj.password || '', []],
            external: [obj.external || false, []],
            command: [obj.command || '', []],
            dependencies: [obj.dependencies || {}, []],
            disabled: [obj.disabled || false, []],
        })
    }

    submit() {
        return new Promise((resolve, reject) => {
            if (this.group.valid) {
                let url = this.id ? `plugin/${this.id}` : `plugin/create`
                this.rs.post(url, this.group.value).subscribe(res => {
                    this.msg.success("保存成功");
                    resolve(true);
                })

                return;
            } else {
                Object.values(this.group.controls).forEach(control => {
                    if (control.invalid) {
                        control.markAsDirty();
                        control.updateValueAndValidity({ onlySelf: true });
                        reject();
                    }
                });

            }
        })
    }

}
