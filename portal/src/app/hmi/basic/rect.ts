import {HmiComponent} from "../hmi";

export let RectComponent: HmiComponent = {
  uuid: "rect",
  name: "矩形",
  icon: "/assets/hmi/rect.svg",
  group: "基础组件",
  drawer: "rect",

  color: true,
  stroke: true,

  create() {
  },

  setup(properties: any): void {

  }
}
