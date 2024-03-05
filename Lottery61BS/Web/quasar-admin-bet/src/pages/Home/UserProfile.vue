<template>
  <q-page class="q-pa-sm">
    <div class="row q-col-gutter-sm">
      <div class="col-lg-8 col-md-8 col-xs-12 col-sm-12">
        <q-card class="card-bg text-white no-shadow" bordered>
          <q-card-section class="text-h6 ">
            <div class="text-h6">修改 我的信息</div>
<!--            <div class="text-subtitle2">Complete your profile</div>-->
          </q-card-section>
          <q-card-section class="q-pa-sm">
            <q-list class="row">
              <q-item class="col-lg-12 col-md-12 col-sm-12 col-xs-12">
                <q-item-section side>
                  <q-avatar size="100px">
                    <img src="https://cdn.quasar.dev/img/boy-avatar.png">
                  </q-avatar>
                </q-item-section>
                <q-item-section>
<!--                  <q-btn label="更改头像" class="text-capitalize" rounded color="info" style="max-width: 120px"></q-btn>-->
                </q-item-section>
              </q-item>

              <q-item class="col-lg-6 col-md-6 col-sm-12 col-xs-12">
                <q-item-section>
                  <q-input dark color="white" disable dense v-model="user_details.UserName" label="用户名"/>
                </q-item-section>
              </q-item>
              <q-item class="col-lg-6 col-md-6 col-sm-12 col-xs-12">
                <q-item-section>
                  <q-input dark color="white" disable dense v-model="user_details.NickName" label="昵称"/>
                </q-item-section>
              </q-item>
            </q-list>
          </q-card-section>
<!--          <q-card-actions align="right">-->
<!--            <q-btn class="text-capitalize bg-info text-white" @click="onClickUpdateSysUserInfo">更新用户信息</q-btn>-->
<!--          </q-card-actions>-->
        </q-card>
      </div>
      <div class="col-lg-8 col-md-8 col-xs-12 col-sm-12">
        <q-card class="card-bg text-white no-shadow" bordered>
          <q-card-section class="text-h6 q-pa-sm">
            <div class="text-h6">修改密码</div>
          </q-card-section>
          <q-card-section class="q-pa-sm row">
            <q-item class="col-lg-4 col-md-4 col-sm-12 col-xs-12">
              <q-item-section  class="text-right">
                原密码:
              </q-item-section>
            </q-item>
            <q-item class="col-lg-8 col-md-8 col-sm-12 col-xs-12">
              <q-item-section>
                <q-input type="password" dark dense outlined color="white" round
                         v-model="password_dict.OldPwd" />
              </q-item-section>
            </q-item>
            <q-item class="col-lg-4 col-md-4 col-sm-12 col-xs-12">
              <q-item-section  class="text-right">
                新密码:
              </q-item-section>
            </q-item>
            <q-item class="col-lg-8 col-md-8 col-sm-12 col-xs-12">
              <q-item-section>
                <q-input type="password" dark dense outlined color="white" round v-model="password_dict.NewPwd"/>
              </q-item-section>
            </q-item>
            <q-item class="col-lg-4 col-md-4 col-sm-12 col-xs-12">
              <q-item-section  class="text-right">
                确认新密码:
              </q-item-section>
            </q-item>
            <q-item class="col-lg-8 col-md-8 col-sm-12 col-xs-12">
              <q-item-section>
                <q-input type="password" dark dense outlined round color="white"
                         v-model="password_dict.NewPwd2" />
              </q-item-section>
            </q-item>
          </q-card-section>
          <q-card-actions align="right">
            <q-btn class="text-capitalize bg-info text-white" @click="onClickModifyPwd">修改密码</q-btn>
          </q-card-actions>

        </q-card>
      </div>
    </div>
  </q-page>
</template>

<script>
import {defineComponent} from 'vue'
import {ModifyPwd, GetSysUserInfo, UpdateSysUser} from "src/api/SysInfo";
export default defineComponent({
  name: "UserProfile",
  data() {
    return {
      user_details: {
        GameId:0,
        StaffId:0,
        NickName:'',
        UserName:'',
      },
      password_dict: {
        OldPwd:'',
        NewPwd:'',
        NewPwd2:''
      }
    }
  },
  methods:{
    reloadData(){
      GetSysUserInfo().then(response => {
        const { code, msg, obj } = response
        if (code==200){
          this.user_details  = obj
        }
      }).catch(error => {
      })

    },
    onClickModifyPwd(){
      if (this.password_dict.NewPwd!=this.password_dict.NewPwd2){
        this.$q.notify("确认新密码不正确")
        return
      }

      let query ={OldPwd:this.password_dict.OldPwd, NewPwd:this.password_dict.NewPwd}
      ModifyPwd(query).then(response => {
        const { code, msg, obj } = response
        if (code==200){
          this.$q.notify("修改成功")
        } else {
          this.$q.notify("修改失败："+msg)
        }
      }).catch(error => {
      })
    },

    onClickUpdateSysUserInfo(){
      UpdateSysUser(this.user_details).then(response => {
        const { code, msg, obj } = response
        if (code==200){
          this.$q.notify("修改成功")
        } else {
          this.$q.notify("修改失败："+msg)
        }
      }).catch(error => {
      })
    },
  },
  mounted() {
    this.reloadData()
  },
  setup() {
    return {
      // user_details: {},
      // password_dict: {
      //   OldPwd:'',
      //   NewPwd:'',
      //   NewPwd2:''
      // }
    }
  }
})
</script>

<style scoped>

.card-bg {
  background-color: #162b4d;
}
</style>
