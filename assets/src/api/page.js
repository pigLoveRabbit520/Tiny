import request from '@/utils/request'

export function getList(params) {
    return request({
        url: '/public_account',
        method: 'get',
        params
    })
}

export function addAccount(data) {
    return request({
        url: '/public_account',
        method: 'post',
        data: data
    })
}

export function updatePushUrls(data) {
    return request({
        url: '/public_account/push_urls',
        method: 'post',
        data: data
    })
}
