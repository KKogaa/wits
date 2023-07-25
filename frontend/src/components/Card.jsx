
export const Card = (props) => {
  return (
    //TODO: add clickable modal functionality, maybe to download the image idk
    <div className="block rounded-lg p-4 shadow-sm shadow-indigo-100">
      <img
        src={props.url}
        className="h-32 w-48 rounded-md object-cover border-4 border-indigo-900"
      />
    </div>
  )
}
