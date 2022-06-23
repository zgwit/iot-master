import {Component, OnInit} from '@angular/core';
import {NzTableQueryParams} from "ng-zorro-antd/table";
import {Router} from "@angular/router";
import {RequestService} from "../../request.service";
import {NzModalService} from "ng-zorro-antd/modal";
import {parseTableQuery} from "../table";

@Component({
  selector: 'app-camera',
  templateUrl: './camera.component.html',
  styleUrls: ['./camera.component.scss']
})
export class CameraComponent implements OnInit {
  datum: any[] = [];

  loading = false;
  total = 1;
  pageSize = 20;
  pageIndex = 1;

  params: any = {filter: {}};
  view: string = 'card';

  constructor(private router: Router, private rs: RequestService, private ms: NzModalService) {

  }

  ngOnInit(): void {
    this.load();
  }

  search(keyword: string) {
    this.pageIndex = 1;
    this.params.skip = 0;
    if (keyword)
      this.params.keyword = {name: keyword, addr: keyword};
    else
      delete this.params.keyword;
    this.load();
  }

  onQuery(params: NzTableQueryParams) {
    parseTableQuery(params, this.params);
    this.load();
  }

  onPageChange(page: number) {
    this.pageIndex = page;
    this.params.skip = (this.pageIndex - 1) * this.pageSize;
    this.load()
  }

  load(): void {
    this.loading = true;
    this.rs.post('camera/list', this.params).subscribe(res => {
      console.log('res', res);
      this.datum = res.data;
      this.total = res.total;
    }).add(() => {
      this.loading = false;
    });
  }

  create(): void {
    this.router.navigate(["admin/camera/create"]);
  }

  open(data: any): void {
    this.router.navigate(['/admin/camera/detail/' + data.id]);
  }

  remove(data: any, i: number) {
    this.rs.get(`camera/${data.id}/delete`).subscribe(res => {
      this.datum.splice(i, 1);
    });
  }

  enable(data: any) {
    this.rs.get(`camera/${data.id}/enable`).subscribe(res => {
      data.disabled = false
    });
  }

  disable(data: any) {
    this.rs.get(`camera/${data.id}/disable`).subscribe(res => {
      data.disabled = true
    });
  }
}
