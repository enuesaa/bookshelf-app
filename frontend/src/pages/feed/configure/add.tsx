import { Header } from '@/components/common/Header'
import { Main } from '@/components/common/Main'
import { ConfigureAdd } from '@/components/feed/ConfigureAdd'

export default function Page() {
  return (
    <>
      <Header />
      <Main>
        <ConfigureAdd />
      </Main>
    </>
  )
}
