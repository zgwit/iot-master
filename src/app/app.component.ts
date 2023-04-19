import {Component} from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  constructor() {
    setTimeout(function () {
      let loadElem = document.querySelector('.preloader');
      if (loadElem)
        loadElem.remove()
      //loadElem.className = 'preloader-hidden';
    }, 500)
  }
}
