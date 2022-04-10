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
        path: 'tunnel', component: ContainerComponent, data: {breadcrumb: "通道"}, children: [
          {path: '', component: TunnelComponent, data: {breadcrumb: "列表"}},
          {path: ':id', component: TunnelDetailComponent, data: {breadcrumb: "详情"}},
          //{path: 'tunnel/edit/:id', component: TunnelEditComponent},
          //{path: 'tunnel/create', component: TunnelEditComponent},
        ]
      },

      {
        path: 'link', component: ContainerComponent, data: {breadcrumb: "链接"}, children: [
          {path: '', component: LinkComponent, data: {breadcrumb: "链接"}},
          {path: ':id', component: LinkDetailComponent, data: {breadcrumb: "详情"}},
        ]
      },

      {
        path: 'device', component: ContainerComponent, data: {breadcrumb: "设备"}, children: [
          {path: '', component: DeviceComponent, data: {breadcrumb: "设备"}},
          {path: 'detail/:id', component: DeviceDetailComponent, data: {breadcrumb: "详情"}},
          {path: 'edit/:id', component: DeviceEditComponent, data: {breadcrumb: "编辑"}},
          {path: 'create', component: DeviceEditComponent, data: {breadcrumb: "创建"}},
        ]
      },

      {
        path: 'element', component: ContainerComponent, data: {breadcrumb: "元件"}, children: [
          {path: '', component: ElementComponent, data: {breadcrumb: "元件"}},
          {path: 'edit/:id', component: ElementEditComponent, data: {breadcrumb: "编辑"}},
          {path: 'create', component: ElementEditComponent, data: {breadcrumb: "创建"}},
        ]
      },

      {
        path: 'project', component: ContainerComponent, data: {breadcrumb: "项目"}, children: [
          {path: '', component: ProjectComponent, data: {breadcrumb: "项目"}},
          {path: 'detail/:id', component: ProjectDetailComponent, data: {breadcrumb: "详情"}},
          {path: 'edit/:id', component: ProjectEditComponent, data: {breadcrumb: "编辑"}},
          {path: 'create', component: ProjectEditComponent, data: {breadcrumb: "创建"}},
        ]
      },

      {
        path: 'template', component: ContainerComponent, data: {breadcrumb: "模板"}, children: [
          {path: '', component: TemplateComponent, data: {breadcrumb: "模板"}},
          {path: 'edit/:id', component: TemplateEditComponent, data: {breadcrumb: "编辑"}},
          {path: 'create', component: TemplateEditComponent, data: {breadcrumb: "创建"}},
        ]
      },

      {
        path: 'extension', component: ContainerComponent, data: {breadcrumb: "扩展"}, children: [
          {path: 'plugin', component: PluginComponent, data: {breadcrumb: "插件"}},
          {path: 'protocol', component: ProtocolComponent, data: {breadcrumb: "协议"}},
        ]
      },

      {
        path: 'setting', component: ContainerComponent, children: [
          {path: '', component: SettingComponent},
          {path: 'user', component: UserComponent},
          {path: 'password', component: PasswordComponent},
        ]
      },

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
