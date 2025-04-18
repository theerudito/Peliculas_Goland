import "../styles/Modal.css";
import { useModal } from "../store/useModal";
import { useMovies } from "../store/useMovies";
import { useData } from "../store/useData";

export const Modal_Movie = () => {
  const { _modal_Movie, OpenModal_Movie } = useModal((state) => state);
  const { form_movie } = useMovies((state) => state);
  const { gender_list, form_gender } = useData((state) => state);
  const handleChangeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    // const { name, value } = e.target;
    // const updatedForm = {
    //   ...form_Movies,
    //   [name]: value,
    // };
  };

  const handleChangeSelect = (e: React.ChangeEvent<HTMLSelectElement>) => {
    // const { name, value } = e.target;
    // const updatedForm = {
    //   ...form_Movies,
    //   [name]: parseInt(value, 10),
    // };
  };

  const SendData = () => { };

  return (
    <div>
      {_modal_Movie && (
        <div className="container_modal">
          <div className="container-modal-body">
            <div className="container-modal-header">
              <p>ADD MOVIE</p>
              <i
                className="bi bi-x-lg"
                onClick={() => OpenModal_Movie(false)}
              ></i>
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
              <button onClick={SendData}>ADD</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
