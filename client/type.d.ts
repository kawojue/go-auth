interface User {
    otp: string,
    email: string,
    userId: string,
    avatar: string | null,
    password: string,
    username: string,
    password2: string,
    loading: boolean,
}

interface UserAuthStates {
    user: User,
    setUser: (user: User) => void
}