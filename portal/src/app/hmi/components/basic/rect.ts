import {HmiComponent} from "../../hmi";

export let RectComponent: HmiComponent = {
  uuid: "rect",
  name: "矩形",
  icon: "/assets/hmi/components/rect.svg",
  group: "基础组件",
  type: "rect",
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
