import{q as t,o as e,c as i,w as s,a,b as l,r as o,e as n,d as r}from"./index.8d4a8305.js";import{_ as d,a as c}from"./uni-list.723f536b.js";import{d as u,_ as m}from"./time.0102cd3d.js";import{_ as f}from"./uni-icons.51822eb1.js";import{_ as h}from"./uni-section.96f3ba19.js";import{r as p}from"./const.9e5d4232.js";import{_ as g}from"./plugin-vue_export-helper.21dcd24c.js";import"./_commonjsHelpers.4e997714.js";var x=g({data:()=>({data:{}}),onLoad(t){this.id=t.id,this.load()},onPullDownRefresh(){this.load()},methods:{format:u,load(){p({url:"tunnel/"+this.id,success:t=>{this.data=t},complete(){uni.stopPullDownRefresh()}})},remove(){uni.showModal({title:"提示",content:"确定删除？",success:e=>{e.confirm&&(t("log","at pages/tunnel/detail.vue:80","用户点击确定"),p({url:"tunnel/"+this.id+"/delete",success:t=>{uni.navigateBack(),uni.showToast({title:"删除成功"})}}))},fail:console.error})}}},[["render",function(t,u,p,g,x,_){const T=o(n("uni-list-item"),d),b=o(n("uni-list"),c),v=o(n("uni-card"),m),j=o(n("uni-icons"),f),k=o(n("uni-section"),h),w=r;return e(),i(w,null,{default:s((()=>[a(v,{title:x.data.name,subTitle:x.data.id,extra:"在线",note:"Tips",thumbnail:"/static/icons/link.svg"},{default:s((()=>[a(b,{border:!1},{default:s((()=>[a(T,{title:"ID",rightText:x.data.id},null,8,["rightText"]),a(T,{title:"网关",rightText:x.data.gateway_id},null,8,["rightText"]),a(T,{title:"地址",rightText:x.data.addr},null,8,["rightText"]),a(T,{title:"远程",rightText:x.data.remote},null,8,["rightText"]),a(T,{title:"创建时间",rightText:_.format(x.data.created)},null,8,["rightText"])])),_:1})])),_:1},8,["title","subTitle"]),a(b,null,{default:s((()=>[a(T,{title:"设备列表",link:"",to:"/pages/device/device?tunnel="+t.id},{header:s((()=>[a(j,{class:"list-icon",customPrefix:"iconfont",type:"icon-list"})])),_:1},8,["to"]),a(T,{title:"编辑通道",link:"",to:"./edit?id="+t.id},{header:s((()=>[a(j,{class:"list-icon",customPrefix:"iconfont",type:"icon-pen"})])),_:1},8,["to"]),a(T,{title:"删除通道",onClick:_.remove,clickable:!0},{header:s((()=>[a(j,{class:"list-icon",customPrefix:"iconfont",type:"icon-dustbin"})])),_:1},8,["onClick"])])),_:1}),a(k,{title:"日志",type:"line"},{default:s((()=>[l(" TODO：Event ")])),_:1})])),_:1})}]]);export{x as default};
