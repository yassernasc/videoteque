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

      for (let i = 0; i < cues.length; i += 1) {
        const cue = cues[i]
        cue.startTime += offset
        cue.endTime += offset
      }

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
