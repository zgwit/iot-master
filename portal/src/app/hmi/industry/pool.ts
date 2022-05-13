import {HmiComponent} from "../hmi";

export let PoolComponent: HmiComponent = {
  uuid: "pool",
  name: "水池",
  icon: "/assets/hmi/pool.svg",
  group: "工业",
  type: "svg",

  properties: [
    {
      label: '液体颜色',
      name: 'color',
      type: 'color',
      default: '#8BBB11'
    },
    {
      label: '内壁背景',
      name: 'back',
      type: 'color',
      default: '#666'
    },
    {
      label: '水池背景',
      name: 'fill',
      type: 'color',
      default: '#ccc'
    },
  ],


  //配置
  create(props: any) {
    // @ts-ignore
    let box = this.$element.bbox();
    let radius = Math.min(50, box.height * 0.5, box.width * 0.5)
    let stroke = radius * 0.2

    // @ts-ignore
    this.clipRect = this.$element.rect().size("100%", "100%").move(0, radius)
    
    // @ts-ignore
    let clipRect = this.$element.clip().add(this.clipRect);
        
    // @ts-ignore
    this.rect = this.$element.rect().size("100%", "100%").radius(radius).clipWith(clipRect)

    // @ts-ignore
    this.back = this.$element.rect().size(box.width - stroke * 2, box.height - stroke * 2).radius(radius).cx(box.cx).cy(box.cy).clipWith(clipRect)

    // @ts-ignore
    this.clipCell = this.$element.rect().size("100%", "100%").move(0, box.cy)

    // @ts-ignore
    let clipCell = this.$element.clip().add(this.clipCell)

    // @ts-ignore
    this.cell = this.$element.rect().size(box.width - stroke * 2, box.height - stroke * 2).radius(radius).cx(box.cx).cy(box.cy).clipWith(clipCell)

  },

  resize() {
    // @ts-ignore
    let box = this.$element.bbox();
    let radius = Math.min(50, box.height * 0.5, box.width * 0.5)
    let stroke = radius * 0.2

    // @ts-ignore
    this.rect.radius(radius)
    
    // @ts-ignore
    this.back.size(box.width - stroke * 2, box.height - stroke * 2).radius(radius).cx(box.cx).cy(box.cy)

    // @ts-ignore
    this.cell.size(box.width - stroke * 2, box.height - stroke * 2).radius(radius).cx(box.cx).cy(box.cy)

    // @ts-ignore
    this.clipRect.move(0, radius)
    
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
  },
}
