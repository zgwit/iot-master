import {Component, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzModalService} from "ng-zorro-antd/modal";
import {NzTableQueryParams} from "ng-zorro-antd/table";
import {parseTableQuery} from "../table";

@Component({
  selector: 'app-tunnel',
  templateUrl: './tunnel.component.html',
  styleUrls: ['./tunnel.component.scss']
})
export class TunnelComponent implements OnInit {
  datum: any[] = [];

  loading = false;
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  filterGender = [
    {text: 'TCP服务器', value: 'tcp-server'},
    {text: 'TCP客户端', value: 'tcp-client'}
  ];

  params: any = {filter: {}};

  constructor(private router: Router, private rs: RequestService, private ms: NzModalService) {

  }

  ngOnInit(): void {
    //this.load();
  }

  search(keyword: string) {
    this.pageIndex = 1;
    this.params.skip = 0;
    if (keyword)
      this.params.keyword = {Name: keyword, Type: keyword, Address: keyword};
    else
      delete this.params.keyword;
    this.load();
  }

  onQuery(params: NzTableQueryParams) {
    parseTableQuery(params, this.params);
    this.load();
  }

  load(): void {
    this.loading = true;
    this.rs.post('tunnel/list', this.params).subscribe(res => {
      console.log('res', res);
      this.datum = res.data;
      this.total = res.total;
    }).add(() => {
      this.loading = false;
    });
  }

  create(): void {
    this.router.navigate(["admin/tunnel/create"]);
  }

  open(data: any): void {
    this.router.navigate(['/admin/tunnel/detail/' + data.id]);
  }

  remove(data: any, i: number) {
    this.rs.delete(`tunnel/${data.id}/delete`).subscribe(res => {
      this.datum.splice(i, 1);
    });
  }

  onEnableChange(data: any, enable: boolean) {
    if (enable) {
      this.rs.post(`tunnel/${data.id}/setting`, {enable}).subscribe(res => {
      });
      return;
    }
    this.ms.confirm({
      nzTitle: "提示",
      nzContent: "确认禁用吗?", //TODO 更丰富、人性 的 提醒
      nzOnOk: () => {
        this.rs.post(`tunnel/${data.id}/setting`, {enable}).subscribe(res => {
        });
      },
      nzOnCancel: () => {
        data.enable = true;
      }
    })
  }

}
