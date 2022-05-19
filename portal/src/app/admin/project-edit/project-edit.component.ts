import {Component, OnInit} from '@angular/core';
import {FormArray, FormBuilder, FormGroup, Validators} from "@angular/forms";
import {ActivatedRoute, Router} from "@angular/router";
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
    "hmi":"",
    "context": {},
    "devices": [],
    "alarms": [],
    "strategies": [],
    "jobs": [],
    "aggregators": [],
  }

  constructor(private fb: FormBuilder, private route: ActivatedRoute, private router: Router, private rs: RequestService, private message: NzMessageService) {
    this.id = route.snapshot.paramMap.get('id');
    if (this.id) this.load();
    Object.assign(this.data, this.route.snapshot.queryParams);
    if (this.data.id) {
      this.data.devices.push({id: this.data.id, station: 1});
      delete this.data.id;
    }
    this.buildForm();
  }

  buildForm(): void {
    this.basicForm = this.fb.group({
      name: [this.data.name, [Validators.required]],
      template_id: [this.data.template_id, []],
      hmi: [this.data.hmi, []],
      disabled: [this.data.disabled, []],

      context: [this.data.context || {}],
      devices: [this.data.devices || []],
      alarms: [this.data.alarms || []],
      jobs: [this.data.jobs || []],
      strategies: [this.data.strategies || []],
      aggregators: [this.data.scripts || []],
    });
  }

  ngOnInit(): void {
  }


  load(): void {
    this.rs.get('project/' + this.id).subscribe(res => {
      this.data = res.data;
      this.buildForm();
    })
  }

  submit(): void {
    this.submitting = true
    const uri = this.id ? 'project/' + this.id : 'project/create';
    this.rs.post(uri, this.basicForm.value).subscribe(res => {
      this.message.success("提交成功");
      this.router.navigate(['/admin/project/detail/' + res.data.id]);
    }).add(() => {
      this.submitting = false;
    })
  }

  change() {
    //console.log('change', e)
    this.data = this.basicForm.value;
  }

}
