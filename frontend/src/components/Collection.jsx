import { useNavigate } from "react-router-dom"
import { Button } from "./Button"
import { Card } from "./Card"

export const Collection = () => {

  const navigate = useNavigate()

  const fakeData = [
    {
      url: "link",
      source: "source",
      title: "title",
    },
    {
      url: "link",
      source: "source",
      title: "title",
    },
  ]

  const uploadImage = async () => {
    navigate("/upload")
  }

  return (
    <div>
      <div className="m-5">
        <button onClick={uploadImage} className="bg-indigo-400 rounded p-1">Upload</button>
      </div>
      <div className="flex flex-row flex-wrap justify-center">
        {fakeData.map((item) => (
          <Card />
        ))}
      </div>


    </div>
  )
}
