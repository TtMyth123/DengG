import request from '../utils/request'


export function GetC(query) {
  return request({
    url: '/dataapi/getc',
    method: 'post',
    params: query
  })
}
export function SaveSiteUserInfo(query) {
  return request({
    url: '/dataapi/savesiteuserinfo',
    method: 'post',
    params: query
  })
}
export function DelSiteUserInfo(query) {
  return request({
    url: '/dataapi/delsiteuserinfo',
    method: 'post',
    params: query
  })
}
export function GetSiteTypeList(query) {
  return request({
    url: '/dataapi/getsitetypelist',
    method: 'get',
    params: query
  })
}


export function GetSiteUserInfoList(query) {
  return request({
    url: '/dataapi/getsiteuserinfolist',
    method: 'get',
    params: query
  })
}


export function SixLogin(query) {
  return request({
    url: '/api/login',
    method: 'post',
    params: query
  })
}
export function SixLoginOut(query) {
  return request({
    url: '/api/loginout',
    method: 'post',
    params: query
  })
}

export function ModifyLoginPwd(query) {
  return request({
    url: '/sysapi/modifyloginpwd',
    method: 'post',
    params: query
  })
}



export function UpdateBaseUrl(query) {
  return request({
    url: '/dataapi/updatebaseurl',
    method: 'post',
    params: query
  })
}


export function Bet(query) {
  return request({
    url: '/dataapi/bet',
    method: 'post',
    params: query
  })
}

export function SetBetState(query) {
  return request({
    url: '/dataapi/setbetstate',
    method: 'post',
    params: query
  })
}


export function SaveSiteType(query) {
  return request({
    url: '/dataapi/savesitetype',
    method: 'post',
    params: query
  })
}


export function GetBetSiteDataList(query) {
  return request({
    url: '/dataapi/getbetsitedatalist',
    method: 'post',
    params: query
  })
}


export function DelBetSiteData(query) {
  return request({
    url: '/dataapi/delbetsitedata',
    method: 'post',
    params: query
  })
}


export function SetLotteryType(query) {
  return request({
    url: '/dataapi/setlotterytype',
    method: 'post',
    params: query
  })
}


export function GetResultHistory(query) {
  return request({
    url: '/dataapi/getresulthistory',
    method: 'post',
    params: query
  })
}
