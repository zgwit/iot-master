import {Component, Input, OnDestroy, OnInit} from '@angular/core';
import {RequestService} from "../../request.service";

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
    })
  }


}
