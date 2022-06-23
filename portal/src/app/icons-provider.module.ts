import {NgModule} from '@angular/core';
import {NZ_ICONS, NzIconModule} from 'ng-zorro-antd/icon';

import {
  UserOutline,
  LockOutline,
  SaveOutline,
  ReloadOutline,
  PlusOutline,
  CloseOutline,
  MenuFoldOutline,
  MenuUnfoldOutline,
  HomeOutline,
  DashboardOutline,
  BlockOutline,
  SettingOutline,
  AppstoreAddOutline,
  ClusterOutline,
  LogoutOutline,
  BellOutline,
  AppstoreOutline,
  BarChartOutline,
  BuildOutline,
  FormOutline, DragOutline, DeleteOutline, LineChartOutline, SyncOutline, ProfileOutline, BookOutline,

  GithubOutline, CustomerServiceOutline,

  TableOutline, UnorderedListOutline, KeyOutline,
} from '@ant-design/icons-angular/icons';

const icons = [
  UserOutline,
  LockOutline,
  SaveOutline,
  ReloadOutline,
  PlusOutline,
  CloseOutline,


  // 菜单相关
  MenuFoldOutline, MenuUnfoldOutline, HomeOutline, DashboardOutline, BlockOutline,
  SettingOutline, AppstoreAddOutline, ClusterOutline, LogoutOutline, UserOutline,
  BellOutline,AppstoreOutline, BarChartOutline,BuildOutline,
  // 表格操作
  FormOutline,
  DragOutline,DeleteOutline,
  LineChartOutline,SyncOutline,

  ProfileOutline,

  BookOutline,
  GithubOutline,CustomerServiceOutline,

  TableOutline, UnorderedListOutline, KeyOutline,
];

@NgModule({
  imports: [NzIconModule],
  exports: [NzIconModule],
  providers: [
    {provide: NZ_ICONS, useValue: icons}
  ]
})
export class IconsProviderModule {
}
