import { useDispatch } from "react-redux";
import { Modal_Auth } from "../modals/Modal_Auth";
import { Modal_Content } from "../modals/Modal_Content";
import { Modal_Movie } from "../modals/Modal_Movie";
import "../styles/App.css";
import { Component_Footer } from "./Component_Footer";
import { Component_Header } from "./Component_Header";
import { openModal } from "../store/slice/ModalSlices";

export const Component_Home = () => {
  const dispatch = useDispatch();
  return (
    <div className="app-container">
      <Component_Header />
      <div className="main-content">
        <Modal_Content />
        <Modal_Movie />
        <Modal_Auth />
        <div style={{ background: "red" }}>
          <button onClick={() => dispatch(openModal(2))}>ADD MOVIE</button>
          <button onClick={() => dispatch(openModal(3))}>ADD CONTENT</button>
        </div>
      </div>

      <Component_Footer />
    </div>
  );
};
