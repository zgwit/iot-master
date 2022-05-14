import {HmiComponent} from "../hmi";

export let PipeComponent: HmiComponent = {
  uuid: "pipe",
  name: "管道",
  icon: "/assets/hmi/pipe.svg",
  group: "工业",
  drawer: "poly",

  color: false,
  stroke: false,
  rotation: false,
  position: false,

  properties: [
    {
      label: '颜色',
      name: 'color',
      type: 'color',
      default: '#fff'
    },
    {
      label: '宽度',
      name: 'stroke',
      type: 'number',
      default: 20
    },
    {
      label: '水流',
      name: 'flow',
      type: 'color',
      default: '#8BBB11'
    },
  ],

  create() {
    //@ts-ignore
    this.element = this.$container.polyline().fill('none')
    //@ts-ignore
    this.flow = this.$container.polyline().fill('none')
  },

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("color") || props.hasOwnProperty("stroke") || props.hasOwnProperty("flow")) {
      //@ts-ignore
      this.element.stroke({color: p.color, width: p.stroke})
      //@ts-ignore
      this.flow.stroke({color:p.flow, width:p.stroke / 2, dasharray: p.stroke + ' ' + Math.floor(p.stroke / 2)})
    }
    if (props.hasOwnProperty("points")) {
      //@ts-ignore
      this.element.plot(p.points)
      //@ts-ignore
      this.flow.plot(p.points)
    }
  }
}
