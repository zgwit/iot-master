import{q as a,o as e,c as l,w as d,a as t,b as o,r as s,e as i,f as n,d as u}from"./index.8d4a8305.js";import{b as r,_ as m,a as p}from"./uni-forms.3c0cbf84.js";import{_ as c}from"./uni-section.96f3ba19.js";import{r as f}from"./const.9e5d4232.js";import{_ as h}from"./plugin-vue_export-helper.21dcd24c.js";import"./uni-icons.51822eb1.js";var V=h({data:()=>({id:"",data:{name:"",type:"",addr:"",protocol:{name:""},devices:[]}}),onLoad(a){this.id=a.id,this.id&&this.load()},methods:{load(){f({url:"server/"+this.id,success:a=>{this.data=a}})},save(){a("log","at pages/server/edit.vue:59",this.id),uni.showLoading({}),f({url:"server/"+(this.id||"create"),method:"POST",data:this.data,success(){uni.navigateBack()},complete(){uni.hideLoading()}})}}},[["render",function(a,f,h,V,_,v){const b=s(i("uni-easyinput"),r),y=s(i("uni-forms-item"),m),g=s(i("uni-section"),c),j=s(i("uni-forms"),p),U=n,k=u;return e(),l(k,{class:"p20"},{default:d((()=>[t(j,null,{default:d((()=>[t(y,{label:"名称",name:"name"},{default:d((()=>[t(b,{modelValue:_.data.name,"onUpdate:modelValue":f[0]||(f[0]=a=>_.data.name=a),placeholder:""},null,8,["modelValue"])])),_:1}),t(y,{label:"类型",name:"type"},{default:d((()=>[t(b,{modelValue:_.data.type,"onUpdate:modelValue":f[1]||(f[1]=a=>_.data.type=a),placeholder:"tcp,udp"},null,8,["modelValue"])])),_:1}),t(y,{label:"地址",name:"addr"},{default:d((()=>[t(b,{modelValue:_.data.addr,"onUpdate:modelValue":f[2]||(f[2]=a=>_.data.addr=a),placeholder:""},null,8,["modelValue"])])),_:1}),t(g,{title:"协议",type:"line"},{default:d((()=>[t(y,{label:"名称",name:"name"},{default:d((()=>[t(b,{modelValue:_.data.protocol.name,"onUpdate:modelValue":f[3]||(f[3]=a=>_.data.protocol.name=a),placeholder:"TODO:改为下拉选择"},null,8,["modelValue"])])),_:1})])),_:1}),t(g,{title:"默认设备",type:"line"})])),_:1}),t(U,{type:"primary",onClick:v.save},{default:d((()=>[o("保存")])),_:1},8,["onClick"])])),_:1})}]]);export{V as default};