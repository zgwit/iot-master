import {HmiComponent} from "../hmi";

export let LightComponent: HmiComponent = {
  uuid: "light",
  name: "指示灯",
  icon: "/assets/hmi/light.svg",
  group: "工业",
  drawer: "circle",

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
    {
      label: '颜色',
      name: 'color',
      type: 'color',
      default: '#8BBB11'
    },
    {
      label: '关闭色',
      name: 'dark',
      type: 'color',
      default: '#eee'
    },
    {
      label: '边框色',
      name: 'fill',
      type: 'color',
      default: '#ccc'
    },
    {
      label: '边框',
      name: 'stroke',
      type: 'number',
      default: 10
    },
  ],

  create() {
    //@ts-ignore
    this.element = this.$container.circle(this.$properties.radius)
    //@ts-ignore
    this.cell = this.$container.circle(0)
  },

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("color"))//@ts-ignore
      this.cell.fill(p.color)
    if (props.hasOwnProperty("fill"))//@ts-ignore
      this.element.fill(p.fill)
    if (props.hasOwnProperty("radius") || props.hasOwnProperty("stroke")) {
      //@ts-ignore
      this.element.radius(p.radius)
      //@ts-ignore
      this.cell.radius(p.radius - p.stroke)
    }
    if (props.hasOwnProperty("x") || props.hasOwnProperty("y")) {
      // @ts-ignore
      this.element.center(p.x, p.y)
      // @ts-ignore
      this.cell.center(p.x, p.y)
    }
  }
}
