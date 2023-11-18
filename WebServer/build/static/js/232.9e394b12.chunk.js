"use strict";(self.webpackChunkwebfile=self.webpackChunkwebfile||[]).push([[232],{7006:function(e,t,n){n.d(t,{lq:function(){return f},D_:function(){return m},oi:function(){return h}});var r=n(8550),a=n(8096),i=n(4925),o=n(5594),u=n(5022),l=n(6934),c=n(9955),s=n(5523),d=(n(2791),n(184)),p=n(4454),h=function(e){var t=e.id,n=e.name,a=e.defaultValue,i=e.placeHolder,o=e.onChange;return(0,d.jsx)(r.Z,{id:t,fullWidth:!0,label:n,placeholder:i,defaultValue:a,onChange:o})},f=function(e){var t=e.id,n=e.name,r=e.options,l=e.defaultValue,c=e.placeHolder,s=e.onChange;return(0,d.jsxs)(a.Z,{fullWidth:!0,children:[(0,d.jsx)(i.Z,{id:"form-layouts-separator-select-label",children:n}),(0,d.jsx)(o.Z,{id:t,label:n,defaultValue:l,placeholder:c,labelId:"form-layouts-separator-select-label",onChange:s,children:r.map((function(e,t){return(0,d.jsx)(u.Z,{value:e.id,children:e.name},t)}))})]})},m=((0,l.ZP)(c.Z)((function(e){var t=e.theme;return{padding:8,"& .MuiSwitch-track":{borderRadius:11,"&:before, &:after":{content:'""',position:"absolute",top:"50%",transform:"translateY(-50%)",width:16,height:16},"&:before":{backgroundImage:'url(\'data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" height="16" width="16" viewBox="0 0 24 24"><path fill="'.concat(encodeURIComponent(t.palette.getContrastText(t.palette.primary.main)),'" d="M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z"/></svg>\')'),left:12},"&:after":{backgroundImage:'url(\'data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" height="16" width="16" viewBox="0 0 24 24"><path fill="'.concat(encodeURIComponent(t.palette.getContrastText(t.palette.primary.main)),'" d="M19,13H5V11H19V13Z" /></svg>\')'),right:12}},"& .MuiSwitch-thumb":{boxShadow:"none",width:16,height:16,margin:2}}})),function(e){var t=e.id,n=e.name,r=e.onChange,a=e.checked;return(0,d.jsx)(s.Z,{id:t,control:(0,d.jsx)(p.Z,{defaultChecked:!0}),label:n,labelPlacement:"start",checked:a,onChange:r})})},232:function(e,t,n){n.r(t),n.d(t,{default:function(){return L}});var r=n(885),a=n(6459),i=n(627),o=n(3767),u=n(6151),l=n(4005),c=n(2791),s=n(5646),d=n(9658),p=n(184),h=function(e){var t=e.content,n=e.duration,a=void 0===n?3e3:n,i=e.type,o=c.useState(!0),u=(0,r.Z)(o,2),l=u[0],h=u[1],f=function(){h(!1)};return(0,p.jsx)(s.Z,{open:l,autoHideDuration:a,anchorOrigin:{vertical:"top",horizontal:"center"},onClose:f,children:(0,p.jsx)(d.Z,{onClose:f,severity:i,children:t})})},f=n(4164),m={success:function(e,t){var n=document.createElement("div"),r=(0,p.jsx)(h,{content:e,duration:t,type:"success"});f.render(r,n),document.body.appendChild(n)},error:function(e,t){var n=document.createElement("div"),r=(0,p.jsx)(h,{content:e,duration:t,type:"error"});f.render(r,n),document.body.appendChild(n)},warning:function(e,t){var n=document.createElement("div"),r=(0,p.jsx)(h,{content:e,duration:t,type:"warning"});f.render(r,n),document.body.appendChild(n)},info:function(e,t){var n=document.createElement("div"),r=(0,p.jsx)(h,{content:e,duration:t,type:"info"});f.render(r,n),document.body.appendChild(n)}},g=n(4030),v=n(1413),y=n(1889),x=n(8550),b=n(890),w=n(3241),Z=n(1562),k=n(6571),C=n(7006),j=function(e){var t=e.items;return(0,p.jsx)(y.ZP,{container:!0,direction:"row",children:t.map((function(e,t){return(0,p.jsx)(y.ZP,{item:!0,xs:6,children:(0,p.jsx)(C.D_,{id:e.id,name:e.name,enabled:e.enabled,checked:e.checked,onChange:e.onChange},t)})}))})},D=function(e){var t=e.items;return(0,p.jsx)(o.Z,{spacing:3,children:null===t||void 0===t?void 0:t.map((function(e,t){switch(e.type){case"text":var n=e.value;return(0,p.jsx)(y.ZP,{item:!0,display:e.display?"block":"none",children:(0,p.jsx)(C.oi,{id:n.id,name:n.name,defaultValue:n.defaultValue,placeHolder:n.placeHolder,onChange:n.onChange},t)},t);case"select":var r=e.value;return(0,p.jsx)(y.ZP,{item:!0,display:e.display?"block":"none",children:(0,p.jsx)(C.lq,{id:r.id,name:r.name,options:r.options,defaultValue:r.defaultValue,placeHolder:r.placeHolder,onChange:r.onChange},t)},t);case"switch":var a=e.value;return(0,p.jsx)(y.ZP,{item:!0,display:e.display?"block":"none",children:(0,p.jsx)(C.D_,{id:a.id,name:a.name,enabled:a.enabled,checked:a.checked,onChange:a.onChange},t)},t);case"title":return(0,p.jsx)(o.Z,{children:(0,p.jsx)(b.Z,{variant:"h4",color:"textPrimary",children:"\u6dfb\u52a0\u89c4\u5219"})},t);case"control":var i=e.value;return(0,p.jsx)(j,{items:i.items});case"time":var u=e.value;return(0,p.jsx)(o.Z,{direction:"column",display:e.display?"flex":"none",children:(0,p.jsx)(k._,{dateAdapter:w.H,children:(0,p.jsx)(Z.x,{label:u.title,value:u.value,onChange:u.onChange,renderInput:function(e){return(0,p.jsx)(x.Z,(0,v.Z)({},e))}})})},t);default:return(0,p.jsx)(p.Fragment,{})}}))})},I=n(6871),H=n(1315),S=function(e){(0,a.Z)(e);var t=(0,l.oR)(g.N),n=(0,I.s0)(),s=(0,I.TH)().state.id,d=(0,c.useState)([]),h=(0,r.Z)(d,2),f=h[0],v=h[1];(0,c.useEffect)((function(){(0,H.lE)().then((function(e){e&&0===e.code?Array.isArray(e.data)&&v(e.data.map((function(e,t){return{id:t.toString(),name:e}}))):m.error("\u83b7\u53d6\u7528\u6237\u5217\u8868\u5931\u8d25")})),t.setTo(new Date)}),[]),console.log(t.role);var y=function(e){switch(e.target.id){case"policyName":t.setPolicyName(e.target.value);break;case"des":t.setDes(e.target.value);break;case"userId":var n=e.target.value.split(/[\uff1b;]/);t.setUserId(n);break;case"url":var r=e.target.value.split(/[\uff1b;]/);t.setUrl(r);break;case"count":t.setCount(parseInt(e.target.value));break;case"period":t.setPeriod(e.target.value);break;case"role":t.setRole(e.target.value)}},x={items:[{type:"text",display:!0,value:{id:"des",name:"\u7b56\u7565\u63cf\u8ff0",placeHolder:"\u8bf7\u8f93\u5165\u5171\u4eab\u63cf\u8ff0",onChange:y}},{type:"title"},{type:"control",value:{items:[{name:"\u9650\u5236\u8bbf\u95ee\u8d26\u53f7",enabled:!0,checked:t.roleDisplay,onChange:function(e){t.setRoleDisplay(e.target.checked)}},{name:"\u9650\u5236\u4f7f\u7528\u6b21\u6570",enabled:!0,checked:t.countDisplay,onChange:function(e){t.setCountDisplay(e.target.checked)}},{name:"\u9650\u5236\u4f7f\u7528\u65f6\u95f4",enabled:!0,checked:t.timeDisplay,onChange:function(e){t.setTimeDisplay(e.target.checked)}}]}},{type:"select",display:t.roleDisplay,value:{id:"role",defaultValue:"",name:"\u7528\u6237\u89d2\u8272",placeHolder:"\u8bf7\u9009\u62e9\u7528\u6237\u89d2\u8272",options:f,onChange:function(e){t.setRole(e.target.value)}}},{type:"text",display:t.userIdDisplay,value:{id:"userId",name:"\u4f7f\u7528\u7528\u6237id",placeHolder:"\u8bf7\u8f93\u5165\u4f7f\u7528\u7528\u6237id\uff0c\u591a\u4e2aid\u7528\u82f1\u6587\u5206\u53f7;\u5206\u5272",onChange:y}},{type:"text",display:t.urlDisplay,value:{id:"url",name:"\u8bbe\u5907url",placeHolder:"\u8bf7\u8f93\u5165\u8bbe\u5907url\uff0c\u591a\u4e2a\u8bbe\u5907url\u7528\u82f1\u6587\u5206\u53f7;\u5206\u5272",onChange:y}},{type:"text",display:t.countDisplay,value:{id:"count",name:"\u7528\u6237\u53ef\u8bbf\u95ee\u6b21\u6570",placeHolder:"\u8bf7\u8f93\u5165\u7528\u6237\u53ef\u8bbf\u95ee\u6b21\u6570",onChange:y}},{type:"text",display:t.peroidDisplay,value:{id:"period",name:"\u7528\u6237\u53ef\u8bbf\u95ee\u7684\u9891\u7387(\u6b21/\u79d2)",placeHolder:"\u8bf7\u8f93\u5165\u7528\u6237\u53ef\u8bbf\u95ee\u7684\u9891\u7387(\u6b21/\u79d2)",onChange:y}},{type:"time",display:t.timeDisplay,value:{id:"to",value:t.to,title:"\u5171\u4eab\u7ed3\u675f\u65f6\u95f4",onChange:function(e){t.setTo(e)}}},{type:"switch",display:t.timeDisplay,value:{id:"delete",checked:t.checked,name:"\u5230\u65f6\u95f4\u662f\u5426\u5220\u9664",enabled:!0,onChange:function(e){t.setCheacked(e.target.checked)}}}]};return(0,p.jsx)(i.Z,{in:!0,children:(0,p.jsxs)(o.Z,{mt:2,mb:2,spacing:4,children:[(0,p.jsx)(D,{items:x.items}),(0,p.jsx)(u.Z,{fullWidth:!0,variant:"contained",sx:{minHeight:50},onClick:function(){if(0!=t.des.length){var e=[];0!==t.userId.length&&t.userIdDisplay&&e.push({type:"user",value:t.userId}),0!==t.url.length&&t.urlDisplay&&e.push({type:"target",value:t.url}),0!==t.count&&t.countDisplay&&e.push({type:"count",value:t.count}),0!==t.period.length&&t.peroidDisplay&&e.push({type:"peroid",value:t.period}),null!=t.from&&t.timeDisplay&&e.push({type:"useTime",value:{from:t.from.getTime(),to:t.to.getTime(),delete:t.checked}}),0!==t.role.length&&t.roleDisplay&&e.push({type:"role",value:t.role}),(0,H.u5)({fileName:s,target:f[t.role].name,expire:t.to.getTime(),useLimit:t.count,isGroup:!0}).then((function(e){0==e.code?(m.success("\u6dfb\u52a0\u6210\u529f"),n("/filemanage")):m.error("\u8bf7\u6c42\u5931\u8d25"+e.message)})).catch((function(e){m.error("\u9519\u8bef"+e)}))}else m.warning("\u63cf\u8ff0\u4e0d\u80fd\u4e3a\u7a7a\uff01")},children:"\u786e\u8ba4\u6dfb\u52a0"})]})})},L=function(){return(0,p.jsx)(S,{})}},1315:function(e,t,n){n.d(t,{Cl:function(){return u},RX:function(){return l},Sv:function(){return f},fD:function(){return s},lE:function(){return d},u5:function(){return p},v3:function(){return h},x4:function(){return c}});var r=n(4165),a=n(5861),i=n(4569),o=n.n(i);function u(e,t){var n=localStorage.getItem("username");return{id:e,name:t.fileName,size:t.fileSize,state:n==t.owner?"owned":"fromShared",owner:t.owner,use:t.use,useLimit:t.useLimit,expire:t.expire}}function l(e,t){var n=localStorage.getItem("username");return{id:e,name:t.fileName,size:0,state:n==t.target?"owned":"fromShared",owner:t.target,use:t.use,useLimit:t.useLimit,expire:t.expire}}var c=function(){var e=(0,a.Z)((0,r.Z)().mark((function e(t,n){var a;return(0,r.Z)().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,o().post("/login",{userName:t,passWord:n});case 2:return a=e.sent,e.abrupt("return",a.data);case 4:case"end":return e.stop()}}),e)})));return function(t,n){return e.apply(this,arguments)}}(),s=function(){var e=(0,a.Z)((0,r.Z)().mark((function e(){var t;return(0,r.Z)().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,o().get("/file",{headers:{Authorization:"Bearer ".concat(localStorage.getItem("token"))}});case 2:return t=e.sent,e.abrupt("return",t.data);case 4:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),d=function(){var e=(0,a.Z)((0,r.Z)().mark((function e(){var t;return(0,r.Z)().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,o().get("/user",{headers:{Authorization:"Bearer ".concat(localStorage.getItem("token"))}});case 2:return t=e.sent,e.abrupt("return",t.data);case 4:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),p=function(){var e=(0,a.Z)((0,r.Z)().mark((function e(t){var n;return(0,r.Z)().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,o().post("/file/share",t,{headers:{Authorization:"Bearer ".concat(localStorage.getItem("token"))}});case 2:return n=e.sent,e.abrupt("return",n.data);case 4:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),h=function(){var e=(0,a.Z)((0,r.Z)().mark((function e(){var t;return(0,r.Z)().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,o().get("/file/share",{headers:{Authorization:"Bearer ".concat(localStorage.getItem("token"))}});case 2:return t=e.sent,e.abrupt("return",t.data);case 4:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),f=function(){var e=(0,a.Z)((0,r.Z)().mark((function e(t){var n;return(0,r.Z)().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,o().get("/file/download?filename=".concat(t),{responseType:"blob",headers:{Authorization:"Bearer ".concat(localStorage.getItem("token"))}});case 2:return n=e.sent,e.abrupt("return",n.data);case 4:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}()}}]);
//# sourceMappingURL=232.9e394b12.chunk.js.map