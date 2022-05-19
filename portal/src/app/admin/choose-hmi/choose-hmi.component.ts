import {Component, forwardRef, Input, HostBinding, OnInit} from '@angular/core';
import {ControlValueAccessor, NG_VALUE_ACCESSOR} from "@angular/forms";
import {ChooseService} from "../choose.service";
import {RequestService} from "../../request.service";

@Component({
  selector: 'app-choose-hmi',
  templateUrl: './choose-hmi.component.html',
  styleUrls: ['./choose-hmi.component.scss'],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => ChooseHmiComponent),
      multi: true
    }
  ]
})
export class ChooseHmiComponent implements OnInit, ControlValueAccessor {
  onChanged: any = () => {}
  onTouched: any = () => {}

  //内容
  @HostBinding('attr.title')
  id = "";
  name = "";

  @Input()
  showClear: any = false;

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
    this.rs.get(`hmi/${this.id}`).subscribe(res=>{
      this.name = res.data.name;
    })
  }

  choose() {
    this.cs.chooseHMI().subscribe(res=>{
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
