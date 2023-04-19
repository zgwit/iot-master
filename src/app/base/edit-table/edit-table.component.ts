import { Component, OnInit, Input, EventEmitter, forwardRef, OnChanges, SimpleChanges } from '@angular/core';
import { FormBuilder, FormArray, ControlValueAccessor, NG_VALUE_ACCESSOR } from "@angular/forms";
import { NzMessageService } from "ng-zorro-antd/message";
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

@Component({
  selector: 'app-edit-table',
  templateUrl: './edit-table.component.html',
  styleUrls: ['./edit-table.component.scss'],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => EditTableComponent),
      multi: true
    }
  ]
})
export class EditTableComponent implements OnInit, ControlValueAccessor {
  group!: any;
  itemObj: object = {};
  constListData: any = [];
  tableData: Array<object> = [];
  onChanged: any = () => { }
  onTouched: any = () => { }

  @Input() data: any = {};
  @Input()
  set listData(dt: Array<{ title: string, type?: any, keyName: string }>) {
    const itemObj: any = {};
    for (let index = 0; index < dt.length; index++) {
      const { keyName } = dt[index];
      itemObj[keyName] = itemObj.defaultValue || '';
    }
    this.itemObj = itemObj;
    this.constListData = dt;
  };
  constructor(
    private msg: NzMessageService,
    private fb: FormBuilder
  ) { }
  ngOnInit(): void {
    this.build();
  }
  writeValue(data: any): void {
    const itemObj = JSON.parse(JSON.stringify(this.itemObj));
    this.tableData = data;
    if (data && data.length) {
      data.forEach((item: any) => {
        const newGroup = this.fb.group(Object.assign(itemObj, item));
        this.aliases.push(newGroup);
      });
    }
  }
  registerOnChange(fn: any): void {
    this.onChanged = fn;
  }
  registerOnTouched(fn: any): void {
    this.onTouched = fn;
  }

  build() {
    this.group = this.fb.group({ keyName: this.fb.array([]) });
  }

  change() {
    this.onChanged(this.tableData);
    this.onTouched();
  }
  handleCopyProperTy(index: number) {
    const oitem = this.group.get('keyName').controls[index].value;
    this.aliases.insert(index, this.fb.group(oitem));
    this.msg.success("复制成功");
  }
  propertyDel(i: number) {
    this.group.get('keyName').controls.splice(i, 1)
  }
  get aliases(): FormArray {
    return this.group.get('keyName') as FormArray;
  }
  drop(event: CdkDragDrop<string[]>): void {
    moveItemInArray(this.group.get('keyName').controls, event.previousIndex, event.currentIndex);
  }
}