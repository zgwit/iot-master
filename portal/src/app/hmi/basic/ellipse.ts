import {HmiComponent} from "../hmi";

export let EllipseComponent: HmiComponent = {
  uuid: "ellipse",
  name: "椭圆",
  icon: "/assets/hmi/ellipse.svg",
  group: "基础组件",
  type: "ellipse",

  color: true,
  stroke: true,

  setup(properties: any): void {

  }
}
