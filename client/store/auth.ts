import { create } from 'zustand'

const initialUserStore: User = {
    otp: '',
    email: '',
    userId: '',
    avatar: null,
    password: '',
    username: '',
    password2: '',
    loading: false,
}

export const userStore = create<UserAuthStates>()((set) => ({
    user: initialUserStore,
    setUser: (user) => set({ user })
}))