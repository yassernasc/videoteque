import useWebSocket from 'react-use-websocket'
import { isClient } from 'utils'

const url = isClient ? `ws://${window.location.host}/ws` : ''

export const useWs = <T>() => {
  const { sendJsonMessage, lastJsonMessage } = useWebSocket(url)

  const message = (lastJsonMessage ?? {}) as T

  return { emit: sendJsonMessage, message }
}
