import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {HistorysComponent} from "./historys/historys.component";
import {AggregatorEditComponent} from "./aggregator-edit/aggregator-edit.component";
import {AggregatorsComponent} from "./aggregators/aggregators.component";
import {PageNotFoundComponent} from "../base/page-not-found/page-not-found.component";

const routes: Routes = [
    {path: '', pathMatch: "full", redirectTo: "list"},
    {path: 'list', component: HistorysComponent},
    {path: 'aggregators', component: AggregatorsComponent},
    {path: 'edit/:id', component: AggregatorEditComponent},
    {path: 'create', component: AggregatorEditComponent},
    {path: '**', component: PageNotFoundComponent}
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class HistoryRoutingModule {
}
