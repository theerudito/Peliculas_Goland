
import { Modal_Movie } from "../modals/Modal_Movie";
import { useContent } from "../store/useContent";
import { useModal } from "../store/useModal";
import "../styles/Styles_Header.css";
import { Link } from "react-router-dom";
import { Modal_Episodes } from "../modals/Modal_Episodes";
import { Modal_Content } from "../modals/Modal_Content";
import { Modal_Login } from "../modals/Modal_Login";
import { Modal_Season } from "../modals/Modal_Season";
import { Modal_Gender } from "../modals/Modal_Gender";
import { useAuth } from "../store/useAuth";

export const Component_Header = () => {

  const { changeType } = useContent((state) => state);
  const { isLogin } = useAuth((state) => state);
  const { OpenModal_Login, OpenModal_Content, OpenModal_Movie, OpenModal_Episodes } = useModal(
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

          <div onClick={() => OpenModal_Episodes(true)}>
            <i className="bi bi-film"></i>
          </div>

          {isLogin === false ? (
            <div onClick={() => OpenModal_Login(true)}>
              <i className="bi bi-person-fill"></i>
            </div>
          ) : (
            <div>
              <i className="bi bi-person-fill-check"></i>
            </div>
          )}



        </div>
      </div >

      <Modal_Episodes />
      <Modal_Content />
      <Modal_Movie />
      <Modal_Login />
      <Modal_Season />
      <Modal_Gender />
    </>
  );
};
