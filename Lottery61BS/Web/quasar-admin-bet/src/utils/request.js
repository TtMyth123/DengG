import axios from 'axios'
import { getToken } from 'src/utils/auth'
import router from '../router';

const token_key = "AdminToken"

const service = axios.create({
   // baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  // timeout: 15000 // request timeout
})

// request interceptor
service.interceptors.request.use(
  config => {
    config.headers[token_key] = getToken()
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
  */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    const res = response.data
    const { code, msg, obj } = res
    if (code==401) {
      console.log("授权有问题1.",msg)
      console.log("授权有问题3")
      router.push({
        path: "/Login",
        query: {msg: msg}
      })

      console.log("授权有问题2")
    }
    return res
  },
  error => {
    console.log('err' + error) // for debug
    return Promise.reject(error)
  }
)

export default service
