import { useState, useEffect, useCallback } from 'react'
import ms from 'ms'

const millisecondsToSeconds = m => m / 1000
const SPEEDS_IN_MILLISENCONDS = [ms('10s'), ms('1m'), ms('5m'), ms('10m')]
// conversion needed because html video duration property is in seconds
const SPEEDS = SPEEDS_IN_MILLISENCONDS.map(millisecondsToSeconds)

export const useSpeed = () => {
  const [counter, setCounter] = useState(0)

  useEffect(() => {
    const clearId = setTimeout(() => setCounter(0), ms('3s'))
    return () => clearTimeout(clearId)
  }, [counter])

  const speed = useCallback(() => {
    setCounter(counter => counter + 1)
    const factor = Math.min(Math.floor(counter / 4), SPEEDS.length - 1)
    return SPEEDS[factor]
  }, [counter])

  return speed
}
