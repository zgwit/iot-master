import {Component, OnInit} from '@angular/core';
import {FormArray, FormBuilder, FormGroup, Validators} from "@angular/forms";
import {ActivatedRoute, Router} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzMessageService} from "ng-zorro-antd/message";

@Component({
  selector: 'app-server-edit',
  templateUrl: './server-edit.component.html',
  styleUrls: ['./server-edit.component.scss']
})
export class ServerEditComponent implements OnInit {
  id: any;
  submitting = false;

  basicForm: FormGroup = new FormGroup({});

  data: any = {
    "name": "新建服务器",
    "type": "tcp",
    "addr": "",
    "disabled": false,
    "register": {
      "regex": '^\\w+$'
    },
    "heartbeat": {
      "enable": false,
      "timeout": 30,
      "text": "",
      "hex": "",
      "regex": '^\\w+$'
    },
    "protocol": {
      "name": "ModbusTCP",
      "options": {}
    },
    "devices": []
  }

  constructor(private fb: FormBuilder, private route: ActivatedRoute, private router: Router, private rs: RequestService, private message: NzMessageService) {
    this.id = route.snapshot.paramMap.get('id');
    if (this.id) this.load();
    this.buildForm();
  }

  buildForm(): void {
    this.basicForm = this.fb.group({
      name: [this.data.name, [Validators.required]],
      type: [this.data.type, [Validators.required]],
      addr: [this.data.addr, [Validators.required]],
      disabled: [this.data.disabled, []],
      register: [this.data.register, []],
      heartbeat: [this.data.heartbeat, []],
      protocol: [this.data.protocol, []],
      devices: [this.data.devices, []],
    });
  }

  ngOnInit(): void {
  }


  load(): void {
    this.rs.get('server/' + this.id).subscribe(res => {
      this.data = res.data;
      this.buildForm();
    })
  }

  submit(): void {
    this.submitting = true
    const uri = this.id ? 'server/' + this.id : 'server/create';
    this.rs.post(uri, this.basicForm.value).subscribe(res => {
      this.message.success("提交成功");
      this.router.navigate(['/admin/server/detail/' + res.data.id]);
    }).add(() => {
      this.submitting = false;
    })
  }

  change() {
    //console.log('change', e)
    this.data = this.basicForm.value;
  }

}
