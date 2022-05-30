import {Component, EventEmitter, forwardRef, HostBinding, Input, OnInit, Output} from '@angular/core';
import {ControlValueAccessor, NG_VALUE_ACCESSOR} from "@angular/forms";
import {ChooseService} from "../choose.service";
import {RequestService} from "../../request.service";

@Component({
  selector: 'app-choose-device',
  templateUrl: './choose-device.component.html',
  styleUrls: ['./choose-device.component.scss'],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => ChooseDeviceComponent),
      multi: true
    }
  ]
})
export class ChooseDeviceComponent implements OnInit, ControlValueAccessor {
  onChanged: any = () => {}
  onTouched: any = () => {}

  //内容
  @HostBinding('attr.title')
  id = "";
  name = "";

  @Input()
  showClear: any = false;

  @Output()
  device: EventEmitter<any> = new EventEmitter<any>()

  constructor(private cs: ChooseService, private rs: RequestService) { }

  ngOnInit(): void {
  }

  registerOnChange(fn: any): void {
    this.onChanged = fn;
  }

  registerOnTouched(fn: any): void {
    this.onTouched = fn;
  }

  writeValue(obj: any): void {
    this.id = obj;
    this.load();
  }

  load() {
    if (!this.id) return;
    this.name = "加载中...";
    this.rs.get(`device/${this.id}`).subscribe(res=>{
      this.name = res.data.name;
      this.device.emit(res.data)
      if (!this.name)
        this.loadProduct(res.data.product_id)
    })
  }

  loadProduct(id: string) {
    this.rs.get(`product/${id}/detail`).subscribe(res=>{
      this.name = res.data.name;
    })
  }

  choose() {
    this.cs.chooseDevice().subscribe(res=>{
      if (res){
        this.id = res;
        this.load();
        this.onChanged(res);
        this.onTouched();
      }
    })
  }

  clear() {
    this.id = '';
    this.name = '';
    this.onChanged('');
    this.onTouched();
  }
}
