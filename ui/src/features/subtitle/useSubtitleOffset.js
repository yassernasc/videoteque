import { useEffect, useRef, useState } from 'react'
import { useWs } from '../../hooks'

const offsetMap = {
  'a bit early': 0.75,
  'a bit late': -0.75,
  'too early': 5,
  'too late': -5,
}

const formatOffset = num => (num >= 0 ? `+${num}` : num.toString())

export const useSubtitleOffset = onNewOffset => {
  const [totalOffset, setTotalOffset] = useState(0)
  const { message } = useWs()
  const callbackRef = useRef()

  useEffect(() => {
    callbackRef.current = onNewOffset
  }, [onNewOffset])

  useEffect(() => {
    if (message?.state) {
      const offset = offsetMap[message.state]

      if (callbackRef.current) {
        callbackRef.current(offset)
      }

      setTotalOffset(old => old + offset)
    }
  }, [message, callbackRef])

  return formatOffset(totalOffset)
}
