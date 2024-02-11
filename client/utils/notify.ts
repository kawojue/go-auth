import { toast } from 'react-hot-toast'

const notify = (message: string, action?: NotificationAction): void => {
    switch (action) {
        case 'success':
            toast.success(message, {
                duration: 3200
            })
            break
        case 'error':
            toast.error(message, {
                duration: 4200
            })
            break
        default:
            toast.custom(message, {
                duration: 3000,
                position: 'top-right'
            })
            break
    }
}

export default notify