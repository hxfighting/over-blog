import axios from '@/libs/api.request'

//获取文章列表
export const getArticleList = ({search,pageSize,pageNum,category_id}) => {
    const data = {
        search,
        pageSize,
        pageNum,
        category_id
    }
    return axios.request({
        url: 'admin/article',
        method: 'get',
        params:data
    })
}

export const uploadImage = ({formdata,token}) => {
  return axios.request({
    url: 'admin/upload',
    method: 'post',
    data:formdata,
    headers: { 'Content-Type': 'multipart/form-data' , 'Authorization':token}
  })
}

//修改文章
export const updateArticle = formdata => {
    return axios.request({
        url: 'admin/article',
        method: 'put',
        data:formdata
    })
}

//添加文章
export const addArticle = formdata => {
    return axios.request({
        url: 'admin/article',
        method: 'post',
        data:formdata
    })
}

//删除文章
export const deleteArticle = ({id}) => {
    const data = {
        id
    }
    return axios.request({
        url: 'admin/article',
        method: 'delete',
        data
    })
}
