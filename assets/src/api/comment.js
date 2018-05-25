import request from '@/utils/request'

export function getList(params) {
    return request({
        url: '/admin/comments',
        method: 'get',
        params
    })
}

export function addComment(data) {
    return request({
        url: '/admin/comments',
        method: 'post',
        data: data
    })
}
