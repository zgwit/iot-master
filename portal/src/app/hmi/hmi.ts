import {
  Circle, Container,
  Ellipse, ForeignObject, Image,
  Line,
  Path,
  Polygon,
  Polyline,
  Rect,
  Svg,
  Text, TextPath
} from "@svgdotjs/svg.js";
//import "@svgdotjs/svg.filter.js";
import {strokeProperties, fillProperties, positionProperties, rotateProperties} from "./properties";
import {GetComponent} from "./components";

export type HmiElement =
  Svg
  | Rect
  | Line
  | Polygon
  | Polyline
  | Ellipse
  | Text
  | Path
  | TextPath
  | Circle
  | Image
  | ForeignObject

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
  //type?: "rect" | "circle" | "ellipse" | "line" | "polyline" | "polygon" | "image" | "path" | "text" | "svg" | "object"
  drawer?: "rect" | "circle" | "line" | "poly"

  //分组（默认 扩展）
  group?: string

  //基础配置
  color?: boolean //填充色
  stroke?: boolean //线条
  rotation?: boolean //旋转
  position?: boolean //位置

  //扩展配置项
  properties?: Array<HmiPropertyItem>

  //事件
  events?: Array<HmiEvent>

  //监听
  watches?: Array<HmiValue>

  //[prop: string]: any

  //初始化
  create(): void

  //写入配置
  setup(props: any): void

  //更新数据
  update?(values: any): void

  //产生变量 data(){return {a:1, b2}}
  data?(): any
}

export function GetComponentGlobalProperties(obj: HmiComponent) {
  let properties = [];
  if (obj.color)
    properties?.unshift(...fillProperties)
  if (obj.stroke)
    properties?.unshift(...strokeProperties)
  if (obj.rotation)
    properties?.unshift(...rotateProperties)
  if (obj.position)
    properties?.unshift(...positionProperties)
  return properties
}

export function GetComponentAllProperties(obj: HmiComponent) {
  //@ts-ignore
  return GetComponentGlobalProperties(obj).concat(obj.properties)
}

export function GetPropertiesDefault(component: HmiComponent): any {
  let obj: any = {};
  let properties = GetComponentAllProperties(component)
  properties.forEach(p => {
    if (p.hasOwnProperty('default'))
      obj[p.name] = p.default
  })
  return obj;
}

export interface HmiEntity {
  name: string
  component: string //uuid

  //属性
  properties: any //{ [name: string]: any }
  //响应
  handlers: any //{ [event: string]: []invokes | script }
  //绑定
  bindings: any //{ [name: string]: any }

  //$element: HmiElement
  $container: Container
  $component: HmiComponent
  $object: any;
}

export function CreateEntityObject(entity: HmiEntity) {
  entity.$object = entity.$component.data?.call(entity) ||  {}
  entity.$object.__proto__ = {
    $container: entity.$container,
    $component: entity.$component,
    $properties: entity.properties,
  }
  entity.$component.create.call(entity.$object)
}

export interface HmiEntities {
  name: string
  items: Array<HmiEntity>
}

export interface HmiProject {
  name: string
  views: Array<HmiEntities>
}
