import { useState } from "react";
import { Modal_Auth } from "../modals/Modal_Auth";
import { Modal_Movie } from "../modals/Modal_Movie";
import { useContent } from "../store/useContent";
import { useModal } from "../store/useModal";
import "../styles/Styles_Header.css";
import { Link } from "react-router-dom";
import { Modal_Episodes } from "../modals/Modal_Episodes";

export const Component_Header = () => {
  const [userIn] = useState(false);
  const { changeType } = useContent((state) => state);
  const { OpenModal_Auth, OpenModal_Content, OpenModal_Movie } = useModal(
    (state) => state
  );

  return (
    <>
      <div className="container-header-header">
        <nav className="container-header-nav-links">
          <Link to="/">Inicio</Link>
          <Link to="/peliculas">Películas</Link>
          <Link to="/series" onClick={() => changeType(2)}>
            Series
          </Link>
          <Link to="/animes" onClick={() => changeType(1)}>
            Anime
          </Link>
        </nav>
        <div className="container-header-search-box">
          <input type="text" placeholder="Buscar..." />
          <button className="container-header-search-btn">
            <i className="bi bi-search"></i>
          </button>
        </div>
        <div className="container-header-buttons">
          <div onClick={() => OpenModal_Movie(true)}>
            <i className="bi bi-camera-reels"></i>
          </div>

          <div onClick={() => OpenModal_Content(true)}>
            <i className="bi bi-film"></i>
          </div>

          {userIn === true ? (
            <div onClick={() => OpenModal_Auth(true)}>
              <i className="bi bi-person-fill"></i>
            </div>
          ) : (
            <div>
              <i className="bi bi-person-fill-check"></i>
            </div>
          )}
        </div>
      </div>
      <Modal_Episodes />
      <Modal_Movie />
      <Modal_Auth />
    </>
  );
};
