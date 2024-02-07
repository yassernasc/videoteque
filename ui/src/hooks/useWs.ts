import useWebSocket from 'react-use-websocket'
import { isClient } from 'utils'

const url = isClient ? `ws://${window.location.host}/ws` : ''

export const useWs = <T>() => {
  const { sendJsonMessage, lastJsonMessage } = useWebSocket(url)
  return { emit: sendJsonMessage, message: lastJsonMessage as T }
}
