import { useSubtitleSettings } from './useSubtitleSettings'
import { useText } from './useText'

export const Subtitle = ({ trackRef }) => {
  const text = useText(trackRef)
  const { color, position, font, size, style } = useSubtitleSettings()

  const display = text === '' ? 'hidden' : 'flex'

  return (
    <div
      className={`${display} ${position} absolute bottom-0 w-full justify-center`}
    >
      <div className="flex w-8/12 justify-center">
        <span
          className={`${style} ${color} ${size} ${font} pointer-events-none inline-block rounded py-1 px-6 text-justify leading-tight [text-align-last:center]`}
          dangerouslySetInnerHTML={{ __html: text }}
          title="subtitle"
        ></span>
      </div>
    </div>
  )
}
