import { useEffect, useState } from "react";
import { useContent } from "../store/useContent";
import "../styles/Styles_Content.css";

export const Component_Content = ({ id }: { id: number }) => {
  const [selectedSeasonIndex, setSelectedSeasonIndex] = useState(0);
  const { getContent_full, list_full_data } = useContent(
    (state) => state
  );

  const handleChangeSelect = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const index = parseInt(e.target.value, 10);
    setSelectedSeasonIndex(index);
  };

  useEffect(() => {
    getContent_full(id);
  }, [getContent_full, id]);

  return (
    <div className="anime-viewer">
      <div className="anime-main">
        <div className="anime-image">
          <img src={list_full_data?.content.content_cover} />
        </div>

        <div className="anime-info">
          <h5>{list_full_data?.content.content_title}</h5>
          <p>{list_full_data?.content.content_year}</p>
          <p>{list_full_data?.content.content_gender}</p>
          <select onChange={handleChangeSelect}>
            {list_full_data?.seasons.map((season, index) => (
              <option key={season.season_id} value={index}>
                {season.season_name}
              </option>
            ))}
          </select>
        </div>
      </div>

      <div className="episode-scroll">
        <div className="episode-grid">
          {list_full_data?.seasons[selectedSeasonIndex].episodes.map(
            (item, index) => (
              <div key={index} className="episode">
                <div className="episode-image">
                  <img src={list_full_data?.content.content_cover} />
                  <div className="play-button">
                    <i className="bi bi-play-circle"></i>
                  </div>
                </div>
                <p className="episode-title">{item.episode_name}</p>
              </div>
            )
          )}
        </div>
      </div>
    </div>
  );
};
