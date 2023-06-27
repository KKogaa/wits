import { Button } from "./Button"
import { Card } from "./Card"

export const Collection = () => {

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

  return (
    <div>
      <div className="m-5">
        <Button text={"Upload"} />
      </div>
      <div className="flex flex-row flex-wrap justify-center">
        {fakeData.map((item) => (
          <Card />
        ))}
      </div>


    </div>
  )
}
