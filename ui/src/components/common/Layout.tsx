import { Container } from '@radix-ui/themes'
import { Header } from './Header'
import { ReactNode } from 'react'

type Props = {
  children: ReactNode
}
export const Layout = ({ children }: Props) => {
  return (
    <>
      <Header />
      <Container size='4' m='3' pt='3'>
        {children}
      </Container>
    </>
  )
}
