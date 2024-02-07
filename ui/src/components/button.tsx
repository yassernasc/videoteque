export const Button = ({ children, ...props }) => (
  <button
    {...props}
    className="border-2 border-black py-1 px-2 text-xs font-semibold hover:bg-yellow-400/40 focus:bg-yellow-400/40 active:bg-yellow-400 active:drop-shadow-md ui-checked:bg-yellow-400/70"
  >
    {children}
  </button>
)
