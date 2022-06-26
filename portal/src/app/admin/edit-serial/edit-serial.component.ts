import {Component, forwardRef, OnInit} from '@angular/core';
import {ControlValueAccessor, FormBuilder, FormGroup, NG_VALUE_ACCESSOR, Validators} from "@angular/forms";
import {RequestService} from "../../request.service";

@Component({
  selector: 'app-edit-serial',
  templateUrl: './edit-serial.component.html',
  styleUrls: ['./edit-serial.component.scss'],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => EditSerialComponent),
      multi: true
    }
  ]
})
export class EditSerialComponent implements OnInit, ControlValueAccessor {
  onChanged: any = () => {
  }
  onTouched: any = () => {
  }

  serials: Array<string> = []
  data: any = {};
  formGroup = new FormGroup({});

  constructor(private fb: FormBuilder, private rs: RequestService) {
  }

  ngOnInit(): void {

    this.rs.get('system/serials').subscribe(res => {
      this.serials = res.data;
    })
    this.buildForm();
  }

  buildForm(): void {
    this.formGroup = this.fb.group({
      port: [this.data.port, []],
      baud_rate: [this.data.baud_rate, []],
      data_bits: [this.data.data_bits, []],
      stop_bits: [this.data.stop_bits, []],
      parity: [this.data.parity, []],
      //rs485: [this.data.rs485, []],
    })
  }


  change() {
    this.formGroup.markAsDirty();
    this.formGroup.updateValueAndValidity();
    this.onChanged(this.formGroup.value);
    this.onTouched();
  }

  registerOnChange(fn: any): void {
    this.onChanged = fn;
  }

  registerOnTouched(fn: any): void {
    this.onTouched = fn;
  }

  writeValue(obj: any): void {
    this.data = obj || {};
    this.buildForm();
  }

}
