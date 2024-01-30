import { useCallback, useEffect, useRef } from 'react'
import { useSubtitleOffset } from './useSubtitleOffset'

export const useSyncSubtitle = (trackRef, refreshCallback) => {
  const callbackRef = useRef()

  const syncSubtitles = useCallback(
    offset => {
      const { track } = trackRef.current
      const { cues } = track

      // track is now in maintenance mode
      track.mode = 'disabled'

      Object.entries(cues)
        .filter(([key]) => key !== 'length')
        .forEach(([, c]) => {
          c.startTime += offset
          c.endTime += offset
        })

      // back to previous state
      track.mode = 'hidden'

      callbackRef.current()
    },
    [trackRef, callbackRef]
  )

  useSubtitleOffset(syncSubtitles)

  useEffect(() => {
    callbackRef.current = refreshCallback
  }, [refreshCallback])
}
