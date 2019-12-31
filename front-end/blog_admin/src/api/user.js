import axios from '@/libs/api.request'

export const login = ({userName, password, captcha, key}) => {
    const data = {
        name: userName,
        password,
        captcha,
        key
    }
    return axios.request({
        url: 'admin/login',
        data,
        method: 'post'
    })
}

//获取个人信息
export const getUserInfo = (token) => {
    return axios.request({
        url: 'admin',
        method: 'get'
    })
}

//修改个人信息
export const updateInfo = ({avatar,name,mobile,email}) => {
    const data = {
        avatar,
        name,
        phone:mobile,
        email
    }
    return axios.request({
        url: 'admin',
        method: 'put',
        data
    })
}

//修改密码
export const updatePassword = ({password,password_confirmation}) => {
    const data = {
        password,
        password_confirmation
    }
    return axios.request({
        url: 'admin/password',
        method: 'put',
        data
    })
}

export const logout = (token) => {
    return axios.request({
        url: 'admin/logout',
        method: 'post'
    })
}

export const getUnreadCount = () => {
    return axios.request({
        url: 'message/count',
        method: 'get'
    })
}

export const getMessage = () => {
    return axios.request({
        url: 'message/init',
        method: 'get'
    })
}

export const getContentByMsgId = msg_id => {
    return axios.request({
        url: 'message/content',
        method: 'get',
        params: {
            msg_id
        }
    })
}

export const hasRead = msg_id => {
    return axios.request({
        url: 'message/has_read',
        method: 'post',
        data: {
            msg_id
        }
    })
}

export const removeReaded = msg_id => {
    return axios.request({
        url: 'message/remove_readed',
        method: 'post',
        data: {
            msg_id
        }
    })
}

export const restoreTrash = msg_id => {
    return axios.request({
        url: 'message/restore',
        method: 'post',
        data: {
            msg_id
        }
    })
}

//获取用户列表
export const getUserList = ({search,pageSize,pageNum}) => {
    const data = {
        search,
        pageSize,
        pageNum
    }
    return axios.request({
        url: 'admin/user',
        method: 'get',
        params:data
    })
}

//删除用户
export const deleteUser = ({id}) => {
    const data = {
        id
    }
    return axios.request({
        url: 'admin/user',
        method: 'delete',
        data
    })
}

//改变用户身份
export const updateUser = ({id,is_admin}) => {
    const data = {
        id,
        is_admin
    }
    return axios.request({
        url: 'admin/user',
        method: 'put',
        data
    })
}

