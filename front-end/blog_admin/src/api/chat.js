import axios from '@/libs/api.request'


//获取说说列表
export const getChatList = ({pageNum,pageSize}) => {
    const data = {
        pageNum,
        pageSize
    }
    return axios.request({
        url: 'admin/chat',
        method: 'get',
        params:data
    })
}

//修改说说
export const updateChat = ({id,content,is_show}) => {
    const data = {
        id,
        content,
        is_show
    }
    return axios.request({
        url: 'admin/chat',
        method: 'put',
        data
    })
}

//添加说说
export const addChat = ({content,is_show}) => {
    const data = {
        content,
        is_show
    }
    return axios.request({
        url: 'admin/chat',
        method: 'post',
        data
    })
}

//删除说说
export const deleteChat = ({id}) => {
    const data = {
        id
    }
    return axios.request({
        url: 'admin/chat',
        method: 'delete',
        data
    })
}