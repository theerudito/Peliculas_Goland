import ReactPlayer from "react-player";
import "../styles/Styles_Player.css";
import { usePlayer } from "../store/usePlayer";
import { useEffect } from "react";


export const Component_Player = () => {
  const { close_player, url } = usePlayer((state) => state);

  useEffect(() => {
    const interval = setInterval(() => {
      const video = document.querySelector("video");
      if (video) {
        video.setAttribute("controlsList", "nodownload");
        clearInterval(interval);
      }
    }, 100);

    return () => clearInterval(interval);
  }, []);

  if (url === "") return null;



  return (
    <div className="container-player-body">
      <button className="close-button" onClick={close_player}>
        <i className="bi bi-x-lg"></i>
      </button>

      <ReactPlayer
        url={url}
        playing={true}
        controls={true}
        width="100%"
        height="100%"
        className="react-player"
      />
    </div>

  );
};
