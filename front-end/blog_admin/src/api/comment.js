import axios from '@/libs/api.request'


//获取评论列表
export const getCommentList = ({article_id,pageNum,pageSize}) => {
    const data = {
        article_id,
        pageNum,
        pageSize
    }
    return axios.request({
        url: 'admin/comment',
        method: 'get',
        params:data
    })
}

//回复评论
export const replyComment = formdata => {
    return axios.request({
        url: 'admin/comment',
        method: 'post',
        data:formdata
    })
}

//删除评论
export const deleteComment = ({id}) => {
    const data = {
        id
    }
    return axios.request({
        url: 'admin/comment',
        method: 'delete',
        data
    })
}