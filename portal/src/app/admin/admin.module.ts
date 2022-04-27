import {NgModule} from '@angular/core';

import {IconsProviderModule} from './icons-provider.module';
import {NzLayoutModule} from 'ng-zorro-antd/layout';
import {NzMenuModule} from 'ng-zorro-antd/menu';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {CommonModule} from '@angular/common';
import {HttpClientModule} from '@angular/common/http';
import {AdminRoutingModule} from './admin-routing.module';

import {NzSpaceModule} from 'ng-zorro-antd/space';
import {NzIconModule} from 'ng-zorro-antd/icon';
import {NzToolTipModule} from 'ng-zorro-antd/tooltip';
import {NzTableModule} from 'ng-zorro-antd/table';
import {NzModalModule} from 'ng-zorro-antd/modal';
import {NzButtonModule} from 'ng-zorro-antd/button';
import {NzCheckboxModule} from 'ng-zorro-antd/checkbox';
import {NzSwitchModule} from 'ng-zorro-antd/switch';
import {NzPopconfirmModule} from 'ng-zorro-antd/popconfirm';
import {NzDividerModule} from 'ng-zorro-antd/divider';
import {NzDrawerModule} from 'ng-zorro-antd/drawer';
import {NzSelectModule} from 'ng-zorro-antd/select';
import {NzInputNumberModule} from 'ng-zorro-antd/input-number';
import {NzStatisticModule} from 'ng-zorro-antd/statistic';
import {NzCollapseModule} from 'ng-zorro-antd/collapse';
import {NzFormModule} from 'ng-zorro-antd/form';
import {NzInputModule} from 'ng-zorro-antd/input';
import {NzTabsModule} from 'ng-zorro-antd/tabs';
import {NzTransferModule} from 'ng-zorro-antd/transfer';
import {NzRadioModule} from 'ng-zorro-antd/radio';
import {NzProgressModule} from 'ng-zorro-antd/progress';
import {NzCardModule} from 'ng-zorro-antd/card';
import {NzUploadModule} from 'ng-zorro-antd/upload';
import {NzDropDownModule} from "ng-zorro-antd/dropdown";
import {NzDatePickerModule} from "ng-zorro-antd/date-picker";
import {NzTimePickerModule} from "ng-zorro-antd/time-picker";
import {DragDropModule} from "@angular/cdk/drag-drop";
import {NgxEchartsModule} from "ngx-echarts";
import {NzGridModule} from "ng-zorro-antd/grid";
import {NgxAmapModule} from "ngx-amap";
// import {NgxEchartsModule} from 'ngx-echarts';
//import * as echarts from 'echarts';

import {AdminComponent} from "./admin.component";

import {WelcomeComponent} from "./welcome/welcome.component";
import {UnknownComponent} from "./unknown/unknown.component";
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
import {TunnelEditComponent} from "./tunnel-edit/tunnel-edit.component";
import {LinkDetailComponent} from "./link-detail/link-detail.component";
import {LinkEditComponent} from "./link-edit/link-edit.component";
import {DeviceDetailComponent} from "./device-detail/device-detail.component";
import {DeviceEditComponent} from "./device-edit/device-edit.component";
import {ElementDetailComponent} from "./element-detail/element-detail.component";
import {ElementEditComponent} from "./element-edit/element-edit.component";
import {ProjectDetailComponent} from "./project-detail/project-detail.component";
import {ProjectEditComponent} from "./project-edit/project-edit.component";
import {TemplateDetailComponent} from "./template-detail/template-detail.component";
import {TemplateEditComponent} from "./template-edit/template-edit.component";
import {EditJobsComponent} from "./edit-jobs/edit-jobs.component";
import {EditPollersComponent} from "./edit-pollers/edit-pollers.component";
import {EditPointsComponent} from "./edit-points/edit-points.component";
import {EditCalculatorsComponent} from "./edit-calculators/edit-calculators.component";
import {EditCommandsComponent} from "./edit-commands/edit-commands.component";
import {EditAggregatorsComponent} from "./edit-aggregators/edit-aggregators.component";
import {EditStrategiesComponent} from "./edit-strategies/edit-strategies.component";
import {HelperModule} from "../helper/helper.module";
import {NzBreadCrumbModule} from "ng-zorro-antd/breadcrumb";
import {ContainerComponent} from "./container/container.component";
import {EditRegisterComponent} from "./edit-register/edit-register.component";
import {EditHeartbeatComponent} from "./edit-heartbeat/edit-heartbeat.component";
import {EditProtocolComponent} from "./edit-protocol/edit-protocol.component";
import {TunnelEditDevicesComponent} from "./tunnel-edit-devices/tunnel-edit-devices.component";
import {EventComponent} from "./event/event.component";
import {LinkDeviceComponent} from "./link-device/link-device.component";
import {TemplateProjectComponent} from "./template-project/template-project.component";
import {ElementDeviceComponent} from "./element-device/element-device.component";
import {ElementBrowserComponent} from "./element-browser/element-browser.component";
import {DeviceBrowserComponent} from "./device-browser/device-browser.component";
import {LinkBrowserComponent} from "./link-browser/link-browser.component";
import {ChooseService} from "./choose.service";
import {PromptComponent} from "./prompt/prompt.component";
import {UserBrowserComponent} from "./user-browser/user-browser.component";
import {AlarmComponent} from "./alarm/alarm.component";
import {UserDetailComponent} from "./user-detail/user-detail.component";
import {TemplateBrowserComponent} from "./template-browser/template-browser.component";
import {ChooseTemplateComponent} from "./choose-template/choose-template.component";
import {EditDevicesComponent} from "./edit-devices/edit-devices.component";
import {ChooseElementComponent} from "./choose-element/choose-element.component";
import {ChooseDeviceComponent} from "./choose-device/choose-device.component";
import {ChooseLinkComponent} from "./choose-link/choose-link.component.component";
import {EditElementsComponent} from "./edit-elements/edit-elements.component";
import {EditAlarmsComponent} from "./edit-alarms/edit-alarms.component";
import {EditDirectivesComponent} from "./edit-directives/edit-directives.component";
import {EditInvokesComponent} from "./edit-invokes/edit-invokes.component";
import {TunnelLinkComponent} from "./tunnel-link/tunnel-link.component";

@NgModule({
  declarations: [
    AdminComponent,
    ContainerComponent,
    WelcomeComponent,
    UnknownComponent,
    DashComponent,
    HomeComponent,
    TunnelComponent, TunnelDetailComponent, TunnelEditComponent,
    EditRegisterComponent, EditHeartbeatComponent, EditProtocolComponent,
    TunnelEditDevicesComponent, TunnelLinkComponent,
    LinkComponent, LinkDetailComponent, LinkEditComponent, LinkDeviceComponent,
    LinkBrowserComponent,
    DeviceComponent, DeviceDetailComponent, DeviceEditComponent,
    DeviceBrowserComponent,
    ElementComponent, ElementDetailComponent, ElementEditComponent,
    ElementDeviceComponent, ElementBrowserComponent,
    ProjectComponent, ProjectDetailComponent, ProjectEditComponent,
    TemplateComponent, TemplateDetailComponent, TemplateEditComponent,
    TemplateProjectComponent, TemplateBrowserComponent,
    EditPointsComponent, EditPollersComponent, EditJobsComponent, EditStrategiesComponent,
    EditDevicesComponent, EditElementsComponent, EditAlarmsComponent,
    EditCalculatorsComponent, EditCommandsComponent, EditAggregatorsComponent,
    EditDirectivesComponent, EditInvokesComponent,
    EventComponent, AlarmComponent,
    PluginComponent,
    ProtocolComponent,
    SettingComponent,
    UserComponent, UserBrowserComponent, UserDetailComponent,
    PasswordComponent,
    PromptComponent, ChooseTemplateComponent, ChooseElementComponent, ChooseDeviceComponent, ChooseLinkComponent,
  ],
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    AdminRoutingModule,
    IconsProviderModule,
    NzIconModule,
    NzGridModule,
    NzLayoutModule,
    NzMenuModule,
    NzToolTipModule,
    NzTableModule,
    NzModalModule,
    NzFormModule,
    NzButtonModule,
    NzInputModule,
    NzCheckboxModule,
    NzSwitchModule,
    NzPopconfirmModule,
    NzDividerModule,
    NzDrawerModule,
    NzSelectModule,
    NzSpaceModule,
    NzInputNumberModule,
    NzStatisticModule,
    NzTabsModule,
    NzCollapseModule,
    NzTransferModule,
    NzRadioModule,
    NzProgressModule,
    NzCardModule,
    NzUploadModule,
    NzDropDownModule,
    NzTimePickerModule,
    NzDatePickerModule,
    DragDropModule,

    NgxEchartsModule.forRoot({echarts: () => import('echarts')}),

    NgxAmapModule.forRoot({apiKey: 'e4c1bd11fe1b25d77dae4cf3993f7034', debug: true}),
    HelperModule,
    NzBreadCrumbModule,
  ],
  bootstrap: [AdminComponent],
  providers: [ChooseService]
})
export class AdminModule {
}
