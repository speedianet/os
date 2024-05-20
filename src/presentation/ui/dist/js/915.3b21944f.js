"use strict";(self["webpackChunksos_dash"]=self["webpackChunksos_dash"]||[]).push([[915],{60884:(e,a,l)=>{l.r(a),l.d(a,{default:()=>pa});var o=l(59835),t=l(60499),n=l(86970),s=l(3746);const r=(0,s.Q_)("crons",{state:()=>({keyCronsTable:0,keyCronsForm:0,formType:"",showDialogForm:!1,showDialogRemove:!1,selectedCron:{}}),getters:{getKeyCronsTable(e){return e.keyCronsTable},getKeyCronsForm(e){return e.keyCronsForm},getFormType(e){return e.formType},getShowDialogForm(e){return e.showDialogForm},getShowDialogRemove(e){return e.showDialogRemove},getSelectedCron(e){return e.selectedCron}},actions:{setKeyCronsTable(e){this.keyCronsTable=e},setKeyCronsForm(e){this.keyCronsForm=e},setFormType(e){this.formType=e},setShowDialogForm(e){this.showDialogForm=e},setShowDialogRemove(e){this.showDialogRemove=e},setSelectedCron(e){this.selectedCron=e}}}),i=(0,o.aZ)({__name:"CronsTableActions",props:{selectedCronJob:{}},setup(e){const a=e,l=r(),t=(0,o.Fl)({get:()=>l.getShowDialogForm,set:e=>l.setShowDialogForm(e)}),s=(0,o.Fl)({get:()=>l.getShowDialogRemove,set:e=>l.setShowDialogRemove(e)}),i=(0,o.Fl)({get:()=>l.getSelectedCron,set:e=>l.setSelectedCron(e)}),u=(0,o.Fl)({get:()=>l.getFormType,set:e=>l.setFormType(e)});function c(){u.value="update",i.value=a.selectedCronJob,t.value=!0}function d(){i.value=a.selectedCronJob,s.value=!0}return(e,a)=>{const l=(0,o.up)("q-tooltip"),t=(0,o.up)("q-btn");return(0,o.wg)(),(0,o.iD)("div",null,[(0,o.Wm)(t,{color:"primary",size:"md",icon:"sym_s_edit",class:"q-ma-md",dense:"",onClick:a[0]||(a[0]=e=>c())},{default:(0,o.w5)((()=>[(0,o.Wm)(l,{class:"bg-primary",style:{"font-size":"14px"},offset:[10,10]},{default:(0,o.w5)((()=>[(0,o.Uk)((0,n.zw)(e.$t("cronsTableActions.btnEditCronjob")),1)])),_:1})])),_:1}),(0,o.Wm)(t,{dense:"",color:"negative",size:"md",icon:"sym_s_delete",onClick:a[1]||(a[1]=e=>d())},{default:(0,o.w5)((()=>[(0,o.Wm)(l,{class:"bg-negative",style:{"font-size":"14px"},offset:[10,10]},{default:(0,o.w5)((()=>[(0,o.Uk)((0,n.zw)(e.$t("cronsTableActions.btnRemoveCronjob")),1)])),_:1})])),_:1})])}}});var u=l(68879),c=l(46858),d=l(69984),m=l.n(d);const p=i,v=p;m()(i,"components",{QBtn:u.Z,QTooltip:c.Z});var h=l(84278),b=l(25121);const y={class:"row justify-center q-mt-md"},g=(0,o.aZ)({__name:"CronsTable",props:{data:{}},setup(e){const a=e,l=(0,b.QT)().t,s=(0,t.iH)(""),i=(0,t.iH)({sortBy:"desc",descending:!1,page:1,rowsPerPage:10}),u=(0,t.iH)([{name:"schedule",label:l("cronsTable.columnSchedule"),align:"left",field:"schedule",classes:"td-main-table",headerClasses:"bg-primary text-white",sortable:!0},{name:"command",label:l("cronsTable.columnCommand"),align:"left",field:"command",style:"width: 50%;",classes:"cron-table-td-command",sortable:!0},{name:"comment",label:l("cronsTable.columnComment"),align:"left",field:"comment",sortable:!0}]),c=(0,o.Fl)((()=>Math.ceil(a.data.length/i.value.rowsPerPage)));function d(){const e=r();e.setKeyCronsForm(e.getKeyCronsForm+1),e.setFormType("create"),e.setShowDialogForm(!0)}return(e,a)=>{const l=(0,o.up)("q-icon"),t=(0,o.up)("q-input"),r=(0,o.up)("q-space"),m=(0,o.up)("q-th"),p=(0,o.up)("q-tr"),b=(0,o.up)("q-td"),g=(0,o.up)("q-table"),w=(0,o.up)("q-pagination");return(0,o.wg)(),(0,o.iD)("div",null,[(0,o.Wm)(g,{rows:e.data,columns:u.value,filter:s.value,pagination:i.value,"onUpdate:pagination":a[2]||(a[2]=e=>i.value=e),"no-data-label":e.$t("cronsTable.notFoundCronjobs"),"row-key":"key",color:"primary",flat:"",bordered:"","hide-pagination":""},{top:(0,o.w5)((()=>[(0,o.Wm)(t,{borderless:"",debounce:"300",color:"primary",modelValue:s.value,"onUpdate:modelValue":a[0]||(a[0]=e=>s.value=e),label:e.$t("cronsTable.inputSearch")},{prepend:(0,o.w5)((()=>[(0,o.Wm)(l,{name:"sym_s_search"})])),_:1},8,["modelValue","label"]),(0,o.Wm)(r),(0,o.Wm)(h.Z,{label:e.$t("cronsTable.btnNewTask"),onClick:a[1]||(a[1]=e=>d()),icon:"sym_s_schedule"},null,8,["label"])])),header:(0,o.w5)((e=>[(0,o.Wm)(p,{props:e},{default:(0,o.w5)((()=>[((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(e.cols,(a=>((0,o.wg)(),(0,o.j4)(m,{key:a.name,props:e,style:{"font-weight":"bold","font-size":"14px"}},{default:(0,o.w5)((()=>[(0,o.Uk)((0,n.zw)(a.label),1)])),_:2},1032,["props"])))),128)),(0,o.Wm)(m)])),_:2},1032,["props"])])),body:(0,o.w5)((e=>[(0,o.Wm)(p,{props:e},{default:(0,o.w5)((()=>[((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(e.cols,(a=>((0,o.wg)(),(0,o.j4)(b,{key:a.name,props:e},{default:(0,o.w5)((()=>[(0,o.Uk)((0,n.zw)(a.value),1)])),_:2},1032,["props"])))),128)),(0,o.Wm)(b,{class:"text-right"},{default:(0,o.w5)((()=>[(0,o.Wm)(v,{selectedCronJob:e.row},null,8,["selectedCronJob"])])),_:2},1024)])),_:2},1032,["props"])])),_:1},8,["rows","columns","filter","pagination","no-data-label"]),(0,o._)("div",y,[(0,o.Wm)(w,{modelValue:i.value.page,"onUpdate:modelValue":a[3]||(a[3]=e=>i.value.page=e),color:"primary",max:c.value,size:"md"},null,8,["modelValue","max"])])])}}});var w=l(84277),$=l(13119),f=l(22857),S=l(90136),D=l(31233),_=l(21682),C=l(67220),F=l(80996);const O=g,k=O;m()(g,"components",{QTable:w.Z,QInput:$.Z,QIcon:f.Z,QSpace:S.Z,QTr:D.Z,QTh:_.Z,QTd:C.Z,QPagination:F.Z});var q={class:"q-pa-md"},W={class:"row"},H={class:"text-left",width:"75px"},Z={class:"text-center",width:"250px"},T=(0,o._)("th",{class:"text-left",width:"150px"},null,-1),x={class:"text-center"},Q={class:"text-center"},V={class:"text-right"};function U(e,a){var l=(0,o.up)("q-skeleton"),t=(0,o.up)("q-space"),n=(0,o.up)("q-markup-table"),s=(0,o.up)("q-card-section"),r=(0,o.up)("q-card");return(0,o.wg)(),(0,o.iD)("div",q,[(0,o.Wm)(r,{flat:""},{default:(0,o.w5)((function(){return[(0,o.Wm)(s,null,{default:(0,o.w5)((function(){return[(0,o._)("div",W,[(0,o.Wm)(l,{type:"QInput",width:"400px"}),(0,o.Wm)(t),(0,o.Wm)(l,{type:"QInput",width:"300px"})]),(0,o.Wm)(n,{flat:""},{default:(0,o.w5)((function(){return[(0,o._)("thead",null,[(0,o._)("tr",null,[(0,o._)("th",H,[(0,o.Wm)(l,{animation:"blink",type:"text"})]),(0,o._)("th",Z,[(0,o.Wm)(l,{animation:"blink",type:"text"})]),T])]),(0,o._)("tbody",null,[((0,o.wg)(),(0,o.iD)(o.HY,null,(0,o.Ko)(10,(function(e){return(0,o._)("tr",{key:e},[(0,o._)("td",x,[(0,o.Wm)(l,{animation:"blink",type:"text",width:"75px"})]),(0,o._)("td",Q,[(0,o.Wm)(l,{animation:"blink",type:"text",width:"250px"})]),(0,o._)("td",V,[(0,o.Wm)(l,{animation:"blink",type:"circle",class:"float-right q-mx-sm",size:"30px"}),(0,o.Wm)(l,{animation:"blink",type:"circle",class:"float-right",size:"30px"})])])})),64))])]})),_:1})]})),_:1})]})),_:1})])}var j=l(11639),I=l(44458),R=l(63190),M=l(57133),P=l(66933);const E={},z=(0,j.Z)(E,[["render",U]]),K=z;m()(E,"components",{QCard:I.Z,QCardSection:R.Z,QSkeleton:M.ZP,QSpace:S.Z,QMarkupTable:P.Z});var Y=l(15521),L=l(74462);class A{constructor(){this.$i18n=L.i18n.global.t,this.optionsMinutes=[{label:this.$i18n("cronOptionsSelects.everyMinute"),value:"*"},{label:this.$i18n("cronOptionsSelects.everyMinutes",{minute:5}),value:"*/5"},{label:this.$i18n("cronOptionsSelects.everyMinutes",{minute:10}),value:"*/10"},{label:this.$i18n("cronOptionsSelects.everyMinutes",{minute:15}),value:"*/15"},{label:this.$i18n("cronOptionsSelects.everyMinutes",{minute:30}),value:"*/30"},{label:this.$i18n("cronOptionsSelects.everyMinutes",{minute:45}),value:"*/45"},{label:this.$i18n("cronOptionsSelects.atMinute",{minute:0}),value:"0"},{label:this.$i18n("cronOptionsSelects.atMinute",{minute:5}),value:"5"},{label:this.$i18n("cronOptionsSelects.atMinute",{minute:10}),value:"10"},{label:this.$i18n("cronOptionsSelects.atMinute",{minute:15}),value:"15"},{label:this.$i18n("cronOptionsSelects.atMinute",{minute:30}),value:"30"},{label:this.$i18n("cronOptionsSelects.atMinute",{minute:45}),value:"45"}],this.optionsHours=[{label:this.$i18n("cronOptionsSelects.everyHour"),value:"*"},{label:this.$i18n("cronOptionsSelects.everyHours",{hour:2}),value:"*/2"},{label:this.$i18n("cronOptionsSelects.everyHours",{hour:6}),value:"*/6"},{label:this.$i18n("cronOptionsSelects.everyHours",{hour:12}),value:"*/12"},{label:this.$i18n("cronOptionsSelects.atMidnight"),value:"0"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:1,standard12hours:"1am"}),value:"1"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:2,standard12hours:"2am"}),value:"2"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:3,standard12hours:"3am"}),value:"3"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:4,standard12hours:"4am"}),value:"4"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:5,standard12hours:"5am"}),value:"5"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:6,standard12hours:"6am"}),value:"6"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:7,standard12hours:"7am"}),value:"7"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:8,standard12hours:"8am"}),value:"8"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:9,standard12hours:"9am"}),value:"9"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:10,standard12hours:"10am"}),value:"10"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:11,standard12hours:"11am"}),value:"11"},{label:this.$i18n("cronOptionsSelects.atMidday"),value:"12"},{label:this.$i18n("cronOptionsSelects.atHour",{standard12hours:13,standard24hours:"1pm"}),value:"13"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:14,standard12hours:"2pm"}),value:"14"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:15,standard12hours:"3pm"}),value:"15"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:16,standard12hours:"4pm"}),value:"16"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:17,standard12hours:"5pm"}),value:"17"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:18,standard12hours:"6pm"}),value:"18"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:19,standard12hours:"7pm"}),value:"19"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:20,standard12hours:"8pm"}),value:"20"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:21,standard12hours:"9pm"}),value:"21"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:22,standard12hours:"10pm"}),value:"22"},{label:this.$i18n("cronOptionsSelects.atHour",{standard24hours:23,standard12hours:"11pm"}),value:"23"}],this.optionsDays=[{label:this.$i18n("cronOptionsSelects.everyDay"),value:"*"},{label:this.$i18n("cronOptionsSelects.weekly"),value:"*/7"},{label:this.$i18n("cronOptionsSelects.everyDays",{day:15}),value:"*/15"},{label:this.$i18n("cronOptionsSelects.everyDays",{day:30}),value:"*/30"},{label:this.$i18n("cronOptionsSelects.atDay",{day:1}),value:"1"},{label:this.$i18n("cronOptionsSelects.atDay",{day:2}),value:"2"},{label:this.$i18n("cronOptionsSelects.atDay",{day:3}),value:"3"},{label:this.$i18n("cronOptionsSelects.atDay",{day:4}),value:"4"},{label:this.$i18n("cronOptionsSelects.atDay",{day:5}),value:"5"},{label:this.$i18n("cronOptionsSelects.atDay",{day:6}),value:"6"},{label:this.$i18n("cronOptionsSelects.atDay",{day:7}),value:"7"},{label:this.$i18n("cronOptionsSelects.atDay",{day:8}),value:"8"},{label:this.$i18n("cronOptionsSelects.atDay",{day:9}),value:"9"},{label:this.$i18n("cronOptionsSelects.atDay",{day:10}),value:"10"},{label:this.$i18n("cronOptionsSelects.atDay",{day:11}),value:"11"},{label:this.$i18n("cronOptionsSelects.atDay",{day:12}),value:"12"},{label:this.$i18n("cronOptionsSelects.atDay",{day:13}),value:"13"},{label:this.$i18n("cronOptionsSelects.atDay",{day:14}),value:"14"},{label:this.$i18n("cronOptionsSelects.atDay",{day:15}),value:"15"},{label:this.$i18n("cronOptionsSelects.atDay",{day:16}),value:"16"},{label:this.$i18n("cronOptionsSelects.atDay",{day:17}),value:"17"},{label:this.$i18n("cronOptionsSelects.atDay",{day:18}),value:"18"},{label:this.$i18n("cronOptionsSelects.atDay",{day:19}),value:"19"},{label:this.$i18n("cronOptionsSelects.atDay",{day:20}),value:"20"},{label:this.$i18n("cronOptionsSelects.atDay",{day:21}),value:"21"},{label:this.$i18n("cronOptionsSelects.atDay",{day:22}),value:"22"},{label:this.$i18n("cronOptionsSelects.atDay",{day:23}),value:"23"},{label:this.$i18n("cronOptionsSelects.atDay",{day:24}),value:"24"},{label:this.$i18n("cronOptionsSelects.atDay",{day:25}),value:"25"},{label:this.$i18n("cronOptionsSelects.atDay",{day:26}),value:"26"},{label:this.$i18n("cronOptionsSelects.atDay",{day:27}),value:"27"},{label:this.$i18n("cronOptionsSelects.atDay",{day:28}),value:"28"},{label:this.$i18n("cronOptionsSelects.atDay",{day:29}),value:"29"},{label:this.$i18n("cronOptionsSelects.atDay",{day:30}),value:"30"},{label:this.$i18n("cronOptionsSelects.atDay",{day:31}),value:"31"}],this.optionsMonths=[{label:this.$i18n("cronOptionsSelects.everyMonth"),value:"*"},{label:this.$i18n("cronOptionsSelects.everyMonths",{month:2}),value:"*/2"},{label:this.$i18n("cronOptionsSelects.everyMonths",{month:3}),value:"*/3"},{label:this.$i18n("cronOptionsSelects.everyMonths",{month:4}),value:"*/4"},{label:this.$i18n("cronOptionsSelects.everyMonths",{month:6}),value:"*/6"},{label:this.$i18n("cronOptionsSelects.january"),value:"1"},{label:this.$i18n("cronOptionsSelects.february"),value:"2"},{label:this.$i18n("cronOptionsSelects.march"),value:"3"},{label:this.$i18n("cronOptionsSelects.april"),value:"4"},{label:this.$i18n("cronOptionsSelects.may"),value:"5"},{label:this.$i18n("cronOptionsSelects.june"),value:"6"},{label:this.$i18n("cronOptionsSelects.july"),value:"7"},{label:this.$i18n("cronOptionsSelects.august"),value:"8"},{label:this.$i18n("cronOptionsSelects.september"),value:"9"},{label:this.$i18n("cronOptionsSelects.october"),value:"10"},{label:this.$i18n("cronOptionsSelects.november"),value:"11"},{label:this.$i18n("cronOptionsSelects.december"),value:"12"}],this.optionsWeekDays=[{label:this.$i18n("cronOptionsSelects.everyDay"),value:"*"},{label:this.$i18n("cronOptionsSelects.mondayToFriday"),value:"1-5"},{label:this.$i18n("cronOptionsSelects.onlyWeekends"),value:"6-0"},{label:this.$i18n("cronOptionsSelects.everyMonday"),value:"1"},{label:this.$i18n("cronOptionsSelects.everyTuesday"),value:"2"},{label:this.$i18n("cronOptionsSelects.everyWednesday"),value:"3"},{label:this.$i18n("cronOptionsSelects.everyThursday"),value:"4"},{label:this.$i18n("cronOptionsSelects.everyFriday"),value:"5"},{label:this.$i18n("cronOptionsSelects.everySaturday"),value:"6"},{label:this.$i18n("cronOptionsSelects.everySunday"),value:"0"}]}}const J={class:"text-center text-h4 q-ma-sm",style:{"overflow-x":"auto","overflow-y":"hidden"}},X={class:"q-pa-md"},B={class:"q-pa-md"},N={class:"q-pa-md"},G={class:"q-pa-md"},ee={class:"q-pa-md"},ae={class:"row q-mb-sm"},le={class:"col-6 q-pr-md"},oe={class:"row"},te={class:"col-2 q-pa-sm"},ne={class:"col-10 q-pa-sm"},se={class:"col-6 q-pl-md"},re={class:"row"},ie={class:"col-2 q-pa-sm"},ue={class:"col-10 q-pa-sm"},ce={class:"row q-mb-sm"},de={class:"col-6 q-pr-md"},me={class:"row"},pe={class:"col-2 q-pa-sm"},ve={class:"col-10 q-pa-sm"},he={class:"col-6 q-pl-md"},be={class:"row"},ye={class:"col-2 q-pa-sm"},ge={class:"col-10 q-pa-sm"},we={class:"row q-mb-sm"},$e={class:"col-6 q-pr-md"},fe={class:"row"},Se={class:"col-2 q-pa-sm"},De={class:"col-10 q-pa-sm"},_e={class:"row q-mb-lg"},Ce={class:"row"},Fe="shadow-1 rounded-borders q-py-xs q-px-md",Oe=(0,o.aZ)({__name:"CronsDialogFormTabCustomFields",props:{minute:{default:"*"},hour:{default:"*"},day:{default:"*"},month:{default:"*"},weekday:{default:"*"},command:{default:""},comment:{default:""}},emits:["update:minute","update:hour","update:day","update:month","update:weekday","update:command","update:comment","update:isValidCron"],setup(e,{emit:a}){const l=e,s=(0,t.iH)(l.minute),r=(0,t.iH)(l.hour),i=(0,t.iH)(l.day),u=(0,t.iH)(l.month),c=(0,t.iH)(l.weekday),d=(0,t.iH)(l.command),m=(0,t.iH)(l.comment),p=new A,v=(0,t.iH)(Fe),h=(0,t.iH)(`${Fe} shadow-primary`),b=(0,t.iH)(Fe),y=(0,t.iH)(`${Fe} shadow-primary`),g=(0,t.iH)(Fe),w=(0,t.iH)(`${Fe} shadow-primary`),$=(0,t.iH)(Fe),f=(0,t.iH)(`${Fe} shadow-primary`),S=(0,t.iH)(Fe),D=(0,t.iH)(`${Fe} shadow-primary`),_=(0,o.Fl)((()=>C.value&&F&&O&&k&&q&&l.command.length>2&&l.command.length<4096)),C=(0,o.Fl)((()=>{const e=new RegExp(/^(\*|(?:\*|(?:[0-9]|(?:[1-5][0-9])))\/(?:[0-9]|(?:[1-5][0-9]))|(?:[0-9]|(?:[1-5][0-9]))(?:(?:\-[0-9]|\-(?:[1-5][0-9]))?|(?:\,(?:[0-9]|(?:[1-5][0-9])))*))$/);return e.test(s.value)})),F=(0,o.Fl)((()=>{const e=new RegExp(/^(\*|(1?[0-9]|2[0-3])(-(1?[0-9]|2[0-3]))?)(\/[1-9][0-9]*)?(,(\*|(1?[0-9]|2[0-3])(-(1?[0-9]|2[0-3]))?)(\/[1-9][0-9]*)?)*$/);return e.test(r.value)})),O=(0,o.Fl)((()=>{const e=new RegExp(/^(\*|([1-9]|[1-2][0-9]?|3[0-1])(-([1-9]|[1-2][0-9]?|3[0-1]))?)(\/[1-9][0-9]*)?(,(\*|([1-9]|[1-2][0-9]?|3[0-1])(-([1-9]|[1-2][0-9]?|3[0-1]))?)(\/[1-9][0-9]*)?)*$/);return e.test(i.value)})),k=(0,o.Fl)((()=>{const e=new RegExp(/^(\*|([1-9]|1[0-2]?)(-([1-9]|1[0-2]?))?)(\/[1-9][0-9]*)?(,(\*|([1-9]|1[0-2]?)(-([1-9]|1[0-2]?))?)(\/[1-9][0-9]*)?)*$/);return e.test(u.value)})),q=(0,o.Fl)((()=>{const e=new RegExp(/^(\*|[0-6](-[0-6])?)(\/[1-9][0-9]*)?(,(\*|[0-6](-[0-6])?)(\/[1-9][0-9]*)?)*$/);return e.test(c.value)}));function W(e){switch(e){case"minuteInput":v.value=`${Fe} shadow-primary`,h.value=Fe;break;case"minuteSelect":v.value=`${Fe}`,h.value=`${Fe} shadow-primary`;break;case"hourInput":b.value=`${Fe} shadow-primary`,y.value=Fe;break;case"hourSelect":b.value=`${Fe}`,y.value=`${Fe} shadow-primary`;break;case"dayInput":g.value=`${Fe} shadow-primary`,w.value=Fe;break;case"daySelect":g.value=`${Fe}`,w.value=`${Fe} shadow-primary`;break;case"monthInput":$.value=`${Fe} shadow-primary`,f.value=Fe;break;case"monthSelect":$.value=`${Fe}`,f.value=`${Fe} shadow-primary`;break;case"weekdayInput":S.value=`${Fe} shadow-primary`,D.value=Fe;break;case"weekdaySelect":S.value=`${Fe}`,D.value=`${Fe} shadow-primary`;break}}return(0,o.YP)(_,(e=>{a("update:isValidCron",e)}),{immediate:!0}),(0,o.YP)(s,(e=>{a("update:minute",e)}),{immediate:!0}),(0,o.YP)(r,(e=>{a("update:hour",e)}),{immediate:!0}),(0,o.YP)(i,(e=>{a("update:day",e)}),{immediate:!0}),(0,o.YP)(u,(e=>{a("update:month",e)}),{immediate:!0}),(0,o.YP)(c,(e=>{a("update:weekday",e)}),{immediate:!0}),(0,o.YP)(d,(e=>{a("update:command",e)}),{immediate:!0}),(0,o.YP)(m,(e=>{a("update:comment",e)}),{immediate:!0}),(e,a)=>{const l=(0,o.up)("q-separator"),_=(0,o.up)("q-input"),H=(0,o.up)("q-icon"),Z=(0,o.up)("q-select");return(0,o.wg)(),(0,o.iD)("div",null,[(0,o._)("div",J,[(0,o._)("span",X,(0,n.zw)(s.value),1),(0,o._)("span",B,(0,n.zw)(r.value),1),(0,o._)("span",N,(0,n.zw)(i.value),1),(0,o._)("span",G,(0,n.zw)(u.value),1),(0,o._)("span",ee,(0,n.zw)(c.value),1)]),(0,o.Wm)(l,{class:"q-my-sm"}),(0,o._)("div",ae,[(0,o._)("div",le,[(0,o._)("div",oe,[(0,o._)("div",te,[(0,o.Wm)(_,{modelValue:s.value,"onUpdate:modelValue":a[0]||(a[0]=e=>s.value=e),debounce:"100",borderless:"",class:(0,n.C_)(v.value),onInput:a[1]||(a[1]=e=>W("minuteInput"))},null,8,["modelValue","class"])]),(0,o._)("div",ne,[(0,o.Wm)(Z,{options:(0,t.SU)(p).optionsMinutes,modelValue:s.value,"onUpdate:modelValue":a[2]||(a[2]=e=>s.value=e),label:e.$t("cronsDialogForm.fieldMinutes"),class:(0,n.C_)(h.value),debounce:"100","emit-value":!0,"map-options":"",borderless:"","stack-label":"","error-message":e.$t("cronsDialogForm.messageInputError"),error:!C.value,onInput:a[3]||(a[3]=e=>W("minuteSelect"))},{prepend:(0,o.w5)((()=>[(0,o.Wm)(H,{name:"fas fa-hourglass-start"})])),_:1},8,["options","modelValue","label","class","error-message","error"])])])]),(0,o._)("div",se,[(0,o._)("div",re,[(0,o._)("div",ie,[(0,o.Wm)(_,{modelValue:r.value,"onUpdate:modelValue":a[4]||(a[4]=e=>r.value=e),debounce:"100",borderless:"",class:(0,n.C_)(b.value),onInput:a[5]||(a[5]=e=>W("hourInput"))},null,8,["modelValue","class"])]),(0,o._)("div",ue,[(0,o.Wm)(Z,{options:(0,t.SU)(p).optionsHours,modelValue:r.value,"onUpdate:modelValue":a[6]||(a[6]=e=>r.value=e),label:e.$t("cronsDialogForm.fieldHours"),class:(0,n.C_)(y.value),debounce:"100","emit-value":!0,"map-options":"",borderless:"","stack-label":"","error-message":e.$t("cronsDialogForm.messageInputError"),error:!F.value,onInput:a[7]||(a[7]=e=>W("hourSelect"))},{prepend:(0,o.w5)((()=>[(0,o.Wm)(H,{name:"far fa-clock"})])),_:1},8,["options","modelValue","label","class","error-message","error"])])])])]),(0,o._)("div",ce,[(0,o._)("div",de,[(0,o._)("div",me,[(0,o._)("div",pe,[(0,o.Wm)(_,{modelValue:i.value,"onUpdate:modelValue":a[8]||(a[8]=e=>i.value=e),debounce:"100",borderless:"",class:(0,n.C_)(g.value),onInput:a[9]||(a[9]=e=>W("dayInput"))},null,8,["modelValue","class"])]),(0,o._)("div",ve,[(0,o.Wm)(Z,{options:(0,t.SU)(p).optionsDays,modelValue:i.value,"onUpdate:modelValue":a[10]||(a[10]=e=>i.value=e),label:e.$t("cronsDialogForm.fieldDays"),class:(0,n.C_)(w.value),debounce:"100","emit-value":!0,"map-options":"",borderless:"","stack-label":"","error-message":e.$t("cronsDialogForm.messageInputError"),error:!O.value,onInput:a[11]||(a[11]=e=>W("daySelect"))},{prepend:(0,o.w5)((()=>[(0,o.Wm)(H,{name:"fas fa-calendar-day"})])),_:1},8,["options","modelValue","label","class","error-message","error"])])])]),(0,o._)("div",he,[(0,o._)("div",be,[(0,o._)("div",ye,[(0,o.Wm)(_,{modelValue:u.value,"onUpdate:modelValue":a[12]||(a[12]=e=>u.value=e),debounce:"100",borderless:"",class:(0,n.C_)($.value),onInput:a[13]||(a[13]=e=>W("monthInput"))},null,8,["modelValue","class"])]),(0,o._)("div",ge,[(0,o.Wm)(Z,{options:(0,t.SU)(p).optionsMonths,modelValue:u.value,"onUpdate:modelValue":a[14]||(a[14]=e=>u.value=e),label:e.$t("cronsDialogForm.fieldMonths"),class:(0,n.C_)(f.value),debounce:"100","emit-value":!0,"map-options":"",borderless:"","stack-label":"","error-message":e.$t("cronsDialogForm.messageInputError"),error:!k.value,onInput:a[15]||(a[15]=e=>W("monthSelect"))},{prepend:(0,o.w5)((()=>[(0,o.Wm)(H,{name:"far fa-calendar-alt"})])),_:1},8,["options","modelValue","label","class","error-message","error"])])])])]),(0,o._)("div",we,[(0,o._)("div",$e,[(0,o._)("div",fe,[(0,o._)("div",Se,[(0,o.Wm)(_,{modelValue:c.value,"onUpdate:modelValue":a[16]||(a[16]=e=>c.value=e),debounce:"100",borderless:"",class:(0,n.C_)(S.value),onInput:a[17]||(a[17]=e=>W("weekdayInput"))},null,8,["modelValue","class"])]),(0,o._)("div",De,[(0,o.Wm)(Z,{options:(0,t.SU)(p).optionsWeekDays,modelValue:c.value,"onUpdate:modelValue":a[18]||(a[18]=e=>c.value=e),label:e.$t("cronsDialogForm.fieldWeekDay"),class:(0,n.C_)(D.value),debounce:"100","emit-value":!0,"map-options":"",borderless:"","stack-label":"","error-message":e.$t("cronsDialogForm.messageInputError"),error:!q.value,onInput:a[19]||(a[19]=e=>W("weekdaySelect"))},{prepend:(0,o.w5)((()=>[(0,o.Wm)(H,{name:"far fa-calendar-check"})])),_:1},8,["options","modelValue","label","class","error-message","error"])])])])]),(0,o.Wm)(l,{class:"q-my-md"}),(0,o._)("div",_e,[(0,o.Wm)(Y.Z,{label:e.$t("cronsDialogForm.fieldCommand"),value:d.value,"onUpdate:value":a[20]||(a[20]=e=>d.value=e),type:"textarea",class:"full-width",maxLength:"4096",icon:"fas fa-code",rules:[()=>""!==d.value||e.$t("cronsDialogForm.messageFieldRequired")]},null,8,["label","value","rules"])]),(0,o._)("div",Ce,[(0,o.Wm)(Y.Z,{label:e.$t("cronsDialogForm.fieldComment"),value:m.value,"onUpdate:value":a[21]||(a[21]=e=>m.value=e),class:"full-width",maxLength:"100",icon:"far fa-comment-alt"},null,8,["label","value"])])])}}});var ke=l(50926),qe=l(32259);const We=Oe,He=We;m()(Oe,"components",{QSeparator:ke.Z,QInput:$.Z,QSelect:qe.Z,QIcon:f.Z});var Ze=l(88900),Te=l(45273),xe=l(69036),Qe=function(e,a,l,o){function t(e){return e instanceof l?e:new l((function(a){a(e)}))}return new(l||(l=Promise))((function(l,n){function s(e){try{i(o.next(e))}catch(a){n(a)}}function r(e){try{i(o["throw"](e))}catch(a){n(a)}}function i(e){e.done?l(e.value):t(e.value).then(s,r)}i((o=o.apply(e,a||[])).next())}))};class Ve extends xe.Z{getCrons(){return Qe(this,void 0,void 0,(function*(){return this.request.get("/v1/cron/")}))}createCron(e){return Qe(this,void 0,void 0,(function*(){return this.request.post("/v1/cron/",e)}))}updateCron(e){return Qe(this,void 0,void 0,(function*(){return this.request.put("/v1/cron/",e)}))}deleteCron(e){return Qe(this,void 0,void 0,(function*(){return this.request.delete(`/v1/cron/${e}/`)}))}}const Ue=(0,o.aZ)({__name:"CronsDialogFormTabCustom",setup(e){const a=r(),l=(0,b.QT)().t,n=(0,t.iH)(!1),s=(0,t.iH)(""),i=(0,t.iH)("*"),u=(0,t.iH)("*"),c=(0,t.iH)("*"),d=(0,t.iH)("*"),m=(0,t.iH)("*"),p=(0,t.iH)(""),v=(0,t.iH)(""),y=(0,o.Fl)({get:()=>a.getKeyCronsTable,set:e=>{a.setKeyCronsTable(e)}}),g=(0,o.Fl)((()=>a.formType)),w=(0,o.Fl)((()=>a.getSelectedCron)),$=(0,o.Fl)({get:()=>a.showDialogForm,set:e=>{a.showDialogForm=e}});function f(){"update"===g.value?D():S()}function S(){(0,Te.Q)();const e=new Ve,a={command:p.value,comment:v.value,schedule:`${i.value} ${u.value} ${c.value} ${d.value} ${m.value}`};e.createCron(a).then((()=>{(0,Ze.LX)(`${l("cronsDialogForm.createCronjobSuccess")}`),y.value++,$.value=!1})).catch((e=>{(0,Ze.s9)(e.response.data,`${l("cronsDialogForm.createCronjobError")}`),console.error(e)})).finally((()=>{(0,Te.Z)()}))}function D(){(0,Te.Q)();const e=new Ve,o={id:s.value,command:p.value,comment:v.value,schedule:`${i.value} ${u.value} ${c.value} ${d.value} ${m.value}`};e.updateCron(o).then((()=>{y.value++,(0,Ze.LX)(`${l("cronsDialogForm.updateCronjobSuccess")}`),a.setShowDialogForm(!1)})).catch((e=>{(0,Ze.s9)(e.response.data,`${l("cronsDialogForm.updateCronjobError")}`)})).finally((()=>{(0,Te.Z)()}))}return(0,o.wF)((()=>{var e,a,l,o,t;"create"!==g.value&&(s.value=w.value.id,p.value=w.value.command,v.value=w.value.comment,w.value.hasOwnProperty("predefined")||(i.value=null!==(e=w.value.minute)&&void 0!==e?e:"*",u.value=null!==(a=w.value.hour)&&void 0!==a?a:"*",c.value=null!==(l=w.value.day)&&void 0!==l?l:"*",d.value=null!==(o=w.value.month)&&void 0!==o?o:"*",m.value=null!==(t=w.value.weekday)&&void 0!==t?t:"*"))})),(e,a)=>{const s=(0,o.up)("q-card-section"),r=(0,o.up)("q-card-actions"),b=(0,o.up)("q-card");return(0,o.wg)(),(0,o.j4)(b,{flat:""},{default:(0,o.w5)((()=>[(0,o.Wm)(s,{class:"q-px-none"},{default:(0,o.w5)((()=>[(0,o.Wm)(He,{minute:i.value,"onUpdate:minute":a[0]||(a[0]=e=>i.value=e),hour:u.value,"onUpdate:hour":a[1]||(a[1]=e=>u.value=e),day:c.value,"onUpdate:day":a[2]||(a[2]=e=>c.value=e),month:d.value,"onUpdate:month":a[3]||(a[3]=e=>d.value=e),weekday:m.value,"onUpdate:weekday":a[4]||(a[4]=e=>m.value=e),command:p.value,"onUpdate:command":a[5]||(a[5]=e=>p.value=e),comment:v.value,"onUpdate:comment":a[6]||(a[6]=e=>v.value=e),isValidCron:n.value,"onUpdate:isValidCron":a[7]||(a[7]=e=>n.value=e)},null,8,["minute","hour","day","month","weekday","command","comment","isValidCron"])])),_:1}),(0,o.Wm)(r,{align:"between",class:"q-pa-none"},{default:(0,o.w5)((()=>[(0,o.Wm)(h.Z,{label:e.$t("cronsDialogForm.btnCancel"),color:"grey-7",onClick:a[8]||(a[8]=e=>$.value=!1)},null,8,["label"]),(0,o.Wm)(h.Z,{disable:!n.value,label:"update"===g.value?(0,t.SU)(l)("cronsDialogForm.updateCronjob"):(0,t.SU)(l)("cronsDialogForm.createCronjob"),color:"primary",icon:"update"===g.value?"sym_s_edit":"sym_s_add",onClick:a[9]||(a[9]=e=>f())},null,8,["disable","label","icon"])])),_:1})])),_:1})}}});var je=l(11821);const Ie=Ue,Re=Ie;m()(Ue,"components",{QCard:I.Z,QCardSection:R.Z,QCardActions:je.Z});var Me=l(94629);const Pe={class:"row q-mb-lg"},Ee={class:"row"},ze=(0,o.aZ)({__name:"CronsDialogFormTabPredefined",setup(e){const a=(0,b.QT)().t,l=r(),n=(0,t.iH)(""),s=(0,t.iH)(""),i=(0,t.iH)(""),u=(0,t.iH)("@daily"),c=[{label:`${a("cronsDialogForm.predefinedHourly")}`,value:"@hourly"},{label:`${a("cronsDialogForm.predefined3Hours")}`,value:"3hours"},{label:`${a("cronsDialogForm.predefined6Hours")}`,value:"6hours"},{label:`${a("cronsDialogForm.predefined12Hours")}`,value:"12hours"},{label:`${a("cronsDialogForm.predefinedDaily")}`,value:"@daily"},{label:`${a("cronsDialogForm.predefinedWeekly")}`,value:"@weekly"},{label:`${a("cronsDialogForm.predefinedMonthly")}`,value:"@monthly"},{label:`${a("cronsDialogForm.predefinedAnnually")}`,value:"@annually"},{label:`${a("cronsDialogForm.predefinedReboot")}`,value:"@reboot"}],d=(0,o.Fl)({get:()=>l.getKeyCronsTable,set:e=>l.setKeyCronsTable(e)}),m=(0,o.Fl)({get:()=>l.getShowDialogForm,set:e=>l.setShowDialogForm(e)}),p=(0,o.Fl)({get:()=>l.getSelectedCron,set:e=>l.setSelectedCron(e)}),v=(0,o.Fl)({get:()=>l.getFormType,set:e=>l.setFormType(e)}),y=(0,o.Fl)((()=>u.value.length>=0&&g.value)),g=(0,o.Fl)((()=>s.value.length>=2&&s.value.length<=4096));function w(){"update"===v.value?f():S()}function $(){const e={command:s.value,comment:i.value,schedule:""};switch(u.value){case"3hours":e.schedule="* */3 * * *";break;case"6hours":e.schedule="* */6 * * *";break;case"12hours":e.schedule="* */12 * * *";break;default:e.schedule=u.value}return e}function f(){(0,Te.Q)();const e=new Ve,l=Object.assign({id:n.value},$());e.updateCron(l).then((()=>{(0,Ze.LX)(`${a("cronsDialogForm.updateCronjobSuccess")}`),d.value++,m.value=!1})).catch((e=>{(0,Ze.s9)(e.response.data,`${a("cronsDialogForm.updateCronjobError")}`)})).finally((()=>{(0,Te.Z)()}))}function S(){(0,Te.Q)();const e=new Ve;e.createCron($()).then((()=>{(0,Ze.LX)(`${a("cronsDialogForm.createCronjobSuccess")}`),d.value++,m.value=!1})).catch((e=>{(0,Ze.s9)(e.response.data,`${a("cronsDialogForm.createCronjobError")}`)})).finally((()=>{(0,Te.Z)()}))}return(0,o.wF)((()=>{var e;"create"!==v.value&&(n.value=p.value.id,s.value=p.value.command,i.value=p.value.comment,p.value.hasOwnProperty("predefined")&&(u.value="@yearly"===p.value.predefined?"@annually":null!==(e=p.value.predefined)&&void 0!==e?e:"@daily"))})),(e,a)=>{const l=(0,o.up)("q-card-section"),t=(0,o.up)("q-card-actions"),n=(0,o.up)("q-card");return(0,o.wg)(),(0,o.j4)(n,{flat:""},{default:(0,o.w5)((()=>[(0,o.Wm)(l,{class:"q-px-none"},{default:(0,o.w5)((()=>[(0,o.Wm)(Me.Z,{icon:"far fa-calendar",label:e.$t("cronsDialogForm.selectExecuteTask"),options:c,selected:u.value,"onUpdate:selected":a[0]||(a[0]=e=>u.value=e)},null,8,["label","selected"])])),_:1}),(0,o.Wm)(l,{class:"q-px-none"},{default:(0,o.w5)((()=>[(0,o._)("div",Pe,[(0,o.Wm)(Y.Z,{label:e.$t("cronsDialogForm.fieldCommand"),value:s.value,"onUpdate:value":a[1]||(a[1]=e=>s.value=e),type:"textarea",class:"full-width",maxLength:"4096",icon:"fas fa-code",rules:[()=>""!==s.value||e.$t("cronsDialogForm.messageFieldRequired")]},null,8,["label","value","rules"])]),(0,o._)("div",Ee,[(0,o.Wm)(Y.Z,{label:e.$t("cronsDialogForm.fieldComment"),value:i.value,"onUpdate:value":a[2]||(a[2]=e=>i.value=e),class:"full-width",maxLength:"100",icon:"far fa-comment-alt"},null,8,["label","value"])])])),_:1}),(0,o.Wm)(t,{align:"between",class:"q-pa-none"},{default:(0,o.w5)((()=>[(0,o.Wm)(h.Z,{label:e.$t("cronsDialogForm.btnCancel"),color:"grey-7",onClick:a[3]||(a[3]=e=>m.value=!1)},null,8,["label"]),(0,o.Wm)(h.Z,{disable:!y.value,label:"update"===v.value?e.$t("cronsDialogForm.updateCronjob"):e.$t("cronsDialogForm.createCronjob"),color:"primary",icon:"update"===v.value?"sym_s_edit":"sym_s_add",onClick:a[4]||(a[4]=e=>w())},null,8,["disable","label","icon"])])),_:1})])),_:1})}}}),Ke=ze,Ye=Ke;m()(ze,"components",{QCard:I.Z,QCardSection:R.Z,QCardActions:je.Z});const Le=(0,o.aZ)({__name:"CronsDialogFormTabs",setup(e){const a=r(),l=(0,t.iH)("predefined"),n=(0,o.Fl)((()=>a.getSelectedCron)),s=(0,o.Fl)((()=>a.getFormType));return(0,o.wF)((()=>{"create"!==s.value&&(n.value.hasOwnProperty("predefined")||(l.value="custom"))})),(e,a)=>{const t=(0,o.up)("q-tab"),n=(0,o.up)("q-tabs"),s=(0,o.up)("q-separator"),r=(0,o.up)("q-tab-panel"),i=(0,o.up)("q-tab-panels");return(0,o.wg)(),(0,o.iD)(o.HY,null,[(0,o.Wm)(n,{modelValue:l.value,"onUpdate:modelValue":a[0]||(a[0]=e=>l.value=e),"active-color":"primary","indicator-color":"primary",align:"justify",class:"text-grey","no-caps":""},{default:(0,o.w5)((()=>[(0,o.Wm)(t,{name:"predefined",class:"cron-tabs",icon:"far fa-calendar",label:e.$t("cronsDialogForm.predefinedForm")},null,8,["label"]),(0,o.Wm)(t,{name:"custom",class:"cron-tabs",icon:"sym_s_edit_calendar",label:e.$t("cronsDialogForm.customForm")},null,8,["label"])])),_:1},8,["modelValue"]),(0,o.Wm)(s),(0,o.Wm)(i,{modelValue:l.value,"onUpdate:modelValue":a[1]||(a[1]=e=>l.value=e),animated:""},{default:(0,o.w5)((()=>[(0,o.Wm)(r,{name:"predefined",class:"q-pa-none"},{default:(0,o.w5)((()=>[(0,o.Wm)(Ye)])),_:1}),(0,o.Wm)(r,{name:"custom",class:"q-pa-none"},{default:(0,o.w5)((()=>[(0,o.Wm)(Re)])),_:1})])),_:1},8,["modelValue"])],64)}}});var Ae=l(47817),Je=l(70900),Xe=l(89800),Be=l(84106);const Ne=Le,Ge=Ne;m()(Le,"components",{QTabs:Ae.Z,QTab:Je.Z,QSeparator:ke.Z,QTabPanels:Xe.Z,QTabPanel:Be.Z});const ea={class:"flex justify-between items-center"},aa={class:"title-dialog"},la=(0,o.aZ)({__name:"CronsDialogForm",setup(e){const a=r(),l=(0,o.Fl)((()=>a.getFormType)),t=(0,o.Fl)({get:()=>a.showDialogForm,set:e=>{a.showDialogForm=e}});return(e,a)=>{const s=(0,o.up)("q-btn"),r=(0,o.up)("q-card-section"),i=(0,o.up)("q-card"),u=(0,o.up)("q-dialog");return(0,o.wg)(),(0,o.j4)(u,{modelValue:t.value,"onUpdate:modelValue":a[1]||(a[1]=e=>t.value=e),persistent:""},{default:(0,o.w5)((()=>[(0,o.Wm)(i,{flat:"",style:{width:"950px","max-width":"90vw"},class:"dialog-card-bg"},{default:(0,o.w5)((()=>[(0,o._)("div",ea,[(0,o._)("div",aa,(0,n.zw)("update"===l.value?e.$t("cronsDialogForm.editCronjob"):e.$t("cronsDialogForm.newCronjob")),1),(0,o.Wm)(s,{icon:"sym_s_close",flat:"",round:"",dense:"",onClick:a[0]||(a[0]=e=>t.value=!1)})]),(0,o.Wm)(r,{class:"q-px-none"},{default:(0,o.w5)((()=>[(0,o.Wm)(Ge)])),_:1})])),_:1})])),_:1},8,["modelValue"])}}});var oa=l(32074);const ta=la,na=ta;m()(la,"components",{QDialog:oa.Z,QCard:I.Z,QBtn:u.Z,QCardSection:R.Z});var sa=l(89906);const ra=(0,o.aZ)({__name:"CronsDialogRemove",setup(e){const a=(0,b.QT)().t,l=r(),t=(0,o.Fl)({get:()=>l.getShowDialogRemove,set:e=>{l.setShowDialogRemove(e)}}),n=(0,o.Fl)((()=>l.getSelectedCron.id)),s=(0,o.Fl)({get:()=>l.getKeyCronsTable,set:e=>{l.setKeyCronsTable(e)}});function i(){const e=new Ve;(0,Te.Q)(),e.deleteCron(n.value).then((()=>{s.value++,(0,Ze.LX)(`${a("cronsDialogRemove.removeCronSuccess")}`),t.value=!1})).catch((e=>{console.error(e),(0,Ze.s9)(e.response.data,`${a("cronsDialogRemove.removeCronError")}`)})).finally((()=>{(0,Te.Z)()}))}return(e,a)=>((0,o.wg)(),(0,o.j4)(sa.Z,{showDeleteDialog:t.value,"onUpdate:showDeleteDialog":a[2]||(a[2]=e=>t.value=e),titleDialog:e.$t("cronsDialogRemove.titleDialogRemove"),messageToDelete:e.$t("cronsDialogRemove.confirmRemoveCron"),warningToDelete:e.$t("cronsDialogRemove.warningRemoveCron")},{"card-actions":(0,o.w5)((()=>[(0,o.Wm)(h.Z,{label:e.$t("cronsDialogRemove.btnCancel"),color:"grey-7",onClick:a[0]||(a[0]=e=>t.value=!1)},null,8,["label"]),(0,o.Wm)(h.Z,{color:"negative",label:e.$t("cronsDialogRemove.btnConfirm"),onClick:a[1]||(a[1]=e=>i())},null,8,["label"])])),_:1},8,["showDeleteDialog","titleDialog","messageToDelete","warningToDelete"]))}}),ia=ra,ua=ia,ca=(0,o.aZ)({__name:"CronsIndex",setup(e){const a=r(),l=(0,b.QT)().t,n=(0,t.iH)([]),s=(0,t.iH)(!1),i=(0,o.Fl)((()=>a.getKeyCronsTable));function u(){s.value=!0,n.value=[];const e=new Ve;e.getCrons().then((e=>{0!==e.data.body.length&&(n.value=e.data.body.map((e=>c(e))))})).catch((e=>{console.error(e),(0,Ze.s9)(e.response.data,`${l("cronsIndex.getCronjobsError")}`)})).finally((()=>{s.value=!1}))}function c(e){var a,l,o,t,n;const s=["@hourly","@daily","@reboot","@weekly","@monthly","@yearly","@annually"];if(s.includes(e.schedule))return Object.assign(Object.assign({},e),{predefined:e.schedule});const r=/^(?<minute>\S+)\s+(?<hour>\S+)\s+(?<day>\S+)\s+(?<month>\S+)\s+(?<weekday>\S+)$/,i=RegExp(r).exec(e.schedule);return i?Object.assign(Object.assign({},e),{minute:null===(a=i.groups)||void 0===a?void 0:a.minute,hour:null===(l=i.groups)||void 0===l?void 0:l.hour,day:null===(o=i.groups)||void 0===o?void 0:o.day,month:null===(t=i.groups)||void 0===t?void 0:t.month,weekday:null===(n=i.groups)||void 0===n?void 0:n.weekday}):Object.assign(Object.assign({},e),{minute:"",hour:"",day:"",month:"",weekday:""})}return(0,o.wF)((()=>{u()})),(0,o.YP)(i,(()=>{u()})),(e,a)=>{const l=(0,o.up)("q-card-section"),t=(0,o.up)("q-card"),r=(0,o.up)("q-page");return(0,o.wg)(),(0,o.j4)(r,{padding:""},{default:(0,o.w5)((()=>[(0,o.Wm)(na),(0,o.Wm)(ua),(0,o.Wm)(t,{flat:""},{default:(0,o.w5)((()=>[(0,o.Wm)(l,null,{default:(0,o.w5)((()=>[!0===s.value?((0,o.wg)(),(0,o.j4)(K,{key:0})):((0,o.wg)(),(0,o.j4)(k,{key:1,data:n.value},null,8,["data"]))])),_:1})])),_:1})])),_:1})}}});var da=l(69885);const ma=ca,pa=ma;m()(ca,"components",{QPage:da.Z,QCard:I.Z,QCardSection:R.Z})}}]);