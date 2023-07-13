import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { DragDropModule } from '@angular/cdk/drag-drop'

import { NzInputModule } from "ng-zorro-antd/input";
import { NzButtonModule } from "ng-zorro-antd/button";
import { NzPaginationModule } from 'ng-zorro-antd/pagination';
import { NzFormModule } from "ng-zorro-antd/form";
import { NzSelectModule } from "ng-zorro-antd/select";
import { NzTableModule } from "ng-zorro-antd/table";
import { NzIconModule } from "ng-zorro-antd/icon";
import { NzSpaceModule } from 'ng-zorro-antd/space';
import { NzInputNumberModule } from 'ng-zorro-antd/input-number';
import { NzDividerModule } from "ng-zorro-antd/divider";
import { PageNotFoundComponent } from "./page-not-found/page-not-found.component";
import { ToolbarComponent } from './toolbar/toolbar.component';
import { DatePipe } from "./date.pipe";
import { EditTableComponent } from './edit-table/edit-table.component';
import { NzUploadModule } from 'ng-zorro-antd/upload';
import { TableOperComponent } from './table-oper/table-oper.component';
import { ModalComponent } from './modal/modal.component';
import { NzTabsModule } from 'ng-zorro-antd/tabs';
import { FullscreamDirective } from './fullscream.directive';
import { CardComponent } from './card/card.component';
import { NzTagModule } from 'ng-zorro-antd/tag';
import { DetailComponent } from './detail/detail.component';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzPopconfirmModule } from 'ng-zorro-antd/popconfirm';
import { NzSwitchModule } from "ng-zorro-antd/switch";
import { SearchFormComponent } from './search-form/search-form.component';
import { BatchBtnComponent } from './batch-btn/batch-btn.component';
@NgModule({
  declarations: [
    DatePipe,
    PageNotFoundComponent,
    ToolbarComponent,
    EditTableComponent,
    TableOperComponent,
    ModalComponent,
    FullscreamDirective,
    CardComponent,
    DetailComponent,
    SearchFormComponent,
    BatchBtnComponent
  ],
  exports: [
    DatePipe,
    PageNotFoundComponent,
    ToolbarComponent,
    EditTableComponent,
    TableOperComponent,
    ModalComponent,
    CardComponent,
    DetailComponent,
    SearchFormComponent,
    BatchBtnComponent
  ],
  imports: [
    CommonModule,
    NzInputModule,
    NzModalModule,
    NzPopconfirmModule,
    NzButtonModule,
    NzTabsModule,
    NzPaginationModule,
    NzTagModule,
    FormsModule,
    ReactiveFormsModule,
    NzFormModule,
    DragDropModule,
    NzSelectModule,
    NzTableModule,
    NzIconModule,
    NzSpaceModule,
    NzInputNumberModule,
    NzUploadModule,
    NzDividerModule,
    NzSwitchModule
  ],
  providers: [

  ]
})
export class BaseModule {
}
