import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PluginRoutingModule } from './plugin-routing.module';
import {NzLayoutModule} from "ng-zorro-antd/layout";
import {NzMenuModule} from "ng-zorro-antd/menu";
import {NzIconModule} from "ng-zorro-antd/icon";
import { PluginsComponent } from './plugins/plugins.component';
import { PluginEditComponent } from './plugin-edit/plugin-edit.component';
import {NzCardModule} from "ng-zorro-antd/card";
import {NzFormModule} from "ng-zorro-antd/form";
import {ReactiveFormsModule} from "@angular/forms";
import {NzInputModule} from "ng-zorro-antd/input";
import {NzInputNumberModule} from "ng-zorro-antd/input-number";
import {NzButtonModule} from "ng-zorro-antd/button";
import {BaseModule} from "../base/base.module";
import {NzSpaceModule} from "ng-zorro-antd/space";
import {NzTableModule} from "ng-zorro-antd/table";
import {NzDividerModule} from "ng-zorro-antd/divider";
import { PluginDetailComponent } from './plugin-detail/plugin-detail.component';
import { NzPopconfirmModule } from 'ng-zorro-antd/popconfirm';
import { NzTagModule } from 'ng-zorro-antd/tag';
import {NzSwitchModule} from "ng-zorro-antd/switch";
import {NzRadioModule} from "ng-zorro-antd/radio";
import {NzCheckboxModule} from "ng-zorro-antd/checkbox";
import { NzModalModule } from 'ng-zorro-antd/modal';
@NgModule({
  declarations: [
    PluginsComponent,
    PluginEditComponent,
    PluginDetailComponent,
  ],
    imports: [
        CommonModule,
        PluginRoutingModule,
        NzLayoutModule,
        NzMenuModule,
        NzIconModule,
        NzCardModule,
        NzModalModule,
        NzFormModule,
        NzPopconfirmModule,
        NzTagModule,
        ReactiveFormsModule,
        NzInputModule,
        NzInputNumberModule,
        NzButtonModule,
        BaseModule,
        NzSpaceModule,
        NzTableModule,
        NzDividerModule,
        NzSwitchModule,
        NzRadioModule,
        NzCheckboxModule
    ]
})
export class PluginModule { }
