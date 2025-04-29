export const Modal_Content = () => {
  return (
    <div>
      <div className="container_modal">
        <div className="container-modal-body">
          <div className="container-modal-header">
            <p>AÑADIR SERIE O ANIME</p>
            <i className="bi bi-x-lg"></i>
          </div>
          <div className="container-modal-input">

            <input
              className="input"
              type="text"
              placeholder="TITULO"
            />

            <select>
              <option value={1}>ANIME</option>
              <option value={2}>SERIE</option>
            </select>

            <input
              className="input"
              type="text"
              placeholder="URL IMAGEN"
            />

            <input
              className="input"
              type="number"
              placeholder="AÑO"
            />

            <div className="contenedor_select">

              <select>
                <option>DRAMA</option>
              </select>

              <div>
                <i className="bi bi-plus-circle"></i>
              </div>

            </div>

            <button><i className="bi bi-floppy"></i> GUARDAR</button>

          </div>
        </div>
      </div>
    </div>
  )
}
