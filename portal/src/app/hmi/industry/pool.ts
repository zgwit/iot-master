import {HmiComponent} from "../hmi";

export let PoolComponent: HmiComponent = {
  uuid: "pool",
  name: "水池",
  icon: "/assets/hmi/pool.svg",
  group: "工业",
  type: "svg",

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


  //配置
  create(props: any) {

    // @ts-ignore
    this.clipRect = this.$element.rect().size("100%", "100%")
    
    // @ts-ignore
    let clipRect = this.$element.clip().add(this.clipRect);
        
    // @ts-ignore
    this.rect = this.$element.rect().size("100%", "100%").clipWith(clipRect)

    // @ts-ignore
    this.back = this.$element.rect().clipWith(clipRect)

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
    let radius = this.$properties.radius

    // @ts-ignore
    let stroke = this.$properties.stroke


    // @ts-ignore
    this.rect.radius(radius)

    // @ts-ignore
    this.back.radius(radius).size(box.width - stroke * 2, box.height - stroke * 2).cx(box.cx).cy(box.cy)

    // @ts-ignore
    this.cell.radius(radius).size(box.width - stroke * 2, box.height - stroke * 2).cx(box.cx).cy(box.cy)

    // @ts-ignore
    this.clipRect.move(0, radius + stroke)
    
    // @ts-ignore
    this.clipCell.move(0, box.cy)
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
    if (props.hasOwnProperty("radius") || props.hasOwnProperty("stroke")) {
      // @ts-ignore
      this.$component.resize.call(this)
    }
  },

  //更新数据
  update(values: any) {

  },

};
