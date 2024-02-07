import { OutOfSyncState } from 'features/subtitle'
import { useWs } from 'hooks'
import { MutableRefObject, useCallback, useEffect, useState } from 'react'

type WsMessage = { state: OutOfSyncState }
type TrackArg = MutableRefObject<HTMLTrackElement>

const offsetMap: Record<OutOfSyncState, number> = {
  [OutOfSyncState.ABitEarly]: 0.75,
  [OutOfSyncState.ABitLate]: -0.75,
  [OutOfSyncState.TooEarly]: 5,
  [OutOfSyncState.TooLate]: -5,
}

const formatOffset = (num: number) => (num >= 0 ? `+${num}` : num.toString())

export const useSyncSubtitle = (trackRef?: TrackArg) => {
  const [totalOffset, setTotalOffset] = useState(0)
  const { message } = useWs<WsMessage>()

  const syncSubtitles = useCallback(
    (offset: number) => {
      if (trackRef?.current) {
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
      }

      setTotalOffset(old => old + offset)
    },
    [trackRef]
  )

  useEffect(() => {
    if (message?.state) {
      const offset = offsetMap[message.state]
      syncSubtitles(offset)
    }
  }, [message, syncSubtitles])

  return formatOffset(totalOffset)
}
