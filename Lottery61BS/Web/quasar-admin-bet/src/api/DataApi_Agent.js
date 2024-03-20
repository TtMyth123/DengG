import request from '../utils/request'

export function GetAgentUser(query) {
  return request({
    url: '/dataapi/getagentuser',
    method: 'post',
    params: query
  })
}

export function GetAgentBetData(query) {
  return request({
    url: '/dataapi/getagentbetdata',
    method: 'post',
    params: query
  })
}


export function UpdateAgentUser(query) {
  return request({
    url: '/dataapi/updateagentuser',
    method: 'post',
    params: query
  })
}
