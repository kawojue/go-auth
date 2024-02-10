"use client"

import { ChangeEvent } from 'react'
import {
    Card, CardContent, CardDescription,
    CardFooter, CardHeader, CardTitle,
} from '@/components/ui/card'
import { userStore } from '@/store/auth'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

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
        <section className="mx-auto w-[95vw] max-w-[600px] mt-28">
            <Card>
                <CardHeader>
                    <CardTitle>Sign Up</CardTitle>
                    <CardDescription>Get Started</CardDescription>
                </CardHeader>
                <CardContent>
                    <article className='flex flex-col gap-2.5'>
                        <div>
                            <Label htmlFor='email' className='text-lg'>Email</Label>
                            <Input
                                value={user.email}
                                id='email'
                                type='email'
                                name='email'
                                placeholder='example@mail.com'
                                onChange={handleOnChange}
                            />
                        </div>
                        <div>
                            <Label htmlFor='password' className='text-lg'>Password</Label>
                            <Input
                                name='password'
                                value={user.password}
                                id='password'
                                type='password'
                                placeholder='******'
                                onChange={handleOnChange}
                            />
                        </div>
                        <div>
                            <Label htmlFor='c-password' className='text-lg'>Confirm Password</Label>
                            <Input
                                value={user.password2}
                                id='password2'
                                type='password'
                                name='password2'
                                placeholder='Your password again'
                                onChange={handleOnChange}
                            />
                        </div>
                    </article>
                </CardContent>
                <CardFooter>
                    Already have an account? Login
                </CardFooter>
            </Card>
        </section>
    )
}

export default page