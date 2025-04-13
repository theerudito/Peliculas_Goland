export interface Movies {
  title: string;
  year: number;
  cover: string;
  url: string;
  gender_id: number;
}
export interface MoviesDTO {
  movie_id: number;
  title: string;
  year: number;
  cover: string;
  url: string;
  gender: string;
}

export const _movies: Movies = {
  title: "",
  year: 0,
  cover: "",
  url: "",
  gender_id: 0,
};

export interface Users {
  user: string;
  password: string;
}

export const _users: Users = {
  user: "",
  password: "",
};

export interface Content_Types {
  content_type_id: number;
  title_content: string;
  cover_content: string;
  year_content: number;
  url_contenct: string;
  gender_id: number;
}

export interface Content_TypesDTO {
  content_type_id: number;
  title_content: string;
  cover_content: string;
  year_content: number;
  url_contenct: string;
  gender_id: string;
}

export const _contents: Content_Types = {
  content_type_id: 0,
  title_content: "",
  cover_content: "",
  year_content: 0,
  url_contenct: "",
  gender_id: 0,
};

export interface Seasons {
  season_id: number;
  content_type_id: number;
  season_number: number;
  title_season: string;
}

export interface SeasonsDTO {
  season_id: number;
  content_type_id: string;
  season_number: number;
  title_season: string;
}

export const _seasons: Seasons = {
  season_id: 0,
  content_type_id: 0,
  season_number: 0,
  title_season: "",
};

export interface Episodes {
  episode_id: number;
  season_id: number;
  episode_number: number;
  title_episode: string;
  url_episode: string;
}

export interface EpisodesDTO {
  season_id: number;
  season: string;
  episode_number: number;
  title_episode: string;
  url_episode: string;
}

export const _episodes: Episodes = {
  episode_id: 0,
  season_id: 0,
  episode_number: 0,
  title_episode: "",
  url_episode: "",
};

export const _episodesList: Episodes[] = [
  {
    episode_id: 0,
    season_id: 0,
    episode_number: 0,
    title_episode: "",
    url_episode: "",
  },
];

export interface FormDataDTO {
  gender_id: number;
  season: Seasons;
  content: Content_Types;
  episode: Episodes;
}
