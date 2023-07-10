import { Routes, Route } from "react-router-dom"
import { SearchBar } from "./SearchBar"
import { Results } from "./Results"
import { Collection } from "./Collection"
import { Upload } from "./Upload"

export const Routed = () => {
  return (
    <div>
      <Routes>
        <Route path="/" element={<SearchBar />} />
        <Route path="/search" element={<Results />} />
        <Route path="/collection" element={<Collection />} />
        <Route path="/upload" element={<Upload />} />
      </Routes>
    </div>
  )
}

