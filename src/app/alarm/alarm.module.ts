import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AlarmRoutingModule } from './alarm-routing.module';
import { AlarmsComponent } from './alarms/alarms.component';
import { BaseModule } from "../base/base.module";
import { NzSpaceModule } from "ng-zorro-antd/space";
import { NzTableModule } from "ng-zorro-antd/table";
import { NzIconModule } from "ng-zorro-antd/icon";
import { NzButtonModule } from "ng-zorro-antd/button";
import { NzDividerModule } from "ng-zorro-antd/divider";
import { NzPopconfirmModule } from 'ng-zorro-antd/popconfirm';
import { NotificationComponent } from "./notification/notification.component";
import { SubscriptionComponent } from "./subscription/subscription.component";
import { SubscriptionEditComponent } from "./subscription-edit/subscription-edit.component";
import { ValidatorsComponent } from "./validators/validators.component";
import { ValidatorEditComponent } from "./validator-edit/validator-edit.component";
import { NzTagModule } from "ng-zorro-antd/tag";
import { NzCardModule } from "ng-zorro-antd/card";
import { ReactiveFormsModule } from "@angular/forms";
import { NzInputNumberModule } from "ng-zorro-antd/input-number";
import { NzFormModule } from "ng-zorro-antd/form";
import { NzSelectModule } from "ng-zorro-antd/select";
import { NzInputModule } from "ng-zorro-antd/input";
import { NzCheckboxModule } from "ng-zorro-antd/checkbox";
import { NzSwitchModule } from "ng-zorro-antd/switch";
import { ChinaDatePipe } from '../china-date.pipe';
@NgModule({
    declarations: [
        AlarmsComponent,
        NotificationComponent,
        SubscriptionComponent,
        SubscriptionEditComponent,
        ValidatorsComponent,
        ValidatorEditComponent,
        ChinaDatePipe
    ],
    imports: [
        CommonModule,
        AlarmRoutingModule,
        BaseModule,
        NzSpaceModule,
        NzTableModule,
        NzPopconfirmModule,
        NzIconModule,
        NzButtonModule,
        NzDividerModule,
        NzTagModule,
        NzCardModule,
        ReactiveFormsModule,
        NzInputNumberModule,
        NzFormModule,
        NzSelectModule,
        NzInputModule,
        NzCheckboxModule,
        NzSwitchModule
    ]
})
export class AlarmModule {
}
