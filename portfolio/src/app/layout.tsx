import './globals.css'
import { Rubik } from 'next/font/google'

const rubik = Rubik({ subsets: ['latin'] })

export const metadata = {
  title: '🏠 Portfolio',
  description: 'a portfolio about me',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html className='bg-zinc-50' lang="en" suppressHydrationWarning={true}>
      <body className={rubik.className}>{children}</body>
    </html>
  )
}
