import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzModalService} from "ng-zorro-antd/modal";
import {ChooseService} from "../choose.service";

@Component({
  selector: 'app-device-detail',
  templateUrl: './device-detail.component.html',
  styleUrls: ['./device-detail.component.scss']
})
export class DeviceDetailComponent implements OnInit {
  id = 0;
  data: any = {};
  context: any = {};
  loading = false;

  constructor(private router: ActivatedRoute,
              private rs: RequestService,
              private ms: NzModalService,
              private cs: ChooseService) {
    this.id = parseInt(router.snapshot.params['id']);
    this.load();
  }

  ngOnInit(): void {
  }

  load(): void {
    this.loading = true;
    this.rs.get(`device/${this.id}`).subscribe(res => {
      this.data = res.data;
      this.loading = false;
    });
    this.loadContext()
  }

  loadContext() {
    this.rs.get(`device/${this.id}/context`).subscribe(res => {
      this.context = res.data;
    });
  }

  exec(cmd: any) {
    if (!cmd.argument) {
      this.rs.post(`device/${this.id}/execute`, {command: cmd.name, arguments: [],}).subscribe(res => {
      })
      return
    }

    this.cs.prompt({
      message: "输入控制命令的参数",
      placeholder: "以逗号间隔"
    }, cmd.name).subscribe(res => {
      let params = res ? eval('[' + res + ']') : []
      this.rs.post(`device/${this.id}/execute`, {command: cmd.name, arguments: params,}).subscribe(res => {

      })
    })
  }


  onEnableChange($event: any) {
    if (!$event) {
      this.rs.get(`device/${this.id}/enable`).subscribe(res => {
      });
      return;
    }
    this.ms.confirm({
      nzTitle: "提示",
      nzContent: "确认禁用吗?", //TODO 更丰富、人性 的 提醒
      nzOnOk: () => {
        this.rs.get(`device/${this.id}/disable`).subscribe(res => {
        });
      },
      nzOnCancel: () => {
        this.data.disabled = false;
      }
    })
  }

  refresh(name: any) {
    this.rs.get(`device/${this.id}/refresh/${name}`).subscribe(res => {
      this.context[name] = res.data;
    })
  }

  refreshAll() {
    this.rs.get(`device/${this.id}/refresh`).subscribe(res => {
      //Object.assign(this.data.values, res.data);
      this.loadContext()
    })
  }
}
