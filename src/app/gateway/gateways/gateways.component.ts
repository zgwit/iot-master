import { Component } from '@angular/core';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { Router } from '@angular/router';
import { RequestService } from '../../request.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzTableQueryParams } from 'ng-zorro-antd/table';
import { ParseTableQuery } from '../../base/table';
import { GatewayEditComponent } from "../gateway-edit/gateway-edit.component";
import {
  tableHeight,
  onAllChecked,
  onItemChecked,
  batchdel,
  refreshCheckedStatus,
} from '../../../public';

@Component({
  selector: 'app-gateways',
  templateUrl: './gateways.component.html',
  styleUrls: ['./gateways.component.scss'],
})
export class GatewaysComponent {
  loading = true;
  datum: any[] = [];
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {};
  checked = false;
  indeterminate = false;
  setOfCheckedId = new Set<number>();
  delResData: any = [];

  constructor(
    private modal: NzModalService,
    private router: Router,
    private rs: RequestService,
    private msg: NzMessageService
  ) { }

  reload() {
    this.datum = [];
    this.load();
  }
  disable(mess: number, id: any) {
    if (mess)
      this.rs.get(`gateway/${id}/disable`).subscribe((res) => {
        this.reload();
      });
    else
      this.rs.get(`gateway/${id}/enable`).subscribe((res) => {
        this.reload();
      });
  }
  load() {
    this.loading = true;
    this.rs
      .post('gateway/search', this.query)
      .subscribe((res) => {
        // console.log(res);
        this.datum = res.data || [];
        this.total = res.total;
        this.setOfCheckedId.clear();
        refreshCheckedStatus(this);
      })
      .add(() => {
        this.loading = false;
      });
  }

  delete(id: number, size?: number) {
    this.rs.get(`gateway/${id}/delete`).subscribe((res) => {
      if (!size) {
        this.msg.success('删除成功');
        this.datum = this.datum.filter((d) => d.id !== id);
      } else if (size) {
        this.delResData.push(res);
        if (size === this.delResData.length) {
          this.msg.success('删除成功');
          this.load();
        }
      }
    });
  }

  onQuery($event: NzTableQueryParams) {
    ParseTableQuery($event, this.query);
    this.load();
  }
  pageIndexChange(pageIndex: number) {
    this.query.skip = pageIndex - 1;
  }
  pageSizeChange(pageSize: number) {
    this.query.limit = pageSize;
  }
  search($event: string) {
    this.query.keyword = {
      name: $event,
    };
    this.query.skip = 0;
    this.load();
  }

  cancel() {
    this.msg.info('取消操作');
  }

  handleEdit(id?: string) {
    const nzTitle = id ? "编辑网关" : "创建网关";
    const modal: NzModalRef = this.modal.create({
      nzTitle,
      nzStyle: { top: '20px' },
      nzContent: GatewayEditComponent,
      nzComponentParams: { id },
      nzMaskClosable: false,
      nzFooter: [
        {
          label: '取消',
          onClick: () => {
            modal.destroy();
          }
        },
        {
          label: '保存',
          type: 'primary',
          onClick: componentInstance => {
            componentInstance!.submit().then(() => {
              modal.destroy();
              this.load();
            }, () => { })
          }
        }
      ]
    });
  }
  getTableHeight() {
    return tableHeight(this);
  }
  handleBatchDel() {
    batchdel(this);
  }
  handleAllChecked(id: any) {
    onAllChecked(id, this);
  }
  handleItemChecked(id: number, checked: boolean) {
    onItemChecked(id, checked, this);
  }
}
