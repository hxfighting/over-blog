import axios from 'axios'
import router from '../router/index'

class HttpRequest {
  constructor (baseUrl = baseURL) {
    this.baseUrl = baseUrl
    this.queue = {}
  }

  getInsideConfig () {
    const config = {
      baseURL: this.baseUrl,
      headers: {
        //
      }
    }
    return config
  }

  destroy (url) {
    delete this.queue[url]
    if (!Object.keys(this.queue).length) {
      // Spin.hide()
    }
  }

  interceptors (instance, url) {
    // 请求拦截
    instance.interceptors.request.use(config => {
      // 添加token到header头
      if (localStorage.getItem('token')) {
        config.headers['Authorization'] = localStorage.getItem('token')
      }
      this.queue[url] = true
      return config
    }, error => {
      return Promise.reject(error)
    })
    // 响应拦截
    instance.interceptors.response.use(res => {
      this.destroy(url)
      const {data, status} = res

      if (res.status == 401) {
        router.push({path: '/index'})
        return res
      }

      if (res.headers.authorization != undefined) {
        localStorage.setItem('token',res.headers.authorization)
      }
      return {data, status}
    }, error => {
      this.destroy(url)
      return Promise.reject(error)
    })
  }

  request (options) {
    const instance = axios.create()
    options = Object.assign(this.getInsideConfig(), options)
    this.interceptors(instance, options.url)
    return instance(options)
  }
}

export default HttpRequest
