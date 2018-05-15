import request from '@/utils/request'

export function getList(params) {
    return request({
        url: '/public_account_article',
        method: 'get',
        params
    })
}
