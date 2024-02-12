"use client"

import Link from 'next/link'
import notify from '@/utils/notify'
import { axios } from '@/app/api/axios'
import {
    Card, CardContent, CardDescription,
    CardFooter, CardHeader, CardTitle,
} from '@/components/ui/card'
import { userStore } from '@/store/auth'
import { useRouter } from 'next/navigation'
import throwError from '@/utils/throwError'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { ChangeEvent, useEffect } from 'react'
import { Button } from '@/components/ui/button'
import { AxiosError, AxiosResponse } from 'axios'

const page = () => {
    const router = useRouter()
    const { user, setUser, resetState } = userStore()

    const handleOnChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target
        setUser({
            ...user,
            [name]: value
        })
    }

    useEffect(() => {
        const userId = localStorage.getItem("username") || ''
        setUser({
            ...user,
            userId,
            username: userId,
        })
    }, [user.username])

    return (
        <form
            onSubmit={e => e.preventDefault()}
            className="mx-auto w-[95vw] max-w-[600px] mt-28">
            <Card>
                <CardHeader>
                    <CardTitle>Login</CardTitle>
                    <CardDescription>{`${user.username ? `Welcome, ${user.username}.` : 'Welcome'}`}</CardDescription>
                </CardHeader>
                <CardContent>
                    <article className='flex flex-col gap-3'>
                        <div>
                            <Label htmlFor='userId' className='text-lg'>Username or Email</Label>
                            <Input
                                id='userId'
                                type='userId'
                                name='userId'
                                placeholder='johndoe'
                                value={user.userId}
                                onChange={handleOnChange}
                            />
                        </div>
                        <div>
                            <Label htmlFor='password' className='text-lg'>Password</Label>
                            <Input
                                name='password'
                                id='password'
                                type='password'
                                placeholder='******'
                                value={user.password}
                                onChange={handleOnChange}
                            />
                        </div>
                    </article>
                    <div className='w-full flex mt-2 justify-end'>
                        <Button onClick={async () => await axios.post(
                            '/auth/login', {
                            ...user
                        }).then(({ data }: AxiosResponse) => {
                            notify(data?.message, 'success')
                            const username = data.data?.username
                            localStorage.setItem("username", username)
                            resetState()
                            setTimeout(() => {
                                router.push(`/${username}`)
                            }, 300)
                        }).catch((err: AxiosError) => throwError(err))}>
                            Login
                        </Button>
                    </div>
                </CardContent>
                <CardFooter className="flex gap-1 text-sm">
                    <p>{`Don't have an account?`}</p>
                    <Link href="/signup">Sign Up</Link>
                </CardFooter>
            </Card>
        </form>
    )
}

export default page