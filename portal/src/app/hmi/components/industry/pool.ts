import {HmiComponent} from "../../hmi";

export let PoolComponent: HmiComponent = {
  id: "pool",
  name: "水池",
  icon: "/assets/hmi/pool.svg",
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
      label: '圆角',
      name: 'radius',
      type: 'number',
      default: 20
    },
    {
      label: '边框',
      name: 'stroke',
      type: 'number',
      default: 10
    },
  ],

  values: [
    {name: "value", label: "内容"},
  ],


  //配置
  create() {

    // @ts-ignore
    this.clipRect = this.$container.rect().size("100%", "100%")

    // @ts-ignore
    this.rect = this.$container.rect()

    // @ts-ignore
    this.back = this.$container.rect()

    // @ts-ignore
    this.clipCell = this.$container.rect().size("100%", "100%")

    // @ts-ignore
    this.cell = this.$container.rect()
  },

  //配置
  setup(props: any) {
    //@ts-ignore
    let p = this.$properties

    if (props.hasOwnProperty("color"))  // @ts-ignore
      this.cell.fill(p.color)
    if (props.hasOwnProperty("back"))  // @ts-ignore
      this.back.fill(p.back)
    if (props.hasOwnProperty("fill"))  // @ts-ignore
      this.rect.fill(p.fill)
    if (props.hasOwnProperty("stroke")
      || props.hasOwnProperty("radius")
      || props.hasOwnProperty("width")
      || props.hasOwnProperty("height")
    ) {
      // @ts-ignore
      this.rect.radius(p.radius).size(p.width, p.height).move(p.x, p.y)
      // @ts-ignore
      this.back.radius(p.radius - p.stroke).size(p.width - p.stroke * 2, p.height - p.stroke * 2)
        .cx(p.x + p.width / 2).cy(p.y + p.height / 2)
      // @ts-ignore
      this.cell.radius(p.radius - p.stroke).size(p.width - p.stroke * 2, p.height - p.stroke * 2)
        .cx(p.x + p.width / 2).cy(p.y + p.height / 2)

      // @ts-ignore
      this.clipRect.move(0, p.y + p.radius + p.stroke)
      // @ts-ignore
      let clipRect = this.$container.clip().add(this.clipRect)
      // @ts-ignore
      this.rect.unclip().clipWith(clipRect)
      // @ts-ignore
      this.back.unclip().clipWith(clipRect)

      // @ts-ignore
      this.clipCell.y(p.y + p.radius + p.stroke + (p.height - p.stroke * 2) * 0.6) //TODO value
      // @ts-ignore
      this.cell.unclip().clipWith(this.$container.clip().add(this.clipCell))
    }
    if (props.hasOwnProperty("x") || props.hasOwnProperty("y")) {
      // @ts-ignore
      this.clipRect.move(0, p.y + p.radius + p.stroke)
      // @ts-ignore
      let clipRect = this.$container.clip().add(this.clipRect)
      // @ts-ignore
      this.rect.unclip().clipWith(clipRect)
      // @ts-ignore
      this.back.unclip().clipWith(clipRect)
      // @ts-ignore
      this.clipCell.y(p.y + p.stroke + (p.height - p.stroke * 2) * 0.6) //TODO value
      // @ts-ignore
      this.cell.unclip().clipWith(this.$container.clip().add(this.clipCell))
    }
  },

  //更新数据
  update(values: any) {

  },

};
