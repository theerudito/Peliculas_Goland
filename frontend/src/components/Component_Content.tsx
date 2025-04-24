import { useEffect } from "react";
import "../styles/Styles_Content.css";
import { useContent } from "../store/useContent";
import { ContentDTO } from "../models/Contents";

export const Component_Content = ({ obj }: { obj: ContentDTO }) => {
  const { getEpisode, getSeason, data_content, listSeason } = useContent(
    (state) => state
  );

  // console.log("data_content", data_content);

  // useEffect(() => {
  //   getSeason(obj.content_title);
  //   getEpisode(obj.content_title, obj.content_id);
  // }, [getSeason, getEpisode, obj.content_title, obj.content_id]);

  return (
    <div className="anime-viewer">
      {/* <div className="anime-main">
        <div className="anime-image">
          <img src={obj.content_cover} />
        </div>

        <div className="anime-info">
          <h5>{obj.content_title}</h5>
          <p>{obj.content_year}</p>
          <p>{obj.content_gender}</p>
          <select>
            {listSeason.map((item, index) => (
              <option key={index} value={item.season_id}>
                {item.season_name}
              </option>
            ))}
          </select>
        </div>
      </div> */}

      {/* <div className="episode-scroll">
        <div className="episode-grid">
          {data_content[0].seasons[0].episodes.map((item, index) => (
            <div key={index} className="episode">
              <div className="episode-image">
                <img src={obj.content_cover} />
                <div className="play-button">
                  <i className="bi bi-play-circle"></i>
                </div>
              </div>
              <p className="episode-title">{item.episode_name}</p>
            </div>
          ))}
        </div>
      </div> */}
    </div>
  );
};
