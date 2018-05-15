import request from '@/utils/request'

export function getData(params) {
    return request({
        url: '/settings',
        method: 'get',
    })
}

export function updateExecTime(time) {
    return request({
        url: '/settings',
        method: 'post',
        data: {
            name: 'exec_time',
            value: time
        }
    })
}
