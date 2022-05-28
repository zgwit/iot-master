import {HmiComponent} from "../../hmi";

export let FanComponent: HmiComponent = {
  id: "fan",
  name: "风机",
  icon: "/assets/hmi/fan.svg",
  group: "工业",

  //color: true,

  values: [
    {name: "open", label: "运行"},
    {name: "speed", label: "速度"},
  ],

  create() {
    //@ts-ignore
    this.element = this.$container.image().load("/assets/hmi/fan.svg")
  },

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("width") || props.hasOwnProperty("height"))//@ts-ignore
      this.element.size(p.width, p.height)
    if (props.hasOwnProperty("color"))//@ts-ignore
      this.element.fill(p.color)

  },

  update(values: any) {
    if (values.speed > 0) {
      //@ts-ignore
      //this.product.animate().ease('-').transform({rotate:90, origin: 'center'}).loop()
      this.element.animate().ease('-').transform({rotate:120}).loop()
    }
  }
}
