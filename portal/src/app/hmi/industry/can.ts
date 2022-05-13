import {HmiComponent} from "../hmi";

export let CanComponent: HmiComponent = {
  uuid: "can",
  name: "水罐",
  icon: "/assets/hmi/can.svg",
  group: "工业",
  type: "svg",

  properties: [
    {
      label: '颜色',
      name: 'color',
      type: 'color',
      default: '#ccc'
    },
    {
      label: '背景',
      name: 'fill',
      type: 'color',
      default: '#8BBB11'
    },
  ],

  //配置
  create(props: any) {
    // @ts-ignore
    let box = this.$element.bbox();
    let radius = Math.min(50, box.height * 0.5, box.width * 0.5)

    // @ts-ignore
    this.cell = this.$element.rect().size(box.width - radius * 2, box.height - radius * 2).cx(box.cx)

    // @ts-ignore
    this.rect = this.$element.rect().size("100%", "100%").stroke({width: radius}).fill("none").radius(radius)

    // @ts-ignore
    this.cell.filterWith(add=>{
      add.offset(0,box.height * 0.5)
    })
  },

  resize() {
    // @ts-ignore
    let box = this.$element.bbox();
    let radius = Math.min(50, box.height * 0.5, box.width * 0.5)

    // @ts-ignore
    this.cell.size(box.width - radius * 2, box.height - radius * 2).cx(box.cx)

    // @ts-ignore
    this.rect.stroke({width: radius}).radius(radius)

    // @ts-ignore
    this.cell.filterWith(add=>{
      add.offset(0,box.height * 0.5)
    })
  },

  //配置
  setup(props: any) {
    if (props.color) { // @ts-ignore
      this.rect.stroke({color:props.color})
    }
    if (props.fill) { // @ts-ignore
      this.cell.fill(props.fill)
    }
  },
}
