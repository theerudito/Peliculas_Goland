import "../styles/App.css";
import { Component_Viewer } from "./Component_Viewer";

export const Component_Home = () => {
  return (
    <div className="app-container">
      <div className="main-content">
        <Component_Viewer />
      </div>
    </div>
  );
};
