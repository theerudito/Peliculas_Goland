import { useModal } from "../store/useModal";
import "../styles/Header.css";
import { useState } from "react";

export const Component_Header = () => {
  const [userIn] = useState(false);
  const { OpenModal_Auth, OpenModal_Content, OpenModal_Movie } = useModal(
    (state) => state
  );

  return (
    <div className="container-header-header">
      <nav className="container-header-nav-links">
        <a href="#">Inicio</a>
        <a href="#">Peliculas</a>
        <a href="#">Series</a>
        <a href="#">Anime</a>
      </nav>
      <div className="container-header-search-box">
        <input type="text" placeholder="Buscar..." />
        <button className="container-header-search-btn">
          <i className="bi bi-search"></i>
        </button>
      </div>
      <div className="container-header-buttons">
        <button onClick={() => OpenModal_Movie(true)}>ADD MOVIES</button>
        <button onClick={() => OpenModal_Content(true)}>ADD CONTENT</button>

        {userIn === true ? (
          <button onClick={() => OpenModal_Auth(true)}>LOGOUT</button>
        ) : (
          <button onClick={() => OpenModal_Auth(true)}>LOGIN</button>
        )}
      </div>
    </div>
  );
};
