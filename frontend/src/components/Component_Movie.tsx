import { useEffect } from "react";
import { Component_Player } from "./Component_Player";
import { useMovies } from "../store/useMovies";
import cover from "../assets/logo.webp";
import { usePlayer } from "../store/usePlayer";
import "../styles/Styles_Serie_Anime_Movies.css";
import { useAuth } from "../store/useAuth";


export const Component_Movie = () => {
  const { list_movies, getMovies } = useMovies((state) => state);
  const { open_player, playing } = usePlayer((state) => state);
  const { isLogin } = useAuth((state) => state);

  useEffect(() => {
    getMovies();
  }, [getMovies]);

  return (
    <div className="app-container">

      <div className="container-header-search-box">
        <input type="text" placeholder="Buscar..." />
        <button className="container-header-search-btn">
          <i className="bi bi-search"></i>
        </button>
      </div>

      <div className="main-content">

        <div className="container-content">
          {playing === false ? (
            <div className="container-body">
              {list_movies.map((item) => (
                <div key={item.movie_id} className="container-card">
                  <img
                    src={item.movie_cover === "" ? cover : item.movie_cover}
                    alt={cover}
                    className="card-background-image"
                  />

                  <div className="card-overlay">
                    <p className="card-year">{item.movie_year}</p>
                    <div className="card-play">
                      <i
                        className="bi bi-play-circle"
                        onClick={() => open_player(item.movie_url)}
                      ></i>
                    </div>

                    <div className="card-container-button">
                      {
                        isLogin && (
                          <>
                            <i className="bi bi-pencil-square"></i>
                            <i className="bi bi-trash"></i>
                          </>
                        )
                      }

                    </div>
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <div className="container-player">
              <Component_Player />
            </div>
          )}
        </div>
      </div>
    </div>
  );
};
