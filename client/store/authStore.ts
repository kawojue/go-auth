import { create } from 'zustand'

const initialUserStore = {
    otp: '',
    email: '',
    token: '',
    userId: '',
    avatar: null,
    password: '',
    username: '',
    password2: '',
    loading: false,
    btnLoading: false,
} as User

export const userStore = create<UserAuthStates>()((set) => ({
    user: initialUserStore,
    setUser: (user) => set({ user }),
    resetState: () => set({ user: initialUserStore }),
}))