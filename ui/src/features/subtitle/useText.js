import { useEffect, useState } from 'react'
import { useSubtitleSettings } from './useSubtitleSettings'

const previewText = 'Bela Lugosi is dead'

export const useText = trackRef => {
  const [text, setText] = useState('')
  const settings = useSubtitleSettings()

  useEffect(() => {
    trackRef.current.oncuechange = e => {
      const { activeCues } = e.target.track
      if (activeCues.length > 0) {
        setText(activeCues[0].text)
      } else {
        setText('')
      }
    }
  }, [trackRef])

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
