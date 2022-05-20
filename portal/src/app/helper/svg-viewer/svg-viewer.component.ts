import {AfterViewInit, Component, Input, OnInit, ViewChild} from '@angular/core';
import {Svg} from '@svgdotjs/svg.js';

@Component({
  selector: 'app-svg-viewer',
  templateUrl: './svg-viewer.component.html',
  styleUrls: ['./svg-viewer.component.scss']
})
export class SvgViewerComponent implements OnInit, AfterViewInit {
  // @ts-ignore
  @ViewChild("element", {static: true}) element: ElementRef;

  // @ts-ignore
  canvas: Svg

  @Input() content: string = `<g><g><ellipse rx="32" ry="30.5" cx="208" cy="84" fill="#cccccc" stroke-width="2" stroke="#ffffff"></ellipse></g><g><circle r="42.04759208325728" cx="461" cy="197" fill="#cccccc" stroke-width="2" stroke="#ffffff"></circle></g></g>`

  constructor() {
  }

  ngOnInit(): void {
  }

  ngAfterViewInit(): void {
    this.element.nativeElement.innerHTML = this.content
  }

}
