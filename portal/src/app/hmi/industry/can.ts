import {HmiComponent} from "../hmi";

export let CanComponent: HmiComponent = {
  uuid: "can",
  name: "水罐",
  icon: "/assets/hmi/can.svg",
  group: "工业",

  properties: [
    {
      label: '颜色',
      name: 'color',
      type: 'color',
      default: '#8BBB11'
    },
    {
      label: '背景',
      name: 'back',
      type: 'color',
      default: '#666'
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


  //配置
  create() {

    // @ts-ignore
    this.rect = this.$element.rect().size("100%", "100%")

    // @ts-ignore
    this.back = this.$element.rect()

    // @ts-ignore
    this.clipCell = this.$element.rect().size("100%", "100%")

    // @ts-ignore
    let clipCell = this.$element.clip().add(this.clipCell)

    // @ts-ignore
    this.cell = this.$element.rect().clipWith(clipCell)

    // @ts-ignore
    this.$component.resize.call(this)
  },

  resize() {
    // @ts-ignore
    let box = this.$element.bbox()
    // @ts-ignore
    let radius = box.width / 2 // this.$properties.radius

    // @ts-ignore
    let stroke = this.$properties.stroke


    // @ts-ignore
    this.rect.radius(radius)

    // @ts-ignore
    this.back.radius(radius - stroke).size(box.width - stroke * 2, box.height - stroke * 2).cx(box.cx).cy(box.cy)

    // @ts-ignore
    this.cell.radius(radius - stroke).size(box.width - stroke * 2, box.height - stroke * 2).cx(box.cx).cy(box.cy)


    // @ts-ignore
    this.clipCell.move(0, stroke + (box.height - stroke * 2) * 0.6) //TODO value
  },

  //配置
  setup(props: any) {
    if (props.color) { // @ts-ignore
      this.cell.fill(props.color)
    }
    if (props.back) { // @ts-ignore
      this.back.fill(props.back)
    }
    if (props.fill) { // @ts-ignore
      this.rect.fill(props.fill)
    }
    if (props.hasOwnProperty("stroke")) {
      // @ts-ignore
      this.$component.resize.call(this)
    }
  },

  //更新数据
  update(values: any) {

  },

};
