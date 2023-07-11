import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { ProductRoutingModule } from './product-routing.module';
import { ProductsComponent } from "./products/products.component";
import { ProductEditComponent } from "./product-edit/product-edit.component";
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
import { NzCollapseModule } from "ng-zorro-antd/collapse";
import { NzSelectModule } from "ng-zorro-antd/select";
import { NzTypographyModule } from "ng-zorro-antd/typography";
import { NzPopconfirmModule } from 'ng-zorro-antd/popconfirm';
import { NzUploadModule } from 'ng-zorro-antd/upload';
import { ProductEditComponentComponent } from './product-edit-component/product-edit-component.component';
import { ProductSelectComponent } from './product-select/product-select.component';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzPaginationModule } from 'ng-zorro-antd/pagination';
@NgModule({
    declarations: [
        ProductsComponent,
        ProductEditComponent,
        ProductEditComponentComponent,
        ProductSelectComponent
    ],
    exports: [
        ProductSelectComponent
    ],
    imports: [
        CommonModule,
        ProductRoutingModule,
        NzLayoutModule,
        NzMenuModule,
        NzPaginationModule,
        NzModalModule,
        NzIconModule,
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
        NzCollapseModule,
        NzSelectModule,
        NzTypographyModule,
        DragDropModule,
        NzUploadModule,
    ]
})
export class ProductModule {
}
