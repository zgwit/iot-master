import {HmiComponent} from "../../hmi";

export let CircleComponent: HmiComponent = {
  uuid: "circle",
  name: "圆形",
  icon: "/assets/hmi/circle.svg",
  group: "基础组件",
  drawer: "circle",

  color: true,
  stroke: true,
  rotation: false,
  position: false,

  properties: [
    {
      label: 'X',
      name: 'x',
      type: 'number',
    },
    {
      label: 'Y',
      name: 'y',
      type: 'number',
    },
    {
      label: '半径',
      name: 'radius',
      type: 'number',
    },
  ],

  create() {
    //@ts-ignore
    this.element = this.$container.circle(this.$properties.radius)
  },

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("fill"))//@ts-ignore
      this.element.fill(p.fill)
    if (props.hasOwnProperty("color") || props.hasOwnProperty("stroke"))//@ts-ignore
      this.element.stroke({color:p.color, width:p.stroke})
    if (props.hasOwnProperty("radius")) //@ts-ignore
      this.element.radius(p.radius)
    if (props.hasOwnProperty("x") || props.hasOwnProperty("y")) {
      // @ts-ignore
      this.element.center(p.x, p.y)
    }
  }
}
