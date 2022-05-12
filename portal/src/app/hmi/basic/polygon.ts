import {HmiComponent} from "../hmi";

export let PolygonComponent: HmiComponent = {
  uuid: "polygon",
  name: "折线",
  icon: "/assets/hmi/polygon.svg",
  group: "基础组件",
  type: "polygon",
  color: true,
  stroke: true,
  rotation: false,

  setup(properties: any): void {

  }
}
