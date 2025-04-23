import { SeasonDTO } from "./Seasons";

export interface Content {
  content_id: number;
  content_title: string;
  content_type_id: number;
  content_cover: string;
  content_year: number;
  content_url: string;
  gender_id: number;
}

export interface ContentDTO {
  content_id: number;
  content_title: string;
  content_type: string;
  content_cover: string;
  content_year: number;
  content_url: string;
  content_gender: string;
}

export const _content: Content = {
  content_id: 0,
  content_title: "",
  content_type_id: 0,
  content_cover: "",
  content_year: 0,
  content_url: "",
  gender_id: 0,
};

export const _contentDTO: ContentDTO = {
  content_id: 0,
  content_title: "",
  content_type: "",
  content_cover: "",
  content_year: 0,
  content_url: "",
  content_gender: "",
};

export interface ContentData {
  content: Content;
  seasons: SeasonDTO;
}
