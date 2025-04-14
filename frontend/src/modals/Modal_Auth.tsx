import { useState } from "react";
import { _users, Users } from "../models/Movies";
import "../styles/Modal.css";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../store/store";
import { closeModal } from "../store/slice/ModalSlices";

export const Modal_Auth = () => {
  const { isOpen } = useSelector((store: RootState) => store.modal);
  const dispatch = useDispatch();

  const [formData, setFormData] = useState<Users>(_users);

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
      {isOpen && (
        <div className="container_modal">
          <div className="container-modal-body">
            <div className="container-modal-header">
              <p>LOGIN USER</p>
              <i
                className="bi bi-x-lg"
                onClick={() => dispatch(closeModal(1))}
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
