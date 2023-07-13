import { Component, Input, Optional } from '@angular/core';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { Router } from '@angular/router';
import { RequestService } from '../../request.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzTableQueryParams } from 'ng-zorro-antd/table';
import { ParseTableQuery } from '../../base/table';
import { DeviceEditComponent } from "../device-edit/device-edit.component"
import {
  isIncludeAdmin,
  readCsv,
  tableHeight,
  onAllChecked,
  onItemChecked,
  batchdel,
  refreshCheckedStatus,
} from '../../../public';

@Component({
  selector: 'app-devices',
  templateUrl: './devices.component.html',
  styleUrls: ['./devices.component.scss'],
})
export class DevicesComponent {
  @Input() chooseGateway = false;

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
    @Optional() protected ref: NzModalRef,
    private modal: NzModalService,
    private router: Router,
    private rs: RequestService,
    private msg: NzMessageService
  ) {
    //this.load();
  }

  reload() {
    this.datum = [];
    this.load();
  }
  disable(mess: number, id: any) {
    if (mess)
      this.rs.get(`device/${id}/disable`).subscribe((res) => {
        this.reload();
      });
    else
      this.rs.get(`device/${id}/enable`).subscribe((res) => {
        this.reload();
      });
  }
  load() {
    //筛选网关
    if (this.chooseGateway) this.query.filter = { type: 'gateway' };

    this.loading = true;
    this.rs
      .post('device/search', this.query)
      .subscribe((res) => {
        this.datum = res.data || [];
        this.datum.filter(
          (item) =>
          (item.disabled =
            item.disabled === undefined ? false : item.disabled)
        );
        this.total = res.total;
        this.setOfCheckedId.clear();
        refreshCheckedStatus(this);
      })
      .add(() => {
        this.loading = false;
      });
  }

  create() {
    let path = '/device/create';
    if (location.pathname.startsWith('/admin')) path = '/admin' + path;
    this.router.navigateByUrl(path);
  }

  delete(id: number, size?: number) {
    this.rs.get(`device/${id}/delete`).subscribe((res) => {
      if (!size) {
        this.msg.success('删除成功');
        this.datum = this.datum.filter((d) => d.id !== id);
        return
      }
      this.delResData.push(res);
      if (size === this.delResData.length) {
        this.msg.success('删除成功');
        this.delResData = [];
        this.load();
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

  open(id: any) {
    if (this.chooseGateway) {
      this.select(id);
      return;
    }
    const path = `${isIncludeAdmin()}/device/detail/${id}`;
    this.router.navigateByUrl(path);
  }
  handleEdit(id?: string) {
    const nzTitle = id ? "编辑设备" : "创建设备";
    const modal: NzModalRef = this.modal.create({
      nzTitle,
      nzStyle: { top: '20px' },
      nzContent: DeviceEditComponent,
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
  select(id: any) {
    this.ref && this.ref.close(id);
  }
  cancel() {
    this.msg.info('取消操作');
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
