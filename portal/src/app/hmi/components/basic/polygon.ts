import {HmiComponent} from "../../hmi";

export let PolygonComponent: HmiComponent = {
  uuid: "polygon",
  name: "折线",
  icon: "/assets/hmi/polygon.svg",
  group: "基础组件",
  drawer: "poly",

  color: true,
  stroke: true,
  rotation: false,
  position: false,

  create() {
    //@ts-ignore
    this.element = this.$container.polygon()
  },

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("fill"))//@ts-ignore
      this.element.fill(p.fill)
    if (props.hasOwnProperty("color") || props.hasOwnProperty("stroke"))//@ts-ignore
      this.element.stroke({color:p.color, width:p.stroke})
    if (props.hasOwnProperty("points")) {
      //@ts-ignore
      this.element.plot(p.points)
    }
  }
}
