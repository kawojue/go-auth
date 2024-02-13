"use client"
import { axios } from '../api/axios'
import { Spin } from '@/components/Loading'
import throwError from '@/utils/throwError'
import { useQuery } from '@tanstack/react-query'
import { AxiosError, AxiosResponse } from 'axios'

const page = ({ params: { username } }: Params) => {
    const { data, isLoading } = useQuery({
        queryKey: ['user'],
        queryFn: async () => {
            return await axios.get(`/${username}`)
                .then(({ data }: AxiosResponse) => {
                    return data
                }).catch((err: AxiosError) => {
                    throwError(err)
                })
        }
    })

    if (isLoading) {
        return <Spin />
    }

    console.log(data)

    return (
        <>
            {username}
        </>
    )
}

export default page