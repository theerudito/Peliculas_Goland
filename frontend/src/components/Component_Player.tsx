import ReactPlayer from "react-player";
import "../styles/Player.css";

export const Component_Player = ({ url }: { url: string }) => {
  if (url == "") return;
  return (
    <div className="container-player-body">
      <div>
        <button className="close-button">
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
