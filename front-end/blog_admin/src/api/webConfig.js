import axios from '@/libs/api.request'

//添加配置
export const addConfig = formData => {
    return axios.request({
        url: 'admin/config',
        data:formData,
        method: 'post'
    })
}

//获取配置列表
export const getList = listData => {
    return axios.request({
        url: 'admin/config',
        method: 'get',
        params:listData
    })
}

//修改配置
export const updateConfig = formData => {
    return axios.request({
        url: 'admin/config',
        data:formData,
        method: 'put'
    })
}

//删除配置
export const deleteConfig = ({id}) => {
    const data = {
        id
    }
    return axios.request({
        url: 'admin/config',
        data,
        method: 'delete'
    })
}
