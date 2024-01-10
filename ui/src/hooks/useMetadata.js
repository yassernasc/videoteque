import useSWR from 'swr'

const fetcher = (...args) => fetch(...args).then(res => res.json())

export const useMetadata = () => {
  const { data } = useSWR('/metadata', fetcher)
  return data
}
