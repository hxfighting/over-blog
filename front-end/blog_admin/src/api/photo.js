import axios from '@/libs/api.request'


//获取照片列表
export const getPhotoList = () => {
    return axios.request({
        url: 'admin/photo',
        method: 'get'
    })
}

//修改照片
export const updatePhoto = ({id,image_url}) => {
    const data = {
        id,
        image_url
    }
    return axios.request({
        url: 'admin/photo',
        method: 'put',
        data
    })
}

//添加照片
export const addPhoto = ({image_url}) => {
    const data = {
        image_url
    }
    return axios.request({
        url: 'admin/photo',
        method: 'post',
        data
    })
}

//删除照片
export const deletePhoto = ({id}) => {
    const data = {
        id
    }
    return axios.request({
        url: 'admin/photo',
        method: 'delete',
        data
    })
}