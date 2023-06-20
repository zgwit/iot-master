import { SubscriptionEditComponent } from './subscription-edit/subscription-edit.component';
import { SubscriptionComponent } from './subscription.component';
import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router'; 
import {PageNotFoundComponent} from "../base/page-not-found/page-not-found.component";

const routes: Routes = [
    {path: '', pathMatch: "full", redirectTo: "list"},
    {path: 'list', component:SubscriptionComponent}, 
    {path: 'edit/:id', component: SubscriptionEditComponent},
    {path: 'create', component: SubscriptionEditComponent},
    {path: '**', component: PageNotFoundComponent}
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class SubscriptionRoutingModule {
}
