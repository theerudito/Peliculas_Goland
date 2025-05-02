import "../styles/Styles_Modal.css";
import { useModal } from "../store/useModal";
import { useMovies } from "../store/useMovies";
import { useData } from "../store/useData";
import { Movies } from "../models/Movies";
import { useEffect } from "react";

export const Modal_Movie = () => {
  const { getMovies, reset } = useMovies((state) => state);
  const { _modal_movie, OpenModal_Movie, OpenModal_Gender } = useModal((state) => state);
  const { form_movie, postMovies } = useMovies((state) => state);
  const { gender_list, getGender, getYear, year_list } = useData((state) => state);


  const handleChangeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    useMovies.setState((state) => ({
      form_movie: {
        ...state.form_movie,
        [name as keyof typeof state.form_movie]: value,
      },
    }));
  };

  const handleChangeSelect = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const { name, value, selectedIndex, options } = e.target;

    useMovies.setState((state) => {
      let selectedValue: string | number = value;

      if (name === "year") {
        selectedValue = Number(options[selectedIndex].text);
      }

      if (name === "gender_id") {
        selectedValue = Number(value);
      }

      return {
        form_movie: {
          ...state.form_movie,
          [name as keyof typeof state.form_movie]: selectedValue,
        },
      };
    });
  };

  function sendData() {

    const { movie_title, movie_year, movie_url, movie_cover, gender_id } = form_movie;

    const currentYear = new Date().getFullYear();

    const obj: Movies = {
      movie_movie_id: 0,
      movie_title,
      movie_year: movie_year === 0 ? currentYear : movie_year,
      movie_url,
      movie_cover,
      gender_id: gender_id === 0 ? 1 : gender_id
    };

    postMovies(obj);
    getMovies();
    handleCloseModal();
  }

  useEffect(() => {
    getGender()
    getYear()
  }, [getGender, getYear])


  function handleCloseModal() {
    reset();
    OpenModal_Movie(false);
  }

  return (
    <div>
      {_modal_movie && (
        <div className="container_modal">
          <div className="container-modal-body">

            <div className="container-modal-header">
              <p>AÑADIR PELICULAS</p>
              <i className="bi bi-x-lg" onClick={() => handleCloseModal()}></i>
            </div>

            <div className="container-modal-input">

              <input
                className="input"
                type="text"
                placeholder="TITULO"
                name="movie_title"
                value={form_movie.movie_title}
                onChange={handleChangeInput}
              />

              <select name="movie_year" onChange={handleChangeSelect}>
                {
                  year_list.map((item) => (
                    <option key={item.id_year} value={item.year}>{item.year}</option>
                  ))
                }

              </select>

              <input
                className="input"
                type="text"
                placeholder="URL IMAGEN"
                name="movie_cover"
                value={form_movie.movie_cover}
                onChange={handleChangeInput}
              />

              <input
                className="input"
                type="text"
                placeholder="URL PELICULA"
                value={form_movie.movie_url}
                name="movie_url"
                onChange={handleChangeInput}
              />

              <div className="contenedor_select">

                <select name="gender_id" onChange={handleChangeSelect}>
                  {gender_list.map((item) => (
                    <option key={item.gender_id} value={item.gender_id}>
                      {item.gender_name}
                    </option>
                  ))}
                </select>

                <div onClick={() => OpenModal_Gender(true)}>
                  <i className="bi bi-plus-circle"></i>
                </div>

              </div>

              <button onClick={() => sendData()}><i className="bi bi-floppy"></i> GUARDAR</button>

            </div>
          </div>
        </div>
      )}
    </div>
  );
};
