import { useEffect } from "react";
import { useContent } from "../store/useContent";
import { useData } from "../store/useData";
import { useModal } from "../store/useModal";
import "../styles/Styles_Modal.css";

export const Modal_Episodes = () => {
  const { list_content, form_content_episode, getContent } =
    useContent((state) => state);
  const { _modal_episodes, OpenModal_Episodes, OpenModal_Season, OpenModal_Content } = useModal((state) => state);
  const { season_list, getSeason } = useData((state) => state);



  const handleChangeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    useContent.setState((state) => ({
      form_content_episode: {
        ...state.form_content_episode,
        [name as keyof typeof state.form_content_episode]: value,
      },
    }));
  };

  const handleChangeSelect = (e: React.ChangeEvent<HTMLSelectElement>) => {

    const { name, value } = e.target;

    useContent.setState((state) => {
      let selectedValue: string | number = value;

      if (name === "content_id" || name === "season_id") {
        selectedValue = Number(value);
      }

      return {
        form_content_episode: {
          ...state.form_content_episode,
          [name as keyof typeof state.form_content_episode]: selectedValue,
        },
      };
    });
  };
  const sendData = () => {
    console.log(form_content_episode)
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
                  value={form_content_episode.episode_name}
                  onChange={handleChangeInput}
                  placeholder="TUTULO CAPITULO" />

                <div>
                  <i className="bi bi-plus-circle"></i>
                </div>

              </div>

              <div className="contenedor_select">

                <input
                  name="episode_url"
                  value={form_content_episode.episode_url}
                  onChange={handleChangeInput}
                  placeholder="URL CAPITULO" />

                <div style={{ background: "red" }}>
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
