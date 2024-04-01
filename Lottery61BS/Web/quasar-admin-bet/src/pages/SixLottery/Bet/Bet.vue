<template>
  <q-card class="table-bg no-shadow bg-white" bordered>
    <q-card-section  class="q-pa-none row"  >
      <q-card class="table-bg bg-white row" bordered>
        <sqan class="content-center"> 代理：{{AgentUserInfo.UserData.UserName}}</sqan>
        <div v-if="AgentUserInfo.UserData.IsLogin==0" >
          <q-btn @click="showAgentLogin(AgentUserInfo.UserData)">登录</q-btn>
          <q-btn  @click="onClickAgentLogout(AgentUserInfo.UserData)">登出</q-btn>
        </div>
        <div  v-else>
          <q-btn  @click="onClickAgentLogout(AgentUserInfo.UserData)">登出</q-btn>
        </div>
         <q-btn @click="onShowEditDialogAgent">修改</q-btn>
      </q-card>
      <q-card   style="margin-left:8px"  class="table-bg bg-white row" bordered>
        <sqan class="text-center">  会员：{{BetUserInfo.UserData.UserName}}</sqan>


        <div class="row" v-if="BetUserInfo.UserData.IsLogin==0" >
          <q-btn   @click="showBetLogin(BetUserInfo.UserData)">登录</q-btn>
          <q-btn   @click="onClickBetLogout(BetUserInfo.UserData)">登出</q-btn>
        </div>
        <div  class="row" v-else>
          <q-btn   @click="onClickBetLogout(BetUserInfo.UserData)">登出</q-btn>
        </div>
        <q-btn @click="onShowEditDialogBet">修改</q-btn>
      </q-card>

      <q-card  style="margin-left:8px"  class="table-bg bg-white" bordered>
        <q-btn @click="onClickR">刷新({{RefreshData.AgentData.MaxSecond-RefreshData.AgentData.Second}})</q-btn>
        <q-checkbox v-model="RefreshData.IsSave">读取代理数据</q-checkbox>
      </q-card>
    </q-card-section>

    <q-card-section class="q-pa-none row">
      <div class="row col q-pa-sm-xs q-gutter-sm" style="max-height: 50px">
        <div class="q-gutter-sm">
          <q-radio @update:model-value ="changeConfigDataBetWay" v-model="ConfigData.BetWay" :val="BetWaySelects.Rate" label="比例" />
          <q-radio @update:model-value ="changeConfigDataBetWay" v-model="ConfigData.BetWay" :val="BetWaySelects.BetM" label="金额" />
        </div>
        <div class="q-gutter-sm">
          <q-radio @update:model-value ="changeConfigDataBetWay" v-model="ConfigData.RateWay" :val="RateWaySelects.Z" label="正" />
          <q-radio @update:model-value ="changeConfigDataBetWay" v-model="ConfigData.RateWay" :val="RateWaySelects.F" label="反" />
        </div>
        <q-input @change="changeConfigDataRate" input-class="text-right" class="col" filled dense mask="########"   v-model.number="ConfigData.Rate">
          <template v-slot:prepend class="bg-white">
            <span  class="input_title">比例：</span>
          </template>
          <template v-slot:append>
            %
          </template>
        </q-input>
        <q-select   class="col"  filled style="width: 200px" @update:model-value="changeConfigDataStock" v-model="ConfigData.Stock" :options="StockSelectItems"
                   dense
                   option-value="Id"
                   option-label="Name"
                   option-disable="inactive"
                   emit-value
                   map-options >
          <template v-slot:prepend>
            <span class="input_title">方式：</span>
          </template>
        </q-select>

        <q-select  class="col"  filled style="width: 200px" @update:model-value="changeConfigDataLotteryType"
                   v-model="ConfigData.LotteryType" :options="LotteryTInfo.ItemTypes"
                   dense
                   option-value="Id"
                   option-label="Name"
                   option-disable="inactive"
                   emit-value
                   map-options >
          <template v-slot:prepend>
            <span class="input_title">彩种：</span>
          </template>
        </q-select>

<!--        <q-input  class="col"  filled dense v-model="ConfigData.BetTime1" mask="time" :rules="['time']">-->
<!--          <template v-slot:prepend class="bg-white">-->
<!--            <span  class="input_title"  >投注时间1：</span>-->
<!--          </template>-->
<!--          <template v-slot:append>-->
<!--            <q-icon name="access_time" class="cursor-pointer">-->
<!--              <q-popup-proxy cover transition-show="scale" transition-hide="scale">-->
<!--                <q-time v-model="ConfigData.BetTime1">-->
<!--                  <div class="row items-center justify-end">-->
<!--                    <q-btn v-close-popup label="Close" color="primary" flat />-->
<!--                  </div>-->
<!--                </q-time>-->
<!--              </q-popup-proxy>-->
<!--            </q-icon>-->
<!--          </template>-->
<!--        </q-input>-->
<!--        <q-input  class="col"  filled dense v-model="ConfigData.BetTime2" mask="time" :rules="['time']">-->
<!--          <template v-slot:prepend class="bg-white">-->
<!--            <span class="input_title "  >投注时间2：</span>-->
<!--          </template>-->
<!--          <template v-slot:append>-->
<!--            <q-icon name="access_time" class="cursor-pointer">-->
<!--              <q-popup-proxy cover transition-show="scale" transition-hide="scale">-->
<!--                <q-time v-model="ConfigData.BetTime1">-->
<!--                  <div class="row items-center justify-end">-->
<!--                    <q-btn v-close-popup label="Close" color="primary" flat />-->
<!--                  </div>-->
<!--                </q-time>-->
<!--              </q-popup-proxy>-->
<!--            </q-icon>-->
<!--          </template>-->
<!--        </q-input>-->
        <div  class="col">
          <q-btn @click="onClickBet" label="直接投注"/>
<!--          <q-btn label="启动"/>-->
        </div>
      </div>
    </q-card-section>
    <q-card-section class="q-pa-none row"  >
      <div class="row col-12">
        <q-input  input-class="text-right" mask="########" v-model.number="InputData.Datas[i]"
                 dense    class="input-row-equal-width q-px-sm col" v-for="i in 5">
          <template v-slot:append>
            <q-btn  @click="onClickSetBetM(i)" label="填充" />
          </template>
          <template v-slot:prepend>
            <q-btn  @click="onClickSetBetMBaseRate(i)" label="比例金额" />
          </template>
        </q-input>
      </div>
      <div class="col-12 q-pa-none example-row-equal-width">
        <div class="row" v-for="i in 10">
          <div class="col row" v-for="j in 5">
            <q-input  input-class="text-right" reverse-fill-mask mask="########" v-if="j*10 - 10 +i !=50"
                     style="width: 100%; background-color:white" class="col-auto"

  @blur="onB(j*10 - 10 +i )"
                     dense v-model.number="layout1[ j*10 - 10 +i].V">
              <template v-slot:prepend>
                <div class="row">
                  <span :class="getNumClass(layout1[ j*10 - 10 +i].Num)">{{ layout1[ j*10 - 10 +i].Num }}</span>
                  <span class="bet_m">{{ layout1[ j*10 - 10 +i].BetM }}</span>
                </div>
              </template>
              <template v-slot:append>
                <div class="row">
                  <span class="bet_m2">{{ layout1[ j*10 - 10 +i].OkBetM }}</span>
                </div>
              </template>
            </q-input>
          </div>

        </div>
      </div>
      <div class="row col-12">
        <q-select   class="col"  filled style="width: 200px"
                    v-model="IndexData.Datas[i-1]"
                    :options="IndexData.SelectItems"
                    @popup-hide = "onIndexDataHide(i-1)"
                    dense
                    option-value="Id"
                    option-label="Name"
                    option-disable="inactive"
                    emit-value
                    map-options
                    v-for="i in 5"
        >

        </q-select>

      </div>
    </q-card-section>

  </q-card>

  <q-card class="table-bg no-shadow bg-white" bordered>
  </q-card>

  <q-dialog    v-model="DialogAgentLogin.Visible"  >
    <q-card style="width: 280px">
      <q-toolbar>
        <q-toolbar-title> {{ DialogAgentLogin.title }}</q-toolbar-title>
        <q-btn flat round dense icon="close" v-close-popup />
      </q-toolbar>
      <q-separator />
      <q-card-section class="q-pa-none">
        <q-item dense>
          <q-item-section>
            <q-item-label>
              <q-img width="100px" :src="DialogAgentLogin.form.CodePicUrl" @click="refreshAgentPic(DialogAgentLogin.SiteUserInfo.Id)"></q-img>
            </q-item-label>
          </q-item-section>
          <q-item-section>
            <q-item-label>
              <q-input dense hint="输入验证码"  v-model="DialogAgentLogin.form.Code"/>
            </q-item-label>
          </q-item-section>
        </q-item>
      </q-card-section>
      <q-card-actions   align="right">
        <q-btn flat @click="onClickAgentLoginCode">确定</q-btn>
        <q-btn flat  @click="DialogAgentLogin.Visible=false">取消</q-btn>
      </q-card-actions>
    </q-card>
  </q-dialog>
  <q-dialog    v-model="DialogBetLogin.Visible"  >
    <q-card style="width: 280px">
      <q-toolbar>
        <q-toolbar-title> {{ DialogBetLogin.title }}</q-toolbar-title>
        <q-btn flat round dense icon="close" v-close-popup />
      </q-toolbar>
      <q-separator />
      <q-card-section class="q-pa-none">
        <q-item dense>
          <q-item-section>
            <q-item-label>
              <q-img width="100px" :src="DialogBetLogin.form.CodePicUrl" @click="refreshBetPic(DialogBetLogin.SiteUserInfo.Id)"></q-img>
            </q-item-label>
          </q-item-section>
          <q-item-section>
            <q-item-label>
              <q-input dense hint="输入验证码"  v-model="DialogBetLogin.form.Code"/>
            </q-item-label>
          </q-item-section>
        </q-item>
      </q-card-section>
      <q-card-actions   align="right">
        <q-btn flat @click="onClickBetLoginCode">确定</q-btn>
        <q-btn flat  @click="DialogBetLogin.Visible=false">取消</q-btn>
      </q-card-actions>
    </q-card>
  </q-dialog>
  <q-dialog    v-model="EditDialogAgent.Visible"  >
    <q-card style="width: 600px">
      <q-toolbar>
        <q-toolbar-title> {{ EditDialogAgent.Title }}</q-toolbar-title>
        <q-btn flat round dense icon="close" v-close-popup />
      </q-toolbar>
      <q-separator />


      <q-item dense>
        <q-item-section top class="col-2 gt-sm">
          <q-item-label class="q-mt-sm text-right">标题:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label><q-input v-model="EditDialogAgent.Form.Title" dense /></q-item-label>
        </q-item-section>
      </q-item>

      <q-item dense>
        <q-item-section top class="col-2 gt-sm">
          <q-item-label class="q-mt-sm text-right">网址:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label><q-input v-model="EditDialogAgent.Form.SiteUrl" dense /></q-item-label>
        </q-item-section>
      </q-item>

      <q-item dense>
        <q-item-section top class="col-2 gt-sm">
          <q-item-label class="q-mt-sm text-right">用户名:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label><q-input v-model="EditDialogAgent.Form.UserName" dense /></q-item-label>
        </q-item-section>
      </q-item>

      <q-item dense>
        <q-item-section top class="col-2 gt-sm">
          <q-item-label class="q-mt-sm text-right">密码:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label><q-input v-model="EditDialogAgent.Form.Pwd" dense /></q-item-label>
        </q-item-section>
      </q-item>


      <q-card-actions   align="right">
        <q-btn flat @click="onClickSaveAgent">确定</q-btn>
        <q-btn flat  @click="EditDialogAgent.Visible=false">取消</q-btn>
      </q-card-actions>
    </q-card>
  </q-dialog>
  <q-dialog    v-model="EditDialogBet.Visible"  >
    <q-card style="width: 600px">
      <q-toolbar>
        <q-toolbar-title> {{ EditDialogBet.Title }}</q-toolbar-title>
        <q-btn flat round dense icon="close" v-close-popup />
      </q-toolbar>
      <q-separator />


      <q-item dense>
        <q-item-section top class="col-2 gt-sm">
          <q-item-label class="q-mt-sm text-right">标题:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label><q-input v-model="EditDialogBet.Form.Title" dense /></q-item-label>
        </q-item-section>
      </q-item>

      <q-item dense>
        <q-item-section top class="col-2 gt-sm">
          <q-item-label class="q-mt-sm text-right">网址:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label><q-input v-model="EditDialogBet.Form.SiteUrl" dense /></q-item-label>
        </q-item-section>
      </q-item>

      <q-item dense>
        <q-item-section top class="col-2 gt-sm">
          <q-item-label class="q-mt-sm text-right">用户名:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label><q-input v-model="EditDialogBet.Form.UserName" dense /></q-item-label>
        </q-item-section>
      </q-item>

      <q-item dense>
        <q-item-section top class="col-2 gt-sm">
          <q-item-label class="q-mt-sm text-right">密码:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label><q-input v-model="EditDialogBet.Form.Pwd" dense /></q-item-label>
        </q-item-section>
      </q-item>


      <q-card-actions   align="right">
        <q-btn flat @click="onClickSaveBet">确定</q-btn>
        <q-btn flat  @click="EditDialogBet.Visible=false">取消</q-btn>
      </q-card-actions>
    </q-card>
  </q-dialog>

</template>

<script>
import {GetC,} from "src/api/Api"
import {GetAgentBetData, GetAgentUser, UpdateAgentUser} from "src/api/DataApi_Agent"
import {GetBetUser,UpdateBetUser} from "src/api/DataApi_Bet"
import {AgentLogin, AgentLoginOut, GetAndSaveNumData} from "src/api/AgentSite";
import {BetLogin, BetLoginOut, Bet} from "src/api/BetSite";
import {GetConfig, SaveConfig} from "src/api/DataApi";
export default {
  name: "Bet",
  data(){
    return{
      StockSelectItems:[
        {Id:0,Name:"全部"},
        {Id:1,Name:"实占货量"},
      ],

      BetWaySelects:{
        Rate:1,//比例
        BetM:2,//金额
      },

      RateWaySelects:{
        F:0,//反
        Z:1,//正
      },
      InputData:{
        Datas:[0,0,0,0,0,0]
      },
      IndexData:{
        SelectItems:[
          {Id:1,Name:"1"},
          {Id:2,Name:"2"},
          {Id:3,Name:"3"},
          {Id:4,Name:"4"},
          {Id:5,Name:"5"},
        ],
         Datas:[1,2,3,4,5],
        Datas2:[1,2,3,4,5]
      },
      InputBetData:{
        Datas:[0,
          0,0,0,0,0,0,0,0,0,0,
          0,0,0,0,0,0,0,0,0,0,
          0,0,0,0,0,0,0,0,0,0,
          0,0,0,0,0,0,0,0,0,0,
          0,0,0,0,0,0,0,0,0,0,
        ]
      },
      AgentBetData:[],
      ConfigData:{
        Id:1,
        RateWay:0,
        IsAuto:0,
        BetWay:1,
        Rate:50,
        Stock:0,//占成方式:0:全部,1:实占
        BetTime1:"10:56",
        BetTime2:"10:56",
        LotteryType:1,
      },
      EditDialogAgent:{
        Title:'',
        Visible:false,
        Form:{Id:0,Title:'',UserName:'',Pwd:'',SiteUrl:''
        },
      },
      EditDialogBet:{
        Title:'',
        Visible:false,
        Form:{Id:0,Title:'',UserName:'',Pwd:'',SiteUrl:''},
      },
      AgentUserInfo:{
        UserData:{Id:1,UserName:'test', IsLogin:0}
      },
      BetUserInfo:{
        UserData:{UserName:'test', IsLogin:0}
      },
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
        IsSave:false ,
        refreshDataTimer:null,
        CurResultSecond:0,
        ReloadSiteUserDataSecond:0,
        AgentData:{
          Second:0,
          MaxSecond:15,
        },
      },
      QuickBetM:'',
      defValue:1,
      layout1:[
        {Num:0,V:0,BetM:0,OkBetM:0},
        {Num:1,V:0,BetM:0,OkBetM:0}
      ],

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
          // { Id:3,Name:"超级6+1六合彩"},
          // { Id:4,Name:"闪电6+1六合彩"}
        ]
      },

      DialogAgentLogin:{
        title:'登录',
        isBet:0,
        SiteUserInfo:{},
        Visible:false,
        form:{
          CodePicUrl:'',
          Code:''
        }
      },

      DialogBetLogin:{
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
    onIndexDataHide(i){
      // let  iiIndex  = this.IndexData.Datas[i]
      // if (this.IndexData.Datas[i]!=this.IndexData.Datas2[i]) {
      //   this.IndexData.Datas[iiIndex-1] = this.IndexData.Datas2[i]
      //   for (let j=0;j<4;j++) {
      //     this.IndexData.Datas2[j] =  this.IndexData.Datas[j]
      //   }
      //
      // }
      this.setLayoutData()
    },
    loadAgentUser(){
      let query = {}
      GetAgentUser(query).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
          this.AgentUserInfo.UserData = obj
        } else {
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },
    loadBetUser(){
      let query = {}
      GetBetUser(query).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
          this.BetUserInfo.UserData = obj
        } else {
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },
    loadAgentBetData(){
      let IsSave = 0
      if (this.RefreshData.IsSave) {
        IsSave  = 1
      }
      let query = {AgentId:this.AgentUserInfo.UserData.Id,IsSave:IsSave}
      GetAndSaveNumData(query).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
          if(obj.Msg!=""){
            this.AgentUserInfo.UserData.IsLogin = false
            this.$q.notify(obj.Msg)
          }
          this.AgentBetData = obj.Data
          this.setLayoutData()
          this.$q.notify("刷新数据完成")
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },
    onB(i){
      if (this.ConfigData.BetWay==2 ){
        this.InputBetData.Datas[i] =  this.layout1[i].V
      }
    },
    setLayoutData(){
      let divisor =100
      let newAgentBetData =  {... this.AgentBetData}
      for (let i=1;i<=49;i++) {
        if (this.ConfigData.BetWay==this.BetWaySelects.Rate){
          let j = 50-i
          if (this.ConfigData.RateWay==1){
            j=i
          }
          newAgentBetData[j].V = Math.floor(( newAgentBetData[i].BetM *  this.ConfigData.Rate)/divisor)
        }else {
           newAgentBetData[i].V = this.InputBetData.Datas[i]
        }

        if ( newAgentBetData[i].V==0) {
          newAgentBetData[i].V = null
        }
      }

      let newAgentBetData2 = JSON.parse(JSON.stringify(newAgentBetData))
      console.log("0:")
      console.log(newAgentBetData)
      for (let i=0;i<5;i++) {
        for (let j=1;j<=10;j++) {
          let i10 = this.IndexData.Datas[i]-1
          let ni = i*10+j
          let ni2 = i10*10+j
          if (ni==50||ni2==50) {

          }

          // console.log(ni,ni2)
          newAgentBetData[ni].V = newAgentBetData2[ni2].V
        }
      }
      console.log("1:")
      console.log(newAgentBetData)

      console.log("2:")
      console.log(newAgentBetData2)


      this.layout1 = newAgentBetData
    },
    onClickSetBetM(v){
      let beginI = (v-1) * 10
      let sI = 0
      for (let i=1;i<=10;i++) {
        sI =beginI+i
        this.layout1[sI].V =this.InputData.Datas[v]
        if (this.ConfigData.BetWay==2) {
          this.InputBetData.Datas[sI] = this.layout1[sI].V
        }
        if (this.layout1[sI].V ==0) {
          this.layout1[sI].V = null
        }
      }
    },
    onClickSetBetMBaseRate(v){
      let divisor =100
      let beginI = (v-1) * 10
      let sI = 0
      for (let i=1;i<=10;i++) {
        sI =beginI+i
        this.layout1[sI].V = Math.floor(this.layout1[sI].BetM * this.ConfigData.Rate / divisor)
        if (this.ConfigData.BetWay==2) {
          this.InputBetData.Datas[sI] = this.layout1[sI].V
        }
        if (this.layout1[sI].V ==0) {
          this.layout1[sI].V = null
        }
      }
    },


    openEditDialog(row){
      this.EditDialogAgent.Title = `修改[${row.SiteName}]`
      this.EditDialogAgent.Form.Id = row.Id
      this.EditDialogAgent.Form.SiteName = row.SiteName
      this.EditDialogAgent.Form.SiteTitle = row.SiteTitle
      this.EditDialogAgent.Form.SiteUrl = row.SiteUrl
      this.EditDialogAgent.Form.SiteDes = row.SiteDes

      this.EditDialogAgent.Visible = true
    },


     getNumClass(i){
      return 'num'.concat(i)
    },
    iniLayout1(){
      let layout = []
      for (let i=0;i<=50;i++) {
        let aNum = {Num:i,V:null, BetM:0,OkBetM:0}
        layout.push(aNum)
      }
      this.layout1 = layout
    },


    showAgentLogin(row){
      this.DialogAgentLogin.title = `登录【${row.UserName}】`
      this.DialogAgentLogin.SiteUserInfo = {...row}
        this.DialogAgentLogin.form.Code = ""
      this.DialogAgentLogin.form.CodePicUrl = this.getAgentCodePicUrl(row.Id, true)
      this.DialogAgentLogin.Visible = true
    },
    showBetLogin(row){
      this.DialogBetLogin.title = `登录【${row.UserName}】`
      this.DialogBetLogin.SiteUserInfo = {...row}
      this.DialogBetLogin.form.Code = ""
      this.DialogBetLogin.form.CodePicUrl = this.getBetCodePicUrl(row.Id, true)
      this.DialogBetLogin.Visible = true
    },
    getAgentCodePicUrl(id, f){
      let t1
      if (f) {
        let t = new Date()
        t1 = t.getTime()
      }
      let codeSrc ='/agentsite/getcodeimage?Id='+id+'&t='+t1
      return codeSrc
    },
    onClickAgentLoginCode(){
      const loginP ={
        Id:this.DialogAgentLogin.SiteUserInfo.Id,
        Code:this.DialogAgentLogin.form.Code,
        IsBet:this.DialogAgentLogin.isBet
      }
      AgentLogin(loginP).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
          this.C.HasLogin = true
          this.DialogAgentLogin.Visible = false
          this.loadAgentUser()
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },

    onClickAgentLogout(UserInfo){
      const query ={
        Id:UserInfo.Id
      }
      AgentLoginOut(query).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
          this.loadAgentUser()
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },
    onClickBetLogout(UserInfo){
      const query ={
        Id:UserInfo.Id
      }
      BetLoginOut(query).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
          this.loadBetUser()
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },


    getBetCodePicUrl(id, f){
      let t1
      if (f) {
        let t = new Date()
        t1 = t.getTime()
      }
      let codeSrc ='/betsite/getcodeimage?Id='+id+'&t='+t1
      return codeSrc
    },
    onClickBetLoginCode(){
      const loginP ={
        Id:this.DialogBetLogin.SiteUserInfo.Id,
        Code:this.DialogBetLogin.form.Code,
      }
      BetLogin(loginP).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
          this.C.HasLogin = true
          this.DialogBetLogin.Visible = false
          this.loadBetUser()
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },

    onClickAgentLoginOut(row){
      const loginP ={
        Id:row.Id,
      }
      AgentLoginOut(loginP).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },

    onClickBetLoginOut(row){
      const loginP ={
        Id:row.Id,
      }
      BetLoginOut(loginP).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },

    refreshBetPic(id){
      this.DialogBetLogin.form.CodePicUrl = this.getBetCodePicUrl(id, true)
    },
    refreshAgentPic(id){
      this.DialogAgentLogin.form.CodePicUrl = this.getAgentCodePicUrl(id, true)
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
        this.$q.notify(error)
      })
    },
    autoRefreshData(){

      this.RefreshData.CurResultSecond++
      if (this.RefreshData.CurResultSecond > 1000000){
        this.RefreshData.CurResultSecond = 1
      }
      if(!this.RefreshData.IsSave) {
        return
      }

      this.RefreshData.AgentData.Second++
      if (this.RefreshData.CurResultSecond % this.RefreshData.AgentData.MaxSecond==0) {
        this.loadAgentBetData()
        this.RefreshData.AgentData.Second = 0
      }
    },
    onClickBet(){
      let arrBetData = []
      for (let i = 1; i <50 ; i++) {
        if (this.layout1[i].V>0){
          arrBetData.push(this.layout1[i])
        }
      }
      if( arrBetData.length==0) {
        this.$q.notify("最少有一个金额大于0")
        return
      }
      let BetData =JSON.stringify(arrBetData)
      let LoConfig =JSON.stringify(this.ConfigData)

      let query = {BetData:BetData,LoConfig}
      Bet(query).then(response => {
        const { code, msg, obj } = response
        if (code==200){
          this.loadAgentBetData()
          this.$q.notify("投注成功")
        } else {
          this.$q.notify("投注失败:"+msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },
    onShowEditDialogAgent(){
      this.EditDialogAgent.Title = `修改代理信息`
      this.EditDialogAgent.Form = {...this.AgentUserInfo.UserData}
      this.EditDialogAgent.Visible = true
    },
    onShowEditDialogBet(){
      this.EditDialogBet.Title = `修改会员信息`
      this.EditDialogBet.Form = {...this.BetUserInfo.UserData}
      this.EditDialogBet.Visible = true
    },
    onClickSaveAgent(){
      UpdateAgentUser(this.EditDialogAgent.Form).then(response => {
        const { code, msg, obj } = response
        if (code==200){
          this.loadAgentUser()
          this.EditDialogAgent.Visible = false
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },

    onClickSaveBet(){
      UpdateBetUser(this.EditDialogBet.Form).then(response => {
        const { code, msg, obj } = response
        if (code==200){
          this.loadBetUser()
          this.EditDialogBet.Visible = false
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },
    changeConfigDataBetWay(value){
      let query = {BetWay:value}
      this.updateConfig(query,"")
    },

    changeConfigDataRateWay(value){
      let query = {RateWay:value}
      this.updateConfig(query,"")
    },
    changeConfigDataRate(value){
      let query = {Rate:value}
      this.updateConfig(query,"")
    },

    changeConfigDataStock(value){
      let query = {Stock:value}
      this.updateConfig(query,"")

    },
    changeConfigDataLotteryType(value){
      let query = {LotteryType:value}
      this.updateConfig(query,"")
    },
    loadConfig(){
      let query = {}
      GetConfig(query).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
          this.ConfigData = obj
        } else {
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },
    onClickR(){
      this.loadAgentBetData()
      this.RefreshData.AgentData.Second= 0
      // this.
    },
    updateConfig(query, msgOK){
      SaveConfig(query).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
          if (msgOK!="") {
            this.$q.notify(msgOK)
          }
        } else {
          this.$q.notify(msg)
        }
        this.loadAgentBetData()
      }).catch(error => {
        this.$q.notify(error)
      })
    },
    updateAllConfig(){
      SaveConfig(this.ConfigData).then(response => {
        const { code, msg, obj } = response
        if (code==200) {
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },
  },
  watch:{
    'C.L':function (v){
    },
    'C.LotteryType':function (v){
    }
  },
  mounted() {
    if (this.RefreshData.refreshDataTimer) {
      clearInterval(this.RefreshData.refreshDataTimer)
      this.RefreshData.refreshDataTimer = null;
    }
    this.RefreshData.refreshDataTimer = setInterval(()=>{this.autoRefreshData()},1000);
  },

  unmounted() {
    clearInterval(this.RefreshData.refreshDataTimer);
    this.RefreshData.refreshDataTimer = null
  },
  created() {
    this.getC()
    this.loadConfig()
    this.loadAgentUser()
    this.loadBetUser()
    this.iniLayout1()
    this.loadAgentBetData()
  }

}
</script>

<style scoped lang="scss">
  .input_title {
    font-size: 14px;
  }
  .example-row-equal-width {
    .row > div{
      border: 1px solid
    }
    .row + .row{
      margin-top: -1px
    }
  }
  .input-row-equal-width{
    border: 1px solid
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

  .bet_m{
    margin: 0  0px;

    padding-bottom: 10px;
    padding-top: 10px;
    font-size: 12px;
    align-content: center;
    text-align: right;
    background-color: #DDD;
    color:black;
    width: 60px
  }

  .bet_m2{
    margin: 0  0px;
    padding-bottom: 10px;
    padding-top: 10px;
    font-size: 12px;
    align-content: center;
    text-align: right;
    background-color: #DDD;
    color:black;
    width: 55px
  }

</style>
