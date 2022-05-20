import {HmiComponent} from "../../hmi";

export let PolylineComponent: HmiComponent = {
  id: "polyline",
  name: "折线",
  icon: "/assets/hmi/polyline.svg",
  group: "基础组件",
  drawer: "poly",

  color: false,
  stroke: true,
  rotation: false,
  position: false,

  create() {
    //@ts-ignore
    this.element = this.$container.polyline().fill('none')
  },

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("color") || props.hasOwnProperty("stroke"))//@ts-ignore
      this.element.stroke({color:p.color, width:p.stroke})
    if (props.hasOwnProperty("points")) {
      //@ts-ignore
      this.element.plot(p.points)
    }
  }
}
