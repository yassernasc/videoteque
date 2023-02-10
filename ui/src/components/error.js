export const Error = ({ message }) => (
  <div className="flex h-screen w-screen items-center justify-center bg-black">
    <span className="text-white">{`Error playing movie >> ${message}`}</span>
  </div>
)
