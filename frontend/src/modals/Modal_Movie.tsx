import "../styles/Styles_Modal.css";
import { useModal } from "../store/useModal";
import { useMovies } from "../store/useMovies";
import { useData } from "../store/useData";
import { Movies } from "../models/Movies";

export const Modal_Movie = () => {
  const { getMovies, reset } = useMovies((state) => state);
  const { _modal_Movie, OpenModal_Movie } = useModal((state) => state);
  const { form_movie, postMovies } = useMovies((state) => state);
  const { gender_list, form_gender } = useData((state) => state);

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
    const { name, value } = e.target;

    useData.setState((state) => ({
      form_gender: {
        ...state.form_gender,
        [name as keyof typeof state.form_gender]: value,
      },
    }));
  };

  function handleCloseModal() {
    reset();
    OpenModal_Movie(false);
  }

  function sendData() {
    const { movie_title, movie_year, movie_url, movie_cover } = form_movie;
    const obj: Movies = {
      movie_movie_id: 0,
      movie_title,
      movie_year: Number(movie_year),
      movie_url,
      movie_cover,
      gender_id: Number(form_gender.gender_id),
    };
    postMovies(obj);
    handleCloseModal();
    getMovies();
  }

  return (
    <div>
      {_modal_Movie && (
        <div className="container_modal">
          <div className="container-modal-body">
            <div className="container-modal-header">
              <p>ADD MOVIE</p>
              <i className="bi bi-x-lg" onClick={() => handleCloseModal()}></i>
            </div>
            <div className="container-modal-input">
              <input
                className="input"
                type="text"
                placeholder="TITLE"
                name="movie_title"
                value={form_movie.movie_title}
                onChange={handleChangeInput}
              />
              <input
                className="input"
                type="number"
                name="movie_year"
                placeholder="YEAR"
                value={form_movie.movie_year}
                onChange={handleChangeInput}
              />
              <input
                className="input"
                type="text"
                placeholder="COVER"
                name="movie_cover"
                value={form_movie.movie_cover}
                onChange={handleChangeInput}
              />
              <input
                className="input"
                type="text"
                placeholder="URL"
                value={form_movie.movie_url}
                name="movie_url"
                onChange={handleChangeInput}
              />
              <select
                name="gender_id"
                onChange={handleChangeSelect}
                value={form_gender.gender_id}
              >
                {gender_list.map((item) => (
                  <option key={item.gender_id} value={item.gender_id}>
                    {item.gender_name}
                  </option>
                ))}
              </select>
              <button onClick={() => sendData()}>ADD</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
