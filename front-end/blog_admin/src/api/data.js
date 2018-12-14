import axios from '@/libs/api.request'

export const getTableData = () => {
  return axios.request({
    url: 'get_table_data',
    method: 'get'
  })
}

export const getDragList = () => {
  return axios.request({
    url: 'get_drag_list',
    method: 'get'
  })
}

//获取错误日志列表
export const getErrorList = ({pageSize,pageNum}) => {
  const data = {
    pageNum,
    pageSize
  }
  return axios.request({
    url: 'admin/error',
    method: 'get',
    params:data
  })
}

//添加错误日志
export const saveErrorLogger = info => {
  return axios.request({
    url: 'admin/error',
    data: info,
    method: 'post'
  })
}

//删除错误日志
export const deleteErrorLog = ({ids}) => {
  const data = {
    ids
  }
  return axios.request({
    url: 'admin/error',
    data: data,
    method: 'delete'
  })
}

export const uploadImg = formData => {
  return axios.request({
    url: 'image/upload',
    data: formData
  })
}

//获取验证码
export const getCaptcha = () => {
    return axios.request({
        url: '/captcha',
        method: 'get'
    })
}
