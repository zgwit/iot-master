import {Component, forwardRef, Input, OnInit} from '@angular/core';
import {ControlValueAccessor, FormArray, FormBuilder, FormGroup, NG_VALUE_ACCESSOR, Validators} from "@angular/forms";

@Component({
  selector: 'app-edit-validators',
  templateUrl: './edit-validators.component.html',
  styleUrls: ['./edit-validators.component.scss'],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => EditValidatorsComponent),
      multi: true
    }
  ]
})
export class EditValidatorsComponent implements OnInit, ControlValueAccessor {
  onChanged: any = () => {
  }
  onTouched: any = () => {
  }

  @Input() points: any;

  items: any[] = [];

  formGroup = new FormGroup({});

  current: any = {};
  showModal = false;


  compareFun = (o1: any | string, o2: any): boolean => {
    if (o1) {
      return typeof o1 === 'string' ? o1 === o2.name : o1.name === o2.name;
    } else {
      return false;
    }
  }

  constructor(private fb: FormBuilder) {
  }

  ngOnInit(): void {
    this.buildForm({});
  }

  buildForm(d: any): void {
    this.formGroup = this.fb.group({
      name: [d.name, [Validators.required]],
      compare: [d.compare, [Validators.required]],
      value: [d.value, [Validators.required]],
      value1: [d.value1, [Validators.required]],
      value2: [d.value2, [Validators.required]],
      delay: [d.delay, [Validators.required]],
      reset_timeout: [d.reset_timeout, []],
      reset_total: [d.reset_total, []],
      code: [d.code, [Validators.required]],
      message: [d.message, [Validators.required]],
      level: [d.level, [Validators.required]],
      disabled: [d.disabled, [Validators.required]],
    })
  }

  copy(i: number) {
    let item = this.items[i]
    item = JSON.parse(JSON.stringify(item))
    this.items.splice(i + 1, 0, item)
  }

  remove(i: number) {
    this.items.splice(i, 1)
    this.change();
  }

  clear() {
    this.items = [];
    this.change();
  }

  change() {
    //this.formGroup.markAsDirty();
    //this.formGroup.updateValueAndValidity();
    this.onChanged(this.items);
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
    //this.buildForm();
  }

  drop($event: any) {
    const item = this.items.splice($event.previousIndex, 1);
    this.items.splice($event.currentIndex, 0, ...item);
    this.change();
  }

  edit(data?: any) {
    if (!data) {
      data = {
        name: '',
        compare: '=',
        value1: 0,
        value2: 0,
        delay: 0,
        reset_timeout: 0,
        reset_total: 0,
        code: '',
        message: '',
        level: 0,
        disabled: false,
        as: '',
        expression: '',
      }
      this.items.push(data)
    }
    this.current = data;
    this.buildForm(data)
    this.showModal = true;
  }

  onOk() {
    this.showModal = false;
    Object.assign(this.current, this.formGroup.value)
  }
}
