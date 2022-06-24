import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {LoginComponent} from './login/login.component';
import {PageNotFoundComponent} from './page-not-found/page-not-found.component';
import {InstallComponent} from "./install/install.component";
import {LoginGuard} from "./login.guard";

const routes: Routes = [
  {path: '', pathMatch: 'full', redirectTo: '/admin'},
  {path: 'login', pathMatch: 'full', component: LoginComponent},
  {path: 'install', pathMatch: 'full', component: InstallComponent},
  {path: 'admin', loadChildren: () => import('./admin/admin.module').then(m => m.AdminModule), canActivate: [LoginGuard]},
  //{path: 'admin', loadChildren: './admin/admin.module#AdminModule'},

  {path: '**', component: PageNotFoundComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
