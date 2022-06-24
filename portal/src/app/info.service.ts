import { Injectable } from '@angular/core';
import {RequestService} from "./request.service";
import {Router} from "@angular/router";

@Injectable({
  providedIn: 'root'
})
export class InfoService {
  public info: any = {}

  constructor(private rs: RequestService, private route: Router) {
    this.rs.get("info").subscribe(res=>{
      this.info = res.data
      if (!this.info.installed) {
        this.route.navigate(['/install']);
      }
    })
  }
}
