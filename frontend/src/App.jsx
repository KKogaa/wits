import { Navbar } from "./components/Navbar"
import { Routed } from "./components/Routed"
import { Helmet } from "react-helmet";
import favicon from "./assets/brain.svg"

function App() {

  return (
    <>
      <Helmet>
        <meta charSet="utf-8" />
        <title>WITS</title>
        <link rel="icon" type="image/png" href={favicon} sizes="16x16" />
      </Helmet>
      <Navbar />
      <Routed />
    </>
  )
}

export default App
