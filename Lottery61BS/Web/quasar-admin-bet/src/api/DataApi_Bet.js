import request from '../utils/request'

export function GetBetUser(query) {
  return request({
    url: '/dataapi/getbetuser',
    method: 'post',
    params: query
  })
} 

export function UpdateBetUser(query) {
  return request({
    url: '/dataapi/updatebetuser',
    method: 'post',
    params: query
  })
}
