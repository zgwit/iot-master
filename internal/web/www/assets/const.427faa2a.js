var r=Object.defineProperty,e=Object.defineProperties,t=Object.getOwnPropertyDescriptors,o=Object.getOwnPropertySymbols,a=Object.prototype.hasOwnProperty,c=Object.prototype.propertyIsEnumerable,s=(e,t,o)=>t in e?r(e,t,{enumerable:!0,configurable:!0,writable:!0,value:o}):e[t]=o;import{y as i,p as n}from"./index.945c150b.js";function p(r){var p,l;i((p=((r,e)=>{for(var t in e||(e={}))a.call(e,t)&&s(r,t,e[t]);if(o)for(var t of o(e))c.call(e,t)&&s(r,t,e[t]);return r})({},r),l={url:"http://localhost:8080/api/"+r.url,success(e){const t=e.data;if(t.error)return r.error&&r.error(t.error),void n({icon:"error",title:t.error});r.success&&r.success(t.data)}},e(p,t(l))))}export{p as r};