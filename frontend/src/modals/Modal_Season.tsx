export const Modal_Season = () => {
  return (
    <div>
      <div className="container_modal">
        <div className="container-modal-body">
          <div className="container-modal-header">
            <p>AÑADIR TEMPORADAS</p>
            <i className="bi bi-x-lg" ></i>
          </div>
          <div className="container-modal-input">
            <input placeholder='TUTULO DE TEMPORADA' />
            <button ><i className="bi bi-floppy"></i> GUARDAR</button>
          </div>
        </div>
      </div>
    </div>
  )
}
