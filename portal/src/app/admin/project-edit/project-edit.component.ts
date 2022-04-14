import {Component, OnInit} from '@angular/core';
import {FormArray, FormBuilder, FormGroup, Validators} from "@angular/forms";
import {ActivatedRoute} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzMessageService} from "ng-zorro-antd/message";

@Component({
  selector: 'app-project-edit',
  templateUrl: './project-edit.component.html',
  styleUrls: ['./project-edit.component.scss']
})
export class ProjectEditComponent implements OnInit {
  id: any;
  submitting = false;

  basicForm: FormGroup = new FormGroup({});

  data: any = {
    "name": "新建项目",
    "disabled": false,
    "template_id":"",
    "commands": [],
    "context": {},
    "devices": [],
    "strategies": [],
    "jobs": [],
    "aggregators": [],
  }

  constructor(private fb: FormBuilder, private route: ActivatedRoute, private rs: RequestService, private message: NzMessageService) {
    this.id = route.snapshot.paramMap.get('id');
    if (this.id) this.load();
    Object.assign(this.data, this.route.snapshot.queryParams);
    if (this.data.device_id) {
      this.data.devices.push({device_id: this.data.device_id, station: 1});
      delete this.data.device_id;
    }
    this.buildForm();
  }

  buildForm(): void {
    this.basicForm = this.fb.group({
      name: [this.data.name, [Validators.required]],
      template_id: [this.data.template_id, []],
      disabled: [this.data.disabled, [Validators.required]],

      commands: [this.data.commands || []],
      context: [this.data.variables || {}],
      devices: [this.data.devices || []],
      jobs: [this.data.jobs || []],
      strategies: [this.data.strategies || []],
      aggregators: [this.data.scripts || []],
    });
  }

  ngOnInit(): void {
  }


  load(): void {
    this.rs.get('project/' + this.id + '/detail').subscribe(res => {
      this.data = res.data;
      this.buildForm();
    })
  }

  submit(): void {
    this.submitting = true
    const uri = this.id ? 'project/' + this.id + '/setting' : 'project/create';
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
