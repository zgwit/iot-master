import { SubscriptionEditComponent } from './subscription-edit/subscription-edit.component';
import { SubscriptionComponent } from './subscription.component';
import { SubscriptionRoutingModule } from './subscription-routing.module';
import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common'; 
import {BaseModule} from "../base/base.module";
import {NzSpaceModule} from "ng-zorro-antd/space";
import {NzIconModule} from "ng-zorro-antd/icon";
import {NzButtonModule} from "ng-zorro-antd/button";
import {NzTableModule} from "ng-zorro-antd/table";
import {NzTagModule} from "ng-zorro-antd/tag";
import {NzDividerModule} from "ng-zorro-antd/divider";
import {NzCardModule} from "ng-zorro-antd/card";
import {NzFormModule} from "ng-zorro-antd/form";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NzSelectModule} from "ng-zorro-antd/select";
import {NzInputModule} from "ng-zorro-antd/input";
import {NzCollapseModule} from "ng-zorro-antd/collapse";
import { NzInputNumberModule } from 'ng-zorro-antd/input-number';
import { NzCheckboxModule } from 'ng-zorro-antd/checkbox';
@NgModule({
    declarations: [
        SubscriptionComponent ,
       SubscriptionEditComponent
    ],
    imports: [
        CommonModule,
        SubscriptionRoutingModule,
        BaseModule,
        NzSpaceModule,
        NzIconModule,
        NzButtonModule,
        NzTableModule,
        NzTagModule,
        NzCheckboxModule ,
        NzInputNumberModule,
        NzDividerModule,
        NzCardModule,
        NzFormModule,
        ReactiveFormsModule,
        NzSelectModule,
        NzInputModule,
        NzCollapseModule,
        FormsModule
    ]
})
export class SubscriptionModule {
}
