import { Html, Head, Main, NextScript } from 'next/document'

const icon =
  '<svg xmlns=%22http://www.w3.org/2000/svg%22 viewBox=%220 0 100 100%22><text y=%22.9em%22 font-size=%2290%22>🐈‍⬛</text></svg>'

const Document = () => (
  <Html>
    <Head>
      <link rel="icon" href={`data:image/svg+xml,${icon}`} />
    </Head>
    <body>
      <Main />
      <NextScript />
    </body>
  </Html>
)

export default Document
