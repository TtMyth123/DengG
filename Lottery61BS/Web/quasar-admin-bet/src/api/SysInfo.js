import request from '../utils/request'


export function ModifyPwd(query){
  return request({
    url: '/sys/modifypwd',
    method: 'post',
    params: query
  })
}

export function GetSysUserInfo(query){
  return request({
    url: '/sys/getsysuserinfo',
    method: 'post',
    params: query
  })
}

export function GetSysMenuList(query){
  return request({
    url: '/sys/getmenulist',
    method: 'post',
    params: query
  })
}

export function GetSysUserList(query){
  return request({
    url: '/sys/getsysuserlist',
    method: 'post',
    params: query
  })
}

export function UpdateSysUser(query){
  return request({
    url: '/sys/updatesysuser',
    method: 'post',
    params: query
  })
}

export function AddSysUser(query){
  return request({
    url: '/sys/addsysuser',
    method: 'post',
    params: query
  })
}

export function DelSysUser(query){
  return request({
    url: '/sys/delsysuser',
    method: 'post',
    params: query
  })
}
export function UpdateSysUserPwd(query){
  return request({
    url: '/sys/updatesysuserpwd',
    method: 'post',
    params: query
  })
}

export function GetAllMenuList(query){
  return request({
    url: '/sys/getallmenulist',
    method: 'post',
    params: query
  })
}


export function SaveMenuPower(query){
  return request({
    url: '/sys/savemenupower',
    method: 'post',
    params: query
  })
}

export function GetUserInfo(query){
  return request({
    url: '/sys/getsysuserinfo',
    method: 'post',
    params: query
  })
}


