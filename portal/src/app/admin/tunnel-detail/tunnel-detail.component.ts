import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzModalService} from "ng-zorro-antd/modal";

@Component({
  selector: 'app-tunnel-detail',
  templateUrl: './tunnel-detail.component.html',
  styleUrls: ['./tunnel-detail.component.scss']
})
export class TunnelDetailComponent implements OnInit {
  id = 0;
  data: any = {};
  loading = false;

  constructor(private router: ActivatedRoute, private rs: RequestService, private ms: NzModalService) {
    this.id = parseInt(router.snapshot.params['id']);
    this.load();
  }

  ngOnInit(): void {
  }

  load(): void {
    this.loading = true;
    this.rs.get(`tunnel/${this.id}`).subscribe(res=>{
      this.data = res.data;
      this.loading = false;
    });
    //TODO 监听
  }

  onEnableChange(disabled: boolean) {
    if (!disabled) {
      this.rs.get(`tunnel/${this.id}/enable`).subscribe(res => {
        this.load()
      });
      return;
    }
    this.ms.confirm({
      nzTitle: "提示",
      nzContent: "确认禁用吗?", //TODO 更丰富、人性 的 提醒
      nzOnOk: () => {
        this.rs.get(`tunnel/${this.id}/disable`).subscribe(res => {
          this.load()
        });
      },
      nzOnCancel: () => {
        this.data.disabled = false;
      }
    })
  }

  start() {
      this.rs.get(`tunnel/${this.id}/start`).subscribe(res => {
        this.load()
      });
  }

  stop() {
    this.rs.get(`tunnel/${this.id}/stop`).subscribe(res => {
      this.load()
    });
  }


}
