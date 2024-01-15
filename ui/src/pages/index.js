import Head from 'next/head'
import { useState, useEffect } from 'react'
import ms from 'ms'
import { Error } from '../components'
import { Player } from '../features'
import { isBrowserModern } from '../utils'

const isBrowserOld = !isBrowserModern()

const PlayerRouter = () => {
  const [error, setError] = useState(null)

  if (isBrowserOld) {
    return <p>Redirecting to the Legacy Player..</p>
  }

  if (error) {
    return <Error message={error} />
  }

  return <Player onError={setError} />
}

const Root = () => {
  useEffect(() => {
    if (isBrowserOld) {
      setTimeout(() => (location.href = 'legacy'), ms('2s'))
    }
  }, [])

  return (
    <>
      <Head>
        <title>Vidéothèque</title>
      </Head>
      <PlayerRouter />
    </>
  )
}

export default Root
