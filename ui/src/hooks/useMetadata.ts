import { useEffect, useState } from 'react'

export const useMetadata = () => {
  const [metadata, setMetadata] = useState(null)

  useEffect(() => {
    fetch('/metadata')
      .then(res => res.json())
      .then(setMetadata)
  }, [])

  return metadata
}
