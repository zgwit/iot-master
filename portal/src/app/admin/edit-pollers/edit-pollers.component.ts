import {Component, forwardRef, Input, OnInit} from '@angular/core';
import {ControlValueAccessor, FormArray, FormBuilder, FormGroup, NG_VALUE_ACCESSOR, Validators} from "@angular/forms";

@Component({
  selector: 'app-edit-pollers',
  templateUrl: './edit-pollers.component.html',
  styleUrls: ['./edit-pollers.component.scss'],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => EditPollersComponent),
      multi: true
    }
  ]
})
export class EditPollersComponent implements OnInit, ControlValueAccessor {
  @Input() codes: any = [];

  onChanged: any = () => {
  }
  onTouched: any = () => {
  }

  items: any[] = [];
  formGroup = new FormGroup({});
  formArray: FormArray = new FormArray([]);

  constructor(private fb: FormBuilder) { }

  ngOnInit(): void {
    this.buildForm();
  }

  buildForm(): void{
    this.formGroup = this.fb.group({
      items: this.formArray = this.fb.array(this.items.map((d: any) => {
        return this.fb.group({
          type: [d.type, [Validators.required]],
          interval: [d.interval, [Validators.required]],
          clock: [d.clock, [Validators.required]],
          crontab: [d.crontab, [Validators.required]],
          code: [d.code, [Validators.required]],
          address: [d.address, [Validators.required]],
          length: [d.length, [Validators.required]],
          disabled: [d.disabled, [Validators.required]],
        })
      }))
    })
  }

  add() {
    this.formArray.push(this.fb.group({
          type: ["interval", [Validators.required]],
          interval: [1000, [Validators.required]],
          clock: [60, [Validators.required]],
          crontab: ['', [Validators.required]],
          code: ['', [Validators.required]],
          address: ['0', [Validators.required]],
          length: [1, [Validators.required]],
          disabled: [false, [Validators.required]],
    }))
    //复制controls，让表格可以刷新
    this.formArray.controls = [...this.formArray.controls];
    this.change();
  }

  copy(i: number) {
    const group = this.formArray.controls[i];

    this.formArray.controls.splice(i, 0, this.fb.group({
      type: [group.get('type')?.value, [Validators.required]],
      interval: [group.get('interval')?.value, []],
      clock: [group.get('clock')?.value, []],
      crontab: [group.get('crontab')?.value, []],
      code: [group.get('code')?.value, []],
      address: [group.get('address')?.value, [Validators.required]],
      length: [group.get('length')?.value, [Validators.required]],
      disabled: [group.get('disabled')?.value, [Validators.required]],
    }))
  }

  remove(i: number) {
    this.formArray.removeAt(i)
    this.change();
  }

  clear() {
    this.formArray.clear();
    this.change();
  }

  change() {
    this.formArray.markAsDirty();
    this.formArray.updateValueAndValidity();
    this.onChanged(this.formArray.value);
    this.onTouched();
  }

  registerOnChange(fn: any): void {
    this.onChanged = fn;
  }

  registerOnTouched(fn: any): void {
    this.onTouched = fn;
  }

  writeValue(obj: any): void {
    this.items = obj;
    this.buildForm();
  }

  drop($event: any) {
    const item = this.formArray.controls.splice($event.previousIndex, 1);
    this.formArray.controls.splice($event.currentIndex, 0, ...item);
    this.change();
  }
}
