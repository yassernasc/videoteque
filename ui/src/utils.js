export const isBrowserModern = () => {
  if (!isClient || typeof CSS === 'undefined' || !CSS.supports) {
    return false
  }

  // test rbg color alpha parameter support, css feature from 2018
  return CSS.supports('background: rgb(255 122 127 / .2)')
}

// window object is not defined at build time
export const isClient = typeof window !== 'undefined'
