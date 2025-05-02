import { useEffect } from "react";
import { useModal } from "../store/useModal";
import { useData } from "../store/useData";

export const Modal_Content = () => {
  const { _modal_content, OpenModal_Content, OpenModal_Gender } = useModal((state) => state);
  const { gender_list, getGender, year_list, getYear, type_list, getType } = useData((state) => state);


  useEffect(() => {
    getGender()
    getYear()
    getType()
  }, [getGender, getYear, getType])

  return (
    <div>
      {
        _modal_content &&

        <div className="container_modal">
          <div className="container-modal-body">
            <div className="container-modal-header">
              <p>AÑADIR SERIE O ANIME</p>
              <i className="bi bi-x-lg" onClick={() => OpenModal_Content(false)}></i>
            </div>
            <div className="container-modal-input">

              <input
                className="input"
                type="text"
                placeholder="TITULO"
              />

              <select>
                {
                  type_list.map((item) => (
                    <option key={item.content_type_id} value={item.content_type_id}>{item.content_type_title}</option>
                  ))
                }
              </select>

              <input
                className="input"
                type="text"
                placeholder="URL IMAGEN"
              />

              <select>
                {
                  year_list.map((item) => (
                    <option key={item.id_year} value={item.year}>{item.year}</option>
                  ))
                }
              </select>

              <div className="contenedor_select">

                <select>
                  {
                    gender_list.map((item) => (
                      <option key={item.gender_id}>{item.gender_name}</option>
                    ))
                  }
                </select>

                <div onClick={() => OpenModal_Gender(true)}>
                  <i className="bi bi-plus-circle"></i>
                </div>

              </div>

              <button><i className="bi bi-floppy"></i> GUARDAR</button>

            </div>
          </div>
        </div>
      }

    </div>
  )
}
