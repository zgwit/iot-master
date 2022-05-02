import {Component, OnInit} from '@angular/core';
import {FormArray, FormBuilder, FormGroup, Validators} from "@angular/forms";
import {ActivatedRoute, Router} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzMessageService} from "ng-zorro-antd/message";

@Component({
  selector: 'app-device-edit',
  templateUrl: './device-edit.component.html',
  styleUrls: ['./device-edit.component.scss']
})
export class DeviceEditComponent implements OnInit {
  id: any;
  submitting = false;

  basicForm: FormGroup = new FormGroup({});

  data: any = {
    "name": "新建设备",
    "element_id": "",
    "link_id": 0,
    "tags": [],
    "icon": "",
    "station": 1,
    "disabled": false,
    "points": [],
    "context": {},
    "commands": [],
    "pollers": [],
    "calculators": [],
    "alarms": [],
  }

  constructor(private fb: FormBuilder, private route: ActivatedRoute, private router: Router, private rs: RequestService, private message: NzMessageService) {
    this.id = route.snapshot.paramMap.get('id');
    if (this.id) this.load();
    Object.assign(this.data, this.route.snapshot.queryParams);
    this.buildForm();
  }

  buildForm(): void {
    this.basicForm = this.fb.group({
      name: [this.data.name, [Validators.required]],
      tags: [this.data.tags, []],
      icon: [this.data.icon, []],
      element_id: [this.data.element_id, [Validators.required]],
      link_id: [this.data.link_id, [Validators.required]],
      station: [this.data.station, [Validators.required]],

      disabled: [this.data.disabled, [Validators.required]],


      points: [this.data.points || []],
      context: [this.data.context || {}],
      commands: [this.data.commands || []],
      pollers: [this.data.pollers || []],
      calculators: [this.data.calculators || []],
      alarms: [this.data.alarms || []],
    });
  }

  ngOnInit(): void {
  }


  load(): void {
    this.rs.get('device/' + this.id).subscribe(res => {
      this.data = res.data;
      this.buildForm();
    })
  }

  submit(): void {
    this.submitting = true
    const uri = this.id ? 'device/' + this.id : 'device/create';
    this.rs.post(uri, this.basicForm.value).subscribe(res => {
      this.message.success("提交成功");
      this.router.navigate(['/admin/device/detail/' + res.data.id]);
    }).add(() => {
      this.submitting = false;
    })
  }

  change() {
    //console.log('change', e)
    this.data = this.basicForm.value;
  }
}
