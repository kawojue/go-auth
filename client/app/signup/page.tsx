"use client"

import Link from 'next/link'
import router from 'next/router'
import { ChangeEvent } from 'react'
import notify from '@/utils/notify'
import { axios } from '@/app/api/axios'
import {
    Card, CardContent, CardDescription,
    CardFooter, CardHeader, CardTitle,
} from '@/components/ui/card'
import { userStore } from '@/store/auth'
import throwError from '@/utils/throwError'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Button } from '@/components/ui/button'
import { AxiosError, AxiosResponse } from 'axios'

const page = () => {
    const { user, setUser } = userStore()

    const handleOnChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target
        setUser({
            ...user,
            [name]: value
        })
    }

    return (
        <form
            onSubmit={e => e.preventDefault()}
            className="mx-auto w-[95vw] max-w-[600px] mt-28">
            <Card>
                <CardHeader>
                    <CardTitle>Sign Up</CardTitle>
                    <CardDescription>Get Started</CardDescription>
                </CardHeader>
                <CardContent>
                    <article className='flex flex-col gap-2.5'>
                        <div className="flex justify-between items-center">
                            <div>
                                <Label htmlFor='email' className='text-lg'>Email</Label>
                                <Input
                                    id='email'
                                    type='email'
                                    name='email'
                                    value={user.email}
                                    onChange={handleOnChange}
                                    placeholder='example@mail.com'
                                />
                            </div>
                            <div>
                                <Label htmlFor='username' className='text-lg'>Username</Label>
                                <Input
                                    id='username'
                                    type='username'
                                    name='username'
                                    placeholder='kawojue'
                                    value={user.username}
                                    onChange={handleOnChange}
                                />
                            </div>
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
                        <div>
                            <Label htmlFor='c-password' className='text-lg'>Confirm Password</Label>
                            <Input
                                id='password2'
                                type='password'
                                name='password2'
                                value={user.password2}
                                onChange={handleOnChange}
                                placeholder='Your password again'
                            />
                        </div>
                    </article>
                    <div className='w-full flex mt-2 justify-end'>
                        <Button onClick={async () => await axios.post(
                            '/auth/signup', {
                            ...user
                        }).then((res: AxiosResponse) => {
                            router.push('/login')
                            notify(res.data?.message, 'success')
                        }).catch((err: AxiosError) => throwError(err))}>
                            Sign up
                        </Button>
                    </div>
                </CardContent>
                <CardFooter className="flex gap-1 text-sm">
                    <p>Already have an account?</p>
                    <Link href="/login">Login</Link>
                </CardFooter>
            </Card>
        </form>
    )
}

export default page