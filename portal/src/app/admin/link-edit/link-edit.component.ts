import {Component, OnInit} from '@angular/core';
import {FormArray, FormBuilder, FormGroup, Validators} from "@angular/forms";
import {ActivatedRoute} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzMessageService} from "ng-zorro-antd/message";

@Component({
  selector: 'app-link-edit',
  templateUrl: './link-edit.component.html',
  styleUrls: ['./link-edit.component.scss']
})
export class LinkEditComponent implements OnInit {
  id: any;
  submitting = false;

  basicForm: FormGroup = new FormGroup({});

  data: any = {
    "name": "",
    "timeout": 60,
    "disabled": true,
  }

  constructor(private fb: FormBuilder, private route: ActivatedRoute, private rs: RequestService, private message: NzMessageService) {
    this.id = route.snapshot.paramMap.get('id');
    if (this.id) this.load();
    this.buildForm();
  }

  buildForm(): void {
    this.basicForm = this.fb.group({
      name: [this.data.name, [Validators.required]],
      company_id: [this.data.company_id, []],
      timeout: [this.data.timeout, [Validators.required]],
      disabled: [this.data.disabled, [Validators.required]],
    });
  }

  ngOnInit(): void {
  }


  load(): void {
    this.rs.get('link/' + this.id).subscribe(res => {
      this.data = res.data;
      this.buildForm();
    })
  }

  submit(): void {
    this.submitting = true
    const uri = 'link/' + this.id;
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
