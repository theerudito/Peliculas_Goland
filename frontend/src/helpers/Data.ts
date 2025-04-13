import image_1 from "../assets/imagen1.jpg";
import image_2 from "../assets/imagen3.avif";
import image_3 from "../assets/imagen2.jpg";

export const Movies_List = [
  {
    movie_id: 1,
    title: "MAN OF HONOR",
    year: 2000,
    cover: image_1,
    url: "https://pixeldrain.com/api/file/kyjhvCUN",
    gender: "CRIMEN",
  },
  {
    movie_id: 2,
    title: "CADENA PERPECTUA",
    year: 1994,
    cover: image_2,
    url: "https://pixeldrain.com/api/file/CkpdetaB",
    gender: "DRAMA",
  },
  {
    movie_id: 3,
    title: "THE GLADIATOR",
    year: 1994,
    cover: image_3,
    url: "https://pixeldrain.com/api/file/kyjhvCUN",
    gender: "DRAMA",
  },
];

export const Gender_List = [
  { gender_id: 1, name: "DRAMA" },
  { gender_id: 2, name: "COMEDI" },
  { gender_id: 3, name: "ACCION" },
];

export const Season_List = [
  { season_id: 1, name: "SEASON 1" },
  { season_id: 2, name: "SEASON 2" },
  { season_id: 3, name: "SEASON 3" },
  { season_id: 4, name: "SEASON 4" },
  { season_id: 5, name: "SEASON 5" },
  { season_id: 6, name: "SEASON 6" },
  { season_id: 7, name: "SEASON 7" },
  { season_id: 8, name: "SEASON 8" },
  { season_id: 9, name: "SEASON 9" },
  { season_id: 10, name: "SEASON 10" },
];

export const Content_List = [
  { content_type_id: 1, name: "ANIME" },
  { content_type_id: 2, name: "SERIE" },
];

export const Episode_List = [
  { episode_id: 1, title: "TITLE 1", episode: "EPISODE 1" },
  { episode_id: 2, title: "TITLE 4", episode: "EPISODE 2" },
  { episode_id: 3, title: "TITLE 3", episode: "EPISODE 3" },
  { episode_id: 4, title: "TITLE 4", episode: "EPISODE 4" },
];
