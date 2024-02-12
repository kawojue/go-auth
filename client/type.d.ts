type NotificationAction = 'error' | 'success'

interface User {
    otp: string
    email: string
    userId: string
    avatar: string | null
    password: string
    username: string
    password2: string
    loading: boolean
    btnLoading: boolean
}

interface UserAuthStates {
    user: User
    setUser: (user: User) => void
    resetState: () => void
}

interface Params {
    params: {
        username: string
    }
}