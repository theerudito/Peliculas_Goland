import { useState } from "react";
import "../styles/Modal.css";
import { Gender_List } from "../helpers/Data";
import { _movies, Movies } from "../models/Movies";

export const Modal_Movie = () => {
  const [isOpen, setIsOpen] = useState(false);
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
  };

  return (
    <div>
      <button className="open-modal-btn" onClick={() => setIsOpen(true)}>
        Abrir Modal
      </button>

      {isOpen && (
        <div className="container_modal">
          <div className="container-modal-body">
            <div className="container-modal-header">
              <p>ADD MOVIE</p>
              <i className="bi bi-x-lg" onClick={() => setIsOpen(false)}></i>
            </div>
            <div className="container-modal-input">
              <input
                type="text"
                placeholder="TITLE"
                name="title"
                value={formData.title}
                onChange={handleChangeInput}
              />
              <input
                type="number"
                name="year"
                placeholder="YEAR"
                value={formData.year}
                onChange={handleChangeInput}
              />
              <input
                type="text"
                placeholder="COVER"
                name="cover"
                value={formData.cover}
                onChange={handleChangeInput}
              />
              <input
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
