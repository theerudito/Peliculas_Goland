import "../styles/Header.css";

export const Component_Header = () => {
  return (
    <div className="header">
      <nav className="nav-links">
        <a href="#">Inicio</a>
        <a href="#">Peliculas</a>
        <a href="#">Series</a>
        <a href="#">Anime</a>
      </nav>
      <div className="search-box">
        <input type="text" placeholder="Buscar..." />
        <button className="search-btn">
          <i className="bi bi-search"></i>
        </button>
      </div>
      <button className="login-btn">LOGIN</button>
    </div>
  );
};
