<template>
  <q-card class="table-bg no-shadow" bordered>
    <q-card-section>
      <div class="text-h6">
        站点用户
      </div>
    </q-card-section>
    <q-separator color="white"/>
    <q-card-section class="q-pa-none">
      <q-table :rows="tableInfo.rows" :columns="tableInfo.columns" separator="cell"
               v-model:pagination="tableInfo.pagination" @request="onRequest">
        <template v-slot:top class="$caption">
          <q-btn color="primary" label="添加" @click="openAddDialog" class="q-mx-sm"/>
        </template>

        <template v-slot:body="props">
          <q-tr :props="props">
            <q-td key="Id" :props="props">
              {{ props.row.Id }}
            </q-td>
            <q-td key="SiteName" :props="props">
              {{ props.row.SiteName }}
            </q-td>
            <q-td key="Title" :props="props">
              {{ props.row.Title }}
            </q-td>
            <q-td key="SiteUrl" :props="props">
              {{ props.row.SiteUrl }}
            </q-td>

            <q-td key="UserName" :props="props"  auto-width>
              {{ props.row.UserName }}
            </q-td>
            <q-td key="Rate" :props="props"  auto-width>
              {{ (props.row.Rate*100).toFixed(0) }}%
            </q-td>
            <q-td key="Operate" :props="props" auto-width>
              <q-btn color="primary" size="sm" label="修改" class="q-mx-xs" @click="showEditDialog(props.row)"/>
              <q-btn color="primary" size="sm" label="删除" class="q-mx-xs" @click="showDelSiteInfoDialog(props.row)"/>
            </q-td>
          </q-tr>

        </template>
      </q-table>
    </q-card-section>
  </q-card>

  <q-dialog    v-model="EditDialog.Visible"  >
    <q-card style="width: 600px">
      <q-toolbar>
        <q-toolbar-title> {{ EditDialog.Title }}</q-toolbar-title>
        <q-btn flat round dense icon="close" v-close-popup />
      </q-toolbar>
      <q-separator />

<!--      <q-item-label header>名称</q-item-label>-->
      <q-item dense>
        <q-item-section top class="col-2 gt-sm">
          <q-item-label class="q-mt-sm text-right">类别:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label>
            <q-select v-model="EditDialog.selectType" :options="EditDialog.SiteTypeList"
                      @update:model-value="changeSiteType"
                      dense
                      option-value="Id"
                      option-label="SiteTitle"
                      option-disable="inactive"
                      emit-value
                      map-options />
          </q-item-label>
        </q-item-section>
      </q-item>

      <q-item dense>
        <q-item-section top class="col-2 gt-sm">
          <q-item-label class="q-mt-sm text-right">标题:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label><q-input v-model="EditDialog.Form.Title" dense /></q-item-label>
        </q-item-section>
      </q-item>

      <q-item dense>
        <q-item-section top class="col-2 gt-sm">
          <q-item-label class="q-mt-sm text-right">网址:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label><q-input v-model="EditDialog.Form.SiteUrl" dense /></q-item-label>
        </q-item-section>
      </q-item>

      <q-item dense>
        <q-item-section top class="col-2 gt-sm">
          <q-item-label class="q-mt-sm text-right">用户名:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label><q-input v-model="EditDialog.Form.UserName" dense /></q-item-label>
        </q-item-section>
      </q-item>

      <q-item dense>
        <q-item-section top class="col-2 gt-sm">
          <q-item-label class="q-mt-sm text-right">密码:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label><q-input v-model="EditDialog.Form.Pwd" dense /></q-item-label>
        </q-item-section>
      </q-item>

      <q-item dense>
        <q-item-section top class="col-2 gt-sm">
          <q-item-label class="q-mt-sm text-right">投注比:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label><q-input v-model.number="EditDialog.Form.Rate" dense /></q-item-label>
        </q-item-section>
      </q-item>

      <q-card-actions   align="right">
        <q-btn flat @click="onClickSaveInfo">确定</q-btn>
        <q-btn flat  @click="EditDialog.Visible=false">取消</q-btn>
      </q-card-actions>
    </q-card>
  </q-dialog>


</template>

<script>
import {GetSiteUserInfoList, GetSiteTypeList , DelSiteUserInfo, SaveSiteUserInfo} from "src/api/Api"
import {DelSysUser} from "src/api/SysInfo";
export default {
  name: "SixSiteType",
  data(){
    return{
      tableInfo:{
        columns:[
          {name:'Id',label:'ID',field:'Id', align: 'center'},
          {name:'SiteName',label:'站点名',field:'SiteName' , align: 'center'},
          {name:'Title',label:'标题',field:'Title', align: 'center'},
          {name:'SiteUrl',label:'网址',field:'SiteUrl' , align: 'center'},
          {name:'UserName',label:'账号',field:'UserName', align: 'center'},
          {name:'Rate',label:'投注比',field:'Rate' , align: 'center'},
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
        }
      },
      EditDialog:{
        Title:'',
        Visible:false,
        Form:{Id:0,SiteId:'',SiteUrl:'',Title:'',UserName:'',Pwd:'',Rate:0.0, IsBet:1},
        selectType:0,
        SiteTypeList:[],
      }
    }
  },
  methods:{
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

      GetSiteUserInfoList(this.tableInfo.query).then(response => {
        const { code, msg, obj } = response
        if (code==200){
          this.tableInfo.pagination.rowsNumber = obj.GroupData.C
          this.tableInfo.rows = obj.ListData
        }
      }).catch(error => {
      })
    },
    showDelSiteInfoDialog(row){
      this.$q.dialog({
        title: '提示',
        message: `确定要删除[${row.Title}]吗`,
        ok:"确定",
        cancel:"取消",
      }).onOk(() => {
        let query = {Id:row.Id}
        DelSiteUserInfo(query).then(response => {
          const { code, msg, obj } = response
          if (code==200){
            this.$q.notify(`删除${row.Title}成功`)
            this.reloadTableData()
          } else {
            this.$q.notify(msg)
          }
        }).catch(error => {
          this.$q.notify(error)
        })
      })
    },
    changeSiteType(value){
      this.EditDialog.Form.SiteId = value
      this.EditDialog.SiteTypeList.forEach(item => {
        if (item.Id ==value){
          this.EditDialog.Form.SiteUrl =  item.SiteUrl
          return
        }
      })
    },
    showEditDialog(row){
      this.EditDialog.Title = `修改[${row.SiteName}]`
      this.EditDialog.selectType = row.SiteId
      this.EditDialog.Form.Id = row.Id
      this.EditDialog.Form.SiteId = row.SiteId
      this.EditDialog.Form.Title = row.Title
      this.EditDialog.Form.SiteUrl = row.SiteUrl
      this.EditDialog.Form.SiteDes = row.SiteDes
      this.EditDialog.Form.UserName = row.UserName
      this.EditDialog.Form.Pwd = row.Pwd
      this.EditDialog.Form.Rate =  (row.Rate*100).toFixed(0)

      this.EditDialog.Visible = true
    },

    openAddDialog(){
      this.EditDialog.Form.Id = 0
      if (this.EditDialog.SiteTypeList.length>0){
        this.EditDialog.selectType =this.EditDialog.SiteTypeList[0].Id
        this.EditDialog.Form.SiteId = this.EditDialog.SiteTypeList[0].Id
        this.EditDialog.Form.SiteUrl = this.EditDialog.SiteTypeList[0].SiteUrl
      } else{
        this.EditDialog.Form.SiteId = 0
        this.EditDialog.Form.SiteUrl = ''
      }

      this.EditDialog.Form.Title = ''
      this.EditDialog.Form.SiteDes = ""
      this.EditDialog.Form.UserName = ""
      this.EditDialog.Form.Pwd = ""
      this.EditDialog.Form.Rate = 100
      this.EditDialog.Title = "添加站点用户信息"

      this.EditDialog.Visible = true
    },

    onClickSaveInfo(){
      this.EditDialog.Form.Rate = this.EditDialog.Form.Rate / 100
      this.EditDialog.Form.SiteId = this.EditDialog.selectType

      SaveSiteUserInfo(this.EditDialog.Form).then(response => {
        const { code, msg, obj } = response
        if (code==200){
          this.reloadTableData()
          this.EditDialog.Visible = false
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },
    getSiteType(){
      GetSiteTypeList({}).then(response => {
        const { code, msg, obj } = response
        this.EditDialog.SiteTypeList = obj.ListData
      }).catch(error => {
      })
    },
  },
  mounted() {
    this.getSiteType()
    this.reloadTableData()
  }

}
</script>

<style scoped>

</style>
