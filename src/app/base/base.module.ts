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

import { PageNotFoundComponent } from "./page-not-found/page-not-found.component";
import { ToolbarComponent } from './toolbar/toolbar.component';
import { SearchBoxComponent } from './search-box/search-box.component';
import { DatePipe } from "./date.pipe";
import { EditTableComponent } from './edit-table/edit-table.component';
import { ImportComponent } from './import/import.component';
import { NzUploadModule } from 'ng-zorro-antd/upload';
@NgModule({
  declarations: [
    DatePipe,
    PageNotFoundComponent,
    ToolbarComponent,
    SearchBoxComponent,
    EditTableComponent,
    ImportComponent,
  ],
  exports: [
    DatePipe,
    PageNotFoundComponent,
    ToolbarComponent,
    SearchBoxComponent,
    EditTableComponent,
    ImportComponent,
  ],
  imports: [
    CommonModule,
    NzInputModule,
    NzButtonModule,
    NzPaginationModule,
    FormsModule,
    ReactiveFormsModule,
    NzFormModule,
    DragDropModule,
    NzSelectModule,
    NzTableModule,
    NzIconModule,
    NzSpaceModule,
    NzInputNumberModule,
    NzUploadModule
  ],
  providers: [

  ]
})
export class BaseModule {
}
