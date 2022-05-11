import {Component, Input, OnDestroy, OnInit} from '@angular/core';
import {RequestService} from "../../request.service";
import * as filesize from "filesize";

@Component({
  selector: 'app-dash-mem',
  templateUrl: './mem.component.html',
  styleUrls: ['./mem.component.scss']
})
export class MemComponent implements OnInit {
  @Input() interval = 30000;

  info: any = {
    used: 0,
    total: 0,
  };
  handle: any;
  options: any = {};

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
    this.rs.get('system/memory').subscribe(res => {
      //console.log('mem info', res)
      this.info = res.data;
      this.options = {
        title: {
          text: '内存',
          left: 'center'
        },
        color: [
          '#fac858',
          '#91cc75',
          '#fc8452',
          '#73c0de',
          '#3ba272',
          '#5470c6',
          '#ee6666',
          '#9a60b4',
          '#ea7ccc'
        ],
        tooltip: {
          //formatter: '{b} {c}',
          valueFormatter: (value: number) => filesize(value)
        },
        series: [
          {
            type: 'pie',
            radius: '65%',
            center: ['50%', '50%'],
            avoidLabelOverlap: false,
            label: {show: false},
            labelLine: {show: false},
            data: [
              {value: this.info.used, name: '已用'},
              {value: this.info.free, name: '闲置'},
            ],
          }
        ]
      };

    })
  }


}
