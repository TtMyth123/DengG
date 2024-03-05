<template>
  <q-card class="table-bg no-shadow" bordered>
    <q-card-section>
      <div class="text-h6">
        站点类型
      </div>
    </q-card-section>
    <q-separator color="white"/>
    <q-card-section class="q-pa-none">
      <q-table :rows="tableInfo.rows" :columns="tableInfo.columns" separator="cell"
               v-model:pagination="tableInfo.pagination" @request="onRequest">


        <template v-slot:body="props">
          <q-tr :props="props">
            <q-td key="Id" :props="props">
              {{ props.row.Id }}
            </q-td>
            <q-td key="SiteName" :props="props">
              {{ props.row.SiteName }}
            </q-td>
            <q-td key="SiteTitle" :props="props">
              {{ props.row.SiteTitle }}
            </q-td>
            <q-td key="SiteDes" :props="props">
              {{ props.row.SiteDes }}
            </q-td>

            <q-td key="SiteUrl" :props="props">
              {{ props.row.SiteUrl }}
            </q-td>
            <q-td key="UpdatedAt" :props="props">
              {{ props.row.UpdatedAt }}
            </q-td>
            <q-td key="Operate" :props="props">
              <q-btn color="primary" size="sm" label="修改" class="q-mx-xs" @click="openEditDialog(props.row)"/>
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
          <q-item-label class="q-mt-sm text-right">名称:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label><q-input v-model="EditDialog.Form.SiteName" dense /></q-item-label>
        </q-item-section>
      </q-item>

      <q-item dense>
        <q-item-section top class="col-2 gt-sm">
          <q-item-label class="q-mt-sm text-right">标题:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label><q-input v-model="EditDialog.Form.SiteTitle" dense /></q-item-label>
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
          <q-item-label class="q-mt-sm text-right">说明:</q-item-label>
        </q-item-section>
        <q-item-section>
          <q-item-label><q-input v-model="EditDialog.Form.SiteDes" dense /></q-item-label>
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
import {GetSiteTypeList, SaveSiteType} from "src/api/Api"
export default {
  name: "SixSiteType",
  data(){
    return{
      tableInfo:{
        columns:[
          {name:'Id',label:'ID',field:'Id', align: 'center'},
          {name:'SiteName',label:'名称',field:'SiteName', align: 'center'},
          {name:'SiteTitle',label:'标题',field:'SiteTitle', align: 'right'},
          {name:'SiteDes',label:'描述',field:'SiteDes', align: 'right'},
          {name:'SiteUrl',label:'网址',field:'SiteUrl' , align: 'center'},
          {name:'UpdatedAt',label:'更新时间',field:'UpdatedAt' , align: 'center'},

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
        Form:{SiteName:'',SiteTitle:'',SiteUrl:'',SiteDes:'',State:0,Id:0},
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

    onClickSaveInfo(){
      SaveSiteType(this.EditDialog.Form).then(response => {
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
  },
  mounted() {
    this.reloadTableData()
  }

}
</script>

<style scoped>

</style>
