import Cookies from 'js-cookie'

const TokenKey = 'Admin-Token'
const tokenValidDays = 7;

export function getToken() {
    return Cookies.get(TokenKey)
}

export function setToken(token) {
    return Cookies.set(TokenKey, token, { expires: tokenValidDays })
}

export function removeToken() {
    return Cookies.remove(TokenKey)
}
