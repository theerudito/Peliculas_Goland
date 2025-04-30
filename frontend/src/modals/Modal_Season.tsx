import { useModal } from "../store/useModal";

export const Modal_Season = () => {
  const { _modal_season, OpenModal_Season } = useModal((state) => state);
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
                <input placeholder='TUTULO DE TEMPORADA' />
                <button ><i className="bi bi-floppy"></i> GUARDAR</button>
              </div>
            </div>
          </div>
        )
      }

    </div>
  )
}
