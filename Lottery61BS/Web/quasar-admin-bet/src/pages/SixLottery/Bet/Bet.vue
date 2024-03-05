<template>
  <q-card class="table-bg no-shadow bg-white" bordered>
    <q-card-section class="q-pa-none row"  >
      <div class="col-9 q-pa-none example-row-equal-width">
        <div class="row" v-for="i in 10">
          <div class="col row" v-for="j in 5">
            <q-input  reverse-fill-mask mask="########" v-if="j*10 - 10 +i !=50"
                     @click="onclickSetV(j*10 - 10 +i)"
                     style="width: 100%; background-color:white" class="col-auto "
                     dense v-model.number="layout1[ j*10 - 10 +i].v">
              <template v-slot:prepend>
                <span  :class="getNumClass(j*10 - 10 +i)">{{ j*10 - 10 +i }}</span>
              </template>
            </q-input>
          </div>

        </div>
      </div>
      <div class="col-3 q-gutter-sm">
        <q-tabs
          v-model="tab"
          dense
          class="text-grey"
          active-color="primary"
          indicator-color="primary"
          align="justify"
          narrow-indicator
        >
          <q-tab name="mails" label="投注操作" />
          <q-tab name="alarms" label="历史记录" />
        </q-tabs>
        <q-separator />
        <q-tab-panels v-model="tab" animated>
          <q-tab-panel name="mails">
            <div class="q-gutter-sm">
              <div class="q-gutter-sm">
                <q-btn @click="onclickSetVByType('Red')" class="but_r" label="红"/>
                <q-btn @click="onclickSetVByType('Blue')" class="but_b" label="蓝"/>
                <q-btn @click="onclickSetVByType('Green')" class="but_g" label="绿"/>
              </div>
              <div class="q-gutter-sm">
                <q-btn @click="onclickSetVByType('Single')" class="" label="单"/>
                <q-btn @click="onclickSetVByType('Double')" class="" label="双"/>
              </div>
              <div class="q-gutter-sm">
                <q-btn @click="onclickSetVByType('Shu')">鼠</q-btn>
                <q-btn @click="onclickSetVByType('Niu')">牛</q-btn>
                <q-btn @click="onclickSetVByType('Hu')">虎</q-btn>
                <q-btn @click="onclickSetVByType('Tu')">兔</q-btn>
              </div>
              <div class="q-gutter-sm">
                <q-btn @click="onclickSetVByType('Long')">龙</q-btn>
                <q-btn @click="onclickSetVByType('She')">蛇</q-btn>
                <q-btn @click="onclickSetVByType('Ma')">马</q-btn>
                <q-btn @click="onclickSetVByType('Yang')">羊</q-btn>
              </div>
              <div class="q-gutter-sm">
                <q-btn @click="onclickSetVByType('Hou')">猴</q-btn>
                <q-btn @click="onclickSetVByType('Ji')">鸡</q-btn>
                <q-btn @click="onclickSetVByType('Gou')">狗</q-btn>
                <q-btn @click="onclickSetVByType('Zhu')">猪</q-btn>
              </div>
            </div>
          </q-tab-panel>
          <q-tab-panel name="alarms"  class="q-pa-none">

              {{ ResultHistoryInfo.Data.CycleId }}
            <div class="row">
              <span class="col-1" v-for="item in ResultHistoryInfo.Data.TMs" :class="getNumClass(item)">{{ item }}</span>
            </div>

          </q-tab-panel>
        </q-tab-panels>
      </div>


    </q-card-section>
    <q-card-section class="q-pa-none row">
      <div class="row q-pa-sm-xs q-gutter-sm" style="max-height: 50px">
        <q-input class="bg-white" outlined mask="########"  dense v-model.number="defValue">
          <template v-slot:prepend class="bg-white">
            <q-btn flat size="sm" @click="onSetDefV">默认金额：</q-btn>
          </template>
        </q-input>
        <q-btn  color="purple" @click="onClickBet"  label="提交投注"/>
        <q-btn color="purple"  @click="onclickClear0" label="清0"/>
        <q-input  style="width: 140px" filled readonly dense v-model="C.NumOrder">
          <template v-slot:prepend>
            <span style=" font-size: 14px">投注号：</span>
          </template>
        </q-input>

        <q-input style="width: 140px" filled readonly dense v-model="C.S">
          <template v-slot:prepend>
            <span style=" font-size: 14px">封盘时间：</span>
          </template>
        </q-input>
        <q-select filled style="width: 200px" @update:model-value="onUpdateLotteryType" v-model="LotteryTInfo.LotteryType" :options="LotteryTInfo.ItemTypes"
                  dense
                  option-value="Id"
                  option-label="Name"
                  option-disable="inactive"
                  emit-value
                  map-options >
          <template v-slot:prepend>
            <span style=" font-size: 14px">彩种：</span>
          </template>
        </q-select>
        <q-input  style="width: 180px" filled readonly dense v-model="C.L">
          <template v-slot:prepend>
            <span style=" font-size: 14px">期号：</span>
          </template>
        </q-input>
      </div>
      <div class="row col-12">
        <q-input  class="bg-white col-12" @update:model-value="onQuickBet"
                  outlined placeholder="按号码=金额的格式，多个用空格分隔。如1=10 12=20"  dense v-model="QuickBetM">
          <template v-slot:prepend class="bg-white">
            <span style="font-size: 14px">快速投注：</span>
          </template>
        </q-input>
      </div>
    </q-card-section>

    <q-card-section class="q-pa-none">
      <q-table :rows="tableInfo.rows" :columns="tableInfo.columns" separator="cell"  class="q-table-low-height"
               v-model:pagination="tableInfo.pagination" @request="onRequest">
        <template v-slot:body="props">
          <q-tr :props="props">
            <q-td key="Id" :props="props">
              {{ props.row.Id }}
            </q-td>
            <q-td key="Title" :props="props">
              {{ props.row.Title }}
            </q-td>
            <q-td key="SiteUrl" :props="props">
              {{ props.row.SiteUrl }}
            </q-td>
            <q-td key="UserName" :props="props">
              {{ props.row.UserName }}
            </q-td>
            <q-td key="DataState" :props="props">
              {{ props.row.DataState }}
            </q-td>
            <q-td key="Rate" :props="props">
              {{ props.row.Rate * 100 }}%
            </q-td>
            <q-td key="BetMoney" :props="props">
              {{ props.row.BetMoney }}
            </q-td>
            <q-td key="ErrBetMoney" :props="props">
              {{ props.row.ErrBetMoney }}
            </q-td>
            <q-td key="IsBet" :props="props">
              <div v-if="(props.row.IsBet==1 && props.row.LoginState==1)" >
                <q-btn size="sm"  color="primary" @click="onClickSetBetState(props.row,0)">启用中</q-btn>
                <q-btn size="sm" @click="showBetDetailDialog(props.row)">投注明细</q-btn>
              </div>
              <div v-else>
                <q-btn size="sm" @click="onClickSetBetState(props.row,1)">禁用中</q-btn>
                <q-btn size="sm" @click="showBetDetailDialog(props.row)">投注明细</q-btn>
              </div>
            </q-td>
            <q-td key="StrLotteryType" :props="props">
              <div v-if="props.row.StrLotteryType==''" >
                无数据
              </div>
              <span v-else>
               {{ props.row.StrLotteryType }}
             </span>
            </q-td>
            <q-td key="LoginState" :props="props">
             <span v-if="props.row.LoginState==1">
               已登录
             </span>
              <span v-else>
               未登录
             </span>
            </q-td>
            <q-td key="Operate" :props="props">
              <div v-if="props.row.LoginState==1" >
                <q-btn color="primary" size="sm" label="下线" class="q-mx-xs" @click="onClickLoginOut(props.row)"/>
              </div>
              <div v-else>
                <q-btn  size="sm" label="登录" class="q-mx-xs" @click="showLogin(props.row)"/>
              </div>


            </q-td>
          </q-tr>
        </template>
      </q-table>
    </q-card-section>

  </q-card>

  <q-card class="table-bg no-shadow bg-white" bordered>
  </q-card>


  <q-dialog    v-model="betDetailDialog.dialogVisible"  >
    <q-card style="width: 90%">
      <q-toolbar>
        <q-toolbar-title> {{ EditDialog.Title }}</q-toolbar-title>
        <q-btn flat round dense icon="close" v-close-popup />
      </q-toolbar>
      <q-separator />

      <q-card-section class="q-pa-none">
        <q-table :rows="betDetailDialog.tableInfo.rows" :columns="betDetailDialog.tableInfo.columns" separator="cell"
                 v-model:pagination="betDetailDialog.tableInfo.pagination" @request="onBetDetailRequest">
          <template v-slot:body="props">
            <q-tr :props="props">
              <q-td key="Id" :props="props">
                {{ props.row.Id }}
              </q-td>
              <q-td key="NumOrder" :props="props">
                {{ props.row.NumOrder }}
              </q-td>
              <q-td key="BetMoney" :props="props">
                {{ props.row.BetMoney }}
              </q-td>
              <q-td key="ErrBetMoney" :props="props">
                {{ props.row.ErrBetMoney }}
              </q-td>
              <q-td key="BetInfo" :props="props">
                {{ props.row.BetInfo }}
              </q-td>

              <q-td key="Operate" :props="props">
                <q-btn color="primary" size="sm" label="删除" class="q-mx-xs" @click="onClickDelBetDetail(props.row)"/>
              </q-td>
            </q-tr>

          </template>
        </q-table>
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog    v-model="loginDialog.Visible"  >
    <q-card style="width: 280px">
      <q-toolbar>
        <q-toolbar-title> {{ loginDialog.title }}</q-toolbar-title>
        <q-btn flat round dense icon="close" v-close-popup />
      </q-toolbar>
      <q-separator />
      <q-card-section class="q-pa-none">
        <q-item dense>
          <q-item-section>
            <q-item-label>
              <q-img width="100px" :src="loginDialog.form.CodePicUrl" @click="refreshPic(loginDialog.SiteUserInfo.Id)"></q-img>
            </q-item-label>
          </q-item-section>
          <q-item-section>
            <q-item-label>
              <q-input dense hint="输入验证码"  v-model="loginDialog.form.Code"/>
            </q-item-label>
          </q-item-section>
        </q-item>
      </q-card-section>
      <q-card-actions   align="right">
        <q-btn flat @click="onClickLoginCode">确定</q-btn>
        <q-btn flat  @click="loginDialog.Visible=false">取消</q-btn>
      </q-card-actions>
    </q-card>
  </q-dialog>

</template>

<script>
import {
  GetSiteTypeList, GetNumComparisonInfo, DelBetSiteData, GetBetSiteDataList,
  GetSiteUserInfoList, SixLogin, SetBetState, SixLoginOut, GetC, Bet, SetLotteryType,
  GetResultHistory
} from "src/api/Api"
export default {
  name: "Bet",
  data(){
    return{
      tab:"mails",
      ResultHistoryInfo:{
        Data:{
          CycleId:"",
          TMs:[]
        },
      },

      C:{
        L:"",
        S:0,
        NumOrder:"0",
        HasLogin:false,
        LotteryType:0,
      },
      refreshDataTimer:null,
      RefreshData:{
        refreshDataTimer:null,
        MaxCurResultSecond:2,
        CountdownSecond:2,
        CurResultSecond:0,
        ReloadSiteUserDataMaxSecond:10,
        ReloadSiteUserDataSecond:0
      },
      QuickBetM:'',
      defValue:1,
      layout1:[
        {num:0,v:0},
        {num:1,v:0}
      ],
      formData :{
        betM:[]},
      NumComparisonInfo:{
        Single:[],
        Double:[],
        Green:[],
        Blue:[],
        Red:[],
      },

      vvv:0,
      LotteryTInfo:{
        LotteryType :0,
        ItemTypes:[
          { Id:1,Name:"澳门六合彩"},
          { Id:2,Name:"新澳门六合彩"},
          { Id:3,Name:"超级6+1六合彩"},
          { Id:4,Name:"闪电6+1六合彩"}
        ]
      },

      tableInfo:{
        columns:[
          {name:'Id',label:'ID',field:'Id', align: 'center'},
          {name:'Title',label:'名称',field:'Title', align: 'center'},
          {name:'SiteUrl',label:'地址',field:'SiteUrl', align: 'right'},
          {name:'UserName',label:'账号',field:'UserName', align: 'right'},
          {name:'DataState',label:'状态',field:'DataState' , align: 'center'},
          {name:'Rate',label:'投注比',field:'Rate' , align: 'center'},
          {name:'BetMoney',label:'投注金额',field:'BetMoney' , align: 'center'},
          {name:'ErrBetMoney',label:'失败金额',field:'ErrBetMoney' , align: 'center'},
          {name:'IsBet',label:'投注状态',field:'IsBet' , align: 'center'},
          {name:'StrLotteryType',label:'彩种状态',field:'StrLotteryType' , align: 'center'},
          {name:'LoginState',label:'登录状态',field:'LoginState' , align: 'center'},

          {name:'Operate',label:'操作',field:'Operate' , align: 'center'},
        ],
        rows:[],
        pagination:{
          page:1,
          rowsPerPage:10,
          rowsNumber:0,
        },
        query:{
          PageIndex:1,
          PageSize:10,
        },
        DataEx:{CurCycleId:"",NumOrder:""},
        MaxSecond:30,
        ReloadSiteUserDataSecond:0,
        ReloadSiteUserDataMaxSecond:5,
        Second:0,
        autoRefresh:false,
        Code:'',
      },
      EditDialog:{
        Title:'',
        Visible:false,
        Form:{SiteName:'',SiteTitle:'',SiteUrl:'',SiteDes:'',State:0,Id:0},
      },

      betDetailDialog:{
        title:'投注明细',
        tableInfo:{
          columns:[
            {name:'Id',label:'ID',field:'Id', align: 'center'},
            {name:'Title',label:'名称',field:'Title', align: 'center'},
            {name:'SiteUrl',label:'地址',field:'SiteUrl', align: 'right'},
            {name:'UserName',label:'账号',field:'UserName', align: 'right'},
            {name:'DataState',label:'状态',field:'DataState' , align: 'center'},
            {name:'Rate',label:'投注比',field:'Rate' , align: 'center'},
            {name:'BetMoney',label:'投注金额',field:'BetMoney' , align: 'center'},
            {name:'ErrBetMoney',label:'失败金额',field:'ErrBetMoney' , align: 'center'},
            {name:'IsBet',label:'投注状态',field:'IsBet' , align: 'center'},
            {name:'LoginState',label:'登录状态',field:'LoginState' , align: 'center'},

            {name:'Operate',label:'操作',field:'Operate' , align: 'center'},
          ],
          rows:[],
          pagination:{
            page:1,
            rowsPerPage:10,
            rowsNumber:0,
          },
          query:{
            SiteId:0,
            PageSize:20,
            PageIndex: 1,
            OrderBy: '',
          },
        },
        dialogVisible:false
      },
      loginDialog:{
        title:'登录',
        isBet:0,
        SiteUserInfo:{},
        Visible:false,
        form:{
          CodePicUrl:'',
          Code:''
        }
      },
    }
  },
  methods:{
    onSetDefV(){
      this.layout1.forEach(item =>{
        if (item.v > 0) {
          item.v = this.defValue
        }
      })

      let arrNum = new Array();
      for (let i=1;i<50;i++){
        if (this.layout1[i].v > 0 ){
          arrNum.push(i+"="+this.layout1[i].v)
        }
      }
      let str = arrNum.join(" ")
      this.QuickBetM = str
    },
    onQuickBet(v){
      let mpNum = new Map()
      let arrNumMs = v.split(" ")
      arrNumMs.forEach(numM =>{
        let aNM = numM.split("=")

         if (aNM.length==2){
           let num = parseInt(aNM[0])
           let Bet = parseInt(aNM[1])
           if (num > 0 && num <50) {

             mpNum.set(num, Bet)
           }
         }
      })
      this.onclickClear0()
      this.layout1.forEach(item =>{
          if (mpNum.has(item.num)) {
            item.v = mpNum.get(item.num)
          }
      })
    },
    onUpdateLotteryType(v){
      let query =  {LotteryType:v}

       this.tableInfo.rows.forEach(item =>{
         item.DataState = "正在重新获取数据..."
       })
      SetLotteryType(query).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
          this.getC()
          this.getSiteInfoList()
        } else {
        }
      }).catch(error => {
      })
    },
    onRequest(props){
      if (props.pagination.rowsPerPage==0) {
        props.pagination.rowsPerPage = 1000
      }
      this.tableInfo.pagination = props.pagination
      this.reloadTableData()
    },
    reloadTableData(){
      this.tableInfo.query.PageIndex = this.tableInfo.pagination.page
      this.tableInfo.query.PageSize = this.tableInfo.pagination.rowsPerPage

      GetSiteTypeList(this.tableInfo.query).then(response => {
        const { code, msg, obj } = response
        if (code==200){
          console.log(obj)
          this.tableInfo.pagination.rowsNumber = obj.GroupData.C
          this.tableInfo.rows = obj.ListData
        }
      }).catch(error => {
      })
    },
    openEditDialog(row){
      this.EditDialog.Title = `修改[${row.SiteName}]`
      this.EditDialog.Form.Id = row.Id
      this.EditDialog.Form.SiteName = row.SiteName
      this.EditDialog.Form.SiteTitle = row.SiteTitle
      this.EditDialog.Form.SiteUrl = row.SiteUrl
      this.EditDialog.Form.SiteDes = row.SiteDes

      this.EditDialog.Visible = true
    },

    // onClickSaveInfo(){
    //   SaveSiteType(this.EditDialog.Form).then(response => {
    //     const { code, msg, obj } = response
    //     if (code==200){
    //       this.reloadTableData()
    //       this.EditDialog.Visible = false
    //     } else {
    //       this.$q.notify(msg)
    //     }
    //   }).catch(error => {
    //     this.$q.notify(error)
    //   })
    // },
    onclickClear0() {
      this.layout1.forEach(item =>{
        item.v = null
      })

    },

    onClickBet(){
      let arrBetInfo = new Object()
      let hasNum = false
      this.layout1.forEach(item =>{
        if (item.v > 0) {
          arrBetInfo[item.num]= parseInt(item.v)
          hasNum = true
        }
      })
      if (!hasNum) {
        this.$q.notify("必须有一个号码金额大于0")
        return
      }
      let BetData =JSON.stringify(arrBetInfo)

      const BetQuery ={
        BetData:BetData,
        LotteryType:this.LotteryTInfo.LotteryType
      }
      Bet(BetQuery).then(response => {
        const { code, msg, obj } = response
        console.log(code, msg, obj)
        if (code==200) {
          let ListData = obj.Data.ListData
          let GroupData = obj.Data.GroupData
          let NumOrder = obj.NumOrder
          this.loadSiteUserData(ListData, GroupData, obj.Data.DataEx)
          this.onclickClear0()
          if (obj.ErrM!="") {
            this.$q.notify(`单号${NumOrder}投注完成,但部分数据有问题：`+obj.ErrM)
          } else {
            this.$q.notify(`单号${NumOrder}投注完成`)
          }
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },
    onclickSetV(i){
      if(this.layout1[i].v == this.defValue) {
        this.layout1[i].v = null
        this.formData.betM[i] = 0
      } else {
        this.layout1[i].v = this.defValue
        this.formData.betM[i] = this.defValue
      }

      let arrNum = new Array();
      for (let i=1;i<50;i++){
         if (this.layout1[i].v > 0 ){
           arrNum.push(i+"="+this.layout1[i].v)
         }
      }
      let str = arrNum.join(" ")
      this.QuickBetM = str
    },

    onclickSetVByType(type) {
      let allG0=true
      let allE0=true
      this.NumComparisonInfo[type].forEach(num =>{
          let i = num
          if (this.layout1[i].v==0 || this.layout1[i].v==null){
            allG0 = false
          } else if (this.layout1[i].v>0){
            allE0 = false
          }
        }
      )
      let setV = ""

      if (!allG0) {
        setV = this.defValue
      }
      this.NumComparisonInfo[type].forEach(num =>{
          let i = num
          this.layout1[i].v=setV
        }
      )

    },

    getNumClass(i){
      return 'num'.concat(i)
    },
    iniLayout1(){
      let layout = []
      for (let i=0;i<=50;i++) {
        let aNum = {num:i,v:null}
        layout.push(aNum)
      }
      this.layout1 = layout
    },
    iniBetM(){
      for (let i=0;i<=49;i++) {
        this.formData.betM[i] = 0
      }
    },

    async getNumComparisonInfo() {
      GetNumComparisonInfo({}).then(response => {
        const { code, msg, obj } = response
        this.NumComparisonInfo = obj
      }).catch(error => {
      })
    },
    loadSiteUserData(ListData, GroupData, DataEx) {
      this.C.NumOrder = DataEx.NumOrder
      this.C.L = DataEx.CurCycleId
      this.tableInfo.DataEx = DataEx
      this.tableInfo.ReloadSiteUserDataSecond = 0

      this.tableInfo.rows = ListData

      this.tableInfo.total = GroupData
    },

    async getSiteInfoList() {
      // this.SiteInfo.loading = true
      this.tableInfo.query.LotteryType = this.LotteryTInfo.LotteryType
      GetSiteUserInfoList(this.tableInfo.query).then(response => {
        const { code, msg, obj } = response
        this.loadSiteUserData(obj.ListData, obj.GroupData, obj.DataEx)
        // this.SiteInfo.loading = false
      }).catch(error => {
        // this.SiteInfo.loading = false
      })
    },

    onClickSetBetState(row, isBet) {
      if (row.LoginState!=1 && isBet==1) {
        this.showLogin(row, isBet)
        return
      }

      const query ={
        Id:row.Id,
        State:isBet,
      }
      SetBetState(query).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
          this.getSiteInfoList()
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },

    onBetDetailRequest(props){
      if (props.pagination.rowsPerPage==0) {
        props.pagination.rowsPerPage = 1000
      }
      this.betDetailDialog.tableInfo.pagination = props.pagination
      this.reloadTableData()
    },
    onClickDelBetDetail(row){
      let hint = "认要删除["+row.NumOrder+"]吗？"
      this.$q.dialog({
        title: '提示',
        message: hint,
        ok:"确定",
        cancel:"取消",
      }).onOk(data  => {
        let query = {Id:row.Id}
        DelBetSiteData(query).then(response => {
          const { code, msg, obj } = response
          if (code==200) {
            this.getBetSiteDataList()
            this.$q.notify("删除成功")
          } else {
            this.$q.notify(msg)
          }
        }).catch(error => {
          this.$q.notify(error)
        })
      })

    },
    loadBetSiteDataList(ListData, GroupData){
      this.betDetailDialog.tableInfo.rows = ListData
      this.betDetailDialog.total = GroupData
    },

    getBetSiteDataList(){
      GetBetSiteDataList(this.betDetailDialog.tableInfo.query).then(response => {
        const { code, msg, obj } = response
        this.loadBetSiteDataList(obj.ListData, obj.GroupData)
      }).catch(error => {
      })
    },

    showLogin(row, isBet){
      this.loginDialog.title = `登录【${row.UserName}】`
      this.loginDialog.isBet = isBet
      this.loginDialog.SiteUserInfo = row
      this.loginDialog.form.Code = ""
      this.loginDialog.form.CodePicUrl = this.getCodePicUrl(row.Id, true)
      this.loginDialog.Visible = true
    },
    getCodePicUrl(id, f){
      let t1
      if (f) {
        let t = new Date()
        t1 = t.getTime()
      }
      let codeSrc ='/api/getcodeimage?Id='+id+'&t='+t1
      return codeSrc
    },
    onClickLoginCode(){
      const loginP ={
        Id:this.loginDialog.SiteUserInfo.Id,
        Code:this.loginDialog.form.Code,
        IsBet:this.loginDialog.isBet
      }
      SixLogin(loginP).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
          this.C.HasLogin = true
          this.getSiteInfoList()
          this.loginDialog.Visible = false
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },

    onClickLoginOut(row){
      const loginP ={
        Id:row.Id,
      }
      SixLoginOut(loginP).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
          this.getSiteInfoList()
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },
    refreshPic(id){
      this.loginDialog.form.CodePicUrl = this.getCodePicUrl(id, true)
    },
    getC(){
      let query = {LotteryType:this.LotteryTInfo.LotteryType}
      GetC(query).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
          this.C = obj
          if (this.LotteryTInfo.LotteryType==0) {
            this.LotteryTInfo.LotteryType = obj.LotteryType
          }

        } else {
        }
      }).catch(error => {
      })
    },
    autoRefreshData(){
      this.RefreshData.CurResultSecond++
      if (this.RefreshData.CurResultSecond > 1000000){
        this.RefreshData.CurResultSecond = 0
      }
      let isGetC = false
      if (this.RefreshData.CurResultSecond % this.RefreshData.MaxCurResultSecond==0) {
        isGetC = true
        this.getC()
      }

      this.RefreshData.ReloadSiteUserDataSecond++
      if (this.RefreshData.ReloadSiteUserDataSecond>this.RefreshData.ReloadSiteUserDataMaxSecond){
        this.RefreshData.ReloadSiteUserDataSecond = 0
        this.getSiteInfoList()
      }

      if (this.RefreshData.CurResultSecond % this.RefreshData.CountdownSecond==0) {
        if (this.C.S>0) {
          this.C.S = this.C.S -1
        }
        if  (this.C.S==0 && isGetC==false&&this.C.HasLogin) {
          this.getC()
        }
      }
    },
    reloadResultHistory(){
      const query ={
        LotteryType:this.LotteryTInfo.LotteryType
      }
      GetResultHistory(query).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
          this.ResultHistoryInfo.Data = obj
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },
  },
  watch:{
    'C.L':function (v){
      this.reloadResultHistory()
    },
    'C.LotteryType':function (v){
      this.reloadResultHistory()
    }
  },
  mounted() {
    if (this.RefreshData.refreshDataTimer) {
      clearInterval(this.RefreshData.refreshDataTimer)
      this.RefreshData.refreshDataTimer = null;
    }
    this.RefreshData.refreshDataTimer = setInterval(()=>{this.autoRefreshData()},500);
  },

  unmounted() {
    clearInterval(this.RefreshData.refreshDataTimer);
    this.RefreshData.refreshDataTimer = null
  },
  created() {
    this.getC()
    this.iniLayout1()
    this.getNumComparisonInfo()
    this.iniBetM()
    this.getSiteInfoList()
  }

}
</script>

<style scoped lang="scss">
  .example-row-equal-width {
    .row > div{
      border: 1px solid
    }
    .row + .row{
      margin-top: -1px
    }
  }

  .num1,.num01, .num2,.num02, .num7,.num07, .num8,.num08, .num12, .num13, .num18, .num19, .num23, .num24, .num29, .num30, .num34, .num35, .num40, .num45, .num46{
    width: 32px;
    margin: 0  2px;
    padding: 3px;
    align-content: center;
    text-align: center;
    //background-color: red;
    //color:white;
    background-color: #DDD;
    border: white;
    color:red;
    font-size: 22px;
    //border-radius: 50%;
  }

  .num3,.num03, .num4,.num04, .num9,.num09, .num10, .num14, .num15, .num20, .num25, .num26, .num31, .num36, .num37, .num41, .num42, .num47, .num48{
    width: 32px;
    margin: 0  2px;
    padding: 3px;
    align-content: center;
    text-align: center;
    //background-color: blue;
    //color:white;
    background-color: #DDD;
    border: white;
    color:blue;
    font-size: 22px;
    //border-radius: 50%;
  }
  .num5,.num05, .num6,.num06, .num11, .num16, .num17, .num21, .num22, .num27, .num28, .num32, .num33, .num38, .num39, .num43, .num44, .num49{
    width: 32px;
    margin: 0  2px;
    padding: 3px;
    align-content: center;
    text-align: center;
    //background-color: green;
    //color:white;
    background-color: #DDD;
    border: white;
    color:#087921;
    font-size: 22px;
    //border-radius: 50%;
  }


  .but_r{
    color: red;
  }
  .aaa{
    margin-left: 10px;
  }

  .but_b{
    color: blue;
  }

  .but_g{
    color: green;
  }

  .q-table-low-height{
    //height: 150px;
    height: calc(100vh - 420px - 80px - 20px );
  }

  .aaabb{

  }

</style>
