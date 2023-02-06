import { useState, useEffect } from 'react'

export const Subtitle = ({ videoRef, trackRef }) => {
  const [text, setText] = useState('')

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
    <div className={`${display} absolute bottom-0 mb-16 w-full justify-center`}>
      <span className="pointer-events-none w-9/12 text-center text-3xl text-yellow-400 drop-shadow-md">
        {text}
      </span>
    </div>
  )
}
