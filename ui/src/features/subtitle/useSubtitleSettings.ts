import {
  Color,
  Font,
  PositionChange,
  SizeChange,
  Style,
} from 'features/subtitle'
import { useWs } from 'hooks'
import ms from 'ms'
import { useCallback, useEffect, useState } from 'react'
import { useTimeoutFn } from 'react-use'

type WsMessage = {
  color: Color
  font: Font
  position: PositionChange
  size: SizeChange
  style: Style
}

const colorMap: Record<Color, string> = {
  [Color.Amber]: 'text-amber-400',
  [Color.White]: 'text-white',
  [Color.Yellow]: 'text-yellow-400',
}
const updateColor = (color: Color) => colorMap[color]

const fontMap: Record<Font, string> = {
  [Font.Futura]: 'font-futura',
  [Font.Georgia]: 'font-georgia',
  [Font.Sans]: 'font-dm',
}
const updateFont = (font: Font) => fontMap[font]

const styleMap: Record<Style, string> = {
  [Style.Bordered]: 'popcorn-shadow',
  [Style.Shadowed]: 'bg-black/60',
}
const updateStyle = (style: Style) => styleMap[style]

const positionScales = ['mb-16', 'mb-24', 'mb-32', 'mb-40', 'mb-80']
const updatePosition = ({
  payload,
  currentPosition,
}: {
  payload: PositionChange
  currentPosition: string
}) => {
  const currentIndex = positionScales.findIndex(s => s === currentPosition)
  const newIndex =
    payload === PositionChange.Upper ? currentIndex + 1 : currentIndex - 1
  return positionScales[newIndex] ?? positionScales[currentIndex]
}

const sizeScales = ['text-xl', 'text-2xl', 'text-3xl', 'text-4xl', 'text-5xl']
const updateSize = ({
  payload,
  currentSize,
}: {
  payload: SizeChange
  currentSize: string
}) => {
  const currentIndex = sizeScales.findIndex(s => s === currentSize)
  const newIndex =
    payload === SizeChange.Bigger ? currentIndex + 1 : currentIndex - 1
  return sizeScales[newIndex] ?? sizeScales[currentIndex]
}

const middle = (list: string[]) => list[Math.floor(list.length / 2)]

export const useSubtitleSettings = () => {
  const [style, setStyle] = useState(() => styleMap[Style.Shadowed])
  const [color, setColor] = useState(() => colorMap[Color.Yellow])
  const [font, setFont] = useState(() => fontMap[Font.Sans])
  const [position, setPosition] = useState(() => middle(positionScales))
  const [size, setSize] = useState(() => middle(sizeScales))

  const [changed, setChanged] = useState(false)
  const [, , reset] = useTimeoutFn(() => setChanged(false), ms('3s'))
  const updateChanged = useCallback(() => {
    setChanged(true)
    reset()
  }, [reset])

  const { message } = useWs<WsMessage>()

  useEffect(() => {
    if (Object.hasOwn(message, 'style')) {
      setStyle(updateStyle(message.style))
      updateChanged()
    }
  }, [message.style, updateChanged])

  useEffect(() => {
    if (Object.hasOwn(message, 'color')) {
      setColor(updateColor(message.color))
      updateChanged()
    }
  }, [message.color, updateChanged])

  useEffect(() => {
    if (Object.hasOwn(message, 'font')) {
      setFont(updateFont(message.font))
      updateChanged()
    }
  }, [message.font, updateChanged])

  useEffect(() => {
    if (Object.hasOwn(message, 'position')) {
      setPosition(currentPosition =>
        updatePosition({ payload: message.position, currentPosition })
      )
    }
  }, [message.position, updateChanged])

  useEffect(() => {
    if (Object.hasOwn(message, 'size')) {
      setSize(currentSize => updateSize({ payload: message.size, currentSize }))
      updateChanged()
    }
  }, [message.size, updateChanged])

  return { color, font, position, size, style, changed }
}
