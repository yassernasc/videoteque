type Props = { message: string }

export const Error = ({ message }: Props) => (
  <div className="flex h-screen w-screen items-center justify-center bg-black">
    <span className="font-bold text-white">{message}</span>
  </div>
)
