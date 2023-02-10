export const H1 = ({ className = '', children }) => (
  <h1 className={`${className} text-2xl font-bold`}>{children}</h1>
)

export const H2 = ({ className = '', children }) => (
  <h2 className={`${className} text-lg font-semibold`}>{children}</h2>
)

export const Caption = ({ className = '', children }) => (
  <span className={`${className} font-georgia text-sm underline decoration-1`}>
    {children}
  </span>
)
