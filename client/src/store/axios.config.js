import axios from 'axios'
import router from '../router'

// axios.defaults.baseURL = 'http://localhost:8000/api/v1'
// axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*'
// axios.defaults.headers.post['Content-Type'] = 'application/json;charset=UTF-8'
axios.defaults.baseURL = '/api/v1'

axios.interceptors.request.use(
  function (config) {
    if (['post', 'put', 'delete'].includes(config.method))
      console.log([config.method, config.url, config.data])
    return config
  },
  function (error) {
    console.log(error)
    return Promise.reject(error)
  }
)

axios.interceptors.response.use(
  function (response) {
    return response.data
  },
  function (error) {
    if (router.currentRoute._value.fullPath != '/login') {
      console.log(error)
      if (error.response.status > 401) {
        // router.replace({ name: 'Login' })
      }
    }
    return Promise.reject(error).catch(() => {})
  }
)
