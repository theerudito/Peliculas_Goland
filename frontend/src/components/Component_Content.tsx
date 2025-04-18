import { useEffect, useState } from "react";
import { GET_Movies } from "../helpers/Fetching";
import { Component_Player } from "./Component_Player";
import "../styles/Content.css";
import { Movies_List } from "../helpers/Data";
import { MoviesDTO } from "../models/Movies";

export const Component_Content = () => {
  const [value, setValue] = useState<MoviesDTO[]>(Movies_List);
  const [playing, setPlaying] = useState(false);
  const [url, setUrl] = useState("");

  const getMovies = async () => {
    const obj = await GET_Movies();

    setValue(obj);
  };

  const playVideo = (url: string) => {
    if (url === "") return;
    setPlaying(true);
    setUrl(url);
  };

  useEffect(() => {
    getMovies();
  }, []);

  return (
    <div className="container-content">
      {playing === false ? (
        <div className="container-body">
          {value.map((item) => (
            <div key={item.movie_movie_id} className="container-card">
              <img
                src={item.movie_cover}
                alt={item.movie_cover}
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
                <p className="card-title">{item.movie_title}</p>
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
  );
};
