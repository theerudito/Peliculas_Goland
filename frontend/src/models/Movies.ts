export interface Movies {
  movie_title: string;
  movie_year: number;
  movie_cover: string;
  movie_url: string;
  gender_id: number;
}
export interface MoviesDTO {
  movie_movie_id: number;
  movie_title: string;
  movie_year: number;
  movie_cover: string;
  movie_url: string;
  movie_gender: string;
}

export const _movies: Movies = {
  movie_title: "",
  movie_year: 0,
  movie_cover: "",
  movie_url: "",
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

export interface GenderDTO {
  gender_id: number;
  name: string;
}

export interface Content_Types {
  content_type_id: number;
  content_title: string;
  content_cover: string;
  content_year: number;
  content_url: string;
  gender_id: number;
}

export interface Content_TypesDTO {
  content_type_id: number;
  content_title: string;
  content_cover: string;
  content_year: number;
  content_url: string;
  gender_id: string;
}

export const _contents: Content_Types = {
  content_type_id: 0,
  content_title: "",
  content_cover: "",
  content_year: 0,
  content_url: "",
  gender_id: 0,
};

export interface Seasons {
  season_id: number;
  content_type_id: number;
  season_number: number;
  season_title: string;
}

export interface SeasonsDTO {
  season_id: number;
  name: string;
}

export const _seasons: Seasons = {
  season_id: 0,
  content_type_id: 0,
  season_number: 0,
  season_title: "",
};

export interface Episodes {
  episode_id: number;
  season_id: number;
  episode_number: number;
  episode_title: string;
  episode_url: string;
}

export interface EpisodesDTO {
  season_id: number;
  season: string;
  episode_number: number;
  episode_title: string;
  episode_url: string;
}

export const _episodes: Episodes = {
  episode_id: 0,
  season_id: 0,
  episode_number: 0,
  episode_title: "",
  episode_url: "",
};

export const _episodesList: Episodes[] = [
  {
    episode_id: 0,
    season_id: 0,
    episode_number: 0,
    episode_title: "",
    episode_url: "",
  },
];

export interface FormDataDTO {
  gender_id: number;
  season: Seasons;
  content: Content_Types;
  episode: Episodes[];
}

export const _form: FormDataDTO = {
  gender_id: 0,
  season: {
    season_id: 0,
    content_type_id: 0,
    season_number: 0,
    season_title: "",
  },
  content: {
    content_type_id: 0,
    content_title: "",
    content_cover: "",
    content_year: 0,
    content_url: "",
    gender_id: 0,
  },
  episode: [
    {
      episode_id: 0,
      season_id: 0,
      episode_number: 0,
      episode_title: "",
      episode_url: "",
    },
  ],
};
