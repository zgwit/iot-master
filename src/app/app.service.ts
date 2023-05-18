import { Injectable } from '@angular/core';
import { RequestService } from "./request.service";

const internals: any[] = [{
    icon: '/assets/product.png',
    nzIcon: 'appstore',
    name: '产品管理',
    internal: true,
    entries: [
        { name: '所有产品', path: '/product/list' },
        // {name: '创建产品', path: '/product/create'},
    ]
}, {
    icon: '/assets/device.png',
    nzIcon: 'block',
    name: '设备管理',
    internal: true,
    entries: [
        { name: '所有设备', path: '/device/list' },
        // {name: '创建设备', path: '/device/create'},
        // {name: '批量创建', path: '/device/batch'},
        // {name: '设备分组', path: '/device/group'},
        // {name: '设备类型', path: '/device/type'},
        // {name: '设备区域', path: '/device/area'}
    ]
},
{
    icon: '/assets/user.png',
    nzIcon: 'user',
    name: '用户管理',
    internal: true,
    entries: [
        { name: '所有用户', path: '/user/list' },
        // {name: '创建用户', path: '/user/create'},
        { name: '角色权限', path: '/user/role' },
    ]
}, {
    icon: '/assets/server.png',
    nzIcon: 'database',
    name: '数据总线',
    internal: true,
    entries: [
        { name: '所有总线', path: '/broker/list' },
        // {name: '创建总线', path: '/broker/create'},
    ]
}, {
    icon: '/assets/gateway.png',
    nzIcon: 'partition',
    name: '网关管理',
    internal: true,
    entries: [
        { name: '所有网关', path: '/gateway/list' },
        { name: '批量创建', path: '/gateway/batch' },
        // {name: '创建网关', path: '/gateway/create'},
    ]
}, {
    icon: '/assets/plugin.png',
    nzIcon: 'appstore-add',
    name: '插件管理',
    internal: true,
    entries: [
        { name: '所有插件', path: '/plugin/list' },
        // {name: '创建插件', path: '/plugin/create'},
    ]
}, {
    icon: '/assets/setting.png',
    nzIcon: 'setting',
    name: '系统设置',
    internal: true,
    entries: [
        { name: '数据备份', path: '/setting/backup' },
        { name: 'Web服务', path: '/setting/web' },
        { name: '数据库', path: '/setting/database' },
        { name: '日志', path: '/setting/log' },
        { name: 'OEM', path: '/setting/oem' },
        { name: 'MQTT', path: '/setting/mqtt' },
    ]
}];

@Injectable({
    providedIn: 'root'
})
export class AppService {

    apps: any[] = internals;

    constructor(private rs: RequestService) {
        this.load()
    }

    load() {
        this.rs.get("apps").subscribe(res => {
            const arr = res.data.sort((a: { id: string; }, b: { id: string; }) => {
                return a.id.length - b.id.length;
            })
            this.apps = internals.concat(arr)
        })
    }
}
