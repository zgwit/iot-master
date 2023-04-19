import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {BrokersComponent} from "./brokers/brokers.component";
import {BrokerEditComponent} from "./broker-edit/broker-edit.component";
import {PageNotFoundComponent} from "../base/page-not-found/page-not-found.component";
import {BrokerDetailComponent} from "./broker-detail/broker-detail.component";

const routes: Routes = [
  {path: '', pathMatch: "full", redirectTo: "list"},
  {path: 'list', component: BrokersComponent},
  {path: 'detail/:id', component: BrokerDetailComponent},
  {path: 'edit/:id', component: BrokerEditComponent},
  {path: 'create', component: BrokerEditComponent},
  {path: '**', component: PageNotFoundComponent}
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BrokerRoutingModule {
}
