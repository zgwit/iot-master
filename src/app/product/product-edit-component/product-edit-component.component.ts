import { Component, Input, OnInit } from '@angular/core';
import { FormBuilder, Validators } from "@angular/forms";
import { Router } from "@angular/router";
import { RequestService } from "../../request.service";
import { NzMessageService } from "ng-zorro-antd/message";
import { isIncludeAdmin } from "../../../public";
import { NzUploadChangeParam } from 'ng-zorro-antd/upload';
import { EditTableItem } from "../../base/edit-table/edit-table.component";

@Component({
    selector: 'app-product-edit-component',
    templateUrl: './product-edit-component.component.html',
    styleUrls: ['./product-edit-component.component.scss']
})
export class ProductEditComponentComponent implements OnInit {
    group!: any;
    listData: EditTableItem[] = [{
        label: '名称(ID)',
        name: 'name'
    }, {
        label: '显示',
        name: 'label'
    }, {
        label: '类型',
        name: 'type',
        type: 'select',
        default: 'int',
        options: [{
            label: '整数',
            value: 'int'
        }, {
            label: '浮点数',
            value: 'float'
        }, {
            label: '布尔型',
            value: 'bool'
        }, {
            label: '文本',
            value: 'text'
        }, {
            label: '枚举',
            value: 'enum'
        }, {
            label: '数组',
            value: 'array'
        }, {
            label: '对象',
            value: 'object'
        }]
    }, {
        label: '单位',
        name: 'unit'
    }, {
        label: '模式',
        name: 'mode',
        type: 'select',
        default: 'rw',
        options: [{
            label: '只读',
            value: 'r'
        }, {
            label: '读写',
            value: 'rw'
        }]
    }]

    parameterslistData: EditTableItem[] = [
        {
            label: '名称(ID)',
            name: 'name'
        }, {
            label: '显示',
            name: 'label'
        }, {
            label: '最大值',
            name: 'max',
            type: 'number',
            default: 0
        }, {
            label: '最小值',
            name: 'min',
            type: 'number',
            default: 0
        }, {
            label: '默认值',
            name: 'default',
            type: 'number',
            default: 0
        }
    ]

    constraintslistData: EditTableItem[] = [
        {
            label: '等级',
            name: 'level'
        },
        {
            label: '标题',
            name: 'label'
        },
        {
            label: '模板',
            name: 'template'
        },
        {
            label: '表达式',
            name: 'expression'
        },
        {
            label: '报警延迟s',
            name: 'delay',
            type: 'number'
        },
        {
            label: '再次提醒',
            name: 'repeat',
            type: 'bool'
        },
        {
            label: '再次提醒延迟s',
            name: 'repeat_delay',
            type: 'number'
        },
        {
            label: '总提醒次数',
            name: 'repeat_total',
            type: 'number'
        }
    ]

    validatorsListData: EditTableItem[] = [
        {
            label: '标题',
            name: 'title'
        },
        {
            label: '等级',
            name: 'level',
            type: 'number',
            default: 0
        },
        {
            label: '模板',
            name: 'template'
        },
        {
            label: '表达式',
            name: 'expression'
        },
        {
            label: '报警延迟s',
            name: 'delay',
            type: 'number',
            default: 5
        },
        {
            label: '再次提醒',
            name: 'repeat',
            type: 'bool',
            default: true
        },
        {
            label: '再次提醒延迟s',
            name: 'repeat_delay',
            type: 'number',
            default: 300 //5分钟
        },
        {
            label: '总提醒次数',
            name: 'repeat_total',
            type: 'number',
            default: 10
        }
    ]
    aggregatorsListData: EditTableItem[] = [
        {
            label: '定时计划',
            name: 'crontab'
        },
        {
            label: '表达式',
            name: 'expression'
        },
        {
            label: '类型',
            name: 'type',
            type: 'select',
            default: 'inc',
            options: [{
                label: '增加',
                value: 'inc'
            }, {
                label: '减少',
                value: 'dec'
            }, {
                label: '平均',
                value: 'avg'
            }, {
                label: '计数',
                value: 'count'
            }, {
                label: '最小',
                value: 'min'
            }, {
                label: '最大',
                value: 'max'
            }, {
                label: '求合',
                value: 'sum'
            }, {
                label: '最后',
                value: 'last'
            }, {
                label: '最前',
                value: 'first'
            }]
        }, {
            label: '赋值',
            name: 'assign'
        }
    ]
    @Input() id!: any;

    constructor(
        private fb: FormBuilder,
        private router: Router,
        private rs: RequestService,
        private msg: NzMessageService
    ) {
    }

    ngOnInit(): void {
        if (this.id) {
            this.rs.get(`product/${this.id}`).subscribe(res => {
                this.setData(res.data || {});
            })
        }
        this.build()
    }
    setData(resData: any) {
        const odata = this.group.value;
        for (const key in odata) {
            if (resData[key]) {
                odata[key] = resData[key];
            }
        }
        this.group.setValue(odata);
    }
    build(obj?: any) {
        obj = obj || {}
        this.group = this.fb.group({
            id: [obj.id || '', []],
            name: [obj.name || '', [Validators.required]],
            desc: [obj.desc || '', []],
            version: [obj.version || '', []],
            properties: [obj.properties || [], []],
            parameters: [obj.parameters || [], []],
            constraints: [obj.constraints || [], []],
            validators: [obj.validators || [], []],
            aggregators: [obj.aggregators || [], []],
        })
    }

    submit() {
        // console.log(this.group.value)
        return new Promise((resolve) => {
            if (this.group.valid) {
                let url = this.id ? `product/${this.id}` : `product/create`;
                const { desc } = this.group.value;
                if (desc.length > 200) {
                    this.msg.warning("【说明】字数过长");
                    return;
                }
                this.rs.post(url, this.group.value).subscribe(res => {
                    this.msg.success("保存成功");
                    resolve(true);
                })
            }
        })

    }

    handleChange(info: NzUploadChangeParam): void {
        if (info.file.status !== 'uploading') {
            // console.log(info.file, info.fileList);
        }
        if (info.file.status === 'done') {
            this.msg.success(`${info.file.name} 文件上传成功`);
        } else if (info.file.status === 'error') {
            this.msg.error(`${info.file.name} 文件上传失败.`);
        }
    }

    handleCancel() {
        const path = `${isIncludeAdmin()}/product/list`;
        this.router.navigateByUrl(path);
    }
}
