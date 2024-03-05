<template>
  <q-card class="table-bg no-shadow" bordered>
    <q-card-section>
      <div class="text-h6">
        系统用户管理
      </div>
    </q-card-section>
    <q-separator color="white"/>
    <q-card-section class="q-pa-none">
      <q-table :rows="tableData" :columns="columns" separator="cell"
               v-model:pagination="pagination" @request="onRequest">
        <template v-slot:top class="$caption">
          <q-input dense v-model="query.UserName" filled class="q-pr-sm">
            <template v-slot:prepend >
              <span class="text-body2">用户名:</span>
            </template>
          </q-input>
          <q-input dense v-model="query.UserId" filled class="q-pr-sm">
            <template v-slot:prepend >
              <span class="text-body2">用户Id:</span>
            </template>
          </q-input>
<!--          <q-input dense v-model="query.BeginDay" type="date" filled class="q-pr-sm">-->
<!--            <template v-slot:prepend >-->
<!--              <span class="text-body2">开始时间:</span>-->
<!--            </template>-->
<!--          </q-input>-->
<!--          <q-input dense v-model="query.EndDay" type="date" filled class="q-pr-sm">-->
<!--            <template v-slot:prepend>-->
<!--              <span class="text-body2">结束时间:</span>-->
<!--            </template>-->
<!--          </q-input>-->
          <q-btn color="primary" label="搜索" @click="handleSearch" class="q-mx-sm" />
          <q-btn color="primary" label="添加" @click="openAddUserDialog" class="q-mx-sm"/>
        </template>

        <template v-slot:body="props">
          <q-tr :props="props">
            <q-td key="Id" :props="props">
              {{ props.row.Id }}
            </q-td>
            <q-td key="UserName" :props="props">
              {{ props.row.UserName }}
            </q-td>
            <q-td key="NickName" :props="props">
              {{ props.row.NickName }}
            </q-td>
            <q-td key="IsSuper" :props="props">
              {{ props.row.IsSuper }}
            </q-td>
            <q-td key="LoginTime" :props="props">
              {{ props.row.LoginTime }}
            </q-td>
            <q-td key="Operate" :props="props">
              <q-btn color="primary" size="sm" label="修改" class="q-mx-xs" @click="openEditUserDialog(props.row)"/>
              <q-btn color="brown-5" size="sm" label="删除" class="q-mx-xs" @click="onclickDelUser(props.row)"/>
              <q-btn color="brown-5" size="sm" label="修改密码" class="q-mx-xs" @click="onclickUpdatePwd(props.row)"/>
              <q-btn color="brown-5" size="sm" label="设置权限" class="q-mx-xs" @click="openPowerUserDialog(props.row)"/>
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
      <q-card-section class="q-pa-sm row">
        <q-input class="col-lg-4 col-md-4 col-sm-12 col-xs-12 q-pr-sm" v-model="EditDialog.Form.UserName" label="用户名"/>
        <q-input class="col-lg-4 col-md-4 col-sm-12 col-xs-12 q-pr-sm" v-model="EditDialog.Form.NickName" label="昵称"/>
        <q-input v-if="EditDialog.T==0" class="col-lg-4 col-md-4 col-sm-12 col-xs-12 q-pr-sm" v-model="EditDialog.Form.Password" label="密码"/>
      </q-card-section>
      <q-card-actions   align="right">
        <q-btn flat @click="onClickSaveSysUserInfo">确定</q-btn>
        <q-btn flat  @click="EditDialog.Visible=false">取消</q-btn>
      </q-card-actions>@
    </q-card>
  </q-dialog>


  <q-dialog    v-model="EditPowerDialog.Visible"  >
    <q-card style="width: 600px">
      <q-toolbar>
        <q-toolbar-title> {{ EditPowerDialog.Title }}</q-toolbar-title>
        <q-btn flat round dense icon="close" v-close-popup />
      </q-toolbar>
      <q-separator />
      <q-card-section class="q-pa-sm row">
        <q-tree :nodes="EditPowerDialog.Data" node-key="Id"  default-expand-all>

          <template v-slot:default-header="prop">
            <div class="row justify-between" style="width: 400px">
              <div class="col-4">{{ prop.node.title }}</div>
              <div class="col-8 row justify-end">
                <q-option-group inline dense v-model="prop.node.power" :options="EditPowerDialog.options" color="primary"/>
              </div>
            </div>
          </template>
        </q-tree>

      </q-card-section>
      <q-card-actions   align="right">
        <q-btn flat @click="savePower">确定</q-btn>
        <q-btn flat  @click="EditPowerDialog.Visible=false">取消</q-btn>
      </q-card-actions>
    </q-card>
  </q-dialog>

</template>

<script>
import { GetSysUserList, UpdateSysUser, AddSysUser, DelSysUser, UpdateSysUserPwd, GetAllMenuList, SaveMenuPower} from "src/api/SysInfo";

export default {
  name: "SysUserList",
  data(){
    return{
      columns:[
        {name:'Id',label:'ID',field:'Id', align: 'center'},
        {name:'UserName',label:'用户名',field:'UserName', align: 'center'},
        {name:'NickName',label:'昵称',field:'NickName', align: 'right'},
        {name:'IsSuper',label:'超级用户',field:'IsSuper', align: 'right'},
        {name:'LoginTime',label:'登录时间',field:'LoginTime' , align: 'center'},
        {name:'Operate',label:'操作',field:'Operate' , align: 'center'},
      ],
      tableData:[],
      pagination:{
        page:1,
        rowsPerPage:10,
        rowsNumber:0,
      },
      query:{
        UserId:null,
        UserName:'',
        PageIndex:1,
        PageSize:10,
        BeginDay:'',
        EndDay:'',
      },

      EditDialog:{
        T:0,
        Visible:false,
        Title:'',
        Form:{
          Id:0,
          UserName:'',
          NickName:'',
          Password:'',
        },
      },


      EditPowerDialog:{
        group:null,
        options:[
          { label: '无', value: 0 },
          { label: '写', value: 2 },
        ],
        T:0,
        Visible:false,
        Title:'',
        Id:0,
        Data:[],
      }
    }
  },
  methods:{
    // getSysStateName(s){
    //   return SysUserStateName(s)
    // },
    getPowerData(arrData){
      let resArr = []
      // console.log("arrData:",arrData)
      arrData.forEach(item=>{
        // console.log("item:",item)
        let p = {Id:item.Id, Power:item.power}
        resArr.push(p)
        if (item.children!=null){
          let arrChildren = this.getPowerData(item.children)
          resArr.push(...arrChildren)
        }
      })
      return resArr
    },
    savePower(){
      let arr = this.getPowerData(this.EditPowerDialog.Data)
      let StrPower = JSON.stringify(arr)
      let query = {Id: this.EditPowerDialog.Id, StrPower:StrPower}
      SaveMenuPower(query).then(response => {
        const { code, msg, obj } = response
        if (code==200){
          this.EditPowerDialog.Visible = false
          this.$q.notify("保存成功")
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })

    },
    openPowerUserDialog(row){
      this.EditPowerDialog.Id =  row.Id
      this.EditPowerDialog.Title = `设置${row.UserName}菜单权限`
      let query = {Id:row.Id}
      GetAllMenuList(query).then(response => {
        const { code, msg, obj } = response
        if (code==200){
          this.EditPowerDialog.Data = obj
          this.EditPowerDialog.Visible = true
        } else {
          this.$q.notify(msg)
        }
      }).catch(error => {
        this.$q.notify(error)
      })
    },

    onClickSaveSysUserInfo(){
      if (this.EditDialog.T==1) {
        this.updateSysUser()
      } else {
        this.addSysUser()
      }
    },
    updateSysUser(){
      let query = {
        Id:this.EditDialog.Form.Id,
        UserName:this.EditDialog.Form.UserName,
        NickName:this.EditDialog.Form.NickName,
      }
      UpdateSysUser(query).then(response => {
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
    addSysUser(){
      let query = {
        Id:this.EditDialog.Form.Id,
        UserName:this.EditDialog.Form.UserName,
        NickName:this.EditDialog.Form.NickName,
        Password:this.EditDialog.Form.Password,
      }
      AddSysUser(query).then(response => {
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

    openEditUserDialog(row){
      this.EditDialog.T = 1 //0：添加。1：修改
      this.EditDialog.Title = `修改[${row.UserName}]`
      this.EditDialog.Form.Id = row.Id
      this.EditDialog.Form.UserName = row.UserName
      this.EditDialog.Form.NickName = row.NickName
      this.EditDialog.Visible = true
    },
    openAddUserDialog(){
      this.EditDialog.T = 0 //0：添加。1：修改
      this.EditDialog.Title = `添加用户`
      this.EditDialog.Form.Id = 0
      this.EditDialog.Form.UserName = ''
      this.EditDialog.Form.NickName = ''
      this.EditDialog.Form.Pwd = ''
      this.EditDialog.Visible = true
    },
    onclickDelUser(row){
      this.$q.dialog({
        title: '提示',
        message: `确定要删除${row.UserName}吗`,
        ok:"确定",
        cancel:"取消",
      }).onOk(() => {
        let query = {Id:row.Id}
        DelSysUser(query).then(response => {
          const { code, msg, obj } = response
          if (code==200){
            this.$q.notify(`删除${row.UserName}成功`)
            this.reloadTableData()
          } else {
            this.$q.notify(msg)
          }
        }).catch(error => {
          this.$q.notify(error)
        })
      })

    },
    onclickUpdatePwd(row){
      this.$q.dialog({
        title: '提示',
        message: `新的密码是什么？`,
        prompt:{
          model: '',
          type: 'text' // optional
        },
        ok:"确定",
        cancel:"取消",
      }).onOk(data => {
        let query = {Id:row.Id,Password:data}
        UpdateSysUserPwd(query).then(response => {
          const { code, msg, obj } = response
          if (code==200){
            this.$q.notify(`修改成功`)
            this.reloadTableData()
          } else {
            this.$q.notify({message:msg})
          }
        }).catch(error => {
          this.$q.notify({message:error})
        })
      })

    },
    onRequest(props){
      if (props.pagination.rowsPerPage==0) {
        props.pagination.rowsPerPage = 1000
      }
      this.pagination = props.pagination
      this.reloadTableData()
    },
    reloadTableData(){
      this.query.PageIndex = this.pagination.page
      this.query.PageSize = this.pagination.rowsPerPage

      GetSysUserList(this.query).then(response => {
        const { code, msg, obj } = response
        if (code==200){
          this.pagination.rowsNumber = obj.C
          this.tableData = obj.DataList
        }
      }).catch(error => {
      })
    },
    handleSearch(){
      this.reloadTableData(this.query.PageIndex , this.query.PageSize)
    },
    onClickAgree(row){
      let p = {Id:row.Id}
      AgreeDrawMoney(p).then(response => {
        const { code, msg, obj } = response
        if (code==200){
          this.pagination.rowsNumber = obj.C
          this.$q.notify({message:row.UserName+'下分['+row.Gold+']成功'})
          this.reloadTableData()
        } else {
          this.$q.notify({message:msg})
        }
      }).catch(error => {
        this.$q.notify({message:error})
      })
    },


  },
  mounted() {
    this.reloadTableData()
  }
}
</script>

<style scoped>
.table-bg {
  /*background-color: #162b4d;*/
}
</style>
