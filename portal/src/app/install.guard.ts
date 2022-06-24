import {Injectable} from '@angular/core';
import {ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree} from '@angular/router';
import {Observable} from 'rxjs';
import {InfoService} from "./info.service";
import {map} from "rxjs/operators";

@Injectable({
  providedIn: 'root'
})
export class InstallGuard implements CanActivate {
  constructor(private router: Router, private is: InfoService) {
  }

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
    if (this.is.info) {
      if (this.is.info.installed)
        return this.router.parseUrl("/admin");
      return true
    }
    return this.is.infoSub.asObservable().pipe(map(res => {
      if (res && res.installed)
        return this.router.parseUrl("/admin");
      return true;
    }))
  }

}
