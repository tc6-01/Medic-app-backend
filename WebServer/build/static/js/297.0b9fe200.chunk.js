"use strict";(self.webpackChunkwebfile=self.webpackChunkwebfile||[]).push([[297],{8297:function(e,n,t){t.r(n),t.d(n,{default:function(){return B}});var i=t(885),r=t(8953),o=t(627),s=t(3767),a=t(2791),c=t(4005),l=t(2001),u=t(2479),f=t(2639),m=t(3456),d=t(9505),x=t(184),Z=function(e){var n=e.items,t=e.onOperationClicked;return(0,x.jsx)(d.Z,{items:n.map((function(e,n){return(0,x.jsx)(f.Z,{fileName:e.name,fileSize:(0,l.N)(e.size),fileTime:(0,u.u)(e.expire),remainUse:e.useLimit-e.use,sharedContent:(0,x.jsx)(m.Z,{state:e.state}),onOperationClicked:function(){return t(n)}},n)}))})},j=t(7442),p=t(3857),h=t(7209),k=t(3195),w=t(4554),g=t(9377),z=t(1186),C=t(4350),v=function(e){var n=e.fileName,t=e.onOperationClicked;return(0,x.jsx)(s.Z,{m:1,children:(0,x.jsxs)(w.Z,{borderRadius:3,children:[(0,x.jsx)(C.Z,{fileName:n}),(0,x.jsx)(j.Z,{title:"\u6587\u4ef6\u64cd\u4f5c",children:(0,x.jsx)(d.Z,{items:[(0,x.jsx)(z.Z,{icon:(0,x.jsx)(g.Z,{fontSize:"large"}),name:"\u6253\u5f00",onclick:function(){return t("open")}}),(0,x.jsx)(z.Z,{icon:(0,x.jsx)(p.Z,{fontSize:"large"}),name:"\u5171\u4eab",onclick:function(){return t("share")}}),(0,x.jsx)(z.Z,{icon:(0,x.jsx)(h.Z,{fontSize:"large"}),name:"\u5220\u9664",onclick:function(){return t("delete")}})]})}),(0,x.jsx)(j.Z,{title:"\u66f4\u591a\u64cd\u4f5c",children:(0,x.jsx)(d.Z,{items:[(0,x.jsx)(z.Z,{icon:(0,x.jsx)(k.Z,{fontSize:"large"}),name:"\u8be6\u7ec6\u4fe1\u606f",onclick:function(){return t("detail")}})]})})]})})},S=t(5320),b=((0,t(8499).Z)(a.createElement("path",{d:"M16 1H4c-1.1 0-2 .9-2 2v14h2V3h12V1zm-1 4l6 6v10c0 1.1-.9 2-2 2H7.99C6.89 23 6 22.1 6 21l.01-14c0-1.1.89-2 1.99-2h7zm-1 7h5.5L14 6.5V12z"}),"FileCopy"),t(3001)),O=t(6871),N=t(1315),D=function(){var e=(0,c.oR)(r.i),n=(0,O.s0)(),t=(b.Z,(0,a.useState)([])),l=(0,i.Z)(t,2),u=l[0],f=l[1];(0,a.useEffect)((function(){(0,N.fD)().then((function(e){console.log(e),f(e.data.map((function(e,n){var t=localStorage.getItem("username");return{id:n,name:e.fileName,size:e.fileSize,state:t==e.owner?"owned":"fromShared",owner:e.owner,use:e.use,useLimit:e.useLimit,expire:e.expire}})))}))}),[]);var m=(0,S.x)();return(0,x.jsx)(o.Z,{in:!0,children:(0,x.jsx)(s.Z,{className:m.root,children:(0,x.jsx)(j.Z,{title:"\u6240\u6709\u6587\u4ef6",children:(0,x.jsx)(Z,{items:u,onOperationClicked:function(t){e.setBottomDrawerOpen(!0),e.setBottomDrawerContent((0,x.jsx)(v,{fileName:u[t].name,onOperationClicked:function(i){return function(t,i){switch(console.log(i),t){case"delete":console.log("delete files at index",i);break;case"open":e.setBottomDrawerOpen(!1),n("/pdfpreview",{state:u[i]});break;case"share":e.setBottomDrawerOpen(!1),n("/policymanage/create",{state:{id:u[i].name}});break;case"detail":console.log("detail files at index",i)}}(i,t)}}))}})})})})},B=function(){return(0,x.jsx)(D,{})}},5320:function(e,n,t){t.d(n,{x:function(){return i}});var i=(0,t(8596).Z)({root:{"& > *":{userSelect:"none"}}})}}]);
//# sourceMappingURL=297.0b9fe200.chunk.js.map