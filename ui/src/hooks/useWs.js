import useWebSocket from 'react-use-websocket'

// window object is not defined at build time
const url =
  typeof window !== 'undefined' ? `ws://${window.location.host}/ws` : ''

export const useWs = () => {
  const { sendJsonMessage, lastJsonMessage } = useWebSocket(url)
  return { emit: sendJsonMessage, message: lastJsonMessage }
}
