import {ActivatedRouteSnapshot, CanMatchFn, ResolveFn, Router, RouterStateSnapshot} from '@angular/router';
import {inject} from "@angular/core";
import {UserService} from "./user.service";
import {Subject} from "rxjs";
import {SmartRequestService} from "@god-jason/smart";

export const ProjectGuard: ResolveFn<any> = (route: ActivatedRouteSnapshot, state: RouterStateSnapshot) => {


}

export const projectGuard: CanMatchFn = () => {
    const us = inject(UserService);
    const rs = inject(SmartRequestService);
    const route = inject(ActivatedRouteSnapshot);
    const router = inject(Router);

    const project = route.paramMap.get("project")

    const sub = new Subject<any>()

    if (us.user) {
        if (us.user.admin)
            return true;
        rs.get(`project/${project}/user/${us.user.id}/exists`).subscribe(res => {
            if (res.data)
                sub.next(true)
            else
                sub.next(router.parseUrl("/select"))
        })
    }

    if (us.getting) {
        us.userSub.subscribe(res => {
            if (res.admin)
                sub.next(true)
            rs.get(`project/${project}/user/${us.user.id}/exists`).subscribe(res => {
                if (res.data)
                    sub.next(true)
                else
                    sub.next(router.parseUrl("/select"))
            })
        })
    }

    return sub.asObservable()
};
