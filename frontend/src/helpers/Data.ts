import image_1 from "../assets/imagen1.jpg";
import image_2 from "../assets/imagen3.avif";
import image_3 from "../assets/imagen2.jpg";
import { Episodes } from "../models/Episodes";
import { Season } from "../models/Seasons";
import { Content_Type } from "../models/Content_Type";
import { MoviesDTO } from "../models/Movies";
import { Genders } from "../models/Gender";

export const Movies_List: MoviesDTO[] = [
  {
    movie_id: 1,
    movie_title: "MAN OF HONOR",
    movie_year: 2000,
    movie_cover: image_1,
    movie_url: "https://pixeldrain.com/api/file/kyjhvCUN",
    movie_gender: "CRIMEN",
  },
  {
    movie_id: 2,
    movie_title: "CADENA PERPECTUA",
    movie_year: 1994,
    movie_cover: image_2,
    movie_url: "https://pixeldrain.com/api/file/CkpdetaB",
    movie_gender: "DRAMA",
  },
  {
    movie_id: 3,
    movie_title: "THE GLADIATOR",
    movie_year: 1994,
    movie_cover: image_3,
    movie_url: "https://pixeldrain.com/api/file/kyjhvCUN",
    movie_gender: "DRAMA",
  },
];

export const Gender_List: Genders[] = [
  { gender_id: 1, gender_name: "DRAMA" },
  { gender_id: 2, gender_name: "COMEDY" },
  { gender_id: 3, gender_name: "ACCION" },
];

export const Season_List: Season[] = [
  { season_id: 1, season_name: "SEASON 1" },
  { season_id: 2, season_name: "SEASON 2" },
  { season_id: 3, season_name: "SEASON 3" },
  { season_id: 4, season_name: "SEASON 4" },
  { season_id: 5, season_name: "SEASON 5" },
  { season_id: 6, season_name: "SEASON 6" },
  { season_id: 7, season_name: "SEASON 7" },
  { season_id: 8, season_name: "SEASON 8" },
  { season_id: 9, season_name: "SEASON 9" },
  { season_id: 10, season_name: "SEASON 10" },
];

export const Content_Type_List: Content_Type[] = [
  { content_type: 1, content_type_title: "ANIME" },
  { content_type: 2, content_type_title: "SERIE" },
];

export const Episode_List: Episodes[] = [
  {
    episode_id: 1,
    season_id: 1,
    episode_number: 1,
    episode_name: "TITLE 1",
    episode_url: "",
  },
  {
    episode_id: 2,
    season_id: 1,
    episode_number: 2,
    episode_name: "TITLE 2",
    episode_url: "",
  },
  {
    episode_id: 3,
    season_id: 1,
    episode_number: 3,
    episode_name: "TITLE 3",
    episode_url: "",
  },
];
