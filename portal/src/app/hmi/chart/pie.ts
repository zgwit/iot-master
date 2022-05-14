import {HmiComponent} from "../hmi";
import * as echarts from "echarts";

export let PieChartComponent: HmiComponent = {
  uuid: "pie-chart",
  name: "饼状图",
  icon: "/assets/hmi/chart-pie.svg",
  group: "图表",
  drawer: "rect",

  data() {
    return {
      options: {
        tooltip: {},
        series: [
          {
            name: "d",
            type: "pie",
            radius: '65%',
            center: ['50%', '50%'],
            data: [
              {name: "cpu", value: 20},
              {name: "mem", value: 60}]
          }
        ]
      }
    }
  },

  create() {
    //@ts-ignore
    this.element = this.$container.foreignObject()
    //@ts-ignore
    this.chart = echarts.init(this.element.node)
    //@ts-ignore
    this.chart.setOption(this.options)
  },

  setup(props: any): void {
    //@ts-ignore
    let p = this.$properties
    if (props.hasOwnProperty("width") || props.hasOwnProperty("height")) {
      //@ts-ignore
      this.element.size(p.width, p.height)
      //@ts-ignore
      this.chart.resize()
    }
  },

  update(values: any) {

  }
}
