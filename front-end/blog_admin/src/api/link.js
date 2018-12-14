import axios from '@/libs/api.request'


//获取友联列表
export const getLinkList = ({name,pageSize,pageNum}) => {
    const data = {
        name,
        pageSize,
        pageNum
    }
    return axios.request({
        url: 'admin/link',
        method: 'get',
        params:data
    })
}

//修改友联
export const updateLink = ({id,url,name,description,order,is_show}) => {
    const data = {
        id,
        url,
        name,
        description,
        order,
        is_show
    }
    return axios.request({
        url: 'admin/link',
        method: 'put',
        data
    })
}

//添加友联
export const addLink = ({url,name,description,order,is_show}) => {
    const data = {
        url,
        name,
        description,
        order,
        is_show
    }
    return axios.request({
        url: 'admin/link',
        method: 'post',
        data
    })
}

//删除友联
export const deleteLink = ({id}) => {
    const data = {
        id
    }
    return axios.request({
        url: 'admin/link',
        method: 'delete',
        data
    })
}