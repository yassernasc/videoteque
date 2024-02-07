type AudioTrack = {
  id: string
  kind: string
  label: string
  language: string
  enabled: boolean
}

class AudioTrackList extends EventTarget {
  readonly length: number

  getTrackById(id: string): AudioTrack | null
  [index: number]: AudioTrack

  onchange: (ev: any) => any
  onaddtrack: (ev: any) => any
  onremovetrack: (ev: any) => any
}

type ExperimentalHTMLVideoElement = HTMLVideoElement & {
  readonly audioTracks: AudioTrackList | null
}
