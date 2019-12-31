import axios from 'axios'
import store from '@/store'
import {getToken, isTokenExpired, setExpire, setToken} from '@/libs/util'
import router from '@/router/index'
import config from '@/config'

const baseUrl = process.env.NODE_ENV === 'development' ? config.baseUrl.dev : config.baseUrl.pro

// import { Spin } from 'iview'
const addErrorLog = errorInfo => {
  const {statusText, status, request: {responseURL}} = errorInfo
  let info = {
    type: 'ajax',
    code: status,
    mes: statusText,
    url: responseURL
  }
  if (!responseURL.includes('save_error_logger')) store.dispatch('addErrorLog', info)
}

//是否有请求正在刷新token
let isRefreshing = false;
/*被挂起的请求数组*/
let refreshSubscribers = []

/*获取刷新token请求的token*/
let getRefreshToken = () => {
  // return JSON.parse(localStorage.auth).refresh_token
}

//push所有请求到数组中
let subscribeTokenRefresh = (cb) => {
  refreshSubscribers.push(cb)
}

//刷新请求(用新的token去请求数据)
let onRrefreshed = (token) => {
  refreshSubscribers.map(cb => cb(token))
}

class HttpRequest {
  constructor(baseUrl = baseURL) {
    this.baseUrl = baseUrl
    this.queue = {}
  }

  getInsideConfig() {
    const config = {
      baseURL: this.baseUrl,
      headers: {
        //
      }
    }
    return config
  }

  destroy(url) {
    delete this.queue[url]
    if (!Object.keys(this.queue).length) {
      // Spin.hide()
    }
  }

  interceptors(instance, url) {
    // 请求拦截
    instance.interceptors.request.use(config => {
      // 添加token到header头
      const token = getToken()
      if (token) {
        console.log(123)
        /*在请求头中添加token类型、token*/
        config.headers.Authorization = token
        /*判断token是否将要过期*/
        if (isTokenExpired()) {
          console.log(456)
          // console.log("token已过期");
          /*判断是否正在刷新*/
          console.log(!isRefreshing)
          if (!isRefreshing) {
            console.log(789)
            // console.log("是否正在刷新");
            /*将刷新token的标志置为true*/
            isRefreshing = true
            /*发起刷新token的请求*/
            axios.get(baseUrl + "admin/token", {
              params: "",
              headers: {"Authorization": token}
            }).then(res => {
              // console.log("刷新token响应");
              // console.log(res);
              let data = res.data;
              if (data.code == -1) {
                setToken("");
                setExpire("");
                router.push({path: '/login'});
              } else {
                isRefreshing = false;
                setToken(data.data.token);
                setExpire(data.data.expire);
                /*更新auth*/
                onRrefreshed(data.data.token);
                /*执行onRefreshed函数后清空数组中保存的请求*/
                refreshSubscribers = []
              }
            }).catch(err => {
              console.log(err)
              setToken("");
              setExpire("");
              router.push({path: '/login'});
            })
          }
          /*把请求(token)=>{....}都push到一个数组中*/
          return new Promise((resolve, reject) => {
            /*(token) => {...}这个函数就是回调函数*/
            subscribeTokenRefresh((token) => {
              config.headers.Authorization = token
              /*将请求挂起*/
              resolve(config)
            })
          })
        }
        return config
      }
      // 添加全局的loading...
      // if (!Object.keys(this.queue).length) {
      //   // Spin.show() // 不建议开启，因为界面不友好
      // }
      console.log(22222)
      return config
    }, error => {
      console.log(error)
      return Promise.reject(error)
    })
    // 响应拦截
    instance.interceptors.response.use(res => {
      this.destroy(url)
      const {data, status} = res

      if (res.data.code == 401) {
        router.push({path: '/login'})
        return res
      }

      if (res.headers.authorization != undefined) {
        setToken(res.headers.authorization)
      }

      return {data, status}
    }, error => {
      this.destroy(url)
      // let errorInfo = error.response
      // if (!errorInfo) {
      //   const {request: {statusText, status}, config} = JSON.parse(JSON.stringify(error))
      //   errorInfo = {
      //     statusText,
      //     status,
      //     request: {responseURL: config.url}
      //   }
      // }
      // addErrorLog(errorInfo)
      return Promise.reject(error)
    })
  }

  request(options) {
    const instance = axios.create()
    options = Object.assign(this.getInsideConfig(), options)
    this.interceptors(instance, options.url)
    return instance(options)
  }
}

export default HttpRequest
