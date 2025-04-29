import { _season } from "../models/Seasons";
import { useContent } from "../store/useContent";
import { useData } from "../store/useData";
import { useModal } from "../store/useModal";
import "../styles/Styles_Modal.css";

export const Modal_Episodes = () => {
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

  const SendData = () => { };

  return (
    <div>
      {_modal_Content && (
        <div className="container_modal">

          <div className="container-modal-body">

            <div className="container-modal-header">

              <p>AÑADIR CAPITULOS</p>

              <i
                className="bi bi-x-lg"
                onClick={() => OpenModal_Content(false)}
              ></i>

            </div>

            <div className="container-modal-input">

              <div className="contenedor_select">

                <select>
                  <option>LOS 100</option>
                </select>

                <div>
                  <i className="bi bi-plus-circle"></i>
                </div>

              </div>

              <div className="contenedor_select">

                <select>
                  <option>TEMPORADA 1</option>
                </select>

                <div>
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
