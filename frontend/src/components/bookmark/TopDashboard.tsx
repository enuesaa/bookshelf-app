import { css, useTheme } from '@emotion/react'
import { TopDashboardItem } from '@/components/bookmark/TopDashboardItem'
import { AiOutlineSwapRight } from 'react-icons/ai'
import Link from 'next/link'

export const TopDashboard = () => {
  const theme = useTheme()

  const styles = {
    main: css({
      margin: '20px',
      padding: '0 10px 10px 10px',
      color: theme.color.main,
    }),
    h2: css(theme.heading, {
      padding: '0 0 0 10px',
      'a': {
        margin: '10px',
        display: 'inline-block',
        color: theme.color.main,
        fontSize: theme.fontSize.large,
      },
    }),
    list: css({
      listStyleType: 'none',
      padding: '0',
    }),
  }

  return (
    <section css={styles.main}>
      <h2 css={styles.h2}>
        Bookmark
        <Link href='/bookmark/configure'><AiOutlineSwapRight /></Link>
      </h2>
      <ul css={styles.list}>
        <TopDashboardItem title='aa' href='https://example.com' />
      </ul>
    </section>
  )
} 