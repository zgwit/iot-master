import {HmiComponent} from "../hmi";

export let PolylineComponent: HmiComponent = {
  uuid: "polyline",
  name: "折线",
  icon: "/assets/hmi/polyline.svg",
  group: "基础组件",
  type: "polyline",
  color: true,
  stroke: true,
  rotation: false,

  setup(properties: any): void {

  }
}
