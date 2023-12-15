import { useCallback, useEffect, useRef, useState } from 'react'
import { Subtitle } from '..'
import { useRemote, useSpeed } from '../../hooks'

export const Player = ({ onError }) => {
  const containerRef = useRef(null)
  const videoRef = useRef(null)
  const trackRef = useRef(null)

  const [immersed, setImmersed] = useState(false)

  const backSpeed = useSpeed()
  const forwardSpeed = useSpeed()

  const handleCommand = useCallback(
    command => {
      const play = () => {
        containerRef.current.requestFullscreen()
        videoRef.current.play()
        setImmersed(true)
      }

      const pause = () => {
        videoRef.current.pause()
        setImmersed(false)
      }

      if (command === 'toogle') {
        videoRef.current.paused ? play() : pause()
      }

      if (command === 'restart') {
        videoRef.current.currentTime = 0
      }

      if (command === 'back') {
        videoRef.current.currentTime -= backSpeed()
      }

      if (command === 'forward') {
        videoRef.current.currentTime += forwardSpeed()
      }
    },
    [backSpeed, forwardSpeed]
  )

  useRemote(handleCommand)

  useEffect(() => {
    const handleError = error => onError(error.message)

    videoRef.current.onerror = handleError
    if (videoRef.current.error) {
      handleError(videoRef.current.error)
    }
  }, [onError])

  const cursor = immersed ? 'cursor-none' : 'cursor-pointer'

  return (
    <div
      ref={containerRef}
      className={`h-screen bg-black ${cursor}`}
      onClick={() => handleCommand('toogle')}
    >
      <video ref={videoRef} className="h-full w-full" src="/movie">
        <track ref={trackRef} default src="/subtitle" />
      </video>

      <Subtitle trackRef={trackRef} />
    </div>
  )
}
