import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NzLayoutModule } from 'ng-zorro-antd/layout';
import { NzMenuModule } from 'ng-zorro-antd/menu';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzCardModule } from 'ng-zorro-antd/card';
import { NzFormModule } from 'ng-zorro-antd/form';
import { ReactiveFormsModule } from '@angular/forms';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzInputNumberModule } from 'ng-zorro-antd/input-number';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { BaseModule } from '../base/base.module';
import { NzSpaceModule } from 'ng-zorro-antd/space';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzDividerModule } from 'ng-zorro-antd/divider';
import { NzPopconfirmModule } from 'ng-zorro-antd/popconfirm';
import { GatewaysComponent } from './gateways/gateways.component';
import { GatewayEditComponent } from './gateway-edit/gateway-edit.component';
import { GatewayRoutingModule } from './gateway-routing.module';
import { GatewayDetailComponent } from './gateway-detail/gateway-detail.component';
import { NzSwitchModule } from 'ng-zorro-antd/switch';  
import { FormsModule } from '@angular/forms';
import { GatewayBatchComponent } from './gateway-batch/gateway-batch.component';
import { NzTagModule } from 'ng-zorro-antd/tag';
@NgModule({
  declarations: [GatewaysComponent, GatewayEditComponent, GatewayDetailComponent, GatewayBatchComponent],
  imports: [
    CommonModule,
    NzLayoutModule,
    NzMenuModule,
    NzIconModule,
    GatewayRoutingModule ,
    NzSwitchModule,
    NzCardModule,
    NzTagModule,
    NzFormModule,
    NzPopconfirmModule,
    ReactiveFormsModule,
    FormsModule,
    NzInputModule,
    NzInputNumberModule,
    NzButtonModule,
    BaseModule,
    NzSpaceModule,
    NzTableModule,
    NzDividerModule,
  ],
})
export class GatewayModule {}
