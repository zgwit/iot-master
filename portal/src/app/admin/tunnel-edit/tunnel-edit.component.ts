import {Component, OnInit} from '@angular/core';
import {FormArray, FormBuilder, FormGroup, Validators} from "@angular/forms";
import {ActivatedRoute, Router} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzMessageService} from "ng-zorro-antd/message";

@Component({
  selector: 'app-tunnel-edit',
  templateUrl: './tunnel-edit.component.html',
  styleUrls: ['./tunnel-edit.component.scss']
})
export class TunnelEditComponent implements OnInit {
  id: any;
  submitting = false;

  basicForm: FormGroup = new FormGroup({});

  data: any = {
    "name": "新建通道",
    "type": "tcp-server",
    "addr": "",
    "disabled": false,
    "register": {
      "enable": true,
      "regex": '^\\w+$'
    },
    "heartbeat": {
      "enable": false,
      "timeout": 30,
      "text": "",
      "hex": "",
      "regex": '^\\w+$'
    },
    retry: {
      "enable": false,
      "timeout": 30,
      "maximum": 0,
    },
    "serial": {
      baud_rate: 9600,
      data_bits: 8,
      stop_bits: 1,
      parity_mode: 0,
      rs485: false,
    },
    "protocol": {
      "name": "",
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
      retry: [this.data.retry, []],
      serial: [this.data.serial, []],
      protocol: [this.data.protocol, []],
      devices: [this.data.devices, []],
    });
  }

  ngOnInit(): void {
  }


  load(): void {
    this.rs.get('tunnel/' + this.id).subscribe(res => {
      this.data = res.data;
      this.buildForm();
    })
  }

  submit(): void {
    this.submitting = true
    const uri = this.id ? 'tunnel/' + this.id : 'tunnel/create';
    this.rs.post(uri, this.basicForm.value).subscribe(res => {
      this.message.success("提交成功");
      this.router.navigate(['/admin/tunnel/detail/' + res.data.id]);
    }).add(() => {
      this.submitting = false;
    })
  }

  change() {
    //console.log('change', e)
    this.data = this.basicForm.value;
  }

}
