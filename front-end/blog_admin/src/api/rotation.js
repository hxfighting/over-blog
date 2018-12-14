import axios from '@/libs/api.request'


//获取轮播图列表
export const getRotationList = () => {
    return axios.request({
        url: 'admin/rotation',
        method: 'get'
    })
}

//修改轮播图
export const updateRotation = ({id,words,image_url}) => {
    const data = {
        id,
        words,
        image_url
    }
    return axios.request({
        url: 'admin/rotation',
        method: 'put',
        data
    })
}

//添加轮播图
export const addRotation = ({words,image_url}) => {
    const data = {
        words,
        image_url
    }
    return axios.request({
        url: 'admin/rotation',
        method: 'post',
        data
    })
}

//删除轮播图
export const deleteRotation = ({id}) => {
    const data = {
        id
    }
    return axios.request({
        url: 'admin/rotation',
        method: 'delete',
        data
    })
}