
import { Component, Input, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from "@angular/forms";
import { ActivatedRoute, Router } from "@angular/router";
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
  @ViewChild('propertyChild') propertyChild: any;
  @ViewChild('parametersChild') parametersChild: any;
  @ViewChild('constraintsChild') constraintsChild: any;

  @Input() id!: any;
  constructor(
    private fb: FormBuilder,
    private router: Router,
    private route: ActivatedRoute,
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

    if (this.group.valid) {
      let url = this.id ? `product/${this.id}` : `product/create`
      const sendData = JSON.parse(JSON.stringify(this.group.value));
      // 属性组件
      const { propertys, parameters, constraints } = this.getValueData();
      sendData.properties = propertys;
      sendData.parameters = parameters;
      sendData.constraints = constraints;
      this.rs.post(url, sendData).subscribe(res => {
        let path = "/product/list"
        if (location.pathname.startsWith("/admin"))
          path = "/admin" + path
        this.router.navigateByUrl(path)
        this.msg.success("保存成功")
      })
    }

  }
  getValueData() {
    const propertyGroup = this.propertyChild.group;
    const propertys = propertyGroup.get('keyName').controls.map((item: { value: any; }) => item.value);
    const parametersGroup = this.parametersChild.group;
    const parameters = parametersGroup.get('keyName').controls.map((item: { value: any; }) => item.value);
    const constraintsGroup = this.constraintsChild.group;
    const constraints = constraintsGroup.get('keyName').controls.map((item: { value: any; }) => item.value);

    return { propertys, parameters, constraints };
  }
  propertyAdd($event: any) {
    $event.stopPropagation();
    if (this.propertyChild) {
      this.propertyChild.group.get('keyName').controls.unshift(
        this.fb.group({
          name: ['', []],
          label: ['', []],
          type: ['int', []],
          unit: ['', []],
          mode: ['rw', []],
        })
      )
    }

  }

  parameterAdd($event: any) {
    $event.stopPropagation()
    if (this.parametersChild) {
      this.parametersChild.group.get('keyName').controls.unshift(
        this.fb.group({
          name: ['', []],
          label: ['', []],
          min: [0, []],
          max: [0, []],
          default: [0, []],
        })
      )
    }
  }

  constraintAdd($event: any) {
    $event.stopPropagation();
    if (this.constraintsChild) {
      this.constraintsChild.group.get('keyName').controls.unshift(
        this.fb.group({
          level: [1, []],
          title: ['', []],
          template: ['', []],
          expression: ['', []],
          delay: [0, []],
          again: [0, []],
          total: [0, []],
        })
      )
    }
  }

  handleCancel() {
    const path = `${isIncludeAdmin()}/product/list`;
    this.router.navigateByUrl(path);
  }
}
