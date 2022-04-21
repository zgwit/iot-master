import {Component, forwardRef, OnInit} from '@angular/core';
import {ControlValueAccessor, FormArray, FormBuilder, FormGroup, NG_VALUE_ACCESSOR, Validators} from "@angular/forms";

@Component({
  selector: 'app-edit-directives',
  templateUrl: './edit-directives.component.html',
  styleUrls: ['./edit-directives.component.scss'],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => EditDirectivesComponent),
      multi: true
    }
  ]
})
export class EditDirectivesComponent implements OnInit, ControlValueAccessor {
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
          address: [d.address, [Validators.required]],
          value: [d.value, [Validators.required]],
          delay: [d.delay, [Validators.required]],
          expression: [d.expression, [Validators.required]],
        })
      }))
    })
  }

  add() {
    this.formArray.push(this.fb.group({
      address: ['', [Validators.required]],
      value: [0, [Validators.required]],
      delay: [0, [Validators.required]],
      expression: ['', [Validators.required]],
    }))
    //复制controls，让表格可以刷新
    this.formArray.controls = [...this.formArray.controls];
    this.change();
  }

  copy(i: number) {
    const group = this.formArray.controls[i];

    this.formArray.controls.splice(i, 0, this.fb.group({
      address: [group.get('address')?.value, [Validators.required]],
      value: [group.get('value')?.value, [Validators.required]],
      delay: [group.get('delay')?.value, [Validators.required]],
      expression: [group.get('expression')?.value, [Validators.required]],
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
