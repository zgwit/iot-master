import { Component, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-search-form',
  templateUrl: './search-form.component.html',
  styleUrls: ['./search-form.component.scss']
})
export class SearchFormComponent {
  @Input() searchText = "搜索"
  @Input() placeholder = "关键字";
  @Output() onSearch = new EventEmitter<string>();
  text = "";
  handleClear() {
    this.text = "";
    this.onSearch.emit('');
  }
}
