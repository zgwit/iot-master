import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {AdminComponent} from './admin.component';
import {WelcomeComponent} from "./welcome/welcome.component";
import {UnknownComponent} from "./unknown/unknown.component";
import {LogoutGuard} from "./logout.guard";
import {DashComponent} from "./dash/dash.component";
import {HomeComponent} from "./home/home.component";
import {TunnelComponent} from "./tunnel/tunnel.component";
import {LinkComponent} from "./link/link.component";
import {DeviceComponent} from "./device/device.component";
import {ElementComponent} from "./element/element.component";
import {ProjectComponent} from "./project/project.component";
import {TemplateComponent} from "./template/template.component";
import {PluginComponent} from "./plugin/plugin.component";
import {ProtocolComponent} from "./protocol/protocol.component";
import {SettingComponent} from "./setting/setting.component";
import {UserComponent} from "./user/user.component";
import {PasswordComponent} from "./password/password.component";

const routes: Routes = [
  {
    path: '',
    component: AdminComponent,
    children: [
      {path: '', component: WelcomeComponent},
      {path: 'dash', component: DashComponent},
      {path: 'home', component: HomeComponent},
      {path: 'tunnel', component: TunnelComponent},
      {path: 'link', component: LinkComponent},
      {path: 'device', component: DeviceComponent},
      {path: 'element', component: ElementComponent},
      {path: 'project', component: ProjectComponent},
      {path: 'template', component: TemplateComponent},
      {path: 'plugin', component: PluginComponent},
      {path: 'protocol', component: ProtocolComponent},
      {path: 'setting', component: SettingComponent},
      {path: 'user', component: UserComponent},
      {path: 'password', component: PasswordComponent},
      {
        path: 'logout',
        canActivate: [LogoutGuard],
        //redirectTo: 'login'
      },

      {path: '**', component: UnknownComponent},
    ]
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AdminRoutingModule {
}
