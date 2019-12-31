import axios from '@/libs/api.request'


//获取分类列表
export const getCategoryList = () => {
    return axios.request({
        url: 'admin/category',
        method: 'get'
    })
}

//修改分类
export const updateCategory = ({id,title}) => {
    const data = {
        id,
        title
    }
    return axios.request({
        url: 'admin/category',
        method: 'put',
        data
    })
}

//添加分类
export const addCategory = ({pid,title}) => {
    const data = {
        pid,
        title
    }
    return axios.request({
        url: 'admin/category',
        method: 'post',
        data
    })
}

//删除分类
export const deleteCategory = ({id}) => {
    const data = {
        id
    }
    return axios.request({
        url: 'admin/category',
        method: 'delete',
        data
    })
}