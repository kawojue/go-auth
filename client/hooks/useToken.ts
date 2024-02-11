import { useEffect } from 'react'
import getCookie from '@/utils/getCookie'

const useToken = () => {
    let access_token = ''

    useEffect(() => {
        access_token = getCookie('access_token')
    }, [access_token])

    return access_token
}

export default useToken