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
