import axios from '@/libs/api.request'


//获取标签列表
export const getTagList = ({pageNum,pageSize}) => {
    const data = {
        pageNum,
        pageSize
    }
    return axios.request({
        url: 'admin/tag',
        method: 'get',
        params:data
    })
}

//修改标签
export const updateTag = ({id,name}) => {
    const data = {
        id,
        name
    }
    return axios.request({
        url: 'admin/tag',
        method: 'put',
        data
    })
}

//添加标签
export const addTag = ({name}) => {
    const data = {
        name
    }
    return axios.request({
        url: 'admin/tag',
        method: 'post',
        data
    })
}

//删除标签
export const deleteTag = ({id}) => {
    const data = {
        id
    }
    return axios.request({
        url: 'admin/tag',
        method: 'delete',
        data
    })
}