import {HmiComponent} from "../../hmi";

export let EllipseComponent: HmiComponent = {
  uuid: "ellipse",
  name: "椭圆",
  icon: "/assets/hmi/components/ellipse.svg",
  group: "基础组件",
  type: "ellipse",
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
