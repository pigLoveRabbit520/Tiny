import request from '@/utils/request'

export function getList(params) {
    return request({
        url: '/admin/pages',
        method: 'get',
        params
    })
}

export function addPage(data) {
    return request({
        url: '/admin/pages',
        method: 'post',
        data: data
    })
}
