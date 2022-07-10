import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {AdminComponent} from './admin.component';
import {UnknownComponent} from "./unknown/unknown.component";
import {LogoutGuard} from "./logout.guard";
import {DashComponent} from "./dash/dash.component";
import {HomeComponent} from "./home/home.component";
import {ServerComponent} from "./server/server.component";
import {TunnelComponent} from "./tunnel/tunnel.component";
import {DeviceComponent} from "./device/device.component";
import {ProductComponent} from "./product/product.component";
import {ProjectComponent} from "./project/project.component";
import {TemplateComponent} from "./template/template.component";
import {PluginComponent} from "./plugin/plugin.component";
import {ProtocolComponent} from "./protocol/protocol.component";
import {SettingComponent} from "./setting/setting.component";
import {UserComponent} from "./user/user.component";
import {ServerDetailComponent} from "./server-detail/server-detail.component";
import {TunnelDetailComponent} from "./tunnel-detail/tunnel-detail.component";
import {DeviceDetailComponent} from "./device-detail/device-detail.component";
import {DeviceEditComponent} from "./device-edit/device-edit.component";
import {ProductEditComponent} from "./product-edit/product-edit.component";
import {ProjectDetailComponent} from "./project-detail/project-detail.component";
import {ProjectEditComponent} from "./project-edit/project-edit.component";
import {TemplateEditComponent} from "./template-edit/template-edit.component";
import {ContainerComponent} from "./container/container.component";
import {ServerEditComponent} from "./server-edit/server-edit.component";
import {TunnelEditComponent} from "./tunnel-edit/tunnel-edit.component";
import {TemplateDetailComponent} from "./template-detail/template-detail.component";
import {ProductDetailComponent} from "./product-detail/product-detail.component";
import {TunnelMonitorComponent} from "./tunnel-monitor/tunnel-monitor.component";
import {TransferComponent} from "./transfer/transfer.component";
import {TransferDetailComponent} from "./transfer-detail/transfer-detail.component";
import {TransferEditComponent} from "./transfer-edit/transfer-edit.component";
import {ComponentComponent} from "./component/component.component";
import {HmiComponent} from "./hmi/hmi.component";
import {HmiDetailComponent} from "./hmi-detail/hmi-detail.component";
import {DeviceValueComponent} from "./device-value/device-value.component";
import {UserDetailComponent} from "./user-detail/user-detail.component";
import {HmiEditContentComponent} from "./hmi-edit-content/hmi-edit-content.component";
import {CameraComponent} from "./camera/camera.component";
import {CameraDetailComponent} from "./camera-detail/camera-detail.component";
import {CameraEditComponent} from "./camera-edit/camera-edit.component";
import {LicenseComponent} from "./license/license.component";
import {HmiEditComponent} from "./hmi-edit/hmi-edit.component";
import {ComponentEditComponent} from "./component-edit/component-edit.component";
import {ComponentEditContentComponent} from "./component-edit-content/component-edit-content.component";
import {ComponentDetailComponent} from "./component-detail/component-detail.component";

const routes: Routes = [
  {
    path: 'hmi-edit/:id',
    component: HmiEditContentComponent,
  },
  {
    path: '',
    component: AdminComponent,
    data: {breadcrumb: "后台"},
    children: [
      //{path: '', component: WelcomeComponent, data: {breadcrumb: "欢迎"}},
      {path: '', component: DashComponent, data: {breadcrumb: "控制台"}},
      {path: 'dash', component: DashComponent, data: {breadcrumb: "控制台"}},
      {path: 'home', component: HomeComponent, data: {breadcrumb: "首页"}},

      {
        path: 'server', component: ContainerComponent, data: {breadcrumb: "服务器"}, children: [
          {path: '', component: ServerComponent, data: {breadcrumb: "服务器"}},
          //{path: ':id', component: TunnelDetailComponent, data: {breadcrumb: "详情"}},
          {path: 'detail/:id', component: ServerDetailComponent, data: {breadcrumb: "详情"}},
          {path: 'edit/:id', component: ServerEditComponent, data: {breadcrumb: "编辑"}},
          {path: 'create', component: ServerEditComponent, data: {breadcrumb: "创建"}},
        ]
      },

      {
        path: 'tunnel', component: ContainerComponent, data: {breadcrumb: "通道"}, children: [
          {path: '', component: TunnelComponent, data: {breadcrumb: "通道"}},
          {path: 'detail/:id', component: TunnelDetailComponent, data: {breadcrumb: "详情"}},
          {path: 'edit/:id', component: TunnelEditComponent, data: {breadcrumb: "编辑"}},
          {path: 'create', component: TunnelEditComponent, data: {breadcrumb: "创建"}},
          {path: 'monitor/:id', component: TunnelMonitorComponent, data: {breadcrumb: "监控"}},
        ]
      },

      {
        path: 'transfer', component: ContainerComponent, data: {breadcrumb: "远程调试"}, children: [
          {path: '', component: TransferComponent, data: {breadcrumb: "远程调试"}},
          {path: 'detail/:id', component: TransferDetailComponent, data: {breadcrumb: "详情"}},
          {path: 'edit/:id', component: TransferEditComponent, data: {breadcrumb: "编辑"}},
          {path: 'create', component: TransferEditComponent, data: {breadcrumb: "创建"}},
        ]
      },

      {
        path: 'device', component: ContainerComponent, data: {breadcrumb: "设备"}, children: [
          {path: '', component: DeviceComponent, data: {breadcrumb: "设备"}},
          {path: 'detail/:id', component: DeviceDetailComponent, data: {breadcrumb: "详情"}},
          {path: 'edit/:id', component: DeviceEditComponent, data: {breadcrumb: "编辑"}},
          {path: 'create', component: DeviceEditComponent, data: {breadcrumb: "创建"}},
          {path: 'value/:id/:name', component: DeviceValueComponent, data: {breadcrumb: "历史"}},
        ]
      },

      {
        path: 'product', component: ContainerComponent, data: {breadcrumb: "产品库"}, children: [
          {path: '', component: ProductComponent, data: {breadcrumb: "产品库"}},
          {path: 'detail/:id', component: ProductDetailComponent, data: {breadcrumb: "详情"}},
          {path: 'edit/:id', component: ProductEditComponent, data: {breadcrumb: "编辑"}},
          {path: 'create', component: ProductEditComponent, data: {breadcrumb: "创建"}},
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
        path: 'template', component: ContainerComponent, data: {breadcrumb: "模板库"}, children: [
          {path: '', component: TemplateComponent, data: {breadcrumb: "模板库"}},
          {path: 'detail/:id', component: TemplateDetailComponent, data: {breadcrumb: "详情"}},
          {path: 'edit/:id', component: TemplateEditComponent, data: {breadcrumb: "编辑"}},
          {path: 'create', component: TemplateEditComponent, data: {breadcrumb: "创建"}},
        ]
      },

      {
        path: 'hmi', component: ContainerComponent, data: {breadcrumb: "组态"}, children: [
          {path: '', component: HmiComponent, data: {breadcrumb: "组态"}},
          {path: 'detail/:id', component: HmiDetailComponent, data: {breadcrumb: "详情"}},
          {path: 'edit/:id', component: HmiEditComponent, data: {breadcrumb: "编辑"}},
          {path: 'create', component: HmiEditComponent, data: {breadcrumb: "创建"}},
          {path: 'edit-content/:id', component: HmiEditContentComponent, data: {breadcrumb: "编辑"}},
        ]
      },

      {
        path: 'component', component: ContainerComponent, data: {breadcrumb: "组件库"}, children: [
          {path: '', component: ComponentComponent, data: {breadcrumb: "组件库"}},
          {path: 'detail/:id', component: ComponentDetailComponent, data: {breadcrumb: "详情"}},
          {path: 'edit/:id', component: ComponentEditComponent, data: {breadcrumb: "编辑"}},
          {path: 'create', component: ComponentEditComponent, data: {breadcrumb: "创建"}},
          {path: 'edit-content/:id', component: ComponentEditContentComponent, data: {breadcrumb: "编辑"}},
        ]
      },

      {
        path: 'extension', component: ContainerComponent, data: {breadcrumb: "扩展"}, children: [
          {path: 'plugin', component: PluginComponent, data: {breadcrumb: "插件"}},
          {path: 'protocol', component: ProtocolComponent, data: {breadcrumb: "协议"}},
        ]
      },

      {
        path: 'camera', component: ContainerComponent, data: {breadcrumb: "摄像头"}, children: [
          {path: '', component: CameraComponent, data: {breadcrumb: "摄像头"}},
          {path: 'detail/:id', component: CameraDetailComponent, data: {breadcrumb: "详情"}},
          {path: 'edit/:id', component: CameraEditComponent, data: {breadcrumb: "编辑"}},
          {path: 'create', component: CameraEditComponent, data: {breadcrumb: "创建"}},
        ]
      },

      {
        path: 'setting', component: ContainerComponent, data: {breadcrumb: "设置"}, children: [
          {path: '', component: SettingComponent, data: {breadcrumb: "系统"}},
          {path: 'user', component: UserComponent, data: {breadcrumb: "用户"}},
          {path: 'license', component: LicenseComponent, data: {breadcrumb: "激活码"}},
        ]
      },

      {
        path: 'user', component: ContainerComponent, data: {breadcrumb: "用户"}, children: [
          {path: '', component: UserComponent, data: {breadcrumb: "用户"}},
          {path: 'detail/:id', component: UserDetailComponent, data: {breadcrumb: "详情"}},
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
