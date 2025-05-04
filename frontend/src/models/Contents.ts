import { EpisodeDTO } from "./Episodes";
import { SeasonDTO } from "./Seasons";

export interface Content {
  content_id: number;
  content_title: string;
  content_type: number;
  content_cover: string;
  content_year: number;
  gender_id: number;
}

export interface ContentDTO {
  content_id: number;
  content_title: string;
  content_type: string;
  content_cover: string;
  content_year: number;
  content_gender: string;
  seasons: SeasonDTO[];
}

export const _content: Content = {
  content_id: 0,
  content_title: "",
  content_type: 0,
  content_cover: "",
  content_year: 0,
  gender_id: 0,
};

export const _contentDTO: ContentDTO = {
  content_id: 0,
  content_title: "",
  content_type: "",
  content_cover: "",
  content_year: 0,
  content_gender: "",
  seasons: [] as SeasonDTO[],
};

export interface ContentDTO_EpisodeDTO {
  content_id: number;
  season_id: number;
  episodes: EpisodeDTO[];
}

export const _content_episodes: ContentDTO_EpisodeDTO = {
  content_id: 0,
  season_id: 0,
  episodes: [] as EpisodeDTO[],
};
