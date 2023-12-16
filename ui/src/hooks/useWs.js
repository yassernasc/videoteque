import useWebSocket from 'react-use-websocket'
import { isClient } from '../utils'

const url = isClient ? `ws://${window.location.host}/ws` : ''

export const useWs = () => {
  const { sendJsonMessage, lastJsonMessage } = useWebSocket(url)
  return { emit: sendJsonMessage, message: lastJsonMessage }
}
