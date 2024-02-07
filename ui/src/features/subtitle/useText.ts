import { MutableRefObject, useCallback, useEffect, useState } from 'react'
import { useSubtitleSettings } from './useSubtitleSettings'
import { useSyncSubtitle } from './useSyncSubtitle'

const previewText = 'Subtitle preview text'

export const useText = (trackRef: MutableRefObject<HTMLTrackElement>) => {
  const [text, setText] = useState('')
  const settings = useSubtitleSettings()
  const totalOffset = useSyncSubtitle(trackRef)

  const refresh = useCallback(() => {
    const { activeCues } = trackRef.current.track

    if (activeCues?.length > 0) {
      const cue = activeCues[0] as VTTCue
      setText(cue.text)
    } else {
      setText('')
    }
  }, [trackRef])

  useEffect(() => {
    trackRef.current.oncuechange = () => refresh()
  }, [trackRef, refresh])

  useEffect(() => {
    refresh()
  }, [totalOffset, refresh])

  useEffect(() => {
    if (settings.changed && text === '') {
      setText(previewText)
    }

    if (!settings.changed && text === previewText) {
      setText('')
    }
  }, [text, settings.changed])

  return text
}
