import {Component, OnInit} from '@angular/core';
import {FormArray, FormBuilder, FormGroup, Validators} from "@angular/forms";
import {ActivatedRoute} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzMessageService} from "ng-zorro-antd/message";

@Component({
  selector: 'app-template-edit',
  templateUrl: './template-edit.component.html',
  styleUrls: ['./template-edit.component.scss']
})
export class TemplateEditComponent implements OnInit {
  id: any;
  submitting = false;

  basicForm: FormGroup = new FormGroup({});

  data: any = {
    "name": "新建模板",
    "commands": [],
    "context": {},
    "elements": [],
    "validators": [],
    "strategies": [],
    "jobs": [],
    "aggregators": [],
  }

  constructor(private fb: FormBuilder, private route: ActivatedRoute, private rs: RequestService, private message: NzMessageService) {
    this.id = route.snapshot.paramMap.get('id');
    if (this.id) this.load();
    this.buildForm();
  }

  buildForm(): void {
    this.basicForm = this.fb.group({
      name: [this.data.name, [Validators.required]],

      context: [this.data.context || {}],
      commands: [this.data.commands || []],
      elements: [this.data.elements || []],
      validators: [this.data.validators || []],
      jobs: [this.data.jobs || []],
      strategies: [this.data.strategies || []],
      aggregators: [this.data.scripts || []],
    });
  }

  ngOnInit(): void {
  }


  load(): void {
    this.rs.get('template/' + this.id + '/detail').subscribe(res => {
      this.data = res.data;
      this.buildForm();
    })
  }

  submit(): void {
    this.submitting = true
    const uri = this.id ? 'template/' + this.id + '/setting' : 'template/create';
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
