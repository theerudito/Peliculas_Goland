import { useData } from "../store/useData";
import { useModal } from "../store/useModal";

export const Modal_Season = () => {
  const { _modal_season, OpenModal_Season } = useModal((state) => state);
  const { form_season, postSeason } = useData((state) => state);

  const handleChangeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    useData.setState((state) => ({
      form_season: {
        ...state.form_season,
        [name]: value,
      },
    }));
  };

  const sendData = async () => {

    if (!form_season.season_name) {
      alert("Debes añadir un titulo de temporada");
      return;
    }

    await postSeason(form_season)
  }

  return (
    <div>
      {
        _modal_season && (
          <div className="container_modal">
            <div className="container-modal-body">
              <div className="container-modal-header">
                <p>AÑADIR TEMPORADAS</p>
                <i className="bi bi-x-lg" onClick={() => OpenModal_Season(false)}></i>
              </div>
              <div className="container-modal-input">
                <input
                  name="season_name"
                  value={form_season.season_name}
                  onChange={handleChangeInput}
                  placeholder='TUTULO DE TEMPORADA' />
                <button onClick={() => sendData()}><i className="bi bi-floppy"></i> GUARDAR</button>
              </div>
            </div>
          </div>
        )
      }

    </div>
  )
}
