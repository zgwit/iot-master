import {Injectable} from '@angular/core';
import {NzModalService} from "ng-zorro-antd/modal";
import {UserBrowserComponent} from "./user-browser/user-browser.component";
import {TemplateBrowserComponent} from "./template-browser/template-browser.component";
import {ElementBrowserComponent} from "./element-browser/element-browser.component";
import {DeviceBrowserComponent} from "./device-browser/device-browser.component";
import {LinkBrowserComponent} from "./link-browser/link-browser.component";
import {PromptComponent} from "./prompt/prompt.component";

@Injectable({
  providedIn: "root"
})
export class ChooseService {

  constructor(private ms: NzModalService) {
  }

  chooseUser(params?: any) {
    const modal = this.ms.create({
      nzTitle: '选择用户',
      nzContent: UserBrowserComponent,
      nzWidth: '80%',
      nzComponentParams: params,
    });
    return modal.afterClose
  }

  chooseDevice(params?: any) {
    const modal = this.ms.create({
      nzTitle: '选择设备',
      nzContent: DeviceBrowserComponent,
      nzWidth: '80%',
      nzComponentParams: params,
    });
    return modal.afterClose
  }

  chooseLink(params?: any) {
    const modal = this.ms.create({
      nzTitle: '选择链接',
      nzContent: LinkBrowserComponent,
      nzWidth: '80%',
      nzComponentParams: params,
    });
    return modal.afterClose
  }

  chooseElement(params?: any) {
    const modal = this.ms.create({
      nzTitle: '选择元件',
      nzContent: ElementBrowserComponent,
      nzWidth: '80%',
      nzComponentParams: params,
    });
    return modal.afterClose
  }

  chooseTemplate() {
    const modal = this.ms.create({
      nzTitle: '选择模板',
      nzContent: TemplateBrowserComponent,
      nzWidth: '80%',
    });
    return modal.afterClose
  }

  prompt(params: any) {
    const modal = this.ms.create({
      nzTitle: '请输入',
      nzContent: PromptComponent,
      nzWidth: '80%',
      nzComponentParams: params,
    });
    return modal.afterClose
  }
}
