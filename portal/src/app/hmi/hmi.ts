import {ElementAlias} from "@svgdotjs/svg.js";

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
  noRotation?: boolean
  noPosition?: boolean

  //扩展配置项
  properties?: Array<HmiPropertyItem>

  [prop: string]: any

  //初始化
  init?(properties: any): void

  //写入配置
  setup(properties: any): void

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

export function CreateComponentObject(component: HmiComponent, element: ElementAlias): any {
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

  $element: ElementAlias
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
