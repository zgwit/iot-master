import{_ as t,a as i}from"./uni-list.e8258a94.js";import{o as a,c as e,w as s,a as l,b as r,r as d,g as n,e as o}from"./index.945c150b.js";import{_ as u}from"./uni-card.2a70e0c0.js";import{_ as c}from"./uni-section.3cb1dac6.js";import{r as m}from"./const.427faa2a.js";import{_ as f}from"./plugin-vue_export-helper.21dcd24c.js";import"./uni-icons.a2708d34.js";var h=f({data:()=>({data:{}}),onLoad(t){this.id=t.id,this.load()},methods:{load(){m({url:"tunnel/"+this.id,success:t=>{this.data=t}})}}},[["render",function(m,f,h,p,x,_){const g=d(n("uni-list-item"),t),T=d(n("uni-list"),i),j=d(n("uni-card"),u),b=d(n("uni-section"),c),v=o;return a(),e(v,null,{default:s((()=>[l(j,{title:x.data.name,note:"Tips"},{default:s((()=>[l(T,{border:!1},{default:s((()=>[l(g,{title:"ID",rightText:x.data.id},null,8,["rightText"]),l(g,{title:"网关",rightText:""}),l(g,{title:"地址",rightText:""}),l(g,{title:"远程",rightText:""}),l(g,{title:"创建时间",rightText:x.data.created},null,8,["rightText"])])),_:1})])),_:1},8,["title"]),l(T,null,{default:s((()=>[l(g,{title:"编辑",link:"",to:"./edit?id="+m.id},null,8,["to"]),l(g,{title:"设备列表",note:""})])),_:1}),l(b,{title:"设备",type:"line"},{default:s((()=>[r(" 内容主体，可自定义内容及样式 ")])),_:1})])),_:1})}]]);export{h as default};