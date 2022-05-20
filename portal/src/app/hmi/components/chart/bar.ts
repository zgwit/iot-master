import {HmiComponent} from "../../hmi";
import * as echarts from "echarts";

export let BarChartComponent: HmiComponent = {
  id: "bar-chart",
  name: "柱状图",
  icon: "/assets/hmi/chart-bar.svg",
  group: "图表",
  drawer: "rect",

  data() {
    return {
      options: {
        tooltip: {},
        xAxis: {data: ['a', 'b', 'c', 'e']},
        yAxis: {},
        series: [
          {
            name: "d",
            type: "bar",
            data: [1, 2, 3, 4]
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
