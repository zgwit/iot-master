import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzModalService} from "ng-zorro-antd/modal";

@Component({
  selector: 'app-user-detail',
  templateUrl: './user-detail.component.html',
  styleUrls: ['./user-detail.component.scss']
})
export class UserDetailComponent implements OnInit {
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
    this.rs.get(`user/${this.id}`).subscribe(res => {
      this.data = res.data;
      this.loading = false;
    });
  }

  onEnableChange($event: any) {
    if (!$event) {
      this.rs.get(`user/${this.id}/enable`).subscribe(res => {
      });
      return;
    }
    this.ms.confirm({
      nzTitle: "提示",
      nzContent: "确认禁用吗?", //TODO 更丰富、人性 的 提醒
      nzOnOk: () => {
        this.rs.get(`user/${this.id}/disable`).subscribe(res => {
        });
      },
      nzOnCancel: () => {
        this.data.disabled = false;
      }
    })
  }

}
