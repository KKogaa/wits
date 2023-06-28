import { createContext, useContext, useState } from 'react'
import axios from 'axios'

const defaultResultState = {
  results: [],
  setResults: () => { },
  getResults: () => { },
  searchTerm: '',
  setSearchTerm: () => { },
}


const ResultContext = createContext(defaultResultState)
export const useResultContext = () => useContext(ResultContext);

const url = "http://localhost:8080/search/text"

export const ResultContextProvider = ({ children }) => {

  const [results, setResults] = useState([])
  const [searchTerm, setSearchTerm] = useState('')


  const getResults = async (text) => {
    const res = await axios.get(`${url}`, { params: { text: `${text}` } }).then(
      (response) => {
        setResults(response.data)
      }
    ).catch((error) => {
      console.log(error)
    })
  }

  return (
    <ResultContext.Provider value={{ getResults, results, searchTerm, setSearchTerm }}>
      {children}
    </ResultContext.Provider>
  )
}

