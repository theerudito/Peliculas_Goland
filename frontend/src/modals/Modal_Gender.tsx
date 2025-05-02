import { useData } from "../store/useData";
import { useModal } from "../store/useModal";

export const Modal_Gender = () => {
  const { _modal_gender, OpenModal_Gender } = useModal((state) => state);
  const { form_gender, postGender } = useData((state) => state);


  const handleChangeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    useData.setState((state) => ({
      form_gender: {
        ...state.form_gender,
        [name]: value,
      },
    }));
  };

  const sendData = async () => {
    await postGender(form_gender)
  }

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
                <input
                  name="gender_name"
                  value={form_gender.gender_name}
                  onChange={handleChangeInput}
                  placeholder='TUTULO DE GENERO' />
                <button onClick={() => sendData()}><i className="bi bi-floppy"></i> GUARDAR</button>
              </div>
            </div>
          </div>
        )
      }

    </div>
  )
}
