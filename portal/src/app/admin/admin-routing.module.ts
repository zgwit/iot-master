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
import {TunnelDetailComponent} from "./tunnel-detail/tunnel-detail.component";
import {LinkDetailComponent} from "./link-detail/link-detail.component";
import {DeviceDetailComponent} from "./device-detail/device-detail.component";
import {DeviceEditComponent} from "./device-edit/device-edit.component";
import {ElementEditComponent} from "./element-edit/element-edit.component";
import {ProjectDetailComponent} from "./project-detail/project-detail.component";
import {ProjectEditComponent} from "./project-edit/project-edit.component";
import {TemplateEditComponent} from "./template-edit/template-edit.component";
import {Container} from "@svgdotjs/svg.js";
import {ContainerComponent} from "./container/container.component";

const routes: Routes = [
  {
    path: '',
    component: AdminComponent,
    data: {breadcrumb: "后台"},
    children: [
      {path: '', component: WelcomeComponent, data: {breadcrumb: "欢迎"}},
      {path: 'dash', component: DashComponent},
      {path: 'home', component: HomeComponent, data: {breadcrumb: "首页"}},

      {
        path: 'tunnel', component: ContainerComponent, data: {breadcrumb: "通道"},
        children: [
          {path: '', component: TunnelComponent, data: {breadcrumb: "列表"}},
          {path: ':id', component: TunnelDetailComponent, data: {breadcrumb: "详情"}},
        ]
      },

      //{path: 'tunnel/edit/:id', component: TunnelEditComponent},
      //{path: 'tunnel/create', component: TunnelEditComponent},
      {path: 'link', component: LinkComponent, data: {breadcrumb: "链接"}},
      {path: 'link/:id', component: LinkDetailComponent, data: {breadcrumb: "详情"}},
      {path: 'device', component: DeviceComponent, data: {breadcrumb: "设备"}},
      {path: 'device/detail/:id', component: DeviceDetailComponent, data: {breadcrumb: "详情"}},
      {path: 'device/edit/:id', component: DeviceEditComponent, data: {breadcrumb: "编辑"}},
      {path: 'device/create', component: DeviceEditComponent, data: {breadcrumb: "创建"}},
      {path: 'element', component: ElementComponent},
      {path: 'element/edit/:id', component: ElementEditComponent},
      {path: 'element/create', component: ElementEditComponent},
      {path: 'project', component: ProjectComponent},
      {path: 'project/detail/:id', component: ProjectDetailComponent},
      {path: 'project/edit/:id', component: ProjectEditComponent},
      {path: 'project/create', component: ProjectEditComponent},
      {path: 'template', component: TemplateComponent},
      {path: 'template/edit/:id', component: TemplateEditComponent},
      {path: 'template/create', component: TemplateEditComponent},
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
