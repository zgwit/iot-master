import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {PageNotFoundComponent} from "../base/page-not-found/page-not-found.component";
import {WebComponent} from "./web/web.component";
import {DatabaseComponent} from "./database/database.component";
import {LogComponent} from "./log/log.component";
import {BackupComponent} from "./backup/backup.component";
import {RebootComponent} from "./reboot/reboot.component";
import { MqttComponent } from './mqtt/mqtt.component';
import { OemComponent } from './oem/oem.component';
const routes: Routes = [
  {path: '', pathMatch: "full", redirectTo: "web"},
  {path: 'web', component: WebComponent},
  {path: 'database', component: DatabaseComponent},
  {path: 'log', component: LogComponent},
  {path: 'backup', component: BackupComponent},
  {path: 'reboot', component: RebootComponent},
  {path: 'mqtt', component: MqttComponent},
  // {path: 'oem', component: OemComponent},
  {path: '**', component: PageNotFoundComponent}
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class SettingRoutingModule {
}
