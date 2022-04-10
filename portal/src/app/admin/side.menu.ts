
export let SideMenu = [
  {
    title: '首页',
    icon: 'home',
    router: 'home',
  },
  {
    title: '连接',
    icon: 'block',
    children: [
      {
        title: '通道',
        router: 'tunnel'
      },
      {
        title: '链接',
        router: 'link'
      },
    ]
  },
  {
    title: '设备',
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
    ]
  },
  {
    title: '项目',
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
    title: '扩展',
    icon: 'appstore-add',
    open: false,
    children: [
      {
        title: '插件',
        router: 'extension/plugin'
      },
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
    title: '设置',
    icon: 'setting',
    open: false,
    children: [
      {
        title: '系统设置',
        router: 'setting'
      },
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
