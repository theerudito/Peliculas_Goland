import { useEffect, useState } from "react";
import "../styles/Styles_Serie_Anime_Movies.css";
import { useContent } from "../store/useContent";
import cover from "../assets/logo.webp";
import { Component_Content } from "./Component_Content";

export const Component_Serie_Anime = () => {
  const { list_content_type, getContent_Type } = useContent((state) => state);

  const [isPlaying, setIsPlaying] = useState(true);

  const [id, setId] = useState(0);

  const playVideo = (id: number) => {
    if (id === null) return;

    setIsPlaying(false);
    setId(id);
  };

  useEffect(() => {
    getContent_Type();
  }, [getContent_Type]);

  return (
    <div className="app-container">
      <div className="main-content">
        <div className="container-content">
          <div className="container-body">
            {isPlaying === true ? (
              <>
                {list_content_type.map((item) => (
                  <div key={item.content_id} className="container-card">
                    <img
                      src={
                        item.content_cover === "" ? cover : item.content_cover
                      }
                      alt={cover}
                      className="card-background-image"
                    />

                    <div className="card-overlay">
                      <p className="card-year">{item.content_year}</p>
                      <div className="card-play">
                        <i
                          className="bi bi-play-circle"
                          onClick={() => playVideo(item.content_id)}
                        ></i>
                      </div>
                      <p className="card-title">{item.content_title}</p>
                    </div>
                  </div>
                ))}
              </>
            ) : (
              <>
                <Component_Content id={id} />
              </>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};
