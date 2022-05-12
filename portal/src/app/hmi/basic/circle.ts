import {HmiComponent} from "../hmi";

export let CircleComponent: HmiComponent = {
  uuid: "circle",
  name: "圆形",
  icon: "/assets/hmi/circle.svg",
  group: "基础组件",
  type: "circle",

  color: true,
  stroke: true,
  rotation: false,

  setup(properties: any): void {

  }
}
