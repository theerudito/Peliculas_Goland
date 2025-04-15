import { useDispatch } from "react-redux";
import { openModal } from "../store/slice/Modal_Slices";
import "../styles/Menu.css";

export const Component_Menu = () => {
  const dispatch = useDispatch();
  return (
    <>
      <div className="container_Menu">
        <div className="menu_buttons">
          <i className="bi bi-x-lg"></i>
          <button onClick={() => dispatch(openModal(3))}>ADD CONTENT</button>
          <button onClick={() => dispatch(openModal(2))}>ADD MOVIE</button>
          <button>CLOSE</button>
        </div>
      </div>
    </>
  );
};
