import {HmiComponent} from "../../hmi";

export let ImageComponent: HmiComponent = {
  uuid: "image",
  name: "图像",
  icon: "/assets/hmi/components/image.svg",
  group: "基础组件",
  type: "image",

  setup(properties: any): void {
    if (properties.stroke) { // @ts-ignore
      this.$element.stroke(properties.stroke)
    }
    if (properties.color) { // @ts-ignore
      this.$element.fill(properties.color)
    }
  }
}
