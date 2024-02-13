import './globals.css'
import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import { Toaster } from 'react-hot-toast'
import QueryProvider from '@/components/QueryProvider'
import { TokenProvider } from '@/components/TokenProvider'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'Go Auth',
  description: 'Testing go-auth API endpoints.',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body className={inter.className}>
        <Toaster
          position="top-center"
          reverseOrder={false} />
        <QueryProvider>
          <TokenProvider>
            {children}
          </TokenProvider>
        </QueryProvider>
      </body>
    </html>
  )
}
