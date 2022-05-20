import {HmiComponent} from "../../hmi";

export let PumpComponent: HmiComponent = {
  id: "pump",
  name: "水泵",
  icon: "/assets/hmi/pump.svg",
  group: "工业",


  //color: true,
  values: [
    {name: "open", label: "运行"},
    {name: "speed", label: "速度"},
  ],

  create() {
    //@ts-ignore
    this.element = this.$container.image().load("/assets/hmi/pump.svg")
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
