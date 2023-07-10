import { useState } from "react"
import { createSearchParams, useNavigate } from "react-router-dom"
import { useResultContext } from "../contexts/ResultContextProvider"

export const SearchBar = () => {

  const navigate = useNavigate()

  const [text, setText] = useState("")

  const { setSearchTerm } = useResultContext();

  const search = () => {
    setSearchTerm(text)
    navigate({ pathname: '/search', search: `?${createSearchParams({ searchTerm: `${text}` })}` })
  }

  return (
    <div className="flex flex-col justify-center items-center">
      <h1 className="font-bold text-5xl m-5 mt-20">What's the sauce?</h1>
      <div className="mx-auto">
        <input value={text}
          type="text"
          className="bg-indigo-100 border-black center text-center text-indigo-900 p-1 pl-20 pr-20"
          placeholder="Search using WITS"
          onChange={(e) => setText(e.target.value)} />
        <button className="bg-indigo-600 rounded p-1 m-1"
          onClick={search}
        >🔍</button>
      </div>
    </div>
  )
}
