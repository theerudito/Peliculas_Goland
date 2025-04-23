import { useEffect } from "react";
import "../styles/Styles_Content.css";
import { useContent } from "../store/useContent";

export const Component_Content = ({ id }: { id: string }) => {

  const { getEpisode, getSeason, data_content } = useContent((state) => state);

  useEffect(() => {
    getSeason();
    getEpisode();
  }, [getSeason, getEpisode]);


  //console.log(data_content[0].content)

  return (
    <div className="anime-viewer">
      <div className="anime-main">
        <div className="anime-image">
          <img src="https://cuevana.biz/_next/image?url=https%3A%2F%2Fimage.tmdb.org%2Ft%2Fp%2Foriginal%2FcoMNNwjHY4BZTqMIsklanMf2Wp7.jpg&w=256&q=75" />
        </div>

        <div className="anime-info">
          <h5>{data_content[0].content.content_title}</h5>
          <p>{data_content[0].content.content_year}</p>
          <p>{data_content[0].content.gender_id}</p>
          <select>
            <option>TEMPORADA 1</option>
          </select>
        </div>
      </div>

      <div className="episode-scroll">
        <div className="episode-grid">
          {/* {


            data_content[0].seasons[0].episodes.map((ep) => (
              <div key={ep.episode_id} className="episode">
                <div className="episode-image">
                  <img src="https://latanime.org/assets/img/serie/imagen/my-hero-academia-vigilantes-1744052990.jpg" />
                  <div className="play-button">
                    <i className="bi bi-play-circle"></i>
                  </div>
                </div>
                <p className="episode-title">{ep.episode_name}</p>
              </div>
            ))} */}
        </div>
      </div>
    </div>
  );
};
