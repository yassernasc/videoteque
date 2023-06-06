import { useState, useEffect, useCallback } from 'react'
import { useTimeoutFn } from 'react-use'
import { useWs } from '../../hooks'

const styleMap = { bordered: 'popcorn-shadow', shadowed: 'bg-black/40' }
const updateStyle = style => styleMap[style]

const colorMap = { yellow: 'text-yellow-400', white: 'text-white' }
const updateColor = color => colorMap[color]

const fontMap = {
  georgia: 'font-georgia',
  futura: 'font-futura',
  sans: 'font-sans',
}
const updateFont = font => fontMap[font]

const positionScales = ['mb-16', 'mb-24', 'mb-32', 'mb-40', 'mb-48']
const updatePosition = ({ payload, currentPosition }) => {
  const currentIndex = positionScales.findIndex(s => s === currentPosition)
  const newIndex = payload === 'upper' ? currentIndex + 1 : currentIndex - 1
  return positionScales[newIndex] ?? positionScales[currentIndex]
}

const sizeScales = ['text-2xl', 'text-3xl', 'text-4xl', 'text-5xl', 'text-6xl']
const updateSize = ({ payload, currentSize }) => {
  const currentIndex = sizeScales.findIndex(s => s === currentSize)
  const newIndex = payload === 'bigger' ? currentIndex + 1 : currentIndex - 1
  return sizeScales[newIndex] ?? sizeScales[currentIndex]
}

const middle = list => list[Math.floor(list.length / 2)]

export const useSubtitleSettings = () => {
  const [style, setStyle] = useState(() => styleMap['shadowed'])
  const [color, setColor] = useState(() => colorMap['yellow'])
  const [font, setFont] = useState(() => fontMap['georgia'])
  const [position, setPosition] = useState(() => middle(positionScales))
  const [size, setSize] = useState(() => middle(sizeScales))

  const [changed, setChanged] = useState(false)
  const [, , reset] = useTimeoutFn(() => setChanged(false), 3000)
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
