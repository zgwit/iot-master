import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {RequestService} from "../../request.service";

@Component({
  selector: 'app-device-detail',
  templateUrl: './device-detail.component.html',
  styleUrls: ['./device-detail.component.scss']
})
export class DeviceDetailComponent implements OnInit {
  id: any = '';
  data: any = {};
  context: any = {};
  loading = false;

  constructor(private router: ActivatedRoute, private rs: RequestService) {
    this.id = router.snapshot.params['id'];
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

  loadContext(){
    this.rs.get(`device/${this.id}/context`).subscribe(res => {
      this.context = res.data;
    });
  }

  exec(cmd: any) {
    let params = [];
    if (cmd.argc > 0) {
      let param = prompt("请输入值，以逗号间隔", "")
      if (!param) return;
      params = eval(`[${param}]`)
    }
    this.rs.post(`device/${this.id}/execute`, {
      command: cmd.name,
      parameters: params,
    }).subscribe(res => {

    })
  }

  disabled($event: any) {

  }

  refresh(name: any) {
    this.rs.get(`device/${this.id}/refresh/${name}`).subscribe(res => {
      this.data.values[name] = res.data;
    })
  }

  refreshAll() {
    this.rs.get(`device/${this.id}/refresh`).subscribe(res => {
      //Object.assign(this.data.values, res.data);
      this.loadContext()
    })
  }
}
