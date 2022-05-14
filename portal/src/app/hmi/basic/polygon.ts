import {HmiComponent} from "../hmi";

export let PolygonComponent: HmiComponent = {
  uuid: "polygon",
  name: "折线",
  icon: "/assets/hmi/polygon.svg",
  group: "基础组件",
  drawer: "poly",

  color: true,
  stroke: true,
  rotation: false,

  create() {
  },

  setup(properties: any): void {

  }
}
