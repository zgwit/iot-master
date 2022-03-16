import {HmiComponent} from "../../hmi";

export let PolylineComponent: HmiComponent = {
  uuid: "polyline",
  name: "折线",
  icon: "/assets/hmi/components/polyline.svg",
  group: "基础组件",
  type: "polyline",
  stroke: true,
  color: true,

  setup(properties: any): void {
    if (properties.stroke) { // @ts-ignore
      this.$element.stroke(properties.stroke)
    }
    if (properties.color) { // @ts-ignore
      this.$element.fill(properties.color)
    } else { // @ts-ignore
      this.$element.fill("none")
    }
  }
}
