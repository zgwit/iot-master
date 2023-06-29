import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators, FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { RequestService } from 'src/app/request.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { isIncludeAdmin } from 'src/public';
@Component({
    selector: 'app-subscription-edit',
    templateUrl: './subscription-edit.component.html',
    styleUrls: ['./subscription-edit.component.scss'],
})
export class SubscriptionEditComponent implements OnInit {
    group!: FormGroup;
    again = 's';
    delay = 's';
    id: any = 0;
    listOfOption: any[] = [];
    validatorList: any[] = [];
    userList: any[] = [];
    deviceList: any[] = [];
    productList: any[] = [];

    checkOptionsThree: any = [
        { label: 'sms', value: 'sms', checked: false },
        { label: 'voice', value: 'voice', checked: false },
        { label: 'email', value: 'email', checked: false },
    ];
    constructor(
        private fb: FormBuilder,
        private router: Router,
        private route: ActivatedRoute,
        private rs: RequestService,
        private msg: NzMessageService
    ) {}

    ngOnInit(): void {
        if (this.route.snapshot.paramMap.has('id')) {
            this.id = this.route.snapshot.paramMap.get('id');
            this.rs.get(`validator/${this.id}`).subscribe((res) => {
                //let data = res.data;
                this.build(res.data);
            });
        }

        this.build();

        this.rs
            .post('product/search', {})
            .subscribe((res) => {
                const data: any[] = [];

                res.data.filter((item: { id: string; name: string }) =>
                    data.push({
                        label: item.id + ' / ' + item.name,
                        value: item.id,
                    })
                );
                this.productList = data;
            })
            .add(() => {});

        this.rs
            .post('device/search', {})
            .subscribe((res) => {
                const data: any[] = [];

                res.data.filter((item: { id: string; name: string }) =>
                    data.push({
                        label: item.id  ,
                        value: item.id,
                    })
                );
                this.deviceList = data;
            })
            .add(() => {});

        this.rs
            .post('user/search', {})
            .subscribe((res) => {
                const data: any[] = [];

                res.data.filter((item: { id: string; name: string }) =>
                    data.push({ label: item.id, value: item.id })
                );
                this.userList = data;
            })
            .add(() => {});
    }

    build(obj?: any) {
        obj = obj || {};
        this.group = this.fb.group({
            // id: [obj.id || '', []],
            product_id: [obj.product_id || '', []],
            device_id: [obj.device_id || '', []],
            user_id: [obj.user_id || '', []],
            level: [obj.level || 0, []],
            channels: [this.checkOptionsThree, []],
        });
    }

    submit() {
        if (this.group.valid) {
            let value = this.group.value;
            let channel: any[] = [];
            if(!value.channels)  this.group.patchValue({channels:[]})
            this.checkOptionsThree.filter((item: any) => {
                if (item.checked) channel.push(item.value);
            });
            this.group.patchValue({channels:channel})
            let url = this.id
                ?    `subscription/${this.id}`
                :    `subscription/create`;
            this.rs.post(url, this.group.value).subscribe((res) => {
                this.handleCancel();
                this.msg.success('保存成功');
            });

            return;
        } else {
            Object.values(this.group.controls).forEach((control) => {
                if (control.invalid) {
                    control.markAsDirty();
                    control.updateValueAndValidity({ onlySelf: true });
                }
            });
        }
    }
    handleCancel() {
        const path = `/subscription`;
        this.router.navigateByUrl(path);
    }
}
