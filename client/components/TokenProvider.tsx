"use client"

import { useEffect, FC, ReactNode } from 'react'
import getCookie from '@/utils/getCookie'
import { userStore } from '@/store/authStore'

export const TokenProvider: FC<{ children: ReactNode }> = ({ children }) => {
    const { user, setUser } = userStore()

    useEffect(() => {
        setUser({
            ...user,
            token: getCookie('access_token')
        })
    }, [])

    return (
        <>
            {children}
        </>
    )
}