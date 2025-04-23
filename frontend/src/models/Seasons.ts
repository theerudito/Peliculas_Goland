import { Episodes } from "./Episodes";

export interface Season {
  season_id: number;
  season_name: string;
}

export const _season: Season = {
  season_id: 0,
  season_name: "",
};

export interface SeasonDTO {
  season_id: number;
  season_name: string;
  episodes: Episodes[];
}

export const _seasonDTO: SeasonDTO = {
  season_id: 0,
  season_name: "",
  episodes: [] as Episodes[],
};
