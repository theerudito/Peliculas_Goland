import { useDispatch } from "react-redux";
import "../styles/Header.css";
import { openModal } from "../store/slice/Modal_Slices";
import { useState } from "react";

export const Component_Header = () => {
  const [userIn] = useState(false);
  const dispatch = useDispatch();

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
        <button onClick={() => dispatch(openModal(2))}>ADD MOVIES</button>
        <button onClick={() => dispatch(openModal(3))}>ADD CONTENT</button>

        {userIn === true ? (
          <button onClick={() => dispatch(openModal(1))}>LOGOUT</button>
        ) : (
          <button onClick={() => dispatch(openModal(1))}>LOGIN</button>
        )}
      </div>
    </div>
  );
};
