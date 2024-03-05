import request from '../utils/request'


export function Login(query){
  return request({
    url: '/loginreg/dologin',
    method: 'post',
    params: query
  })
}

export function Logout(query){
  return request({
    url: '/loginreg/dologout',
    method: 'post',
    params: query
  })
}
