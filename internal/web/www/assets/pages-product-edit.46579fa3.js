import{v as a,H as e,N as l,J as s,o as t,c as d,w as o,a as r,b as u,r as i,g as n,h as m,e as p}from"./index.945c150b.js";import{_ as c,a as f,b as h}from"./uni-forms.82256a5d.js";import{r as v}from"./const.427faa2a.js";import{_ as V}from"./plugin-vue_export-helper.21dcd24c.js";import"./uni-icons.a2708d34.js";var _=V({data:()=>({id:"",data:{name:"",manufacturer:"",version:"",points:[],pollers:[]}}),onLoad(a){this.id=a.id,this.id&&this.load()},methods:{load(){v({url:"product/"+this.id,success:a=>{this.data=a}})},save(){a("log","at pages/product/edit.vue:50",this.id),e({}),v({url:"product/"+(this.id||"create"),method:"POST",data:this.data,success(){l()},complete(){s()}})}}},[["render",function(a,e,l,s,v,V){const _=i(n("uni-easyinput"),c),b=i(n("uni-forms-item"),f),j=i(n("uni-forms"),h),g=m,y=p;return t(),d(y,{class:"p10"},{default:o((()=>[r(j,null,{default:o((()=>[r(b,{label:"名称",name:"name"},{default:o((()=>[r(_,{modelValue:v.data.name,"onUpdate:modelValue":e[0]||(e[0]=a=>v.data.name=a),placeholder:""},null,8,["modelValue"])])),_:1}),r(b,{label:"厂商",name:"type"},{default:o((()=>[r(_,{modelValue:v.data.manufacturer,"onUpdate:modelValue":e[1]||(e[1]=a=>v.data.manufacturer=a),placeholder:""},null,8,["modelValue"])])),_:1}),r(b,{label:"版本",name:"version"},{default:o((()=>[r(_,{modelValue:v.data.version,"onUpdate:modelValue":e[2]||(e[2]=a=>v.data.version=a),placeholder:""},null,8,["modelValue"])])),_:1}),u(" TODO：编辑点位 ")])),_:1}),r(g,{type:"primary",onClick:V.save},{default:o((()=>[u("保存")])),_:1},8,["onClick"])])),_:1})}]]);export{_ as default};