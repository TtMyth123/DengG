import request from '../utils/request'
export function AgentLogin(query) {
  return request({
    url: '/agentsite/login',
    method: 'post',
    params: query
  })
}


export function AgentLoginOut(query) {
  return request({
    url: '/agentsite/loginout',
    method: 'post',
    params: query
  })
}

export function GetAndSaveNumData(query) {
  return request({
    url: '/agentsite/getandsavenumdata',
    method: 'post',
    params: query
  })
}

