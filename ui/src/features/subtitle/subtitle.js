import { useState, useEffect } from 'react'
import { useSubtitleSettings } from './useSubtitleSettings'

export const Subtitle = ({ videoRef, trackRef }) => {
  const [text, setText] = useState('')
  const { color, position, font, size } = useSubtitleSettings()

  useEffect(() => {
    // hide browser default subtitle
    if (videoRef.current.textTracks[0]) {
      videoRef.current.textTracks[0].mode = 'hidden'
    }
  }, [])

  useEffect(() => {
    trackRef.current.oncuechange = e => {
      const { activeCues } = e.target.track
      if (activeCues.length > 0) {
        setText(activeCues[0].text)
      } else {
        setText('')
      }
    }
  }, [])

  const display = text === '' ? 'hidden' : 'flex'

  return (
    <div
      className={`${display} ${position} absolute bottom-0 w-full justify-center`}
    >
      <div className="flex w-8/12 justify-center">
        <span
          className={`${color} ${size} ${font} pointer-events-none inline-block rounded bg-black/40 py-1 px-6 text-justify leading-tight [text-align-last:center]`}
          dangerouslySetInnerHTML={{ __html: text }}
        ></span>
      </div>
    </div>
  )
}
