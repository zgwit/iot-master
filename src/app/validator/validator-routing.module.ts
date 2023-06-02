import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {ValidatorsComponent} from "./validators/validators.component";
import {ValidatorEditComponent} from "./validator-edit/validator-edit.component";
import {PageNotFoundComponent} from "../base/page-not-found/page-not-found.component";

const routes: Routes = [
  {path: '', pathMatch: "full", redirectTo: "list"},
  {path: 'list', component: ValidatorsComponent},
  {path: 'edit/:id', component: ValidatorEditComponent},
  {path: 'create', component: ValidatorEditComponent},
  {path: '**', component: PageNotFoundComponent}
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ValidatorRoutingModule { }
