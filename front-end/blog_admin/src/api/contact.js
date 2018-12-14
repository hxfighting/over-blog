import axios from '@/libs/api.request'


//获取留言列表
export const getContactList = ({search,pageSize,pageNum}) => {
    const data = {
        search,
        pageSize,
        pageNum
    }
    return axios.request({
        url: 'admin/contact',
        method: 'get',
        params:data
    })
}

//删除留言
export const deleteContact = ({id}) => {
    const data = {
        id
    }
    return axios.request({
        url: 'admin/contact',
        method: 'delete',
        data
    })
}

//回复留言
export const replyContact = ({id,reply_content}) => {
    const data = {
        id,
        reply_content
    }
    return axios.request({
        url: 'admin/contact/reply',
        method: 'post',
        data
    })
}