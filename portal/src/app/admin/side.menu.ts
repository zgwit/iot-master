
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
        title: '连接',
        router: 'link'
      },
      {
        title: '透传',
        router: 'pipe'
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
        title: '元件库',
        router: 'element'
      },
      {
        title: '在线元件库',
        router: 'online-element',
        disabled: true,
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
      {
        title: '在线模板库',
        router: 'online-template',
        disabled: true,
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
      {
        title: '在线组件库',
        router: 'online-component',
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
        title: '修改密码',
        router: 'setting/password'
      },
    ]
  },
  {
    title: '退出',
    icon: 'logout',
    router: 'logout'
  }
];
