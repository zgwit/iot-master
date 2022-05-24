import {Component, Input, OnInit, Output, EventEmitter, AfterViewInit, ElementRef} from '@angular/core';
import {CreateEntityObject, HmiEntity} from "../hmi";
import {SVG, Svg, G} from "@svgdotjs/svg.js";
import {GetComponent} from "../components";
import {GetFieldDeeply} from "../lib";

@Component({
  selector: 'hmi-viewer',
  templateUrl: './viewer.component.html',
  styleUrls: ['./viewer.component.scss']
})
export class ViewerComponent implements OnInit, AfterViewInit {
  // @ts-ignore
  canvas: Svg;

  // @ts-ignore
  grid: Rect;

  entities: Array<HmiEntity> = []


  width = 800
  height = 600
  @Output() invoke = new EventEmitter<any>();

  @Input() set hmi(hmi: any) {
    this.canvas?.size(hmi.width, hmi.height)
    this.entities = hmi.entities
    //绘制
    this.renderEntities()
  }

  @Input() set values(values: any) {
    this.entities.forEach(e => {
      let obj: any = {}
      let has = false
      //找到数据绑定的组件，并传入数据
      Object.keys(e.bindings).forEach((k: string) => {
        let val = GetFieldDeeply(values, e.bindings[k])
        if (val !== undefined) obj[k] = val
      })
      if (has) e.$component.update?.call(e.$object, obj)
    })
  }

  constructor(private ref: ElementRef) {
  }

  ngOnInit(): void {
  }

  ngAfterViewInit(): void {
    // @ts-ignore
    this.canvas = SVG().addTo(this.ref.nativeElement).size(this.width, this.height);

    //绘制
    this.renderEntities()
  }

  renderEntities() {
    if (!this.canvas) return;
    this.entities?.forEach(entity => {
      entity.$container = new G()
      entity.$component = GetComponent(entity.component)
      CreateEntityObject(entity)
      this.appendEntity(entity);
      entity.$component.setup.call(entity.$object, entity.properties)
      entity.$object.__proto__.$emit = function (event: string, data: any) {
        entity.handlers[event]?.forEach((handler: Function) => handler(data))
      }
    })
  }

  appendEntity(entity: HmiEntity) {
    // @t-s-ignore
    this.canvas.add(entity.$container)
    //传入全局配置
    this.setupEntity(entity, entity.properties)

    this.entities.push(entity)
  }

  setupEntity(entity: HmiEntity, props: any) {
    Object.assign(entity.properties, props)

    //旋转
    if (props.hasOwnProperty("rotate"))
      entity.$container.transform({rotate: props.rotate})
    if (props.hasOwnProperty("x"))
      entity.$container.x(props.x)
    if (props.hasOwnProperty("y"))
      entity.$container.y(props.y)

    entity.$component.setup.call(entity.$object, props)
  }


}
