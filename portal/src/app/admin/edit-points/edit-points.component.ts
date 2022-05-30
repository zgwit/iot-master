import {Component, forwardRef, Input, OnInit} from '@angular/core';
import {ControlValueAccessor, FormArray, FormBuilder, FormGroup, NG_VALUE_ACCESSOR, Validators} from "@angular/forms";

@Component({
  selector: 'app-edit-points',
  templateUrl: './edit-points.component.html',
  styleUrls: ['./edit-points.component.scss'],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => EditPointsComponent),
      multi: true
    }
  ]
})
export class EditPointsComponent implements OnInit, ControlValueAccessor {
  @Input() codes: any = [];

  onChanged: any = () => {
  }
  onTouched: any = () => {
  }

  items: any[] = [];
  formGroup = new FormGroup({});
  formArray: FormArray = new FormArray([]);

  constructor(private fb: FormBuilder) {
  }

  ngOnInit(): void {
    this.buildForm();
  }

  buildForm(): void {
    this.formGroup = this.fb.group({
      items: this.formArray = this.fb.array(this.items.map((d: any) => {
        return this.fb.group({
          name: [d.name, [Validators.required]],
          label: [d.label, []],
          code: [d.code, [Validators.required]],
          address: [d.address, [Validators.required]],
          type: [d.type, [Validators.required]],
          le: [d.le, [Validators.required]],
          precision: [d.precision, [Validators.required]],
          store: [d.store, [Validators.required]],
          unit: [d.unit, []],
        })
      }))
    })
  }

  add() {
    this.formArray.push(this.fb.group({
      name: ['', [Validators.required]],
      label: ['', []],
      code: ['', [Validators.required]],
      address: ['', [Validators.required]],
      type: ['', [Validators.required]],
      le: [false, [Validators.required]],
      precision: [0, [Validators.required]],
      store: [true, [Validators.required]],
      unit: ["", []],
    }))
    //复制controls，让表格可以刷新
    this.formArray.controls = [...this.formArray.controls];
    this.change();
  }

  copy(i: number) {
    const group = this.formArray.controls[i];

    this.formArray.controls.splice(i, 0, this.fb.group({
      name: [group.get('name')?.value, [Validators.required]],
      label: [group.get('label')?.value, []],
      code: [group.get('code')?.value, []],
      address: [group.get('address')?.value, [Validators.required]],
      type: [group.get('type')?.value, [Validators.required]],
      le: [group.get('le')?.value, [Validators.required]],
      precision: [group.get('precision')?.value, [Validators.required]],
      store: [group.get('store')?.value, [Validators.required]],
      unit: [group.get('unit')?.value, [Validators.required]],
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
