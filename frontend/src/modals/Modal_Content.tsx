import { useEffect, useState } from "react";
import "../styles/Modal.css";
import {
  Content_List,
  Episode_List,
  Gender_List,
  Season_List,
} from "../helpers/Data";
import {
  _contents,
  _episodes,
  _seasons,
  Content_Types,
  Episodes,
  Seasons,
} from "../models/Movies";

export const Modal_Content = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [episode, setEpisode] = useState<Episodes>(_episodes);
  const [season, setSeason] = useState<Seasons>(_seasons);
  const [content, setContent] = useState<Content_Types>(_contents);
  const [gender, setGender] = useState({ gender_id: 0 });
  const [episodeList, setEpisodeList] = useState<Episodes[]>(Episode_List);

  //const [formData, setFormData] = useState<FormDataDTO>();

  const handleChangeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setEpisode((prevData) => ({
      ...prevData,
      [name]: value,
    }));

    setSeason((prevData) => ({
      ...prevData,
      [name]: value,
    }));

    setContent((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleChangeSelect = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const { name, value } = e.target;

    setGender((prevData) => ({
      ...prevData,
      [name]: parseInt(value, 10),
    }));
  };

  const SendData = () => {


    
  };

  const AddEpisode = () => {
    const newEpisode: Episodes = {
      ...episode,
      episode_id:
        episodeList.length > 0
          ? episodeList[episodeList.length - 1].episode_id + 1
          : 1,
    };

    setEpisodeList([...episodeList, newEpisode]);
  };

  const RemoveEpisode = (id: number) => {
    const updatedList = episodeList.filter((item) => item.episode_id !== id);
    setEpisodeList(updatedList);
  };

  useEffect(() => {}, []);

  return (
    <div>
      <button className="open-modal-btn" onClick={() => setIsOpen(true)}>
        Abrir Modal
      </button>

      {isOpen && (
        <div className="container_modal">
          <div className="container-modal-body">
            <div className="container-modal-header">
              <p>ADD ANIME OR SERIE</p>
              <i className="bi bi-x-lg" onClick={() => setIsOpen(false)}></i>
            </div>
            <div className="container-modal-input">
              <input
                type="text"
                placeholder="TITLE EPISODE"
                name="episode_title"
                value={episode.episode_title}
                onChange={handleChangeInput}
              />
              <input
                type="text"
                placeholder="URL EPISODE"
                name="episode_url"
                value={episode.episode_url}
                onChange={handleChangeInput}
              />

              <button onClick={AddEpisode}>1 ADD EPISODE</button>
              <div className="container-modal-table">
                <table>
                  <thead>
                    <tr>
                      <th>TITLE</th>
                      <th>EPISODE</th>
                      <th>A</th>
                    </tr>
                  </thead>
                  <tbody>
                    {episodeList.map((item) => (
                      <tr key={item.episode_id}>
                        <td>{item.episode_title}</td>
                        <td>EPISODE {item.episode_number}</td>
                        <td>
                          <i
                            className="bi bi-trash"
                            onClick={() => RemoveEpisode(item.episode_id)}
                          ></i>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>

              <select
                name="season_id"
                value={season.season_id}
                onChange={handleChangeSelect}
              >
                {Season_List.map((item) => (
                  <option key={item.season_id} value={item.season_id}>
                    {item.name}
                  </option>
                ))}
              </select>

              <input
                type="text"
                placeholder="TITLE SEASON"
                name="season_title"
                value={season.season_title}
                onChange={handleChangeInput}
              />

              <select
                name="content_type_id"
                onChange={handleChangeSelect}
                value={content.content_type_id}
              >
                {Content_List.map((item) => (
                  <option
                    key={item.content_type_id}
                    value={item.content_type_id}
                  >
                    {item.name}
                  </option>
                ))}
              </select>

              <input
                type="text"
                placeholder="DESCRIPCION"
                name="content_title"
                value={content.content_title}
                onChange={handleChangeInput}
              />
              <input
                type="text"
                placeholder="COVER"
                name="content_cover"
                value={content.content_cover}
                onChange={handleChangeInput}
              />
              <input
                type="number"
                placeholder="YEAR"
                name="content_year"
                value={content.content_year}
                onChange={handleChangeInput}
              />

              <select
                name="gender_id"
                value={gender.gender_id}
                onChange={handleChangeSelect}
              >
                {Gender_List.map((item) => (
                  <option key={item.gender_id} value={item.gender_id}>
                    {item.name}
                  </option>
                ))}
              </select>

              <button onClick={SendData}> 2 ADD CONTENT</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
