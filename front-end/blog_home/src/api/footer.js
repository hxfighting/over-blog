import axios from '../libs/api.request'

//获取首页底部内容
export const getFooterList = () => {
    return axios.request({
        url: 'home/footer',
        method: 'get'
    })
}
