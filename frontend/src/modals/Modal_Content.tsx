import { _season } from "../models/Seasons";
import { useContent } from "../store/useContent";
import { useData } from "../store/useData";
import { useModal } from "../store/useModal";
import "../styles/Styles_Modal.css";

export const Modal_Content = () => {
  const {
    gender_list,
    form_gender,
    season_list,
    form_season,
    type_list,
    form_type,
  } = useData((state) => state);
  const { listEpisode, form_Episode, form_content, addEpisode, removeEpisode } =
    useContent((state) => state);
  const { OpenModal_Content, _modal_Content } = useModal((state) => state);

  const handleChangeInput = async (e: React.ChangeEvent<HTMLInputElement>) => {
    // const { name, value } = e.target;
    // setEpisode((prevData) => ({
    //   ...prevData,
    //   [name]: value,
    // }));
    // setSeason((prevData) => ({
    //   ...prevData,
    //   [name]: value,
    // }));
    // setContent((prevData) => ({
    //   ...prevData,
    //   [name]: value,
    // }));
  };

  const handleChangeSelect = (e: React.ChangeEvent<HTMLSelectElement>) => {
    // const { name, value } = e.target;
    // setGender((prevData) => ({
    //   ...prevData,
    //   [name]: parseInt(value, 10),
    // }));
    // setSeason((prevData) => ({
    //   ...prevData,
    //   [name]: parseInt(value, 10),
    // }));
    // setContent((prevData) => ({
    //   ...prevData,
    //   [name]: parseInt(value, 10),
    // }));
  };

  const SendData = () => {};

  return (
    <div>
      {_modal_Content && (
        <div className="container_modal">
          <div className="container-modal-body">
            <div className="container-modal-header">
              <p>ADD ANIME OR SERIE</p>
              <i
                className="bi bi-x-lg"
                onClick={() => OpenModal_Content(false)}
              ></i>
            </div>
            <div className="container-modal-input">
              <input
                className="input"
                type="text"
                placeholder="TITLE EPISODE"
                name="episode_title"
                value={form_Episode.episode_title}
                onChange={handleChangeInput}
              />
              <input
                className="input"
                type="text"
                placeholder="URL EPISODE"
                name="episode_url"
                value={form_Episode.episode_url}
                onChange={handleChangeInput}
              />

              <button onClick={() => addEpisode(form_Episode)}>
                1 ADD EPISODE
              </button>
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
                    {listEpisode.map((item) => (
                      <tr key={item.episode_id}>
                        <td>{item.episode_title}</td>
                        <td>EPISODE {item.episode_number}</td>
                        <td>
                          <i
                            className="bi bi-trash"
                            onClick={() => removeEpisode(item.episode_id)}
                          ></i>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>

              <select
                name="season_id"
                value={form_season.season_id}
                onChange={handleChangeSelect}
              >
                {season_list.map((item) => (
                  <option key={item.season_id} value={item.season_id}>
                    {item.season_name}
                  </option>
                ))}
              </select>

              <input
                className="input"
                type="text"
                placeholder="TITLE SEASON"
                name="season_name"
                value={_season.season_name}
                onChange={handleChangeInput}
              />

              <select
                name="content_type_id"
                onChange={handleChangeSelect}
                value={form_type.content_type_id}
              >
                {type_list.map((item) => (
                  <option
                    key={item.content_type_id}
                    value={item.content_type_id}
                  >
                    {item.content_type_title}
                  </option>
                ))}
              </select>

              <input
                className="input"
                type="text"
                placeholder="DESCRIPCION"
                name="content_title"
                value={form_content.content_title}
                onChange={handleChangeInput}
              />
              <input
                className="input"
                type="text"
                placeholder="COVER"
                name="content_cover"
                value={form_content.content_cover}
                onChange={handleChangeInput}
              />
              <input
                className="input"
                type="number"
                placeholder="YEAR"
                name="content_year"
                value={form_content.content_year}
                onChange={handleChangeInput}
              />

              <select
                name="gender_id"
                value={form_gender.gender_id}
                onChange={handleChangeSelect}
              >
                {gender_list.map((item) => (
                  <option key={item.gender_id} value={item.gender_id}>
                    {item.gender_name}
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
