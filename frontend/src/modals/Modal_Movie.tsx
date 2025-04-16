import "../styles/Modal.css";
import { Gender_List } from "../helpers/Data";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../store/store";
import { closeModal } from "../store/slice/Modal_Slices";
import { addMovies } from "../store/slice/Movies_Slices";

export const Modal_Movie = () => {
  const { openModal_Movie } = useSelector((store: RootState) => store.modal);
  const { form_Movies } = useSelector((store: RootState) => store.movie);

  const dispatch = useDispatch();

  const handleChangeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    const updatedForm = {
      ...form_Movies,
      [name]: value,
    };

    dispatch(addMovies(updatedForm));
  };

  const handleChangeSelect = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const { name, value } = e.target;
    const updatedForm = {
      ...form_Movies,
      [name]: parseInt(value, 10),
    };

    dispatch(addMovies(updatedForm));
  };

  const SendData = () => {
    console.log(form_Movies);
    Reset_Field();
  };

  function Reset_Field() {
    // setFormData(formData);
  }

  return (
    <div>
      {openModal_Movie && (
        <div className="container_modal">
          <div className="container-modal-body">
            <div className="container-modal-header">
              <p>ADD MOVIE</p>
              <i
                className="bi bi-x-lg"
                onClick={() => dispatch(closeModal(2))}
              ></i>
            </div>
            <div className="container-modal-input">
              <input
                className="input"
                type="text"
                placeholder="TITLE"
                name="movie_title"
                value={form_Movies.movie_title}
                onChange={handleChangeInput}
              />
              <input
                className="input"
                type="number"
                name="movie_year"
                placeholder="YEAR"
                value={form_Movies.movie_year}
                onChange={handleChangeInput}
              />
              <input
                className="input"
                type="text"
                placeholder="COVER"
                name="movie_cover"
                value={form_Movies.movie_cover}
                onChange={handleChangeInput}
              />
              <input
                className="input"
                type="text"
                placeholder="URL"
                value={form_Movies.movie_url}
                name="movie_url"
                onChange={handleChangeInput}
              />
              <select
                name="gender_id"
                onChange={handleChangeSelect}
                value={form_Movies.gender_id}
              >
                {Gender_List.map((item) => (
                  <option key={item.gender_id} value={item.gender_id}>
                    {item.name}
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
