import { Link } from "react-router-dom"

export const Navbar = () => {

  return (
    <div
      className="mx-auto flex h-16 items-center gap-8 px-4 bg-indigo-100 sm:px-6 lg:px-8"
    >
      <Link className="bg-indigo-400 rounded p-3 text-indigo-900 font-bold" to={"/"}>WITS 🧠</Link>

      <div className="flex flex-1 items-center justify-end md:justify-between">
        <ul className="flex items-center gap-6 text-sm">
          <Link className="text-indigo-900 font-bold" to={"/collection"}>Collection</Link>
        </ul>

        {/* <div class="flex items-center gap-4"> */}
        {/*   <div class="sm:flex sm:gap-4"> */}
        {/*     <a */}
        {/*       class="block rounded-md bg-teal-600 px-5 py-2.5 text-sm font-medium text-white transition hover:bg-teal-700" */}
        {/*       href="/" */}
        {/*     > */}
        {/*       Login */}
        {/*     </a> */}

        {/*     <a */}
        {/*       class="hidden rounded-md bg-gray-100 px-5 py-2.5 text-sm font-medium text-teal-600 transition hover:text-teal-600/75 sm:block" */}
        {/*       href="/" */}
        {/*     > */}
        {/*       Register */}
        {/*     </a> */}
        {/*   </div> */}
        {/* </div> */}
      </div>
    </div >
  )
}
