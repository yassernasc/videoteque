import { useState, useEffect, useCallback } from 'react'
import { useTimeoutFn } from 'react-use'
import ms from 'ms'
import { useWs } from '../../hooks'

const styleMap = { bordered: 'popcorn-shadow', shadowed: 'bg-black/40' }
const updateStyle = style => styleMap[style]

const colorMap = { yellow: 'text-yellow-400', white: 'text-white' }
const updateColor = color => colorMap[color]

const fontMap = {
  sans: 'font-sans',
  futura: 'font-futura',
  georgia: 'font-georgia',
}
const updateFont = font => fontMap[font]

const positionScales = [
  'mb-[5vh]',
  'mb-[10vh]',
  'mb-[15vh]',
  'mb-[20vh]',
  'mb-[25vh]',
]
const updatePosition = ({ payload, currentPosition }) => {
  const currentIndex = positionScales.findIndex(s => s === currentPosition)
  const newIndex = payload === 'upper' ? currentIndex + 1 : currentIndex - 1
  return positionScales[newIndex] ?? positionScales[currentIndex]
}

const sizeScales = [
  'text-[3.5vh]',
  'text-[4vh]',
  'text-[4.5vh]',
  'text-[5vh]',
  'text-[5.5vh]',
]
const updateSize = ({ payload, currentSize }) => {
  const currentIndex = sizeScales.findIndex(s => s === currentSize)
  const newIndex = payload === 'bigger' ? currentIndex + 1 : currentIndex - 1
  return sizeScales[newIndex] ?? sizeScales[currentIndex]
}

const middle = list => list[Math.floor(list.length / 2)]

export const useSubtitleSettings = () => {
  const [style, setStyle] = useState(() => styleMap['shadowed'])
  const [color, setColor] = useState(() => colorMap['yellow'])
  const [font, setFont] = useState(() => fontMap['sans'])
  const [position, setPosition] = useState(() => middle(positionScales))
  const [size, setSize] = useState(() => middle(sizeScales))

  const [changed, setChanged] = useState(false)
  const [, , reset] = useTimeoutFn(() => setChanged(false), ms('3s'))
  const updateChanged = useCallback(() => {
    setChanged(true)
    reset()
  }, [reset])

  const { message } = useWs()

  useEffect(() => {
    if (message?.style) {
      setStyle(updateStyle(message.style))
      updateChanged()
    }

    if (message?.color) {
      setColor(updateColor(message.color))
      updateChanged()
    }

    if (message?.font) {
      setFont(updateFont(message.font))
      updateChanged()
    }

    if (message?.position) {
      setPosition(currentPosition =>
        updatePosition({ payload: message.position, currentPosition })
      )
      updateChanged()
    }

    if (message?.size) {
      setSize(currentSize => updateSize({ payload: message.size, currentSize }))
      updateChanged()
    }
  }, [message, updateChanged])

  return { color, font, position, size, style, changed }
}
