import{_ as a,a as e}from"./uni-fab.c776ca19.js";import{o as t,c as i,w as s,a as o,j as d,u as n,F as l,r,e as u,d as m,i as c}from"./index.8d4a8305.js";import{_ as h,a as p}from"./uni-list.723f536b.js";import{r as f}from"./const.9e5d4232.js";import{_ as w}from"./plugin-vue_export-helper.21dcd24c.js";import"./uni-icons.51822eb1.js";var k=w({data:()=>({keyword:"",limit:20,datum:[]}),onPullDownRefresh(){this.datum=[],this.skip=0},onReachBottom(){this.load()},onLoad(a){this.gateway=a.gateway,this.load()},methods:{load(){uni.showLoading({}),f({url:"server/search",method:"POST",data:{skip:this.datum.length,limit:this.limit,keyword:this.keyword?{id:this.keyword,name:this.keyword}:{},filter:this.gateway?{gateway_id:this.gateway}:{}},success:a=>{this.datum=this.datum.concat(a)},complete(){uni.hideLoading(),uni.stopPullDownRefresh()}})},onSearch(){this.datum=[],this.load()},create(){uni.navigateTo({url:"./edit"})}}},[["render",function(f,w,k,y,g,_){const b=r(u("uni-search-bar"),a),j=r(u("uni-fab"),e),v=c,C=r(u("uni-list-item"),h),F=r(u("uni-list"),p),x=m;return t(),i(x,null,{default:s((()=>[o(b,{onConfirm:_.onSearch,onInput:w[0]||(w[0]=()=>{}),placeholder:"ID 名称",modelValue:g.keyword,"onUpdate:modelValue":w[1]||(w[1]=a=>g.keyword=a)},null,8,["onConfirm","modelValue"]),o(j,{onFabClick:_.create},null,8,["onFabClick"]),o(F,null,{default:s((()=>[(t(!0),d(l,null,n(g.datum,((a,e)=>(t(),i(C,{key:e,title:a.name,note:a.id,link:"",to:"./detail?id="+a.id},{header:s((()=>[o(v,{class:"icon",src:"/assets/server.c4705aac.svg",mode:"aspectFit"})])),_:2},1032,["title","note","to"])))),128))])),_:1})])),_:1})}],["__scopeId","data-v-8bb2e8e4"]]);export{k as default};