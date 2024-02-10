import { isProd } from '@/utils/prod'
import getCookie from '@/utils/getCookie'
import axios, { AxiosInstance } from 'axios'

const BASE_URL = isProd ? '' : 'http://localhost:3030'

let refSubs = []
let isRef = false

const axiosInstance: AxiosInstance = axios.create({
    baseURL: BASE_URL,
    withCredentials: true,
    headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${getCookie('access_token')}`
    }
})

axiosInstance.interceptors.response.use(
    (response) => {
        return response
    },
    async (error) => {
        if (error.response.status === 401) {
            const req = error.config

            if (!isRef) {
                isRef = true
                try {
                    await axios.post('/auth/refresh')
                    return axios(req)
                } catch (err) {
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
    headers: {
        'Content-Type': 'application/json',
    }
})

export { axiosInstance as axios }