import { useEffect } from 'react'

const codeMap = {
  ArrowLeft: 'back',
  ArrowRight: 'forward',
  Space: 'toogle',
}

// tizen browser codeKeys
const codeKeyMap = {
  10252: 'toogle', // MediaPlayPause
  403: 'restart', // ColorF0Red
}

export const useRemote = callback => {
  useEffect(() => {
    const handleCode = ({ code, keyCode }) => {
      const command = codeMap[code] || codeKeyMap[keyCode]
      if (command) {
        callback(command)
      }
    }

    window.addEventListener('keydown', handleCode)
    return () => window.removeEventListener('keydown', handleCode)
  }, [callback])
}
