import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AlarmRoutingModule } from './alarm-routing.module';
import { AlarmsComponent } from './alarms/alarms.component';
import {BaseModule} from "../base/base.module";
import {NzSpaceModule} from "ng-zorro-antd/space";
import {NzTableModule} from "ng-zorro-antd/table";
import {NzIconModule} from "ng-zorro-antd/icon";
import {NzButtonModule} from "ng-zorro-antd/button";
import {NzDividerModule} from "ng-zorro-antd/divider";
import { NzPopconfirmModule } from 'ng-zorro-antd/popconfirm';

@NgModule({
  declarations: [
    AlarmsComponent
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
    NzDividerModule
  ]
})
export class AlarmModule { }
