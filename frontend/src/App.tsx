import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Component_Home } from "./components/Component_Home";
import { Component_Header } from "./components/Component_Header";
import { Component_Footer } from "./components/Component_Footer";
import { Component_NotFound } from "./components/Component_NotFound";
import { Component_Movie } from "./components/Component_Movie";
import { Component_Content } from "./components/Component_Content";
import { Component_Viewer } from "./components/Component_Viewer";

function App() {
  return (
    <>
      <div>
        <BrowserRouter>
          <Component_Header />
          <Routes>
            <Route path="/" element={<Component_Home />} />
            <Route path="/peliculas" element={<Component_Movie />} />
            <Route path="/series" element={<Component_Content />} />
            <Route path="/animes" element={<Component_Content />} />
            <Route path="/viewer" element={<Component_Viewer />} />


            <Route path="*" element={<Component_NotFound />} />
          </Routes>
          <Component_Footer />
        </BrowserRouter>

      </div>
    </>
  );
}

export default App;
