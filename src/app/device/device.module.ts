import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { DeviceRoutingModule } from './device-routing.module';
import { DevicesComponent } from "./devices/devices.component";
import { DeviceEditComponent } from "./device-edit/device-edit.component";
import { NzLayoutModule } from "ng-zorro-antd/layout";
import { NzMenuModule } from "ng-zorro-antd/menu";
import { NzFormModule } from "ng-zorro-antd/form";
import { NzButtonModule } from "ng-zorro-antd/button";
import { NzTableModule } from "ng-zorro-antd/table";
import { NzIconModule } from "ng-zorro-antd/icon";
import { NzDividerModule } from "ng-zorro-antd/divider";
import { ReactiveFormsModule } from "@angular/forms";
import { NzCardModule } from "ng-zorro-antd/card";
import { NzInputNumberModule } from "ng-zorro-antd/input-number";
import { BaseModule } from "../base/base.module";
import { NzInputModule } from "ng-zorro-antd/input";
import { NzSpaceModule } from "ng-zorro-antd/space";
import { DeviceDetailComponent } from './device-detail/device-detail.component';
import { NzStatisticModule } from "ng-zorro-antd/statistic";
import { NzCollapseModule } from "ng-zorro-antd/collapse";
import { NzTabsModule } from "ng-zorro-antd/tabs";
import { NzSwitchModule } from "ng-zorro-antd/switch";
import { NzSelectModule } from "ng-zorro-antd/select";
import { BatchComponent } from './batch/batch.component';
import { DeviceGroupComponent } from './device-group/device-group.component';
import { DeviceGroupEditComponent } from './device-group-edit/device-group-edit.component';
import { NzModalModule } from "ng-zorro-antd/modal";
import { NzPopconfirmModule } from 'ng-zorro-antd/popconfirm';
import { NzSpinModule } from 'ng-zorro-antd/spin';
import { NzDropDownModule } from 'ng-zorro-antd/dropdown';
import { DeviceTypeComponent } from './device-type/device-type.component';
import { DeviceTypeEditComponent } from './device-type-edit/device-type-edit.component';
import { FormsModule } from '@angular/forms';
import { NzUploadModule } from 'ng-zorro-antd/upload';
import { DeviceAreaComponent } from './device-area/device-area.component';
import { DeviceAreaEditComponent } from './device-area-edit/device-area-edit.component';
import { NzTagModule } from 'ng-zorro-antd/tag';
import { ProductModule } from "../product/product.module";
@NgModule({
  declarations: [
    DevicesComponent,
    DeviceEditComponent,
    DeviceDetailComponent,
    BatchComponent,
    DeviceGroupComponent,
    DeviceGroupEditComponent,
    DeviceTypeComponent,
    DeviceTypeEditComponent,
    DeviceAreaComponent,
    DeviceAreaEditComponent
  ],
  imports: [
    CommonModule,
    DeviceRoutingModule,
    NzLayoutModule,
    NzMenuModule,
    NzTagModule,
    NzIconModule,
    NzDropDownModule,
    FormsModule,
    NzUploadModule,
    NzFormModule,
    NzInputModule,
    NzButtonModule,
    NzPopconfirmModule,
    NzTableModule,
    NzDividerModule,
    BaseModule,
    ReactiveFormsModule,
    NzInputNumberModule,
    NzCardModule,
    NzSpaceModule,
    NzStatisticModule,
    NzCollapseModule,
    NzTabsModule,
    NzSwitchModule,
    NzSelectModule,
    NzModalModule,
    NzSpinModule,
    ProductModule
  ]
})
export class DeviceModule {
}
