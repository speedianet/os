"use strict";(self["webpackChunksos_dash"]=self["webpackChunksos_dash"]||[]).push([[437],{67437:(e,t,a)=>{a.r(t),a.d(t,{default:()=>ce});var s=a(59835),l=a(86970),r=a(28339),n=a(25121),o=a(60499),i=a(94629),u=a(85118),c=a(20503);const d=(0,s.aZ)({__name:"RefreshRateSelect",setup(e){const t=(0,u.V)(),a=new c.Z,l=[5,20,40,60,120,180].map((e=>({label:e+"s",value:e}))),n=(0,s.Fl)((()=>{const e=(0,r.yj)();return"/overview"===e.path})),d=(0,s.Fl)({get:()=>t.getSelectedRefreshRate,set:e=>t.setSelectedRefreshRate(e)});return(0,s.bv)((()=>{d.value=a.getRefreshRate()})),(0,s.YP)(d,(e=>{a.setRefreshRate(e)})),(e,t)=>n.value?((0,s.wg)(),(0,s.j4)(i.Z,{key:0,modelValue:d.value,"onUpdate:modelValue":t[0]||(t[0]=e=>d.value=e),options:(0,o.SU)(l),label:e.$t("refreshRateSelect.selectRefreshRate"),class:"refresh-rate-select"},null,8,["modelValue","options","label"])):(0,s.kq)("",!0)}}),m=d,p=m,g={class:"row justify-start items-center"},v={class:"title-h4 top-bar-text"},_=(0,s.aZ)({__name:"TopBar",setup(e){const t=(0,n.QT)().t;function a(){const e=(0,r.yj)();return e.meta.title?t(e.meta.title.toString()):""}const o=(0,s.Fl)((()=>{const e=(0,r.yj)();return e.meta.icon?e.meta.icon.toString():""}));return(e,t)=>{const r=(0,s.up)("q-icon"),n=(0,s.up)("q-space"),i=(0,s.up)("q-toolbar"),u=(0,s.up)("q-header");return(0,s.wg)(),(0,s.j4)(u,{class:"bg-transparent q-px-xs q-py-sm"},{default:(0,s.w5)((()=>[(0,s.Wm)(i,null,{default:(0,s.w5)((()=>[(0,s._)("div",g,[(0,s.Wm)(r,{color:"primary",name:o.value,size:"md",class:"q-mr-sm"},null,8,["name"]),(0,s._)("div",v,(0,l.zw)(a()),1)]),(0,s.Wm)(n),(0,s.Wm)(p)])),_:1})])),_:1})}}});var h=a(16602),f=a(51663),b=a(22857),y=a(90136),w=a(69984),x=a.n(w);const U=_,M=U;x()(_,"components",{QHeader:h.Z,QToolbar:f.Z,QIcon:b.Z,QSpace:y.Z});var q=a(37747);const W={class:"row justify-between items-center"},Z=(0,s.aZ)({__name:"ProfileCard",setup(e){function t(){const e=new q.Z;e.logout()}return(e,a)=>{const r=(0,s.up)("q-btn"),n=(0,s.up)("q-tooltip");return(0,s.wg)(),(0,s.iD)("div",W,[(0,s._)("div",null,[(0,s.Wm)(r,{dense:"",color:"primary",icon:"sym_s_settings",size:"sm",to:"/settings"}),(0,s.Wm)(n,{anchor:"bottom middle",class:"bg-primary",style:{"font-size":"14px"}},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.$t("profileCard.btnSettings")),1)])),_:1})]),(0,s._)("div",null,[(0,s.Wm)(r,{dense:"",color:"grey-7",icon:"sym_s_logout",size:"sm",onClick:a[0]||(a[0]=e=>t())},{default:(0,s.w5)((()=>[(0,s.Wm)(n,{anchor:"bottom middle",class:"bg-grey-7",style:{"font-size":"14px"}},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.$t("profileCard.btnLogout")),1)])),_:1})])),_:1})])])}}});var z=a(68879),k=a(46858);const Q=Z,I=Q;x()(Z,"components",{QBtn:z.Z,QTooltip:k.Z});var R=a(19302);const S={style:{"text-align":"center",padding:"0 8px"}},B=["src"],P={class:"column items-center"},j={class:"col",style:{"font-size":"12px"}},$=(0,s.aZ)({__name:"SideBarMenu",setup(e){const t=(0,n.QT)().t,a={right:"4px",borderRadius:"5px",backgroundColor:"#c97350",width:"5px",opacity:.75},i={right:"2px",borderRadius:"9px",backgroundColor:"#c97350",width:"9px",opacity:.2},u=(0,o.iH)(),c=[{title:t("sideBarMenu.apps"),icon:"sym_s_apps",path:"/apps",disabled:!0,isMenuItem:!0},{title:t("sideBarMenu.domains"),icon:"sym_s_language",path:"/domains",disabled:!0,isMenuItem:!0},{title:t("sideBarMenu.backups"),icon:"sym_s_backup",path:"/backups",disabled:!0,isMenuItem:!0},{title:t("sideBarMenu.webServers"),icon:"sym_s_web",path:"/web-servers",disabled:!0,isMenuItem:!0},{title:t("sideBarMenu.metrics"),icon:"sym_s_bar_chart",path:"/metrics",disabled:!0,isMenuItem:!0},{title:t("sideBarMenu.logs"),icon:"sym_s_receipt",path:"/logs",disabled:!0,isMenuItem:!1},{title:t("sideBarMenu.imageOptimizer"),icon:"sym_s_image",path:"/image-optimizer",disabled:!0,isMenuItem:!0},{title:t("sideBarMenu.security"),icon:"sym_s_security",path:"/security",disabled:!0,isMenuItem:!0},{title:t("sideBarMenu.terminal"),icon:"sym_s_terminal",path:"/terminal",disabled:!0,isMenuItem:!0}],d=(0,s.Fl)((()=>{const e=(0,R.Z)();return e.dark.isActive?"/_/assets/os-logo-dark.svg":"/_/assets/os-logo-light.svg"}));function m(){const e=(0,r.tv)().getRoutes();let a=[];return e.forEach((e=>{const s=e.children||[];s.forEach((e=>{var s,l,r,n,o,i,u,c,d,m;!1!==(null===(s=e.meta)||void 0===s?void 0:s.isMenuItem)&&a.push({title:null!==(r=t(`${null===(l=e.meta)||void 0===l?void 0:l.title}`))&&void 0!==r?r:"",icon:null!==(i=null===(o=null===(n=e.meta)||void 0===n?void 0:n.icon)||void 0===o?void 0:o.toString())&&void 0!==i?i:"",path:e.path,disabled:null!==(c=null===(u=e.meta)||void 0===u?void 0:u.disabled)&&void 0!==c&&c,useHref:null!==(m=null===(d=e.meta)||void 0===d?void 0:d.useHref)&&void 0!==m&&m})}))})),a.concat(c)}return(0,s.wF)((()=>{u.value=m()})),(e,r)=>{const n=(0,s.up)("router-link"),c=(0,s.up)("q-avatar"),m=(0,s.up)("q-tooltip"),p=(0,s.up)("q-item"),g=(0,s.up)("q-list"),v=(0,s.up)("q-scroll-area"),_=(0,s.up)("q-drawer"),h=(0,s.Q2)("ripple");return(0,s.wg)(),(0,s.j4)(_,{mini:!0,side:"left",bordered:"","show-if-above":"","mini-width":100},{default:(0,s.w5)((()=>[(0,s._)("div",S,[(0,s.Wm)(n,{to:"/"},{default:(0,s.w5)((()=>[(0,s._)("img",{src:d.value,alt:"Speedia OS",style:{"margin-top":"20px",height:"1.82rem"}},null,8,B)])),_:1})]),(0,s.Wm)(I,{class:"q-pt-md q-px-md"}),(0,s.Wm)(v,{"thumb-style":a,"bar-style":i,style:{height:"calc(100% - 100px)","margin-top":"10px"}},{default:(0,s.w5)((()=>[(0,s.Wm)(g,{style:{padding:"0 4px","text-align":"center","overflow-x":"hidden"}},{default:(0,s.w5)((()=>[((0,s.wg)(!0),(0,s.iD)(s.HY,null,(0,s.Ko)(u.value,((e,a)=>(0,s.wy)(((0,s.wg)(),(0,s.j4)(p,(0,s.dG)({key:a},e.useHref?{href:e.path}:{to:e.path},{disable:e.disabled,clickable:""}),{default:(0,s.w5)((()=>[(0,s._)("div",P,[(0,s.Wm)(c,{icon:e.icon,class:"icon-menu-bg"},null,8,["icon"]),(0,s._)("div",j,(0,l.zw)(e.title),1)]),e.disabled?((0,s.wg)(),(0,s.j4)(m,{key:0,anchor:"center end",class:"bg-primary",style:{"font-size":"14px"}},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)((0,o.SU)(t)("sideBarMenu.disabled")),1)])),_:1})):(0,s.kq)("",!0)])),_:2},1040,["disable"])),[[h]]))),128))])),_:1})])),_:1})])),_:1})}}});var F=a(10906),C=a(66663),H=a(13246),T=a(490),L=a(61357),V=a(51136);const A=$,D=A;x()($,"components",{QDrawer:F.Z,QScrollArea:C.Z,QList:H.Z,QItem:T.Z,QAvatar:L.Z,QTooltip:k.Z}),x()($,"directives",{Ripple:V.Z});var Y=a(61957),E=a(87178),O=a(16397);const G={class:"flex justify-end items-center"},K={class:"absolute-full flex flex-center"},J={class:"absolute-full flex flex-center"},N={class:"absolute-full flex flex-center"},X=(0,s.aZ)({__name:"FooterBar",setup(e){const t=(0,r.yj)(),a=(0,u.V)(),n=(0,O.n)(),i=new E.Z,c=(0,o.iH)(),d=(0,s.Fl)({get:()=>n.getSystemInfo,set:e=>n.setSystemInfo(e)}),m=(0,s.Fl)((()=>a.getSelectedRefreshRate)),p=(0,s.Fl)((()=>t.path));function g(e){return e<50?"green":e<80?"orange":"red"}function v(){c.value&&clearInterval(c.value),c.value=setInterval((()=>{i.getSystemInfo().then((e=>{d.value=e.data.body})).catch((e=>{console.error(e)}))}),1e3*m.value)}return(0,s.Ah)((()=>{clearInterval(c.value)})),(0,s.YP)(p,(()=>{clearInterval(c.value),"/overview"!==p.value&&v()})),(e,t)=>{const a=(0,s.up)("q-tooltip"),r=(0,s.up)("q-icon"),n=(0,s.up)("q-badge"),o=(0,s.up)("q-linear-progress"),i=(0,s.up)("q-footer");return(0,s.wg)(),(0,s.j4)(i,{bordered:"",class:"q-px-lg",style:{"min-height":"30px"}},{default:(0,s.w5)((()=>[(0,s.wy)((0,s._)("div",G,[(0,s.Wm)(r,{name:"sym_s_terminal",size:"1.618rem",class:"disabled q-mr-md"},{default:(0,s.w5)((()=>[(0,s.Wm)(a,{class:"bg-primary",style:{"font-size":"14px"}},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.$t("footerBar.disabled")),1)])),_:1})])),_:1}),(0,s.Wm)(r,{name:"sym_s_speed",size:"sm",class:"q-mr-xs"}),(0,s.Wm)(o,{stripe:"",rounded:"",size:"20px",class:"q-mr-md",value:Math.trunc(d.value.currentUsage.cpuUsagePercent)/100,color:g(Math.trunc(d.value.currentUsage.cpuUsagePercent)),label:`${Math.trunc(d.value.currentUsage.cpuUsagePercent)}%`,style:{width:"100px"}},{default:(0,s.w5)((()=>[(0,s._)("div",K,[(0,s.Wm)(n,{color:"white","text-color":"grey-10",label:`${Math.trunc(d.value.currentUsage.cpuUsagePercent)}%`},null,8,["label"])]),(0,s.Wm)(a,{class:"bg-primary",style:{"font-size":"14px"}},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.$t("footerBar.cpuUsage",{cpuUsage:Math.trunc(d.value.currentUsage.cpuUsagePercent)})),1)])),_:1})])),_:1},8,["value","color","label"]),(0,s.Wm)(r,{name:"sym_s_memory",size:"sm",class:"q-mr-xs"}),(0,s.Wm)(o,{stripe:"",rounded:"",class:"q-mr-md",size:"20px",value:Math.trunc(d.value.currentUsage.memUsagePercent)/100,color:g(Math.trunc(d.value.currentUsage.memUsagePercent)),label:`${Math.trunc(d.value.currentUsage.memUsagePercent)}%`,style:{width:"100px"}},{default:(0,s.w5)((()=>[(0,s._)("div",J,[(0,s.Wm)(n,{color:"white","text-color":"grey-10",label:`${Math.trunc(d.value.currentUsage.memUsagePercent)}%`},null,8,["label"])]),(0,s.Wm)(a,{class:"bg-primary",style:{"font-size":"14px"}},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.$t("footerBar.ramUsage",{ramUsage:Math.trunc(d.value.currentUsage.memUsagePercent)})),1)])),_:1})])),_:1},8,["value","color","label"]),(0,s.Wm)(r,{name:"sym_s_sd_card",size:"sm",class:"q-mr-xs"}),(0,s.Wm)(o,{stripe:"",rounded:"",size:"20px",value:Math.trunc(d.value.currentUsage.storageUsage)/100,color:g(Math.trunc(d.value.currentUsage.storageUsage)),label:`${Math.trunc(d.value.currentUsage.storageUsage)}%`,style:{width:"100px"}},{default:(0,s.w5)((()=>[(0,s._)("div",N,[(0,s.Wm)(n,{color:"white","text-color":"grey-10",label:`${Math.trunc(d.value.currentUsage.storageUsage)}%`},null,8,["label"])]),(0,s.Wm)(a,{class:"bg-primary",style:{"font-size":"14px"}},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.$t("footerBar.storageUsage",{storageUsage:Math.trunc(d.value.currentUsage.storageUsage)})),1)])),_:1})])),_:1},8,["value","color","label"])],512),[[Y.F8,d.value.hostname]])])),_:1})}}});var ee=a(11639),te=a(71378),ae=a(8289),se=a(20990);const le=(0,ee.Z)(X,[["__scopeId","data-v-12090e64"]]),re=le;x()(X,"components",{QFooter:te.Z,QIcon:b.Z,QTooltip:k.Z,QLinearProgress:ae.Z,QBadge:se.Z});const ne=(0,s.aZ)({__name:"MainLayout",setup(e){return(e,t)=>{const a=(0,s.up)("router-view"),l=(0,s.up)("q-page-container"),r=(0,s.up)("q-layout");return(0,s.wg)(),(0,s.j4)(r,{view:"lhh LpR lFf"},{default:(0,s.w5)((()=>[(0,s.Wm)(M),(0,s.Wm)(D),(0,s.Wm)(l,null,{default:(0,s.w5)((()=>[(0,s.Wm)(a)])),_:1}),(0,s.Wm)(re)])),_:1})}}});var oe=a(20249),ie=a(12133);const ue=ne,ce=ue;x()(ne,"components",{QLayout:oe.Z,QPageContainer:ie.Z})}}]);