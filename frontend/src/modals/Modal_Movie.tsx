import { useState } from "react";
import "../styles/Modal.css";
import { Gender_List } from "../helpers/Data";
import { _movies, Movies } from "../models/Movies";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../store/store";
import { closeModal } from "../store/slice/ModalSlices";

export const Modal_Movie = () => {
  const { isOpen } = useSelector((store: RootState) => store.modal);
  const dispatch = useDispatch();

  const [formData, setFormData] = useState<Movies>(_movies);

  const handleChangeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    setFormData((prev) => ({
      ...prev,
      [name]: name === "year" ? Number(value) : value,
    }));
  };

  const handleChangeSelect = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const { name, value } = e.target;

    setFormData((prev) => ({
      ...prev,
      [name]: parseInt(value, 10),
    }));
  };
  const SendData = () => {
    console.log(formData);
    Reset_Field();
  };

  function Reset_Field() {
    setFormData(formData);
  }

  return (
    <div>
      {isOpen && (
        <div className="container_modal">
          <div className="container-modal-body">
            <div className="container-modal-header">
              <p>ADD MOVIE</p>
              <i
                className="bi bi-x-lg"
                onClick={() => dispatch(closeModal(2))}
              ></i>
            </div>
            <div className="container-modal-input">
              <input
                className="input"
                type="text"
                placeholder="TITLE"
                name="title"
                value={formData.title}
                onChange={handleChangeInput}
              />
              <input
                className="input"
                type="number"
                name="year"
                placeholder="YEAR"
                value={formData.year}
                onChange={handleChangeInput}
              />
              <input
                className="input"
                type="text"
                placeholder="COVER"
                name="cover"
                value={formData.cover}
                onChange={handleChangeInput}
              />
              <input
                className="input"
                type="text"
                placeholder="URL"
                value={formData.url}
                name="url"
                onChange={handleChangeInput}
              />
              <select
                name="gender_id"
                onChange={handleChangeSelect}
                value={formData.gender_id}
              >
                {Gender_List.map((item) => (
                  <option key={item.gender_id} value={item.gender_id}>
                    {item.name}
                  </option>
                ))}
              </select>
              <button onClick={SendData}>ADD</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
