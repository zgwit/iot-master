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

  codes: any = [];

  data: any = {
    "name": "新建设备",
    "product_id": "",
    "tunnel_id": 0,
    "hmi":"",
    "tags": [],
    "station": 1,
    "disabled": false,
    "points": [],
    "commands": [],
    "pollers": [],
    "calculators": [],
    "alarms": [],
  }

  constructor(private fb: FormBuilder, private route: ActivatedRoute, private router: Router, private rs: RequestService, private message: NzMessageService) {
    this.id = route.snapshot.paramMap.get('id');
    if (this.id) this.load();
    Object.assign(this.data, this.router.getCurrentNavigation()?.extras.state);
    this.buildForm();
  }

  buildForm(): void {
    this.basicForm = this.fb.group({
      name: [this.data.name, [Validators.required]],
      hmi: [this.data.hmi, []],
      tags: [this.data.tags, []],
      product_id: [this.data.product_id, [Validators.required]],
      tunnel_id: [this.data.tunnel_id, [Validators.required]],
      station: [this.data.station, [Validators.required]],

      disabled: [this.data.disabled, [Validators.required]],


      points: [this.data.points || []],
      pollers: [this.data.pollers || []],
      commands: [this.data.commands || []],
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

  onTunnel($event: any) {
    //$event.protocol.name
    this.rs.get("system/protocol/" + $event.protocol.name).subscribe(res=>{
      this.codes = res.data.codes;
    })
  }
}
