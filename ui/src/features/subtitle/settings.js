import { RadioGroup } from '@headlessui/react'
import { useWs } from '../../hooks'
import { H1, H2, Caption, Button } from '../../components'

const Radio = ({ label, items, handleChange }) => (
  <div>
    <Caption>{label}</Caption>
    <RadioGroup onChange={handleChange}>
      <div className="flex gap-2">
        {items.map(item => (
          <RadioGroup.Option key={item} value={item}>
            <Button>{item}</Button>
          </RadioGroup.Option>
        ))}
      </div>
    </RadioGroup>
  </div>
)
const Group = ({ label, items, handleChange }) => (
  <div>
    <Caption>{label}</Caption>
    <div className="flex gap-2">
      {items.map(item => (
        <Button key={item} onClick={() => handleChange(item)}>
          {item}
        </Button>
      ))}
    </div>
  </div>
)

const colors = ['yellow', 'white']
const fonts = ['georgia', 'futura', 'sans']
const sizes = ['bigger', 'smaller']
const positions = ['upper', 'lower']

export const SubtitleSettings = () => {
  const { emit } = useWs()

  return (
    <div className="mx-16 mt-9 flex flex-col items-center">
      <H1>Settings</H1>
      <div className="mt-10">
        <H2 className="mb-0.5">Subtitles</H2>
        <div className="flex flex-col gap-2.5">
          <Radio
            label="color"
            items={colors}
            handleChange={color => emit({ color })}
          />

          <Radio
            label="font"
            items={fonts}
            handleChange={font => emit({ font })}
          />

          <Group
            label="size"
            items={sizes}
            handleChange={size => emit({ size })}
          />

          <Group
            label="positioning"
            items={positions}
            handleChange={position => emit({ position })}
          />
        </div>
      </div>
    </div>
  )
}
