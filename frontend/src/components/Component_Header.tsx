import { useState } from "react";
import { Modal_Auth } from "../modals/Modal_Auth";
import { Modal_Content } from "../modals/Modal_Content";
import { Modal_Movie } from "../modals/Modal_Movie";
import { useContent } from "../store/useContent";
import { useModal } from "../store/useModal";
import "../styles/Header.css";
import { Link } from "react-router-dom";

export const Component_Header = () => {
  const [userIn] = useState(false);
  const { changeType } = useContent((state) => state)
  const { OpenModal_Auth, OpenModal_Content, OpenModal_Movie } = useModal(
    (state) => state
  );

  return (
    <>
      <div className="container-header-header">
        <nav className="container-header-nav-links">
          <Link to="/">Inicio</Link>
          <Link to="/peliculas">Películas</Link>
          <Link to="/series" onClick={() => changeType(2)}>Series</Link>
          <Link to="/animes" onClick={() => changeType(1)}>Anime</Link>
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
      <Modal_Content />
      <Modal_Movie />
      <Modal_Auth />
    </>

  );
};
