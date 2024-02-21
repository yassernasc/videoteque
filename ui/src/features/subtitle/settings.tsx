import { RadioGroup } from '@headlessui/react'
import { Button } from 'components'
import {
  Color,
  Font,
  OutOfSyncState,
  PositionChange,
  SizeChange,
  Style,
} from 'features/subtitle'
import { useWs } from 'hooks'

type GroupArgs = {
  label: string
  items: { label: string; value: number }[]
  handleChange: (item: number) => void
}

const styles = [
  { label: 'shadowed', value: Style.Shadowed },
  { label: 'bordered', value: Style.Bordered },
]
const colors = [
  { label: 'yellow', value: Color.Yellow },
  { label: 'amber', value: Color.Amber },
  { label: 'white', value: Color.White },
]
const fonts = [
  { label: 'sans', value: Font.Sans },
  { label: 'futura', value: Font.Futura },
  { label: 'georgia', value: Font.Georgia },
]
const sizes = [
  { label: 'bigger', value: SizeChange.Bigger },
  { label: 'smaller', value: SizeChange.Smaller },
]
const positions = [
  { label: 'upper', value: PositionChange.Upper },
  { label: 'lower', value: PositionChange.Lower },
]
const outOfSyncStates = [
  { label: 'too early', value: OutOfSyncState.TooEarly },
  { label: 'too late', value: OutOfSyncState.TooLate },
  { label: 'a bit early', value: OutOfSyncState.ABitEarly },
  { label: 'a bit late', value: OutOfSyncState.ABitLate },
]

const Radio = ({ label, items, handleChange }: GroupArgs) => (
  <div>
    <span className="font-serif text-sm underline decoration-1">{label}</span>
    <RadioGroup onChange={handleChange}>
      <div className="flex gap-2">
        {items.map(({ label, value }) => (
          <RadioGroup.Option key={value} value={value}>
            <Button>{label}</Button>
          </RadioGroup.Option>
        ))}
      </div>
    </RadioGroup>
  </div>
)

const Group = ({ label, items, handleChange }: GroupArgs) => (
  <div>
    <span className="font-serif text-sm underline decoration-1">{label}</span>
    <div
      className={
        items.length > 3 ? 'grid grid-cols-2 gap-y-1 gap-x-2' : 'flex gap-2'
      }
    >
      {items.map(({ label, value }) => (
        <Button key={value} onClick={() => handleChange(value)}>
          {label}
        </Button>
      ))}
    </div>
  </div>
)

export const SubtitleSettings = () => {
  const { emit } = useWs()

  return (
    <div className="mx-16 mt-9 flex flex-col items-center">
      <h1 className="text-xl font-bold">Settings</h1>
      <div className="mt-10">
        <h2 className="mb-0.5 text-lg font-semibold">Subtitles</h2>
        <div className="flex flex-col gap-2.5">
          <Radio
            label="style"
            items={styles}
            handleChange={style => emit({ style })}
          />

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

          <Group
            label="out of sync"
            items={outOfSyncStates}
            handleChange={state => emit({ state })}
          />
        </div>
      </div>
    </div>
  )
}
