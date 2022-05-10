import {Component, Input, OnDestroy, OnInit} from '@angular/core';
import {RequestService} from "../../request.service";

@Component({
  selector: 'app-dash-cpu',
  templateUrl: './cpu.component.html',
  styleUrls: ['./cpu.component.scss']
})
export class CpuComponent implements OnInit, OnDestroy {
  @Input() interval = 5000;

  options: any = {}

  last = {
    busy: 0,
    total: 0,
  };

  info: any = {
    usage: 0,
  };
  handle: any;

  constructor(private rs: RequestService) {
    this.load();
  }

  ngOnInit(): void {
    this.handle = setInterval(() => {
      this.load()
    }, this.interval);
  }

  ngOnDestroy(): void {
    clearInterval(this.handle)
  }

  load(): void {
    this.rs.get('system/cpu').subscribe(res => {
      //console.log('cpu info', res)
      let busy = res.data.user + res.data.system;
      let total = busy + res.data.idle;

      let usage = (busy - this.last.busy) / (total - this.last.total) * 100;
      this.last.busy = busy
      this.last.total = total

      console.log(busy, total, usage)

      this.info.usage = usage
      //this.options.series[0].data[0].value = usage.toFixed(2);
      this.options = {
        tooltip: {
          formatter: '{b} : {c}%'
        },
        series: [
          {
            name: 'CPU',
            type: 'gauge',
            data: [
              {
                value: Math.round(usage),
                name: 'CPU'
              }
            ]
          }
        ]
      };
    })
  }


}
