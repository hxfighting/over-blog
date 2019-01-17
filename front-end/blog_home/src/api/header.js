import axios from '../libs/api.request'

//获取轮播图
export const getRotation = () => {
    return axios.request({
        url: 'home/rotation',
        method: 'get'
    })
}

//获取分类
export const getCategory = () => {
    return axios.request({
        url: 'home/category',
        method: 'get'
    })
}
