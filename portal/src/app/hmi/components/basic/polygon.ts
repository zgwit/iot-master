import {HmiComponent} from "../../hmi";

export let PolygonComponent: HmiComponent = {
  uuid: "polygon",
  name: "折线",
  icon: "/assets/hmi/components/polygon.svg",
  group: "基础组件",
  type: "polygon",
  stroke: true,
  color: true,

  setup(properties: any): void {
    if (properties.stroke) { // @ts-ignore
      this.$element.stroke(properties.stroke)
    }
    if (properties.color) { // @ts-ignore
      this.$element.fill(properties.color)
    }
  }
}
