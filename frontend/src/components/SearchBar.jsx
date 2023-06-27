import { useState } from "react"

export const SearchBar = () => {

  const [text, setText] = useState("")


  return (
    <div className="flex flex-col justify-center items-center">
      <h1 className="font-bold text-5xl m-5 mt-20">What's the sauce?</h1>
      <div className="mx-auto">
        {/* <
        <Results
          utton>Upload</button> */}
        <input value={text}
          type="text"
          className="bg-indigo-100 border-black center text-center text-indigo-900 p-1 pl-20 pr-20"
          placeholder="Search using WITS"
          onChange={(e) => setText(e.target.value)} />
        <button className="bg-indigo-600 rounded p-1 m-1">🔍</button>
      </div>
    </div>
  )
}
