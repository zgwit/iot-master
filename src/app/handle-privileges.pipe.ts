import { Pipe, PipeTransform } from '@angular/core';
import { RequestService } from './request.service';

@Pipe({
  name: 'handlePrivileges'
})
export class HandlePrivilegesPipe implements PipeTransform {
  range = 40000
  constructor(
    private rs: RequestService,
  ) {
    const roleObj = localStorage.getItem('roleObj');
    if (!roleObj || !this.judgelocalStorage('roleObj')) {
      this.rs
        .get('privileges')
        .subscribe((res) => {
          const data = res.data || [];
          this.setlocalStorage('roleObj', data);
        })
    }
  }
  transform(value: any, ...args: unknown[]): unknown {
    const arr: any = [];
    if (localStorage.getItem('roleObj')) {
      const data = localStorage.getItem('roleObj') || '';
      const roleObj = JSON.parse(data);
      value.forEach((item: string) => {
        if (roleObj[item]) {
          arr.push(roleObj[item]);
        }
      });
    }
    return arr.join(',');
  }

  setlocalStorage(key: string, value: any) {
    const date = new Date().getTime();
    value.date = date;
    localStorage.setItem(key, JSON.stringify(value))
  }
  // 判断是否过期半小时
  judgelocalStorage(key: any) {
    const dateNow = new Date().getTime();
    const storage = JSON.parse(localStorage.getItem(key) || '')
    if (dateNow - storage.time > this.range) {
      storage.removeItem(key);
      return null
    } else {
      return storage.key;
    }
  }

}
