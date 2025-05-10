import ReactPlayer from "react-player";
import "../styles/Styles_Player.css";
import { usePlayer } from "../store/usePlayer";

export const Component_Player = () => {
  const { close_player, url } = usePlayer((state) => state);

  if (url === "") return;

  return (
    <div className="container-player-body">
      <div>
        <button className="close-button" onClick={close_player}>
          <i className="bi bi-x-lg"></i>
        </button>
      </div>
      <ReactPlayer
        url={url}
        playing={true}
        controls={true}
        light={true}
        className="ReactPlayer"
      />
    </div>
  );
};
