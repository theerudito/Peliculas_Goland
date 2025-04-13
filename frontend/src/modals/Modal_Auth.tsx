import { useState } from "react";
import { _user, Users } from "../models/Movies";
import "../styles/Modal.css";

export const Modal_Auth = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [formData, setFormData] = useState<Users>(_user);

  const handleChange = (e: { target: { name: string; value: string } }) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
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
              <p>LOGIN USER</p>
              <i className="bi bi-x-lg" onClick={() => setIsOpen(false)}></i>
            </div>
            <div className="container-modal-input">
              <input
                type="text"
                name="user"
                placeholder="USER"
                value={formData.user}
                onChange={handleChange}
              />
              <input
                type="password"
                name="password"
                placeholder="PASSWORD"
                value={formData.password}
                onChange={handleChange}
              />
              <button onClick={SendData}>LOGIN</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
