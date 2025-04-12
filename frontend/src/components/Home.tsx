import { useEffect, useState } from "react";
import "../styles/Home.css";
import { GET_Movies } from "../helpers/Fetching";
import { Movies } from "../models/Movies";

export const Home = () => {
  const [value, setValue] = useState<Movies[]>([]);

  const getMovies = async () => {
    const obj = await GET_Movies();

    setValue(obj);
  };

  useEffect(() => {
    getMovies();
  }, []);

  return (
    <div className="container-home">
      <div className="container-home-body">
        {value.map((item) => (
          <div key={item.movie_id} className="container-home-card">
            <img
              src={item.cover}
              alt={item.title}
              className="card-background-image"
            />

            <div className="card-overlay">
              <p className="card-year">{item.year}</p>
              <div className="card-play">
                <i className="bi bi-play-circle"></i>
              </div>
              <p className="card-title">{item.title}</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};
