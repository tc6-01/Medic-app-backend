"use strict";(self.webpackChunkwebfile=self.webpackChunkwebfile||[]).push([[878],{3878:function(e,n,i){i.r(n),i.d(n,{default:function(){return B}});var t=i(2791),s=i(6871),r=i(2639),o=i(9505),l=i(4005),a=i(1413),c=i(885),f=i(8830),u=i(1914);function d(){var e=(0,l.oR)(u.v),n=(0,l.oR)(f.c),i=(0,t.useState)([]),s=(0,c.Z)(i,2),r=s[0],o=s[1];return(0,a.Z)((0,a.Z)((0,a.Z)({},e),n),{},{files:r,setFiles:o})}var x=i(3456),Z=i(8953),j=i(7442),m=i(4350),h=i(3767),k=i(4554),p=i(1186),v=i(9377),w=i(7209),b=i(3195),g=i(184),C=function(e){var n=e.fileName,i=e.onOperationClicked;return(0,g.jsx)(h.Z,{m:1,children:(0,g.jsxs)(k.Z,{borderRadius:3,children:[(0,g.jsx)(m.Z,{fileName:n}),(0,g.jsx)(j.Z,{title:"\u6587\u4ef6\u64cd\u4f5c",children:(0,g.jsx)(o.Z,{items:[(0,g.jsx)(p.Z,{icon:(0,g.jsx)(v.Z,{fontSize:"large"}),name:"\u6253\u5f00",onclick:function(){return i("open")}}),(0,g.jsx)(p.Z,{icon:(0,g.jsx)(w.Z,{fontSize:"large"}),name:"\u5220\u9664",onclick:function(){return i("delete")}})]})}),(0,g.jsx)(j.Z,{title:"\u66f4\u591a\u64cd\u4f5c",children:(0,g.jsx)(o.Z,{items:[(0,g.jsx)(p.Z,{icon:(0,g.jsx)(b.Z,{fontSize:"large"}),name:"\u8be6\u7ec6\u4fe1\u606f",onclick:function(){return i("detail")}})]})})]})})},N=i(627),z=i(1315),R=i(2479),S=i(2001),O=function(){var e,n=(0,s.s0)(),i=(0,l.oR)(d),a=(0,l.oR)(Z.i);(0,t.useEffect)((function(){(0,z.v3)().then((function(e){e&&0===e.code&&(Array.isArray(e.data)?(i.setFiles(e.data),i.setLoading(!1)):i.setFiles([e.data]))}))}),[]);var c=function(e){a.setBottomDrawerOpen(!0),a.setBottomDrawerContent((0,g.jsx)(C,{fileName:i.files[e].fileName,onOperationClicked:function(t){return function(e,t){switch(e){case"delete":console.log("delete files at index",t);break;case"open":a.setBottomDrawerOpen(!1),console.log(i.files),n("/pdfpreview",{state:(0,z.RX)(t,i.files[t])});break;case"share":console.log("share files at index",t);break;case"detail":console.log("detail files at index",t)}}(t,e)}}))};return(0,g.jsx)(N.Z,{in:!0,children:(0,g.jsx)(k.Z,{children:(0,g.jsx)(o.Z,{items:null===(e=i.files)||void 0===e?void 0:e.map((function(e,n){return(0,g.jsx)(r.Z,{fileName:e.fileName,fileTime:(0,R.u)(e.expire),fileSize:(0,S.N)(e.fileSize),remainUse:e.useLimit-e.use,sharedContent:(0,g.jsx)(x.Z,{state:"shared"}),onOperationClicked:function(){return c(n)}},n)})),onItemClicked:c})})})},B=function(){return(0,g.jsx)(l.zt,{of:d,children:(0,g.jsx)(O,{})})}}}]);
//# sourceMappingURL=878.67e1d77e.chunk.js.map