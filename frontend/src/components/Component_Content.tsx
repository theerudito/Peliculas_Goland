import { useState } from "react";
import "../styles/Styles_Content.css";

export const Component_Content = () => {
  const [season] = useState("1");

  const episodes = Array(100).fill(`S ${season} - E1`);

  return (
    <div className="anime-viewer">
      <div className="anime-main">
        <div className="anime-image">
          <img src="https://cuevana.biz/_next/image?url=https%3A%2F%2Fimage.tmdb.org%2Ft%2Fp%2Foriginal%2FcoMNNwjHY4BZTqMIsklanMf2Wp7.jpg&w=256&q=75" />
        </div>

        <div className="anime-info">
          <h5>JUJUTSU KAISEN</h5>
          <p>2025</p>
          <p>Anime</p>
          <select>
            <option>TEMPORADA 1</option>
          </select>
        </div>
      </div>

      <div className="episode-scroll">
        <div className="episode-grid">
          {episodes.map((ep, index) => (
            <div key={index} className="episode">
              <div className="episode-image">
                <img src="https://latanime.org/assets/img/serie/imagen/my-hero-academia-vigilantes-1744052990.jpg" />
                <div className="play-button">
                  <i className="bi bi-play-circle"></i>
                </div>
              </div>
              <p className="episode-title">{ep}</p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};
