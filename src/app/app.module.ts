import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { NZ_I18N } from 'ng-zorro-antd/i18n';
import { zh_CN } from 'ng-zorro-antd/i18n';
import { registerLocaleData } from '@angular/common';
import zh from '@angular/common/locales/zh';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LoginComponent } from './login/login.component';
import { PageNotFoundComponent } from './base/page-not-found/page-not-found.component';
import { RouterModule, RouterOutlet, Routes } from "@angular/router";
import { NzTabsModule } from "ng-zorro-antd/tabs";
import { NzFormModule } from "ng-zorro-antd/form";
import { NzInputModule } from "ng-zorro-antd/input";
import { NzButtonModule } from "ng-zorro-antd/button";
import { NzCheckboxModule } from "ng-zorro-antd/checkbox";
import { NzMessageModule } from "ng-zorro-antd/message";
import { NzLayoutModule } from "ng-zorro-antd/layout";
import { NzMenuModule } from "ng-zorro-antd/menu";
import { NzTableModule } from "ng-zorro-antd/table";
import { NzDividerModule } from "ng-zorro-antd/divider";
import { NzModalModule } from "ng-zorro-antd/modal";
import { DesktopComponent } from './desktop/desktop.component';
import { NzDrawerModule } from "ng-zorro-antd/drawer";
import { WindowComponent } from './window/window.component';
import { AdminComponent } from './admin/admin.component';
import { NzIconModule } from "ng-zorro-antd/icon";
import { NzDropDownModule } from "ng-zorro-antd/dropdown";
import { NzNotificationModule } from "ng-zorro-antd/notification";
import { authGuard } from "./auth.guard";
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzSwitchModule } from 'ng-zorro-antd/switch';
import { IMqttServiceOptions, MqttModule } from 'ngx-mqtt';
registerLocaleData(zh);

//declare var window: Window;

const MQTT_SERVICE_OPTIONS: IMqttServiceOptions = {
    hostname: window.location.hostname,
    port: parseInt(window.location.port),
    path: '/mqtt',
};

const pages: Routes = [
    {
        path: 'broker',
        canActivate: [authGuard],
        loadChildren: () => import('./broker/broker.module').then(m => m.BrokerModule)
    }, {
        path: 'device',
        canActivate: [authGuard],
        loadChildren: () => import('./device/device.module').then(m => m.DeviceModule)
    },
    // {
    //   path: 'alarm',
    //   canActivate: [authGuard],
    //   loadChildren: () => import('./alarm/alarm.module').then(m => m.AlarmModule)
    // },
    {
        path: 'setting',
        canActivate: [authGuard],
        loadChildren: () => import('./setting/setting.module').then(m => m.SettingModule)
    }, {
        path: 'user',
        canActivate: [authGuard],
        loadChildren: () => import('./user/user.module').then(m => m.UserModule)
    }, {
        path: 'product',
        canActivate: [authGuard],
        loadChildren: () => import('./product/product.module').then(m => m.ProductModule)
    }, {
        path: 'plugin',
        canActivate: [authGuard],
        loadChildren: () => import('./plugin/plugin.module').then(m => m.PluginModule)
    },
    {
        path: 'gateway',
        canActivate: [authGuard],
        loadChildren: () => import('./gateway/gateway.module').then(m => m.GatewayModule)
    },
]

const routes: Routes = [
    { path: '', redirectTo: 'desktop', pathMatch: 'full' },
    { path: 'login', component: LoginComponent },
    // {
    //     path: 'admin',
    //     component: AdminComponent,
    //     canActivate: [authGuard],
    //     children: [
    //         { path: '', pathMatch: "full", redirectTo: "device" },
    //         ...pages
    //     ]
    // },
    { path: 'desktop', component: DesktopComponent, canActivate: [authGuard] },
    ...pages,
    { path: '**', component: PageNotFoundComponent }
]

@NgModule({
    declarations: [AppComponent, LoginComponent, DesktopComponent, WindowComponent, AdminComponent],
    imports: [
        RouterModule.forRoot(routes),
        BrowserModule,
        FormsModule,
        HttpClientModule,
        BrowserAnimationsModule,
        RouterOutlet,
        NzTabsModule,
        ReactiveFormsModule,
        NzFormModule,
        NzInputModule,
        NzSwitchModule,
        NzButtonModule,
        NzSelectModule,
        NzCheckboxModule,
        NzMessageModule,
        NzNotificationModule,
        NzLayoutModule,
        NzMenuModule,
        NzTableModule,
        NzDividerModule,
        NzModalModule,
        NzDrawerModule,
        NzIconModule,
        NzDropDownModule,
        MqttModule.forRoot(MQTT_SERVICE_OPTIONS)
    ],
    providers: [{ provide: NZ_I18N, useValue: zh_CN }],
    bootstrap: [AppComponent],
})
export class AppModule {
}
