import { Pipe, PipeTransform } from '@angular/core';
import { DatePipe } from '@angular/common';

@Pipe({
    name: 'chinaDate'
})
export class ChinaDatePipe extends DatePipe implements PipeTransform {
    override transform(value: any, ...args: any[]): any {
        let format = 'yyyy-MM-dd HH:mm:ss';
        if (args.length > 0) {
            if (args[0] === 'date') {
                format = 'yyyy-MM-dd';
            } else if (args[0] === 'datetime') {
                format = 'yyyy-MM-dd HH:mm:ss';
            } else {
                format = args[0];
            }
        }
        return super.transform(value, format, 'UTC+8');
    }
}
