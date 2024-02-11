import notify from './notify'
import { AxiosError } from 'axios'

const throwError = (err: AxiosError): void => {
    // @ts-ignore
    const error = err.response?.data?.error
    if (error) {
        notify(error, 'error')
    } else {
        notify(err.message, 'error')
    }
}

export default throwError