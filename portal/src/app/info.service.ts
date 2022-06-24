import { Injectable } from '@angular/core';
import {RequestService} from "./request.service";
import {Router} from "@angular/router";
import {Subject} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class InfoService {
  public info: any;
  public infoSub = new Subject<any>();

  constructor(private rs: RequestService, private route: Router) {
    this.rs.get("info").subscribe(res=>{
      this.info = res.data
      this.infoSub.next(res.data)

      if (!this.info.installed) {
        this.route.navigate(['/install']);
      }
    }, error=>{
      this.infoSub.next(undefined)
    })
  }
}
