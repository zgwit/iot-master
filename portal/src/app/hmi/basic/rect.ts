import {HmiComponent} from "../hmi";

export let RectComponent: HmiComponent = {
  uuid: "rect",
  name: "矩形",
  icon: "/assets/hmi/rect.svg",
  group: "基础组件",
  type: "rect",
  color: true,
  stroke: true,

  setup(properties: any): void {

  }
}
