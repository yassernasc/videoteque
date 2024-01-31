// custom polyfills for Intl.DisplayNames
import '@formatjs/intl-getcanonicallocales/polyfill'
import '@formatjs/intl-locale/polyfill'
import '@formatjs/intl-displaynames/polyfill'
import '@formatjs/intl-displaynames/locale-data/en'

import '../global.css'

const App = ({ Component, pageProps }) => <Component {...pageProps} />
export default App
