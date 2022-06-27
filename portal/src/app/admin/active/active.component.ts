import {Component, HostListener, OnInit} from '@angular/core';
import {DomSanitizer} from "@angular/platform-browser";
import {RequestService} from "../../request.service";
import {NzModalRef} from "ng-zorro-antd/modal";
import {NzMessageService} from "ng-zorro-antd/message";

@Component({
  selector: 'app-active',
  templateUrl: './active.component.html',
  styleUrls: ['./active.component.scss']
})
export class ActiveComponent implements OnInit {
  url: any = ""

  constructor(private ds: DomSanitizer, private rs: RequestService, private mr: NzModalRef, private ms: NzMessageService) {
    //this.url = ds.bypassSecurityTrustResourceUrl("https://license.zgwit.com/active/iot-master-ce")
    //this.url = ds.bypassSecurityTrustResourceUrl("http://localhost:4200/active/iot-master-ce")
    rs.get("system/machine").subscribe(res=>{
      let url = "https://license.zgwit.com/active/iot-master-ce"
      //let url = "http://localhost:4200/active/iot-master-ce"
      url += "?cpu=" + encodeURIComponent(res.data.cpu)
      url += "&mac=" + encodeURIComponent(res.data.mac)
      url += "&sn=" + encodeURIComponent(res.data.sn)
      url += "&uuid=" + encodeURIComponent(res.data.uuid)
      this.url = ds.bypassSecurityTrustResourceUrl(url)
    })
  }

  ngOnInit(): void {
  }

  @HostListener('window:message', ['$event'])
  messageListener(event: any) {
    console.log(event)
    const res = JSON.parse(event.data)
    if (res.type == "active") {
      this.rs.post("license", {license: res.data}).subscribe(res=>{
        this.ms.success("激活成功")
        this.mr.close()
      })
    }
  }

}
