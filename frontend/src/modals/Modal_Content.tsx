import "../styles/Modal.css";
import { Content_List } from "../helpers/Data";
import { useDispatch, useSelector } from "react-redux";
import { closeModal } from "../store/slice/Modal_Slices";
import { RootState } from "../store/store";

export const Modal_Content = () => {
  const { openModal_Content } = useSelector((store: RootState) => store.modal);
  const content = useSelector((store: RootState) => store.content);
  const values = useSelector((store: RootState) => store.data);

  const form = content.form_Content;
  const dispatch = useDispatch();

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

  const SendData = () => {
    console.log(form);
  };

  const AddEpisode = () => {
    // const newEpisode: Episodes = {
    //   ...episode,
    //   episode_number:
    //     episodeList.length > 0
    //       ? episodeList[episodeList.length - 1].episode_id + 1
    //       : 1,
    //   episode_id:
    //     episodeList.length > 0
    //       ? episodeList[episodeList.length - 1].episode_id + 1
    //       : 1,
    // };
    // setEpisodeList([...episodeList, newEpisode]);
    // Reset_Field(1);
  };

  const RemoveEpisode = (id: number) => {
    // const updatedList = episodeList.filter((item) => item.episode_id !== id);
    // setEpisodeList(updatedList);
  };

  function Reset_Field(action: number) {
    // switch (action) {
    //   case 1:
    //     setEpisode(_episodes);
    //     break;
    //   case 2:
    //     setSeason(_seasons);
    //     break;
    //   case 3:
    //     setContent(_contents);
    //     break;
    // }
  }

  return (
    <div>
      {openModal_Content && (
        <div className="container_modal">
          <div className="container-modal-body">
            <div className="container-modal-header">
              <p>ADD ANIME OR SERIE</p>
              <i
                className="bi bi-x-lg"
                onClick={() => dispatch(closeModal(3))}
              ></i>
            </div>
            <div className="container-modal-input">
              <input
                className="input"
                type="text"
                placeholder="TITLE EPISODE"
                name="episode_title"
                value={form.episode.episode_title}
                onChange={handleChangeInput}
              />
              <input
                className="input"
                type="text"
                placeholder="URL EPISODE"
                name="episode_url"
                value={form.episode.episode_url}
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
                    {content.episodes_list.map((item) => (
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
                value={content.form_Content.gender_id}
                onChange={handleChangeSelect}
              >
                {values.season_list.map((item) => (
                  <option key={item.season_id} value={item.season_id}>
                    {item.name}
                  </option>
                ))}
              </select>

              <input
                className="input"
                type="text"
                placeholder="TITLE SEASON"
                name="season_title"
                value={form.season.season_title}
                onChange={handleChangeInput}
              />

              <select
                name="content_type_id"
                onChange={handleChangeSelect}
                value={content.form_Content.content.content_type_id}
              >
                {values.type_list.map((item) => (
                  <option
                    key={item.content_type_id}
                    value={item.content_type_id}
                  >
                    {item.name}
                  </option>
                ))}
              </select>

              <input
                className="input"
                type="text"
                placeholder="DESCRIPCION"
                name="content_title"
                value={form.content.content_title}
                onChange={handleChangeInput}
              />
              <input
                className="input"
                type="text"
                placeholder="COVER"
                name="content_cover"
                value={form.content.content_cover}
                onChange={handleChangeInput}
              />
              <input
                className="input"
                type="number"
                placeholder="YEAR"
                name="content_year"
                value={form.content.content_year}
                onChange={handleChangeInput}
              />

              <select
                name="gender_id"
                value={content.form_Content.gender_id}
                onChange={handleChangeSelect}
              >
                {values.gender_list.map((item) => (
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
