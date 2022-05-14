import {HmiComponent} from "../hmi";

export let PumpComponent: HmiComponent = {
  uuid: "pump",
  name: "水泵",
  icon: "/assets/hmi/pump.svg",
  group: "工业",


  //color: true,

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
