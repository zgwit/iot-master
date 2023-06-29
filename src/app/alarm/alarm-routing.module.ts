import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {PageNotFoundComponent} from "../base/page-not-found/page-not-found.component";
import {AlarmsComponent} from "./alarms/alarms.component";
import {SubscriptionComponent} from "./subscription/subscription.component";
import {SubscriptionEditComponent} from "./subscription-edit/subscription-edit.component";
import {NotificationComponent} from "./notification/notification.component";
import {ValidatorsComponent} from "./validators/validators.component";
import {ValidatorEditComponent} from "./validator-edit/validator-edit.component";

const routes: Routes = [
    {path: '', pathMatch: "full", redirectTo: "list"},
    {path: 'list', component: AlarmsComponent},
    {path: 'notification', component: NotificationComponent},
    {path: 'subscription', component: SubscriptionComponent},
    {path: 'subscription/edit/:id', component: SubscriptionEditComponent},
    {path: 'subscription/create', component: SubscriptionEditComponent},
    {path: 'validator', component: ValidatorsComponent},
    {path: 'validator/edit/:id', component: ValidatorEditComponent},
    {path: 'validator/create', component: ValidatorEditComponent},
    {path: '**', component: PageNotFoundComponent}
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class AlarmRoutingModule {
}
