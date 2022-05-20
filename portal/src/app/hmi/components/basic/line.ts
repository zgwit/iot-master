import {HmiComponent} from "../../hmi";

export let LineComponent: HmiComponent = {
  id: "line",
  name: "直线",
  icon: "/assets/hmi/line.svg",
  group: "基础组件",
  drawer: "line",

  color: false,
  stroke: true,
  rotation: false,
  position: false,

  properties: [
    {
      label: 'x1',
      name: 'x1',
      type: 'number',
    },
    {
      label: 'y1',
      name: 'y1',
      type: 'number',
    }, {
      label: 'x2',
      name: 'x2',
      type: 'number',
    },
    {
      label: 'y2',
      name: 'y2',
      type: 'number',
    },
  ],

  create() {
    //@ts-ignore
    this.element = this.$container.line()
  },

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("color") || props.hasOwnProperty("stroke"))//@ts-ignore
      this.element.stroke({color:p.color, width:p.stroke})
    if (props.hasOwnProperty("x1") || props.hasOwnProperty("y1")
      || props.hasOwnProperty("x2") || props.hasOwnProperty("y2")) {
      //@ts-ignore
      this.element.plot(p.x1, p.y1, p.x2, p.y2)
    }
  }
}
