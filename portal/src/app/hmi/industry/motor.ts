import {HmiComponent} from "../hmi";

export let MotorComponent: HmiComponent = {
  uuid: "motor",
  name: "电机",
  icon: "/assets/hmi/motor.svg",
  group: "工业",

  //color: true,

  create() {
    //@ts-ignore
    this.element = this.$container.image().load("/assets/hmi/motor.svg")
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
