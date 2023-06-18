import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ValidatorRoutingModule } from './validator-routing.module';
import {ValidatorsComponent} from "./validators/validators.component";
import {ValidatorEditComponent} from "./validator-edit/validator-edit.component";
import {NzDividerModule} from "ng-zorro-antd/divider";
import {BaseModule} from "../base/base.module";
import {NzSpaceModule} from "ng-zorro-antd/space";
import {NzIconModule} from "ng-zorro-antd/icon";
import {NzButtonModule} from "ng-zorro-antd/button";
import {NzTableModule} from "ng-zorro-antd/table";
import {NzTagModule} from "ng-zorro-antd/tag";
import {NzFormModule} from "ng-zorro-antd/form";
import {NzInputNumberModule} from "ng-zorro-antd/input-number";
import {NzSelectModule} from "ng-zorro-antd/select";
import {NzCardModule} from "ng-zorro-antd/card";
import {ReactiveFormsModule} from "@angular/forms";
import {NzPopconfirmModule} from "ng-zorro-antd/popconfirm";


@NgModule({
    declarations: [
        ValidatorsComponent,
        ValidatorEditComponent,
    ],
    imports: [
        CommonModule,
        ValidatorRoutingModule,
        NzDividerModule,
        BaseModule,
        NzSpaceModule,
        NzIconModule,
        NzButtonModule,
        NzTableModule,
        NzTagModule,
        NzFormModule,
        NzInputNumberModule,
        NzSelectModule,
        NzCardModule,
        ReactiveFormsModule,
        NzPopconfirmModule
    ]
})
export class ValidatorModule { }
