import Head from 'next/head'
import Script from 'next/script'
import { PlayerLegacy } from '../features'

const Legacy = () => (
  <>
    <Head>
      <title>Lugosi - Legacy Player</title>
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
