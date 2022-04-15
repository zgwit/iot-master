import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {ActivatedRoute} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzMessageService} from "ng-zorro-antd/message";

@Component({
  selector: 'app-element-edit',
  templateUrl: './element-edit.component.html',
  styleUrls: ['./element-edit.component.scss']
})
export class ElementEditComponent implements OnInit {
  id: any;
  submitting = false;

  basicForm: FormGroup = new FormGroup({});
  data: any = {
    "name": "新建元件",
    "tags": [],
    "icon": "",
    "manufacturer": "",
    "version": "",
    "points": [],
    "context": {},
    "commands": [],
    "pollers": [],
    "jobs": [],
  }

  constructor(private fb: FormBuilder, private route: ActivatedRoute, private rs: RequestService, private message: NzMessageService) {
    this.id = route.snapshot.paramMap.get('id');
    if (this.id) this.load();
    this.buildForm();
  }

  buildForm(): void {
    this.basicForm = this.fb.group({
      name: [this.data.name, [Validators.required]],
      tags: [this.data.tags, []],
      icon: [this.data.icon, []],
      manufacturer: [this.data.manufacturer, []],
      version: [this.data.version, []],

      points: [this.data.points || []],
      context: [this.data.context || {}],
      commands: [this.data.commands || []],
      pollers: [this.data.collectors || []],
      jobs: [this.data.jobs || []],
    });
  }

  ngOnInit(): void {
  }


  load(): void {
    this.rs.get('element/' + this.id + '/detail').subscribe(res => {
      this.data = res.data;
      this.buildForm();
    })
  }

  submit(): void {
    this.submitting = true
    const uri = this.id ? 'element/' + this.id + '/setting' : 'element/create';
    this.rs.post(uri, this.basicForm.value).subscribe(res => {
      this.message.success("提交成功");
    }).add(() => {
      this.submitting = false;
    })
  }

  change() {
    //console.log('change', e)
    this.data = this.basicForm.value;
  }
}
