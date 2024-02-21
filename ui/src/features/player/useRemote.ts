import { RemoteCommand } from 'features/player'
import { useEffect } from 'react'

type CodeMap = { [code: string]: RemoteCommand }
type Callback = (command: RemoteCommand) => void

const codeMap: CodeMap = {
  ArrowLeft: RemoteCommand.Back,
  ArrowRight: RemoteCommand.Forward,
  Space: RemoteCommand.Toogle,
}

// tizen browser codeKeys
const codeKeyMap: CodeMap = {
  10252: RemoteCommand.Toogle, // MediaPlayPause
  403: RemoteCommand.Restart, // ColorF0Red
}

const useRemote = (callback: Callback) => {
  useEffect(() => {
    const handleCode = ({ code, keyCode }) => {
      const command = codeMap[code] ?? codeKeyMap[keyCode]

      if (command) {
        callback(command)
      }
    }

    window.addEventListener('keydown', handleCode)
    return () => window.removeEventListener('keydown', handleCode)
  }, [callback])
}

export { useRemote }
