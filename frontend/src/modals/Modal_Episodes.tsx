import { useEffect } from "react";
import { useContent } from "../store/useContent";
import { useData } from "../store/useData";
import { useModal } from "../store/useModal";
import "../styles/Styles_Modal.css";

export const Modal_Episodes = () => {
  const { list_content, getContent } =
    useContent((state) => state);
  const { _modal_episodes, OpenModal_Episodes, OpenModal_Season, OpenModal_Content } = useModal((state) => state);
  const { season_list, getSeason } = useData((state) => state);



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

  const SendData = () => { };


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

                <select>
                  {
                    list_content.map((item) => (
                      <option key={item.content_id}>{item.content_title}</option>
                    ))
                  }

                </select>

                <div onClick={() => OpenModal_Content(true)}>
                  <i className="bi bi-plus-circle"></i>
                </div>

              </div>

              <div className="contenedor_select">

                <select>
                  {
                    season_list.map((item) => (
                      <option key={item.season_id}>{item.season_name}</option>
                    ))
                  }

                </select>

                <div onClick={() => OpenModal_Season(true)}>
                  <i className="bi bi-plus-circle"></i>
                </div>

              </div>

              <div className="contenedor_select">

                <input placeholder="TUTULO CAPITULO" />

                <div>
                  <i className="bi bi-plus-circle"></i>
                </div>

              </div>

              <div className="contenedor_select">

                <input placeholder="URL CAPITULO" />

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

              <button><i className="bi bi-floppy"></i> GUARDAR</button>

            </div>
          </div>
        </div>
      )}
    </div>
  );
};
