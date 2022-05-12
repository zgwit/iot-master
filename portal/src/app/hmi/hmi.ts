import {
  Circle,
  ClipPath,
  Dom, Element,
  ElementAlias,
  Ellipse, ForeignObject, G, Gradient, Image,
  Line,
  Path,
  Polygon,
  Polyline,
  Rect,
  Svg,
  Text, TextPath,
  Use
} from "@svgdotjs/svg.js";

export type SvgElement = Svg | Rect | Line | Polygon | Polyline | Ellipse | Text | Path | TextPath | Circle | Image | ForeignObject

export interface HmiPropertyItem {
  label: string
  name: string
  type: string | 'boolean' | 'number' | 'text' | 'color' | 'date' | 'time' | 'datetime' | 'font' | 'fontsize'
  unit?: string
  min?: number
  max?: number
  default?: boolean | number | string

  [prop: string]: any
}

export interface HmiEvent {
  event: string
  label: string
}

export interface HmiValue {
  value: string
  label: string
}

export interface HmiComponent {
  uuid: string
  icon: string //url: svg png jpg ...
  name: string

  //类型（默认 svg）
  type?: "rect" | "circle" | "ellipse" | "line" | "polyline" | "polygon" | "image" | "path" | "text" | "svg" | "object"

  //分组（默认 扩展）
  group?: string

  //模板 svg
  svg?: string

  //基础配置
  color?: boolean
  stroke?: boolean
  rotation?: boolean
  position?: boolean

  //扩展配置项
  properties?: Array<HmiPropertyItem>

  //事件
  events?: Array<HmiEvent>

  //监听
  watches?: Array<HmiValue>

  //[prop: string]: any

  //初始化
  create?(props: any): void

  //写入配置
  setup(props: any): void

  //更新数据
  update?(values: any): void

  //产生变量 data(){return {a:1, b2}}
  data?(): any
}

export function basicProperties() {
  return {
    line: false,
    fill: false,
    rotate: true,
    position: true
  }
}

export function GetDefaultProperties(component: HmiComponent): any {
  let obj: any = {};

  component.properties?.forEach(p => {
    if (p.hasOwnProperty('default'))
      obj[p.name] = p.default
  })
  return obj;
}

export function CreateComponentObject(component: HmiComponent, element: SvgElement): any {
  let obj = component.data ? component.data() : {}
  obj.__proto__ = {
    //$name: entity.name,
    $element: element
  }
  return obj
}

export interface HmiEntity {
  name: string
  component: string //uuid
  properties: any //{ [name: string]: any }
  triggers: any
  bindings: any

  $element: SvgElement
  $component: HmiComponent
  $object: any;

  //TODO
  //参数绑定
  //事件响应
  //脚本
}

export interface HmiEntities {
  name: string
  items: Array<HmiEntity>
}

export interface HmiProject {
  name: string
  views: Array<HmiEntities>
}
