import { FC } from 'react'
import ReactLoading from 'react-loading'

const Loading: FC<LoadingProps> = (props) => {
    return <ReactLoading {...props} />
}

export const Spin = () => {
    return <Loading type='spin' />
}