import { Login } from 'src/api/RegLogin'
import { getToken, setToken, removeToken } from 'src/utils/auth'
import {GetSysMenuList, GetUserInfo} from "src/api/SysInfo";

let user = {
  namespaced:true,
  state : {
    token:getToken(),
    UserInfo:{
    },
    sysMenuList:[]
  },
  mutations : {
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_MenuList: (state, sysMenuList) => {
      state.sysMenuList = sysMenuList
    },
    SET_UserInfo: (state, UserInfo) => {
      if (UserInfo!=null) {
        state.UserInfo = UserInfo
      }
    },
  },
  // getters:{
  //   getSysMenuList: state=>{
  //     return state.sysMenuList
  //   }
  // },

  actions : {
    login({ commit }, userInfo) {
      return new Promise((resolve, reject) => {
        Login(userInfo).then(response => {
          const { code, msg, obj } = response
          if (code==200){
            commit('SET_TOKEN', obj.Token)
            commit('SET_MenuList', obj.MenuList)
            setToken(obj.Token)

          }
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // get user info
    getUserInfo({ commit, state }) {
      return new Promise((resolve, reject) => {
        GetUserInfo().then(response => {
          const { code, msg, obj } = response
          if (code!=200) {
            reject(msg)
          }
          commit('SET_UserInfo', obj)
          resolve(obj)
        }).catch(error => {
          reject(error)
        })
      })
    },


    loadMenu({ commit }, info){
      return new Promise((resolve, reject) => {
        GetSysMenuList(info).then(response => {
          const { code, msg, obj } = response
          if (code==200){
            commit('SET_MenuList', obj)
          }
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    }
  }

}
export default user;
