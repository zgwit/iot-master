import { Component, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-table-oper',
  templateUrl: './table-oper.component.html',
  styleUrls: ['./table-oper.component.scss']
})
export class TableOperComponent {
  @Input() data: any;
  @Output() edit = new EventEmitter<string>();
  @Output() delete = new EventEmitter<number>();
  cancel() { }
}
