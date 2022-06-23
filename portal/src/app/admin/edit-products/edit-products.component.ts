import {Component, forwardRef, OnInit} from '@angular/core';
import {ControlValueAccessor, FormArray, FormBuilder, FormGroup, NG_VALUE_ACCESSOR, Validators} from "@angular/forms";
import {ChooseService} from "../choose.service";

@Component({
  selector: 'app-edit-products',
  templateUrl: './edit-products.component.html',
  styleUrls: ['./edit-products.component.scss'],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => EditProductsComponent),
      multi: true
    }
  ]
})
export class EditProductsComponent implements OnInit, ControlValueAccessor {
  onChanged: any = () => {
  }
  onTouched: any = () => {
  }

  items: any[] = [];

  formGroup = new FormGroup({});

  current: any = {};
  showModal = false;

  constructor(private fb: FormBuilder) {
  }

  ngOnInit(): void {
    this.buildForm({});
  }

  buildForm(d:any): void {
    this.formGroup = this.fb.group({
          id: [d.id, [Validators.required]],
          name: [d.name, [Validators.required]],
    })
  }

  // addMore() {
  //   this.cs.chooseProduct({multiple: true}).subscribe(devices=>{
  //     if (devices.length) {
  //       devices.forEach((d: string)=>{
  //         this.formArray.push(this.fb.group({
  //           id: [d, [Validators.required]],
  //           name: ['', [Validators.required]],
  //         }))
  //       });
  //
  //       //复制controls，让表格可以刷新
  //       this.formArray.controls = [...this.formArray.controls];
  //       this.change();
  //     }
  //   })
  // }

  copy(i: number) {
    let item = this.items[i]
    item = JSON.parse(JSON.stringify(item))
    this.items.splice(i+1, 0, item)
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
        id: '',
        name: '',
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
