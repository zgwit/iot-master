import {HmiComponent} from "../../hmi";

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
    this.element = this.$container.ellipse(this.$properties.width, this.$properties.height)
  },

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("fill"))//@ts-ignore
      this.element.fill(p.fill)
    if (props.hasOwnProperty("color") || props.hasOwnProperty("stroke"))//@ts-ignore
      this.element.stroke({color:p.color, width:p.stroke})
    if (props.hasOwnProperty("width") || props.hasOwnProperty("height"))//@ts-ignore
      this.element.size(p.width, p.height)
  }
}
