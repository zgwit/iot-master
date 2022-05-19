import {HmiComponent} from "./hmi";
import {LineComponent} from "./components/basic/line";
import {CircleComponent} from "./components/basic/circle";
import {PolylineComponent} from "./components/basic/polyline";
import {PolygonComponent} from "./components/basic/polygon";
import {RectComponent} from "./components/basic/rect";
import {EllipseComponent} from "./components/basic/ellipse";
import {TextComponent} from "./components/basic/text";
import {ImageComponent} from "./components/basic/image";
import {ButtonComponent} from "./components/control/button";
import {InputComponent} from "./components/control/input";
import {SwitchComponent} from "./components/control/switch";
import {SliderComponent} from "./components/control/slider";
import {ValueComponent} from "./components/control/value";
import {AlarmComponent} from "./components/industry/alarm";
import {CanComponent} from "./components/industry/can";
import {FanComponent} from "./components/industry/fan";
import {LightComponent} from "./components/industry/light";
import {PipeComponent} from "./components/industry/pipe";
import {PoolComponent} from "./components/industry/pool";
import {PumpComponent} from "./components/industry/pump";
import {ValveComponent} from "./components/industry/valve";
import {MotorComponent} from "./components/industry/motor";
import {BarChartComponent} from "./components/chart/bar";
import {GaugeChartComponent} from "./components/chart/gauge";
import {LineChartComponent} from "./components/chart/line";
import {PieChartComponent} from "./components/chart/pie";
import {RadarChartComponent} from "./components/chart/radar";
import {ProgressComponent} from "./components/control/progress";

export let GroupedComponents: Array<Group> = [];

let indexedComponents: { [name: string]: HmiComponent } = {}
let indexedGroupComponents: { [name: string]: Array<HmiComponent> } = {}

export interface Group {
  name: string,
  components: Array<HmiComponent>
}

export function GetComponent(id: string): HmiComponent {
  return indexedComponents[id];
}


export function LoadComponent(obj: HmiComponent) {
  let base = {
    color: false,
    stroke: false,
    rotation: true,
    position: true,
    group: "扩展",
    properties: [],
  }
  obj = Object.assign(base, obj)

  //if (indexedComponents.hasOwnProperty(obj.uuid))
  indexedComponents[obj.uuid] = obj;

  // @ts-ignore
  let group = indexedGroupComponents[obj.group]
  if (!group) {
    // @ts-ignore
    group = indexedGroupComponents[obj.group] = [];
    // @ts-ignore
    GroupedComponents.push({name: obj.group, components: group});
  }
  group.push(obj)
}

//export function loadIntervalComponents() {
let internalComponents = [
  //基础
  LineComponent, CircleComponent, EllipseComponent, RectComponent,
  PolylineComponent, PolygonComponent, TextComponent, ImageComponent,
  //控件
  ButtonComponent, InputComponent, SwitchComponent, SliderComponent,
  ValueComponent, ProgressComponent,
  //ClockComponent, CameraComponent, WeatherComponent,
  //工业
  //AlarmComponent,
  PipeComponent, ValveComponent, CanComponent, PoolComponent,
  PumpComponent, MotorComponent, FanComponent, LightComponent,
  //图表
  GaugeChartComponent, BarChartComponent, LineChartComponent, PieChartComponent
  //, RadarChartComponent
]
internalComponents.forEach(c => LoadComponent(c))
//}

