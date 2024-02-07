import { Transition } from '@headlessui/react'
import { useSyncSubtitle } from 'features/subtitle'
import ms from 'ms'
import { useEffect, useState } from 'react'
import { usePrevious, useTimeoutFn } from 'react-use'

export const Toast = () => {
  const [show, setShow] = useState(false)
  const [, , reset] = useTimeoutFn(() => setShow(false), ms('2s'))
  const offset = useSyncSubtitle()
  const prevOffset = usePrevious(offset)

  useEffect(() => {
    const isInitialState = prevOffset === undefined
    const offsetChanged = offset !== prevOffset

    if (!isInitialState && offsetChanged) {
      setShow(true)
      reset()
    }
  }, [offset, prevOffset, reset])

  return (
    <Transition
      show={show}
      enter="transition-opacity duration-50 ease-in"
      enterFrom="opacity-0"
      enterTo="opacity-100"
      leave="transition-opacity duration-200 ease-out"
      leaveFrom="opacity-100"
      leaveTo="opacity-0"
    >
      <div className="absolute top-0 right-0 mx-5 my-6 flex w-1/4 flex-col gap-0.5 rounded border border-zinc-200 bg-white p-3.5 shadow-md">
        <span className="font-semibold">Subtitle Updated</span>
        <span>
          Current Offset: <span className="text-lg font-bold">{offset}</span>
        </span>
      </div>
    </Transition>
  )
}
