import { filter } from 'rxjs/operators';
import { Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { RequestService } from '../../request.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzModalService } from 'ng-zorro-antd/modal';
import { DevicesComponent } from '../devices/devices.component';
import { isIncludeAdmin } from '../../../public';
@Component({
  selector: 'app-products-edit',
  templateUrl: './device-edit.component.html',
  styleUrls: ['./device-edit.component.scss'],
})
export class DeviceEditComponent implements OnInit {
  group!: FormGroup;
  id: any = 0;
  typeID: any[] = [];
  groupID: any[] = [];
  @ViewChild('childTag') childTag: any;
  constructor(
    private fb: FormBuilder,
    private router: Router,
    private route: ActivatedRoute,
    private rs: RequestService,
    private ms: NzModalService,
    private msg: NzMessageService
  ) {}

  ngOnInit(): void {
    if (this.route.snapshot.paramMap.has('id')) {
      this.id = this.route.snapshot.paramMap.get('id');

      this.rs.get(`device/${this.id}`).subscribe((res) => {
        //let data = res.data;
        if (this.childTag) {
          // 给子组件设值
          const { product_id, group_id } = res.data || {};
          const IdObj = {
            product_id: product_id || '',
            group_id: group_id || '',
          };
          this.childTag.IdObj = JSON.parse(JSON.stringify(IdObj));
        }
        this.build(res.data);
      });
    }

    this.build();
    this.selectId();
  }
  selectId() {
    this.rs.get('device/type/list').subscribe((res) => {
      const data: any[] = [];
      res.data.filter((item: { name: any; desc: any; id: any }) =>
        data.push({
          value: item.id,
          label: item.id + ' / ' + item.name,
        })
      );
      this.typeID = data;
    });

    this.rs
      .post('device/group/search', {})
      .subscribe((res) => {
        const data: any[] = [];

        res.data.filter((item: { id: string; name: string }) =>
          data.push({ label: item.id + ' / ' + item.name, value: item.id })
        );
        this.groupID = data;
      })
      .add(() => {});
  }
  build(obj?: any) {
    obj = obj || {};
    this.group = this.fb.group({
      id: [obj.id || '', []],
      product_id: [obj.product_id || '', []],
      gateway_id: [obj.gateway_id || '', []],
      group_id: [obj.group_id || '', []],
      type_id: [obj.type_id || '', []],
      name: [obj.name || '', [Validators.required]],
      desc: [obj.desc || '', []],
      // disabled: [obj.disabled || false, []],
    });
  }

  submit() {
    if (this.group.valid) {
      const { IdObj } = this.childTag;
      const sendData = Object.assign({}, this.group.value, IdObj);
      let url = this.id ? `device/${this.id}` : `device/create`;
      this.rs.post(url, sendData).subscribe((res) => {
        const path = `${isIncludeAdmin()}/device/list`;
        this.router.navigateByUrl(path);
        this.msg.success('保存成功');
      });
    } else {
      Object.values(this.group.controls).forEach((control) => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  chooseGateway() {
    this.ms
      .create({
        nzTitle: '选择网关',
        nzContent: DevicesComponent,
        nzComponentParams: {
          chooseGateway: true,
          showAddBtn: false,
        },
        nzFooter: null,
      })
      .afterClose.subscribe((res) => {
        if (res) {
          this.group.patchValue({ gateway_id: res });
        }
      });
  }

  handleCancel() {
    const path = `${isIncludeAdmin()}/device/list`;
    this.router.navigateByUrl(path);
  }
}
