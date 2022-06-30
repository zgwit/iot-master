import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {ActivatedRoute, Router} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzMessageService} from "ng-zorro-antd/message";
import cryptoRandomString from "crypto-random-string";

@Component({
  selector: 'app-product-edit',
  templateUrl: './product-edit.component.html',
  styleUrls: ['./product-edit.component.scss']
})
export class ProductEditComponent implements OnInit {
  id: any;
  submitting = false;

  codes: Array<any> = [];

  basicForm: FormGroup = new FormGroup({});
  data: any = {
    "name": "新建产品",
    "hmi":"",
    "tags": [],
    "manufacturer": "",
    "version": "",

    "protocol": {
      "name": "ModbusRTU",
      "options": {}
    },
    "points": [],
    "commands": [],
    "pollers": [],
    "calculators": [],
    "alarms": [],
  }

  constructor(private fb: FormBuilder, private route: ActivatedRoute, private router: Router, private rs: RequestService, private message: NzMessageService) {
    this.id = route.snapshot.paramMap.get('id');
    if (this.id) this.load();
    this.buildForm();
  }

  buildForm(): void {
    this.basicForm = this.fb.group({
      name: [this.data.name, [Validators.required]],
      hmi: [this.data.hmi, []],
      tags: [this.data.tags, []],
      manufacturer: [this.data.manufacturer, []],
      version: [this.data.version, []],

      protocol: [this.data.protocol, []],
      points: [this.data.points || []],
      commands: [this.data.commands || []],
      pollers: [this.data.pollers || []],
      calculators: [this.data.calculators || []],
      alarms: [this.data.alarms || []],
    });
  }

  ngOnInit(): void {
  }


  load(): void {
    this.rs.get('product/' + this.id).subscribe(res => {
      this.data = res.data;
      this.buildForm();
    })
  }

  submit(): void {
    this.submitting = true
    const data = this.basicForm.value;
    //生成随机ID
    if (!data.id) data.id = cryptoRandomString({length: 20, type: 'alphanumeric'})
    const uri = this.id ? 'product/' + this.id : 'product/create';
    this.rs.post(uri, data).subscribe(res => {
      this.message.success("提交成功");
      this.router.navigate(['/admin/product/detail/' + res.data.id]);
    }).add(() => {
      this.submitting = false;
    })
  }

  change() {
    //console.log('change', e)
    this.data = this.basicForm.value;
  }

  onCodes($event: any) {
    console.log("onCodes", $event)
    this.codes = $event;
  }
}
