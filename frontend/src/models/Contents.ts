export interface Content {
  content_id: number;
  content_title: string;
  content_cover: string;
  content_year: number;
  content_url: string;
  gender_id: number;
}

export interface ContentDTO {
  content_id: number;
  content_title: string;
  content_cover: string;
  content_year: number;
  content_url: string;
  content_gender: string;
}
