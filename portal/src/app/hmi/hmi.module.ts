import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {EditorComponent} from './editor/editor.component';
import {ViewerComponent} from './viewer/viewer.component';
import {AttachmentComponent} from "./attachment/attachment.component";
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
import {NzTabsModule} from "ng-zorro-antd/tabs";
import {NzDropDownModule} from "ng-zorro-antd/dropdown";
import {NzUploadModule} from "ng-zorro-antd/upload";
import {NzButtonModule} from "ng-zorro-antd/button";
import {NzPopconfirmModule} from "ng-zorro-antd/popconfirm";
import {NzModalModule} from "ng-zorro-antd/modal";


@NgModule({
  declarations: [
    EditorComponent,
    ViewerComponent,
    AttachmentComponent,
  ],
  exports: [
    EditorComponent,
    ViewerComponent,
  ],
  imports: [
    CommonModule,
    FormsModule,
    NzLayoutModule,
    NzCollapseModule,
    NzIconModule,
    NzDividerModule,
    NzTableModule,
    NzSelectModule,
    NzInputModule,
    NzSwitchModule,
    IconsProviderModule,
    ColorPickerModule,
    NzTabsModule,
    NzDropDownModule,
    NzUploadModule,
    NzButtonModule,
    NzPopconfirmModule,
    NzModalModule,
  ],
  providers: []
})
export class HmiModule {
}
