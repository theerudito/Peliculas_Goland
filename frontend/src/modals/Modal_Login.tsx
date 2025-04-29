export const Modal_Login = () => {
  return (
    <div>
      <div className="container_modal">
        <div className="container-modal-body">
          <div className="container-modal-header">
            <p>LOGIN</p>
            <i className="bi bi-x-lg" ></i>
          </div>
          <div className="container-modal-input">
            <input placeholder='USARIO' />
            <input placeholder='CONTRASEÑA' />
            <button ><i className="bi bi-floppy"></i> GUARDAR</button>
          </div>
        </div>
      </div>
    </div>
  )
}
