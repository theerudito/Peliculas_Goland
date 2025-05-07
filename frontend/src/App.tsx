import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Component_Home } from "./components/Component_Home";
import { Component_Header } from "./components/Component_Header";
import { Component_Footer } from "./components/Component_Footer";
import { Component_NotFound } from "./components/Component_NotFound";
import { Component_Movie } from "./components/Component_Movie";
import { Component_Serie_Anime } from "./components/Component_Serie_Anime";
import { Component_Content } from "./components/Component_Content";

function App() {
  return (
    <>
      <div>
        <BrowserRouter>
          <Component_Header />
          <Routes>
            <Route path="/" element={<Component_Home />} />
            <Route path="/peliculas" element={<Component_Movie />} />
            <Route path="/series" element={<Component_Serie_Anime />} />
            <Route path="/animes" element={<Component_Serie_Anime />} />
            <Route path="/content" element={<Component_Content id={0} />} />

            <Route path="*" element={<Component_NotFound />} />
          </Routes>
          <Component_Footer />
        </BrowserRouter>
      </div>
    </>
  );
}

export default App;
