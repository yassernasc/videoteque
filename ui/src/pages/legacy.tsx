import { PlayerLegacy } from 'features/player'
import Head from 'next/head'
import Script from 'next/script'

const Legacy = () => (
  <>
    <Head>
      <title>Vidéothèque 📽️</title>
      <link
        href="https://unpkg.com/video.js/dist/video-js.min.css"
        rel="stylesheet"
      />
    </Head>
    <Script src="https://unpkg.com/video.js/dist/video.min.js" />
    <PlayerLegacy />
  </>
)

export default Legacy
