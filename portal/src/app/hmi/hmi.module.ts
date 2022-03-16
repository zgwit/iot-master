import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { EditorComponent } from './editor/editor.component';
import { ViewComponent } from './view/view.component';
import {HmiRoutingModule} from "./hmi-routing.module";
import {NzLayoutModule} from "ng-zorro-antd/layout";
import {NzCollapseModule} from "ng-zorro-antd/collapse";
import {NzIconModule} from "ng-zorro-antd/icon";
import {NzDividerModule} from "ng-zorro-antd/divider";
import {IconsProviderModule} from "./icons-provider.module";
import {NzTableModule} from "ng-zorro-antd/table";
import {ColorPickerModule} from "ngx-color-picker";
import {NzSelectModule} from "ng-zorro-antd/select";
import {FormsModule} from "@angular/forms";
import {NzInputModule} from "ng-zorro-antd/input";
import {NzSwitchModule} from "ng-zorro-antd/switch";



@NgModule({
  declarations: [
    EditorComponent,
    ViewComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
    HmiRoutingModule,
    NzLayoutModule,
    NzCollapseModule,
    NzIconModule,
    NzDividerModule,
    NzTableModule,
    IconsProviderModule,
    ColorPickerModule,
    NzSelectModule,
    NzInputModule,
    NzSwitchModule,
  ]
})
export class HmiModule { }
