import { useState } from "react";
import "../styles/Modal.css";

export const Modal_Content = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [formData, setFormData] = useState({
    title: "",
    year: "",
    cover: "",
    url: "",
    gender: "DRAMA",
  });

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  return (
    <div>
      <button className="open-modal-btn" onClick={() => setIsOpen(true)}>
        Abrir Modal
      </button>

      {isOpen && (
        <div className="container_modal">
          <div className="container-modal-body">
            <div className="container-modal-header">
              <p>ADD ANIME OR SERIE</p>
              <i className="bi bi-x-lg" onClick={() => setIsOpen(false)}></i>
            </div>
            <div className="container-modal-input">
              <input type="text" placeholder="TITLE EPISODE" />
              <input type="text" placeholder="URL EPISODE" />

              <button>1 ADD EPISODE</button>
              <div className="container-modal-table">
                <table>
                  <thead>
                    <tr>
                      <th>TITLE</th>
                      <th>EPISODE</th>
                      <th>A</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>JUJUTSUN KAISEN</td>
                      <td>EPISODE 1</td>
                      <td>
                        {" "}
                        <i className="bi bi-trash"></i>
                      </td>
                    </tr>
                    <tr>
                      <td>JUJUTSUN KAISEN</td>
                      <td>EPISODE 2</td>
                      <td>
                        <i className="bi bi-trash"></i>
                      </td>
                    </tr>
                    <tr>
                      <td>JUJUTSUN KAISEN</td>
                      <td>EPISODE 2</td>
                      <td>
                        <i className="bi bi-trash"></i>
                      </td>
                    </tr>
                    <tr>
                      <td>JUJUTSUN KAISEN</td>
                      <td>EPISODE 2</td>
                      <td>
                        <i className="bi bi-trash"></i>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>

              <input type="text" placeholder="TITLE SEASON" />
              <select>
                <option>SEASON 1</option>
                <option>SEASON 2</option>
                <option>SEASON 3</option>
              </select>

              <input type="text" placeholder="TITLE CONTENT" />
              <input type="text" placeholder="DESCRIPCION" />
              <input type="text" placeholder="COVER" />
              <input type="date" placeholder="YEAR" />

              <select>
                <option>DRAMA</option>
                <option>CRIMEN</option>
                <option>ACCION</option>
              </select>
              <button> 2 ADD CONTENT</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
