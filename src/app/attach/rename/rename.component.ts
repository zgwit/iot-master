import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-rename',
  templateUrl: './rename.component.html',
  styleUrls: ['./rename.component.scss']
})
export class RenameComponent {
  @Input() currentName: string = '';
  name: string = ''
}
