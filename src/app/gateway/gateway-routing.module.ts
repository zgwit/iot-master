import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router'; 
import {PageNotFoundComponent} from "../base/page-not-found/page-not-found.component"; 
import { GatewayEditComponent } from './gateway-edit/gateway-edit.component';
import { GatewaysComponent } from './gateways/gateways.component';
import { GatewayBatchComponent } from './gateway-batch/gateway-batch.component';
const routes: Routes = [
  {path: '', pathMatch: "full", redirectTo: "list"}, 
  {path: 'list', component: GatewaysComponent}, 
  {path: 'edit/:id', component: GatewayEditComponent},
  {path: 'create', component: GatewayEditComponent},
  {path: 'batch', component: GatewayBatchComponent},
  {path: '**', component: PageNotFoundComponent}
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class GatewayRoutingModule {
}
