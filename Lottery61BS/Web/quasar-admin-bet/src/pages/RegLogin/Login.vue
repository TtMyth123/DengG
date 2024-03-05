<template>
  <q-layout>
    <q-page-container>
      <q-page class="flex bg-image flex-center">
        <q-card v-bind:style="$q.screen.lt.sm?{'width': '80%'}:{'width':'30%'}">
          <q-card-section>
            <q-avatar size="103px" class="absolute-center shadow-10">
              <img src="profile.svg">
            </q-avatar>
          </q-card-section>
          <q-card-section>
            <div class="text-center q-pt-lg">
              <div class="col text-h6 ellipsis">
                {{ $store.state.sys.title }}登录
              </div>
            </div>
          </q-card-section>
          <q-card-section>
            <q-form class="q-gutter-md" >
              <q-input filled v-model="loginData.username" label="用户名" lazy-rules/>
              <q-input type="password" filled v-model="loginData.password" label="密码" lazy-rules />
              <div>

              </div>
            </q-form>
          </q-card-section>

          <q-card-actions class="row justify-between">
            <q-btn label="登录"    color="primary" @click="onClickLogin"/>
            <q-checkbox  style="margin: 0px" v-model="rememberMe" label="记住我的账号" color="teal" />
          </q-card-actions>
        </q-card>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>
import {
  GetCookie
} from "src/api/cdndefend_js_cookie"
import {defineComponent} from 'vue'
import {ref} from 'vue'
const keyName = "loginData.UserName"
const keyPwd = "loginData.Pwd"
export default defineComponent({
  name: "login",
  data() {
    return{
      rememberMe:true,
      loginData:{
        username:'',
        password:'',
      }
    }
  },
  mounted() {
    this.loginData.username = this.$q.localStorage.getItem(keyName)
    this.loginData.password = this.$q.localStorage.getItem(keyPwd)

    if (this.$route.query.msg!==undefined) {
      this.$q.notify(this.$route.query.msg)
    }

  },

  methods:{
    onClickLogin(){
      let bb = GetCookie("00A99084001B4A82F4E45837B217ED20AFECF4EF");
      this.$q.dialog({
        title: '提示',
        message: bb,
        ok: "确定"
      })
      return
      let _this = this
      this.$store.dispatch('user/login',this.loginData)
        .then((res)=>{
          const { code, msg, obj } = res
          if (code==200) {
            if (this.rememberMe) {
              this.$q.localStorage.set(keyName,this.loginData.username)
              this.$q.localStorage.set(keyPwd,this.loginData.password)
            }else{
              this.$q.localStorage.remove(keyName)
              this.$q.localStorage.remove(keyPwd)
            }
            _this.$router.push("/")
          } else {
            _this.$q.dialog({
              title: '提示',
              message: msg,
              ok: "确定"
            })
          }
        })
        .catch(err => {
          this.$q.dialog({
            title: '提示',
            message: err,
            ok: "确定"
          })
        })
    }
  },
  setup() {
    return {
      username: ref('admin'),
      password: ref('admin')
    }
  },
})
</script>

<style>

.bg-image {
  background-image: linear-gradient(135deg, #7028e4 0%, #e5b2ca 100%);
}
</style>
