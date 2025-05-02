import { useEffect } from "react";
import { Component_Player } from "./Component_Player";
import { useMovies } from "../store/useMovies";
import cover from "../assets/logo.webp";
import { usePlayer } from "../store/usePlayer";
import "../styles/Styles_Serie_Anime_Movies.css";

export const Component_Movie = () => {
  const { list_movies, getMovies } = useMovies((state) => state);
  const { open_player, playing, url } = usePlayer((state) => state);

  const playVideo = (url: string) => {
    if (url === "") return;
    open_player(url);
  };

  useEffect(() => {
    getMovies();
  }, [getMovies]);

  return (
    <div className="app-container">
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
                        onClick={() => playVideo(item.movie_url)}
                      ></i>
                    </div>
                    <p className="card-title"></p>
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <div className="container-player">
              <Component_Player url={url} />
            </div>
          )}
        </div>
      </div>
    </div>
  );
};
