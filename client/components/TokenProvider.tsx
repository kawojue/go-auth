"use client"

import getCookie from '@/utils/getCookie'
import { userStore } from '@/store/authStore'
import { useEffect, FC, ReactNode } from 'react'

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