import { useState } from "react";
import "../styles/Modal.css";

export const Modal_Movie = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [formData, setFormData] = useState({
    title: "",
    year: "",
    cover: "",
    url: "",
    gender: "DRAMA",
  });

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
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
              <input type="text" placeholder="TITLE" />
              <input type="date" placeholder="YEAR" />
              <input type="text" placeholder="COVER" />
              <input type="text" placeholder="URL" />
              <select>
                <option>DRAMA</option>
                <option>CRIMEN</option>
              </select>
              <button>ADD</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
