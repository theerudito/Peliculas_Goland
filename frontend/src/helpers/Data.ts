import image_1 from "../assets/logo.webp";
import { Episodes } from "../models/Episodes";
import { MoviesDTO } from "../models/Movies";

export const Movies_List: MoviesDTO[] = [
  {
    movie_id: 1,
    movie_title: "MAN OF HONOR",
    movie_year: 2000,
    movie_cover: image_1,
    movie_url: "https://pixeldrain.com/api/file/kyjhvCUN",
    movie_gender: "CRIMEN",
  },
];

export const Contents_List: MoviesDTO[] = [
  {
    movie_id: 1,
    movie_title: "MAN OF HONOR",
    movie_year: 2000,
    movie_cover: image_1,
    movie_url: "https://pixeldrain.com/api/file/kyjhvCUN",
    movie_gender: "CRIMEN",
  },
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
];
