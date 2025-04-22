import { useEffect } from "react";
import "../styles/Styles_Serie_Anime_Movies.css";
import { useContent } from "../store/useContent";
import cover from "../assets/logo.webp";

export const Component_Serie_Anime = () => {
  const { list_content, getContent } = useContent((state) => state);

  const playVideo = (url: string) => {
    if (url === "") return;

    //setUrl(url);
  };

  useEffect(() => {
    getContent();
  }, [getContent]);

  return (
    <div className="app-container">
      <div className="main-content">
        <div className="container-content">
          <div className="container-body">
            {list_content.map((item) => (
              <div key={item.content_id} className="container-card">
                <img
                  src={item.content_cover === "" ? cover : item.content_cover}
                  alt={cover}
                  className="card-background-image"
                />

                <div className="card-overlay">
                  <p className="card-year">{item.content_year}</p>
                  <div className="card-play">
                    <i
                      className="bi bi-play-circle"
                      onClick={() => playVideo(item.content_id.toString())}
                    ></i>
                  </div>
                  <p className="card-title">{item.content_title}</p>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};
