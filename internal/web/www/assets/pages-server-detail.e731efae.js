import{_ as t,a as i}from"./uni-list.e8258a94.js";import{o as a,c as e,w as r,a as s,r as l,g as d,e as o}from"./index.945c150b.js";import{_ as n}from"./uni-card.2a70e0c0.js";import{r as u}from"./const.427faa2a.js";import{_ as c}from"./plugin-vue_export-helper.21dcd24c.js";import"./uni-icons.a2708d34.js";var m=c({data:()=>({data:{}}),onLoad(t){this.id=t.id,this.load()},methods:{load(){u({url:"server/"+this.id,success:t=>{this.data=t}})}}},[["render",function(u,c,m,f,h,p){const x=l(d("uni-list-item"),t),g=l(d("uni-list"),i),_=l(d("uni-card"),n),T=o;return a(),e(T,null,{default:r((()=>[s(_,{title:h.data.name,note:"Tips"},{default:r((()=>[s(g,{border:!1},{default:r((()=>[s(x,{title:"ID",rightText:h.data.id},null,8,["rightText"]),s(x,{title:"网关",rightText:""}),s(x,{title:"地址",rightText:""}),s(x,{title:"创建时间",rightText:h.data.created},null,8,["rightText"])])),_:1})])),_:1},8,["title"]),s(g,null,{default:r((()=>[s(x,{title:"编辑",link:"",to:"./edit?id="+u.id},null,8,["to"]),s(x,{title:"通道列表",note:""})])),_:1})])),_:1})}]]);export{m as default};