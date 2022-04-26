import {Component, forwardRef, OnInit} from '@angular/core';
import {ControlValueAccessor, FormArray, FormBuilder, FormGroup, NG_VALUE_ACCESSOR, Validators} from "@angular/forms";

@Component({
  selector: 'app-edit-alarms',
  templateUrl: './edit-alarms.component.html',
  styleUrls: ['./edit-alarms.component.scss'],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => EditAlarmsComponent),
      multi: true
    }
  ]
})
export class EditAlarmsComponent implements OnInit, ControlValueAccessor {
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
          name: [d.name, [Validators.required]],
          condition: [d.condition, [Validators.required]],
          delay: [d.delay, [Validators.required]],
          reset_timeout: [d.reset_timeout, []],
          reset_total: [d.reset_total, []],
          code: [d.code, [Validators.required]],
          message: [d.message, [Validators.required]],
          level: [d.level, [Validators.required]],
          disabled: [d.disabled, [Validators.required]],
        })
      }))
    })
  }

  add() {
    this.formArray.push(this.fb.group({
      name: ['', [Validators.required]],
      condition: ['', [Validators.required]],
      delay: [0, [Validators.required]],
      reset_timeout: [0, []],
      reset_total: [0, []],
          code: ['', [Validators.required]],
          message: ['', [Validators.required]],
          level: [0, [Validators.required]],
      disabled: [false, [Validators.required]],
    }))
    //复制controls，让表格可以刷新
    this.formArray.controls = [...this.formArray.controls];
    this.change();
  }

  copy(i: number) {
    const group = this.formArray.controls[i];

    this.formArray.controls.splice(i, 0, this.fb.group({
      name: [group.get('name')?.value, [Validators.required]],
      condition: [group.get('condition')?.value, [Validators.required]],
      delay: [group.get('delay')?.value, [Validators.required]],
      reset_timeout: [group.get('reset_timeout')?.value, []],
      reset_total: [group.get('reset_total')?.value, []],
      code: [group.get('code')?.value, [Validators.required]],
      message: [group.get('message')?.value, [Validators.required]],
      level: [group.get('level')?.value, [Validators.required]],
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
