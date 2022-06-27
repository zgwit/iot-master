
export let SideMenu = [
  {
    title: '首页',
    icon: 'home',
    router: 'dash',
  },
  {
    title: '连接管理',
    icon: 'block',
    children: [
      {
        title: '通道',
        router: 'tunnel'
      },
      {
        title: '服务器',
        router: 'server'
      },
      {
        title: '远程调试',
        router: 'transfer'
      },
    ]
  },
  {
    title: '设备管理',
    icon: 'appstore',
    children: [
      {
        title: '设备',
        router: 'device'
      },
      {
        title: '产品库',
        router: 'product'
      },
    ]
  },
  {
    title: '项目管理',
    icon: 'cluster', //project
    children: [
      {
        title: '项目',
        router: 'project'
      },
      {
        title: '模板库',
        router: 'template'
      },
    ]
  },
  {
    title: '组态管理',
    icon: 'build', //project
    children: [
      {
        title: '组态',
        router: 'hmi'
      },
      {
        title: '组件库',
        router: 'component',
        disabled: true,
      },
    ]
  },
  {
    title: '扩展管理',
    icon: 'appstore-add',
    open: false,
    children: [
      // {
      //   title: '插件',
      //   router: 'extension/plugin'
      // },
      {
        title: '协议',
        router: 'extension/protocol'
      },
      {
        title: '摄像头',
        router: 'camera'
      },
      // {
      //   title: '接口',
      //   router: 'api'
      // },
    ]
  },
  {
    title: '系统设置',
    icon: 'setting',
    open: false,
    children: [
      // {
      //   title: '系统设置',
      //   router: 'setting'
      // },
      {
        title: '用户管理',
        router: 'setting/user'
      },
      {
        title: '激活码',
        router: 'setting/license'
      },
    ]
  },
  {
    title: '退出',
    icon: 'logout',
    router: 'logout'
  }
];
