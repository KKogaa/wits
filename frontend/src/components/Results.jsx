import { useResultContext } from "../contexts/ResultContextProvider"
import { useEffect } from "react"
import { Card } from "./Card";
import { useSearchParams } from "react-router-dom";

export const Results = () => {

  const { setSearchTerm, getResults, results } = useResultContext();
  const [searchParams, setSearchParams] = useSearchParams()

  useEffect(() => {
    const searchTerm = searchParams.get('searchTerm')
    getResults(searchTerm)
  }, [])

  return (
    <div className="flex flex-row flex-wrap justify-center">
      {results.map((result) => (
        <div key={result.path}>
          <Card url={result.path} />
        </div>
      ))}
    </div>
  )
}
