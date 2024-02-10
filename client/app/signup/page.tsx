"use client"

import {
    Card, CardContent, CardDescription,
    CardFooter, CardHeader, CardTitle,
} from '@/components/ui/card'

const page = () => {
    return (
        <section className="mx-auto w-[95vw] max-w-[600px] mt-36">
            <Card>
                <CardHeader>
                    <CardTitle>Sign Up</CardTitle>
                    <CardDescription>Get Started</CardDescription>
                </CardHeader>
                <CardContent>

                </CardContent>
                <CardFooter>
                    Already have an account? Login
                </CardFooter>
            </Card>
        </section>
    )
}

export default page