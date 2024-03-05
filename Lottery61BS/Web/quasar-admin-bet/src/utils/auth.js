import { Cookies } from 'quasar'
const TokenKey = 'AdminToken'
const UserInfoKey = 'UserInfoKey'

export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}




export function getUserInfo() {
  return Cookies.get(UserInfoKey)
}

export function setUserInfo(UserInfo) {
  return Cookies.set(UserInfoKey, UserInfo)
}

export function removeUserInfo() {
  return Cookies.remove(UserInfoKey)
}
