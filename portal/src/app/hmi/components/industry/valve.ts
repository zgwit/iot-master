import {HmiComponent} from "../../hmi";

export let ValveComponent: HmiComponent = {
  uuid: "valve",
  name: "阀门",
  icon: "/assets/hmi/valve.svg",
  group: "工业",
  drawer: "rect",

  //color: true,

  create() {
    //@ts-ignore
    this.element = this.$container.image().load("/assets/hmi/valve.svg")
  },

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("width") || props.hasOwnProperty("height"))//@ts-ignore
      this.element.size(p.width, p.height)
    if (props.hasOwnProperty("color"))//@ts-ignore
      this.element.fill(p.color)
  }
}
