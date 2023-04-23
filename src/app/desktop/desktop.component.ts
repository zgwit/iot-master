import { Component } from '@angular/core';
import { Router } from "@angular/router";
import { RequestService } from "../request.service";
import { NzModalService } from "ng-zorro-antd/modal";
import { WindowComponent } from "../window/window.component";
import { OemService } from "../oem.service";
import { AppService } from "../app.service";

declare var window: any;

@Component({
  selector: 'app-desktop',
  templateUrl: './desktop.component.html',
  styleUrls: ['./desktop.component.scss']
})
export class DesktopComponent {

  drawVisible: any;
  constructor(
    private router: Router,
    private rs: RequestService,
    private ms: NzModalService,
    protected os: OemService,
    protected _as: AppService) {
    localStorage.setItem("main", "/desktop")
  }

  open(app: any) {
    if (window.innerWidth < 800) {
      this.router.navigate([app.entries[0].path])
      return;
    }

    this.ms.create({
      nzTitle: app.name,
      nzFooter: null,
      //nzMask: false,
      nzMaskClosable: false,
      nzWidth: "90%",
      //nzStyle: {height: "90%"},
      nzBodyStyle: { padding: "0", overflow: "hidden" },
      nzContent: WindowComponent,
      nzData: {
        entries: app.entries || []
      }
    })
  }
}
