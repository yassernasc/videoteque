import { Toast } from 'components'
import { RemoteCommand, useOriginalAudio, useRemote } from 'features/player'
import { Subtitle } from 'features/subtitle'
import { useMetadata, useSpeed } from 'hooks'
import { useCallback, useEffect, useRef, useState } from 'react'
import { useFullscreen } from 'react-use'

export const Player = ({ onError }) => {
  const containerRef = useRef<HTMLElement>(null)
  const videoRef = useRef<ExperimentalHTMLVideoElement>(null)
  const trackRef = useRef<HTMLTrackElement>(null)

  const [immersed, setImmersed] = useState(false)

  const metadata = useMetadata()
  useOriginalAudio(videoRef)

  const backSpeed = useSpeed()
  const forwardSpeed = useSpeed()

  const handleCommand = useCallback(
    (command: RemoteCommand) => {
      const play = () => {
        videoRef.current.play()
        setImmersed(true)
      }

      const pause = () => {
        videoRef.current.pause()
        setImmersed(false)
      }

      if (command === RemoteCommand.Pause) {
        pause()
      }

      if (command === RemoteCommand.Toogle) {
        videoRef.current.paused ? play() : pause()
      }

      if (command === RemoteCommand.Restart) {
        videoRef.current.currentTime = 0
      }

      if (command === RemoteCommand.Back) {
        videoRef.current.currentTime -= backSpeed()
      }

      if (command === RemoteCommand.Forward) {
        videoRef.current.currentTime += forwardSpeed()
      }
    },
    [backSpeed, forwardSpeed]
  )
  useRemote(handleCommand)

  const onFullscreenExit = useCallback(
    (error?: Error) => {
      if (!error) {
        handleCommand(RemoteCommand.Pause)
        setImmersed(false)
      }
    },
    [handleCommand]
  )
  useFullscreen(containerRef, immersed, { onClose: onFullscreenExit })

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

  const fit = immersed ? 'object-contain' : 'object-cover'
  const cursor = immersed ? 'cursor-none' : 'cursor-pointer'

  return (
    <main
      ref={containerRef}
      className={`h-screen w-screen bg-black ${cursor}`}
      onClick={() => handleCommand(RemoteCommand.Toogle)}
    >
      <Toast />
      <video
        className={`h-full w-full ${fit}`}
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
