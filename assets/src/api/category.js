import request from '@/utils/request'

export function getList(params) {
    return request({
        url: '/admin/categories',
        method: 'get',
        params
    })
}

export function addCategory(data) {
    return request({
        url: '/admin/categories',
        method: 'post',
        data: data
    })
}
