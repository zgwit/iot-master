import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {ActivatedRoute, Router} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzMessageService} from "ng-zorro-antd/message";
import cryptoRandomString from "crypto-random-string";
import {NzModalService} from "ng-zorro-antd/modal";

@Component({
  selector: 'app-component-edit',
  templateUrl: './component-edit.component.html',
  styleUrls: ['./component-edit.component.scss']
})
export class ComponentEditComponent implements OnInit {
  id: any;
  submitting = false;

  basicForm: FormGroup = new FormGroup({});

  data: any = {
    "name": "新建组件",
    "group": "扩展",
    "version": "1.0.0",
  }

  constructor(private fb: FormBuilder,
              private route: ActivatedRoute,
              private router: Router,
              private rs: RequestService,
              private message: NzMessageService,
              private ms: NzModalService,
  ) {
    this.id = route.snapshot.paramMap.get('id');
    if (this.id) this.load();
    else this.data.id = cryptoRandomString({length: 20, type: 'alphanumeric'})

    this.buildForm();
  }

  buildForm(): void {
    this.basicForm = this.fb.group({
      id: [this.data.id, []],
      group: [this.data.group, [Validators.required]],
      name: [this.data.name, [Validators.required]],
      version: [this.data.version, [Validators.required]],
    });
  }

  ngOnInit(): void {
  }


  load(): void {
    this.rs.get('component/' + this.id).subscribe(res => {
      this.data = res.data;
      this.buildForm();
    })
  }

  submit(): void {
    this.submitting = true
    const uri = this.id ? 'component/' + this.id : 'component/create';
    this.rs.post(uri, this.basicForm.value).subscribe(res => {
      this.message.success("提交成功");
      this.router.navigate(['/admin/component/detail/' + res.data.id]);

      this.ms.confirm({
        nzContent: "现在编辑组件？",
        nzOnOk: () => {
          this.router.navigate(['/admin/component-edit/' + res.data.id]);
        }
      })

    }).add(() => {
      this.submitting = false;
    })
  }

  change() {
    //console.log('change', e)
    this.data = this.basicForm.value;
  }

}
