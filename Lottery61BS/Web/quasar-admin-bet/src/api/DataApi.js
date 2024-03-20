import request from '../utils/request'

export function GetConfig(query) {
  return request({
    url: '/dataapi/getconfig',
    method: 'post',
    params: query
  })
}

export function SaveConfig(query) {
  return request({
    url: '/dataapi/saveconfig',
    method: 'post',
    params: query
  })
}

