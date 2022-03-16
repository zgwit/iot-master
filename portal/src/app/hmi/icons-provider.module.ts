import {NgModule} from '@angular/core';
import {NZ_ICONS, NzIconModule} from 'ng-zorro-antd/icon';

import {
  SaveOutline,
  ExportOutline,
  RedoOutline,
  UndoOutline,
  CopyOutline,
  ScissorOutline,
  SnippetsOutline,
  DeleteOutline,
  AlignLeftOutline,
  AlignCenterOutline,
  AlignRightOutline,
  VerticalAlignTopOutline,
  VerticalAlignMiddleOutline,
  VerticalAlignBottomOutline,
  UpOutline,
  DownOutline,
  LeftOutline,
  RightOutline,
  VerticalLeftOutline,
  VerticalRightOutline,
  GroupOutline,
  UngroupOutline,
} from '@ant-design/icons-angular/icons';
import {CommonModule} from '@angular/common';

const icons = [
  SaveOutline,ExportOutline,RedoOutline,UndoOutline,
  CopyOutline,ScissorOutline, SnippetsOutline, DeleteOutline,
  AlignLeftOutline,AlignCenterOutline,AlignRightOutline,
  VerticalAlignTopOutline,VerticalAlignMiddleOutline,VerticalAlignBottomOutline,
  UpOutline,DownOutline,LeftOutline,RightOutline,
  VerticalLeftOutline,VerticalRightOutline,
  GroupOutline,UngroupOutline
];

@NgModule({
  imports: [CommonModule, NzIconModule.forChild(icons)],
  exports: [NzIconModule],
  providers: [
    {provide: NZ_ICONS, useValue: icons}
  ]
})
export class IconsProviderModule {
}
