import { useCallback, useEffect, useState } from 'react'
import { useSubtitleSettings } from './useSubtitleSettings'
import { useSyncSubtitle } from './useSyncSubtitle'

const previewText = 'Subtitle preview text'

export const useText = trackRef => {
  const [text, setText] = useState('')
  const settings = useSubtitleSettings()

  const refresh = useCallback(() => {
    const { activeCues } = trackRef.current.track

    if (activeCues.length > 0) {
      setText(activeCues[0].text)
    } else {
      setText('')
    }
  }, [trackRef])

  useSyncSubtitle(trackRef, refresh)

  useEffect(() => {
    trackRef.current.oncuechange = () => refresh()
  }, [trackRef, refresh])

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
