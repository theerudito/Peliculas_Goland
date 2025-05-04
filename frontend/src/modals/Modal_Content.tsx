import { useEffect } from "react";
import { useModal } from "../store/useModal";
import { useData } from "../store/useData";
import { useContent } from "../store/useContent";
import { Content } from "../models/Contents";

export const Modal_Content = () => {
  const { _modal_content, OpenModal_Content, OpenModal_Gender } = useModal((state) => state);
  const { gender_list, getGender, year_list, getYear, type_list, getType } = useData((state) => state);
  const { form_content, postContent } = useContent((state) => state);

  const handleChangeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    useContent.setState((state) => ({
      form_content: {
        ...state.form_content,
        [name as keyof typeof state.form_content]: value,
      },
    }));
  };

  const handleChangeSelect = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const { name, value, selectedIndex, options } = e.target;

    useContent.setState((state) => {
      let selectedValue: string | number = value;

      if (name === "content_year") {
        selectedValue = Number(options[selectedIndex].text);
      }

      if (name === "gender_id" || name === "content_type") {
        selectedValue = Number(value);
      }

      return {
        form_content: {
          ...state.form_content,
          [name as keyof typeof state.form_content]: selectedValue,
        },
      };
    });
  };

  useEffect(() => {
    getGender()
    getYear()
    getType()
  }, [getGender, getYear, getType])


  const sendData = () => {

    if (form_content.content_type === 0 || form_content.gender_id === 0) {
      alert("Debes seleccionar un genero y tipo de contenido");
      return;
    }

    const { content_title, content_cover, content_type, content_year, gender_id } = form_content

    const currentYear = new Date().getFullYear();

    const obj: Content = {
      content_id: 0,
      content_title,
      content_cover,
      content_type: Number(content_type),
      content_year: content_year === 0 ? currentYear : content_year,
      gender_id: gender_id === 0 ? 1 : gender_id
    };

    postContent(obj);
  }

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
                name="content_title"
                value={form_content.content_title}
                onChange={handleChangeInput}
                placeholder="TITULO"
              />

              <select name="content_type" onChange={handleChangeSelect}>
                <option value="0">SELECIONA UN TIPO CONTENIDO</option>
                {
                  type_list.map((item) => (
                    <option key={item.content_type} value={item.content_type}>{item.content_type_title}</option>
                  ))
                }
              </select>

              <input
                className="input"
                type="text"
                name="content_cover"
                value={form_content.content_cover}
                onChange={handleChangeInput}
                placeholder="URL IMAGEN"
              />

              <select name="content_year" onChange={handleChangeSelect}>
                <option value="0">SELECIONA UN AÑO</option>
                {
                  year_list.map((item) => (
                    <option key={item.year_id} value={item.year}>{item.year}</option>
                  ))
                }
              </select>

              <div className="contenedor_select">

                <select name="gender_id" onChange={handleChangeSelect}>
                  <option value="0">SELECIONA UN GENERO</option>
                  {
                    gender_list.map((item) => (
                      <option key={item.gender_id} value={item.gender_id}>{item.gender_name}</option>
                    ))
                  }
                </select>

                <div onClick={() => OpenModal_Gender(true)}>
                  <i className="bi bi-plus-circle"></i>
                </div>

              </div>

              <button onClick={() => sendData()}><i className="bi bi-floppy"></i> GUARDAR</button>

            </div>
          </div>
        </div>
      }

    </div>
  )
}
