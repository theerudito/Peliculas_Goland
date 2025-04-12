import { Modal_Content } from "../modals/Modal_Content";
import "../styles/App.css";
import { Component_Content } from "./Component_Content";
import { Component_Footer } from "./Component_Footer";
import { Component_Header } from "./Component_Header";

export const Component_Home = () => {
  return (
    <div className="app-container">
      <Component_Header />

      <div className="main-content">
        {/* <Component_Content /> */}
        <Modal_Content />
      </div>

      <Component_Footer />
    </div>
  );
};
