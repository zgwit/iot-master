import {HmiComponent} from "../../hmi";

export let CircleComponent: HmiComponent = {
  uuid: "circle",
  name: "圆形",
  icon: "/assets/hmi/components/circle.svg",
  group: "基础组件",
  type: "circle",
  stroke: true,
  color: true,
  noRotation: true,

  setup(properties: any): void {
    if (properties.stroke) { // @ts-ignore
      this.$element.stroke(properties.stroke)
    }
    if (properties.color) { // @ts-ignore
      this.$element.fill(properties.color)
    }
  }
}
