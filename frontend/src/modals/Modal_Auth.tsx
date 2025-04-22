import { useState } from "react";
import "../styles/Styles_Modal.css";
import { _users, Users } from "../models/Auth";
import { useModal } from "../store/useModal";

export const Modal_Auth = () => {
  const [formData, setFormData] = useState<Users>(_users);
  const { _modal_Auth, OpenModal_Auth } = useModal((state) => state);

  const handleChange = (e: { target: { name: string; value: string } }) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
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
      {_modal_Auth && (
        <div className="container_modal">
          <div className="container-modal-body">
            <div className="container-modal-header">
              <p>LOGIN USER</p>
              <i
                className="bi bi-x-lg"
                onClick={() => OpenModal_Auth(false)}
              ></i>
            </div>
            <div className="container-modal-input">
              <input
                className="input"
                type="text"
                name="user"
                placeholder="USER"
                value={formData.user}
                onChange={handleChange}
              />
              <input
                className="input"
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
