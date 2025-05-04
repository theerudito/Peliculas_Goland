import { useEffect } from "react";
import { useContent } from "../store/useContent";
import { useData } from "../store/useData";
import { useModal } from "../store/useModal";
import "../styles/Styles_Modal.css";
import { ContentDTO_EpisodeDTO } from "../models/Contents";

export const Modal_Episodes = () => {
  const { form_episode, form_content_episodes, list_content, getContent, list_episodes, add_episodes, remove_episodes, obtener_episode, postEpisodes } = useContent((state) => state);
  const { _modal_episodes, OpenModal_Episodes, OpenModal_Season, OpenModal_Content } = useModal((state) => state);
  const { season_list, getSeason } = useData((state) => state);

  const handleChangeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    useContent.setState((state) => ({
      form_episode: {
        ...state.form_episode,
        [name as keyof typeof state.form_episode]: value,
      },
    }));
  };

  const handleChangeSelect = (e: React.ChangeEvent<HTMLSelectElement>) => {

    const { name, value } = e.target;

    useContent.setState((state) => {
      const selectedValue = (name === "content_id" || name === "season_id")
        ? Number(value)
        : value;

      return {
        form_content_episodes: {
          ...state.form_content_episodes,
          [name]: selectedValue,
        },
      };
    });
  };


  const sendData = () => {

    const { content_id, season_id } = form_content_episodes;

    if (!content_id || content_id === 0 || !season_id || season_id === 0) {
      alert("Debes seleccionar un título y una temporada.");
      return;
    }

    const obj: ContentDTO_EpisodeDTO = {
      content_id,
      season_id,
      episodes: list_episodes
    };

    postEpisodes(obj);
  };


  useEffect(() => {
    getSeason()
    getContent()
  }, [getSeason, getContent])

  return (
    <div>
      {_modal_episodes && (
        <div className="container_modal">

          <div className="container-modal-body">

            <div className="container-modal-header">

              <p>AÑADIR CAPITULOS</p>

              <i
                className="bi bi-x-lg"
                onClick={() => OpenModal_Episodes(false)}
              ></i>

            </div>

            <div className="container-modal-input">

              <div className="contenedor_select">

                <select name="content_id" onChange={handleChangeSelect}>
                  <option value="0">SELECIONA UN TITULO</option>
                  {
                    list_content.map((item) => (
                      <option key={item.content_id} value={item.content_id}>{item.content_title}</option>
                    ))
                  }

                </select>

                <div onClick={() => OpenModal_Content(true)}>
                  <i className="bi bi-plus-circle"></i>
                </div>

              </div>

              <div className="contenedor_select">

                <select name="season_id" onChange={handleChangeSelect}>
                  <option value="0">SELECIONA UNA TEMPORADA</option>
                  {
                    season_list.map((item) => (
                      <option key={item.season_id} value={item.season_id}>{item.season_name}</option>
                    ))
                  }

                </select>

                <div onClick={() => OpenModal_Season(true)}>
                  <i className="bi bi-plus-circle"></i>
                </div>

              </div>

              <div className="contenedor_select">

                <input
                  name="episode_name"
                  value={form_episode.episode_name}
                  onChange={handleChangeInput}
                  placeholder="TUTULO CAPITULO" />

                <div onClick={() => add_episodes(form_episode)}>
                  <i className="bi bi-plus-circle"></i>
                </div>

              </div>

              <div className="contenedor_select">

                <input
                  name="episode_url"
                  value={form_episode.episode_url}
                  onChange={handleChangeInput}
                  placeholder="URL CAPITULO" />

                <div style={{ background: "red" }} onClick={() => remove_episodes()}>
                  <i className="bi bi-trash"></i>
                </div>

              </div>

              <div className="container-modal-table">

                <table>

                  <thead>
                    <tr>
                      <th>#</th>
                      <th>CAPITULO</th>
                    </tr>
                  </thead>

                  <tbody>
                    {
                      list_episodes.map((item) => (
                        <tr key={item.episode_id} onClick={() => obtener_episode(item.episode_id)}>
                          <td>{item.episode_id}</td>
                          <td>{item.episode_name}</td>
                        </tr>
                      ))
                    }
                  </tbody>

                </table>

              </div>

              <button onClick={() => sendData()}><i className="bi bi-floppy"></i> GUARDAR</button>

            </div>
          </div>
        </div>
      )}
    </div>
  );
};
