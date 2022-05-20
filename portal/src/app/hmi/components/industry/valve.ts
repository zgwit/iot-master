import {HmiComponent} from "../../hmi";

export let ValveComponent: HmiComponent = {
  id: "valve",
  name: "阀门",
  icon: "/assets/hmi/valve.svg",
  group: "工业",
  drawer: "rect",

  //color: true,
  values: [
    {name: "open", label: "打开"},
  ],

  events: [
    {name: "click", label: "点击"},
    {name: "change", label: "变化"},
  ],

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
