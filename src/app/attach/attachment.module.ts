import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AttachmentRoutingModule } from './attachment-routing.module';
import { AttachmentComponent } from './attachment.component';
import { UploadComponent } from './upload/upload.component';
import { RenameComponent } from './rename/rename.component';
import { MoveComponent } from './move/move.component';
import { FileUploadModule } from 'ng2-file-upload';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzFormModule } from "ng-zorro-antd/form";
import { ReactiveFormsModule, FormsModule } from "@angular/forms";
import { NzInputModule } from 'ng-zorro-antd/input';
import { BaseModule } from '../base/base.module';
import { NzSpaceModule } from 'ng-zorro-antd/space';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzImageModule } from 'ng-zorro-antd/image';
import { NzPopconfirmModule } from 'ng-zorro-antd/popconfirm';
import { NzDividerModule } from 'ng-zorro-antd/divider';

@NgModule({
  declarations: [
    AttachmentComponent,
    UploadComponent,
    RenameComponent,
    MoveComponent
  ],
  imports: [
    CommonModule,
    AttachmentRoutingModule,
    FileUploadModule,
    NzButtonModule,
    NzFormModule,
    ReactiveFormsModule,
    NzInputModule,
    FormsModule,
    BaseModule,
    NzSpaceModule,
    NzIconModule,
    NzTableModule,
    NzImageModule,
    NzPopconfirmModule,
    NzDividerModule
  ]
})
export class AttachmentModule { }
