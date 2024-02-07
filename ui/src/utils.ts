export const isBrowserModern = () => {
  if (!isClient || typeof CSS === 'undefined' || !CSS.supports) {
    return false
  }

  // test css variables (custom properties) support, css feature from 2016
  return CSS.supports('color: var(--text-color)')
}

// window object is not defined at build time
export const isClient = typeof window !== 'undefined'
