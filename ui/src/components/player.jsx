import { useEffect, useRef, useState } from 'react'

export const Player = ({ onError }) => {
  const containerRef = useRef(null)
  const videoRef = useRef(null)

  const [immersed, setImmersed] = useState(false)

  useEffect(() => {
    const handleError = error => onError(error.message)

    videoRef.current.onerror = handleError
    if (videoRef.current.error) {
      handleError(videoRef.current.error)
    }
  }, [onError])

  const play = () => {
    containerRef.current.requestFullscreen()
    videoRef.current.play()
    setImmersed(true)
  }

  const pause = () => {
    videoRef.current.pause()
    setImmersed(false)
  }

  const toogle = () => (videoRef.current.paused ? play() : pause())

  const cursor = immersed ? 'cursor-none' : 'cursor-pointer'

  return (
    <div
      ref={containerRef}
      className={`h-screen bg-black ${cursor}`}
      onClick={toogle}
    >
      <video ref={videoRef} className="h-full w-full" src="/movie"></video>
    </div>
  )
}
