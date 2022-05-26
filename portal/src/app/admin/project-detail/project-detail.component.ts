import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzModalService} from "ng-zorro-antd/modal";

@Component({
  selector: 'app-project-detail',
  templateUrl: './project-detail.component.html',
  styleUrls: ['./project-detail.component.scss']
})
export class ProjectDetailComponent implements OnInit {
  id: any = '';
  data: any = {};
  loading = false;

  constructor(private router: ActivatedRoute, private rs: RequestService, private ms: NzModalService) {
    this.id = router.snapshot.params['id'];
    this.load();
  }

  ngOnInit(): void {
  }

  load(): void {
    this.loading = true;
    this.rs.get(`project/${this.id}`).subscribe(res=>{
      this.data = res.data;
      this.loading = false;
    });
  }

  exec(cmd: any) {
    this.rs.post(`project/${this.id}/execute`, {
      command: cmd.name
    }).subscribe(res=>{

    })
  }


  onEnableChange($event: any) {
    if (!$event) {
      this.rs.get(`project/${this.id}/enable`).subscribe(res => {
      });
      return;
    }
    this.ms.confirm({
      nzTitle: "提示",
      nzContent: "确认禁用吗?", //TODO 更丰富、人性 的 提醒
      nzOnOk: () => {
        this.rs.get(`project/${this.id}/disable`).subscribe(res => {
        });
      },
      nzOnCancel: () => {
        this.data.disabled = false;
      }
    })
  }
}
