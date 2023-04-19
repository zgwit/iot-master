import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {PageNotFoundComponent} from "../base/page-not-found/page-not-found.component";
import {MeComponent} from "./me/me.component";
import {PasswordComponent} from "./password/password.component";
import {UsersComponent} from "./users/users.component";
import {UserEditComponent} from "./user-edit/user-edit.component";
import {UserDetailComponent} from "./user-detail/user-detail.component";
import {RoleComponent} from "./role/role.component";
import { RoleEditComponent } from './role-edit/role-edit.component';
const routes: Routes = [
  {path: '', pathMatch: "full", redirectTo: "list"},
  {path: 'list', component: UsersComponent},
  {path: 'detail/:id', component: UserDetailComponent},
  {path: 'edit/:id', component: UserEditComponent},
  {path: 'create', component: UserEditComponent},
  {path: 'role/create', component: RoleEditComponent},
  {path: 'role', component: RoleComponent},
  {path: 'privillege/:id', component: RoleEditComponent },
  {path: 'me', component: MeComponent},
  {path: 'password', component: PasswordComponent},
  {path: '**', component: PageNotFoundComponent}
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UserRoutingModule {
}
