import { Component, forwardRef, HostBinding, Input, OnInit } from '@angular/core';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from "@angular/forms";
import { RequestService } from "../../request.service";
import { ProductsComponent } from "../products/products.component";
import { NzModalService } from "ng-zorro-antd/modal";

@Component({
  selector: 'app-product-select',
  templateUrl: './product-select.component.html',
  styleUrls: ['./product-select.component.scss'],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => ProductSelectComponent),
      multi: true
    }
  ]
})
export class ProductSelectComponent implements OnInit, ControlValueAccessor {
  onChanged: any = () => { }
  onTouched: any = () => { }
  //内容
  @HostBinding('attr.title')

  _id = "";
  name = "";

  @Input()
  showClear: any = false;

  constructor(private ms: NzModalService, private rs: RequestService) {
  }

  ngOnInit(): void {
  }

  registerOnChange(fn: any): void {
    this.onChanged = fn;
  }

  registerOnTouched(fn: any): void {
    this.onTouched = fn;
  }

  writeValue(obj: any): void {
    this._id = obj;
    this.load();
  }

  load() {
    if (!this._id) return;
    this.name = "加载中...";
    this.rs.get(`product/${this._id}`).subscribe(res => {
      this.name = res.data.name;
    })
  }

  choose() {
    this.ms.create({
      nzTitle: '选择产品',
      nzWidth: '700px',
      nzContent: ProductsComponent,
      nzComponentParams: {
        showAddBtn: false,
      },
      nzFooter: null,
    }).afterClose.subscribe((obj) => {
      const { id, name } = obj || {};
      this._id = id;
      this.name = name
      this.load();
      this.onChanged(id);
      this.onTouched();
    });
  }

  clear() {
    this._id = '';
    this.name = '';
    this.onChanged('');
    this.onTouched();
  }
}
