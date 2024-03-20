import request from '../utils/request'
export function BetLogin(query) {
  return request({
    url: '/betsite/login',
    method: 'post',
    params: query
  })
}

export function BetLoginOut(query) {
  return request({
    url: '/betsite/loginout',
    method: 'post',
    params: query
  })
}


export function Bet(query) {
  return request({
    url: '/betsite/bet',
    method: 'post',
    params: query
  })
}
