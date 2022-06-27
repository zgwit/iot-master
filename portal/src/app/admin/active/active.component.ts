import {Component, HostListener, OnInit} from '@angular/core';
import {DomSanitizer} from "@angular/platform-browser";
import {RequestService} from "../../request.service";
import {NzModalRef} from "ng-zorro-antd/modal";

@Component({
  selector: 'app-active',
  templateUrl: './active.component.html',
  styleUrls: ['./active.component.scss']
})
export class ActiveComponent implements OnInit {
  url: any = ""

  constructor(private ds: DomSanitizer, private rs: RequestService, private mr: NzModalRef) {
    this.url = ds.bypassSecurityTrustResourceUrl("https://license.zgwit.com/active/iot-master-ce")
  }

  ngOnInit(): void {
  }

  @HostListener('window:message', ['$event'])
  messageListener(event: any) {
    const res = event.data
    if (res.type == "active") {
      console.log(res)
      this.rs.post("/license", {license: res.data}).subscribe(res=>{
        this.mr.close()
      })
    }
  }

}
