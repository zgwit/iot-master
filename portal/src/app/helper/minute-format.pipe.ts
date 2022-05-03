import { Pipe, PipeTransform } from '@angular/core';
import * as dayjs from "dayjs";

@Pipe({
  name: 'minuteFormat'
})
export class MinuteFormatPipe implements PipeTransform {

  transform(value: number): Date {
    const date = new Date()
    if (value < 0) value = 0;
    else if (value > 1439) value = 1439;
    date.setHours(Math.floor(value / 60), value % 60, 0, 0);
    // @ts-ignore
    return dayjs(date).format('HH:mm')
  }

}
