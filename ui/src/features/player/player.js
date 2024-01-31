import { useCallback, useEffect, useRef, useState } from 'react'
import { useFullscreen } from 'react-use'
import { Subtitle } from '..'
import { Toast } from '../../components'
import { useMetadata, useRemote, useSpeed } from '../../hooks'
import { useOriginalAudio } from './useOriginalAudio'

export const Player = ({ onError }) => {
  const containerRef = useRef(null)
  const videoRef = useRef(null)
  const trackRef = useRef(null)

  const [immersed, setImmersed] = useState(false)
  useFullscreen(containerRef, immersed)

  const metadata = useMetadata()
  useOriginalAudio(videoRef)

  const backSpeed = useSpeed()
  const forwardSpeed = useSpeed()

  const handleCommand = useCallback(
    command => {
      const play = () => {
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
    const messages = {
      1: 'download aborted',
      2: 'network error, please try to restart the server and refresh this page',
      3: 'video format is not supported or the video file is broken',
      4: 'video format is not supported, please try another browser',
    }

    videoRef.current.onerror = () => {
      const { code } = videoRef.current.error
      const message = messages[code] || 'unknown error'
      onError(message)
    }
  }, [onError])

  const cursor = immersed ? 'cursor-none' : 'cursor-pointer'

  return (
    <main
      ref={containerRef}
      className={`h-screen w-screen bg-black ${cursor}`}
      onClick={() => handleCommand('toogle')}
    >
      <Toast />
      <video
        className="h-full w-full"
        poster={metadata?.Tmdb?.Backdrop}
        preload="auto"
        ref={videoRef}
        src="/movie"
      >
        <track ref={trackRef} default src="/subtitle" kind="metadata" />
      </video>
      <Subtitle trackRef={trackRef} />
    </main>
  )
}
