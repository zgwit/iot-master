import {Component, Input, forwardRef, OnInit} from '@angular/core';
import {ControlValueAccessor, NG_VALUE_ACCESSOR} from "@angular/forms";
import {NzSizeLDSType} from "ng-zorro-antd/core/types";

@Component({
  selector: 'app-input-script',
  templateUrl: './input-script.component.html',
  styleUrls: ['./input-script.component.scss'],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => InputScriptComponent),
      multi: true
    }
  ]
})
export class InputScriptComponent implements OnInit, ControlValueAccessor {
  onChanged: any = () => {}
  onTouched: any = () => {}

  //内容
  _js = "";
  get js() {
    return this._js
  }
  set js(y) {
    console.log('js page-editor', y)
    this._js = y;

    this.onChanged(y);
    this.onTouched();
  }
  
  isVisible = false;
  
  @Input()
  nzSize: NzSizeLDSType = "default";

  constructor() {
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
    this._js = obj;
  }

}
