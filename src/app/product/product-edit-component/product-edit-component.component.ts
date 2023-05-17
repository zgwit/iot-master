
import { Component, Input, OnInit } from '@angular/core';
import { FormBuilder, Validators } from "@angular/forms";
import { Router } from "@angular/router";
import { RequestService } from "../../request.service";
import { NzMessageService } from "ng-zorro-antd/message";
import { isIncludeAdmin } from "../../../public";

@Component({
  selector: 'app-product-edit-component',
  templateUrl: './product-edit-component.component.html',
  styleUrls: ['./product-edit-component.component.scss']
})
export class ProductEditComponentComponent implements OnInit {
  group!: any;
  allData: { properties: Array<object> } = { properties: [] };
  listData = [{
    title: '名称(ID)',
    keyName: 'name'
  }, {
    title: '显示',
    keyName: 'label'
  }, {
    title: '类型',
    keyName: 'type',
    type: 'select',
    defaultValue: 'int',
    listOfOption: [{
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
    title: '单位',
    keyName: 'unit'
  }, {
    title: '模式',
    keyName: 'mode',
    type: 'select',
    defaultValue: 'rw',
    listOfOption: [{
      label: '只读',
      value: 'r'
    }, {
      label: '读写',
      value: 'rw'
    }]
  }]
  parameterslistData = [
    {
      title: '名称(ID)',
      keyName: 'name'
    }, {
      title: '显示',
      keyName: 'label'
    }, {
      title: '最大值',
      keyName: 'max',
      type: 'number',
      defaultValue: 0
    }, {
      title: '最小值',
      keyName: 'min',
      type: 'number',
      defaultValue: 0
    }, {
      title: '默认值',
      keyName: 'default',
      type: 'number',
      defaultValue: 0
    }
  ]
  constraintslistData = [
    {
      title: '等级',
      keyName: 'level'
    },
    {
      title: '标题',
      keyName: 'title'
    },
    {
      title: '模板',
      keyName: 'template'
    },
    {
      title: '表达式',
      keyName: 'expression'
    },
    {
      title: '延迟',
      keyName: 'delay',
      type: 'number'
    },
    {
      title: '再次提醒',
      keyName: 'again',
      type: 'number'
    },
    {
      title: '总提醒次数',
      keyName: 'total',
      type: 'number'
    }
  ]

  @Input() id!: any;
  constructor(
    private fb: FormBuilder,
    private router: Router,
    private rs: RequestService,
    private msg: NzMessageService
  ) { }

  ngOnInit(): void {
    if (this.id) {
      this.rs.get(`product/${this.id}`).subscribe(res => {
        this.allData = res.data || {};
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
      desc: [obj.desc || '', []],
      version: [obj.version || '', []],
      properties: [obj.properties || [], []],
      parameters: [obj.parameters || [], []],
      constraints: [obj.constraints || [], []],
    })
  }

  submit() {
   // console.log(this.group.value)
    return new Promise((resolve) => {
      if (this.group.valid) {
        let url = this.id ? `product/${this.id}` : `product/create`
        this.rs.post(url, this.group.value).subscribe(res => {
          this.msg.success("保存成功");
          resolve(true);
        })
      }
    })

  }

  handleCancel() {
    const path = `${isIncludeAdmin()}/product/list`;
    this.router.navigateByUrl(path);
  }
}
