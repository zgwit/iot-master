import {Pipe, PipeTransform} from '@angular/core';
import * as dayjs from "dayjs";
import * as relativeTime from "dayjs/plugin/relativeTime";
import "dayjs/locale/zh-cn";

dayjs.locale("zh-cn");
dayjs.extend(relativeTime);

@Pipe({
  name: 'fromNow'
})
export class FromNowPipe implements PipeTransform {

  transform(date: any): any {
    let d = dayjs(date)
    if (dayjs().diff(d, "day") < 7)
      return d.fromNow()
    return d.format("YYYY-MM-DD HH:mm:ss")
  }

}
