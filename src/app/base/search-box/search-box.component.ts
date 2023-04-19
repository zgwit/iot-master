import {Component, EventEmitter, Input, Output} from '@angular/core';

@Component({
  selector: 'app-search-box',
  templateUrl: './search-box.component.html',
  styleUrls: ['./search-box.component.scss']
})
export class SearchBoxComponent {
  @Input() searchText = "搜索"
  @Input() placeholder = "关键字";
  @Input() text = "";
  @Output() onSearch = new EventEmitter<string>();


}
