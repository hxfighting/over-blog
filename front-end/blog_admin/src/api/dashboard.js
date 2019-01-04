import axios from '@/libs/api.request'


//获取dashboard统计数据
export const getCountData = () => {
    return axios.request({
        url: 'admin/count',
        method: 'get',
    })
}
