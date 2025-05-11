import { useEffect } from "react";
import "../styles/Styles_Serie_Anime_Movies.css";
import { useContent } from "../store/useContent";
import cover from "../assets/logo.webp";
import { useAuth } from "../store/useAuth";
import { Component_Content } from "./Component_Content";

export const Component_Serie_Anime = () => {
  const { list_content_type, getContent_Type, openContent, openContent_Type } = useContent((state) => state);
  const { isLogin } = useAuth((state) => state);

  useEffect(() => {
    getContent_Type();
  }, [getContent_Type]);

  return (
    <div className="app-container">

      {openContent === false ? (

        <div>
          <div className="container-header-search-box">
            <input type="text" placeholder="Buscar..." />
            <button className="container-header-search-btn">
              <i className="bi bi-search"></i>
            </button>
          </div>

          <div className="main-content">
            <div className="container-content">
              <div className="container-body">


                {list_content_type.map((item) => (
                  <div key={item.content_id} className="container-card">
                    <img
                      src={
                        item.content_cover === "" ? cover : item.content_cover
                      }
                      alt={cover}
                      className="card-background-image"
                    />

                    <div className="card-overlay">
                      <p className="card-year">{item.content_year}</p>
                      <div className="card-play">
                        <i
                          className="bi bi-play-circle"
                          onClick={() => openContent_Type(true, item.content_id)}
                        ></i>
                      </div>
                      <div className="card-container-button">
                        {
                          isLogin && (
                            <>
                              <i className="bi bi-pencil-square"></i>
                              <i className="bi bi-trash"></i>
                            </>
                          )
                        }

                      </div>
                    </div>
                  </div>
                ))}


              </div>
            </div>
          </div>

        </div>
      ) : (<Component_Content />)
      }
    </div >
  );
};
