import { useModal } from "../store/useModal";

export const Modal_Login = () => {
  const { _modal_login, OpenModal_Login } = useModal((state) => state);
  return (
    <div>
      {
        _modal_login &&
        <div className="container_modal">
          <div className="container-modal-body">
            <div className="container-modal-header">
              <p>LOGIN</p>
              <i className="bi bi-x-lg" onClick={() => OpenModal_Login(false)}></i>
            </div>
            <div className="container-modal-input">
              <input placeholder='USARIO' />
              <input placeholder='CONTRASEÑA' />
              <button ><i className="bi bi-floppy"></i> INICIAR SECCION</button>
            </div>
          </div>
        </div>

      }

    </div>
  )
}
