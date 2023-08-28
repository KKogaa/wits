import { useNavigate } from "react-router-dom"
import { Button } from "./Button"
import { Card } from "./Card"
import { useEffect, useState } from "react"
import axios from "axios"

export const Collection = () => {

  const [data, setData] = useState([])
  const navigate = useNavigate()

  useEffect(() => {
    const fetchData = async () => {
      try {

        const url = "http://localhost:8080/vectors"
        const res = await axios.get(`${url}`).then(
          (response) => {
            setData(response.data)
          }
        ).catch((error) => {
          console.log(error)
        })

      } catch (error) {
        console.error("Error fetching data:", error)
      }
    }

    fetchData()
  }, [])

  const uploadImage = async () => {
    navigate("/upload")
  }

  return (
    <div>
      <div className="flex flex-row">
        <h1 className="text-2xl text-indigo-900 font-bold m-6 p-1">My Collection</h1>
        <div className="m-6">
          <button onClick={uploadImage} className="bg-indigo-400 rounded p-1">Upload</button>
        </div>
      </div>
      <hr />

      <div className="flex flex-row flex-wrap justify-center">
        {data.map((item) => (
          <Card url={item.path} />
        ))}
      </div>


    </div>
  )
}
