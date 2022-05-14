import {HmiComponent} from "../hmi";

export let EllipseComponent: HmiComponent = {
  uuid: "ellipse",
  name: "椭圆",
  icon: "/assets/hmi/ellipse.svg",
  group: "基础组件",
  drawer: "rect",

  color: true,
  stroke: true,

  create() {

    //@ts-ignore
    this.ellipse = this.$container.ellipse(this.$properties.width, this.$properties.height).center(this.$properties.x, this.$properties.y)
  },

  setup(properties: any): void {

  }
}
