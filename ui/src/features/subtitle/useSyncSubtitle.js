import { useEffect, useRef } from 'react'
import { useWs } from '../../hooks'

const outOfSyncStateMap = {
  'a bit early': 0.25,
  'a bit late': -0.25,
  'too early': 2,
  'too late': -2,
}

export const useSyncSubtitle = (trackRef, refreshCallback) => {
  const { message } = useWs()
  const callbackRef = useRef()

  useEffect(() => {
    callbackRef.current = refreshCallback
  }, [refreshCallback])

  useEffect(() => {
    if (!message?.state) {
      return
    }

    const { track } = trackRef.current
    const { cues } = track
    const offset = outOfSyncStateMap[message.state]

    // track is now in maintenance mode
    track.mode = 'disabled'

    Object.entries(cues)
      .filter(([key]) => key !== 'length')
      .forEach(([, c]) => {
        c.startTime += offset
        c.endTime += offset
      })

    // track can go back to live now
    track.mode = 'hidden'

    callbackRef.current()
  }, [message, trackRef, callbackRef])
}
