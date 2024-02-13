import { isProd } from '@/utils/prod'
import axios, { AxiosInstance } from 'axios'

const BASE_URL = isProd ? '' : 'http://localhost:3030'

let refSubs = []
let isRef = false
let isTokenRefreshFailed = false

const axiosInstance: AxiosInstance = axios.create({
    baseURL: BASE_URL,
    withCredentials: true,
    headers: {
        'Content-Type': 'application/json'
    }
})

export const setTokenRefreshFailure = () => {
    isTokenRefreshFailed = true
}

export const resetTokenRefreshFailure = () => {
    isTokenRefreshFailed = false
}

axiosInstance.interceptors.response.use(
    (response) => response,
    async (error) => {
        if (error.response && error.response.status === 401) {
            const req = error.config
            if (!isRef) {
                isRef = true
                try {
                    await axios.post('/auth/refresh')
                    resetTokenRefreshFailure()
                    return axios(req)
                } catch (err) {
                    refSubs = []
                    setTokenRefreshFailure()
                    return Promise.reject(err)
                } finally {
                    isRef = false
                }
            } else {
                return new Promise((resolve) => {
                    refSubs.push(() => {
                        resolve(axios(req))
                    })
                })
            }
        }
        return Promise.reject(error)
    }
)

export const axiosReq: AxiosInstance = axios.create({
    baseURL: BASE_URL,
    withCredentials: true,
    headers: {
        'Content-Type': 'application/json',
    },
})

export { axiosInstance as axios }