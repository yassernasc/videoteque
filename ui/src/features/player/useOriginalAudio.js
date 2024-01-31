import { useEffect } from 'react'
import { useMetadata } from '../../hooks'

const langNames = new Intl.DisplayNames(['en'], { type: 'language' })
const isSameLanguage = (a, b) => langNames.of(a) === langNames.of(b)

export const useOriginalAudio = videoRef => {
  const metadata = useMetadata()

  useEffect(() => {
    const video = videoRef.current

    const exec = () => {
      const originalLang = metadata?.Tmdb?.Language
      const { audioTracks } = video
      const hasMultipleTracks = audioTracks?.length > 1

      if (hasMultipleTracks && originalLang) {
        let found = false

        for (let i = 0; i < audioTracks.length; i += 1) {
          const track = audioTracks[i]

          if (found) {
            track.enabled = false
            continue
          }

          if (isSameLanguage(track.language, originalLang)) {
            found = true
            track.enabled = true
          } else {
            track.enabled = false
          }
        }

        if (!found) {
          audioTracks[0].enabled = true
        }
      }
    }

    video.addEventListener('loadedmetadata', exec)
    return () => video.removeEventListener('loadedmetadata', exec)
  }, [metadata, videoRef])
}
