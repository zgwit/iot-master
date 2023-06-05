import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AttachmentRoutingModule } from './attachment-routing.module';
import { AttachmentComponent } from './attachment.component';
import { UploadComponent } from './upload/upload.component';
import { RenameComponent } from './rename/rename.component';
import { MoveComponent } from './move/move.component';


@NgModule({
  declarations: [
    AttachmentComponent,
    UploadComponent,
    RenameComponent,
    MoveComponent
  ],
  imports: [
    CommonModule,
    AttachmentRoutingModule
  ]
})
export class AttachmentModule { }
