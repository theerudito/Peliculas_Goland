import "../styles/App.css";


export const Component_Home = () => {
  return (
    <div className="app-container">

      <div className="container-header-search-box">
        <input type="text" placeholder="Buscar..." />
        <button className="container-header-search-btn">
          <i className="bi bi-search"></i>
        </button>
      </div>

      <div className="main-content">

      </div>
    </div>
  );
};
