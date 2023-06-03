import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {HistoryRoutingModule} from './history-routing.module';
import {HistorysComponent} from "./historys/historys.component";
import {AggregatorsComponent} from "./aggregators/aggregators.component";
import {AggregatorEditComponent} from "./aggregator-edit/aggregator-edit.component";
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


@NgModule({
    declarations: [
        HistorysComponent,
        AggregatorsComponent,
        AggregatorEditComponent,
    ],
    imports: [
        CommonModule,
        HistoryRoutingModule,
        BaseModule,
        NzSpaceModule,
        NzIconModule,
        NzButtonModule,
        NzTableModule,
        NzTagModule,
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
export class HistoryModule {
}
