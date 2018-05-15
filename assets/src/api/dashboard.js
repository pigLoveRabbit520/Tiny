import request from '@/utils/request'

export function getList(params) {
    return request({
        url: '/dashboard',
        method: 'get',
        params
    })
}
