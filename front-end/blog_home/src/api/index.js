import axios from '../libs/api.request'

//获取文章列表
export const getArticleList = ({tag_id,category_id,pageSize,pageNum}) => {
  const data = {
    pageNum,pageSize,category_id,tag_id
  }
  return axios.request({
    url: 'home/article',
    method: 'get',
    params:data
  })
}
