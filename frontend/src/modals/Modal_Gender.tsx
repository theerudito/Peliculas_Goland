import { useModal } from "../store/useModal";

export const Modal_Gender = () => {
  const { _modal_gender, OpenModal_Gender } = useModal((state) => state);

  return (
    <div>
      {
        _modal_gender && (
          <div className="container_modal">
            <div className="container-modal-body">
              <div className="container-modal-header">
                <p>AÑADIR GENERO</p>
                <i className="bi bi-x-lg" onClick={() => OpenModal_Gender(false)}></i>
              </div>
              <div className="container-modal-input">
                <input placeholder='TUTULO DE GENERO' />
                <button ><i className="bi bi-floppy"></i> GUARDAR</button>
              </div>
            </div>
          </div>
        )
      }

    </div>
  )
}
