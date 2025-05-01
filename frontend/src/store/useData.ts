import { create } from "zustand";
import { _gender, Genders } from "../models/Gender";
import { _season, Season } from "../models/Seasons";
import { Content_Type } from "../models/Content_Type";
import { Gender_List, Season_List } from "../helpers/Data";
import { GET_Gender, GET_Season } from "../helpers/Fetching_Data";
import { YearDTO } from "../models/Years";

type Data = {
  gender_list: Genders[];
  season_list: Season[];
  year_list: YearDTO[];
  type_list: Content_Type[];

  getSeason: () => void;
  getGender: () => void;
  getYear: () => void;
  getType: () => void;
  reset: () => void;
};

export const useData = create<Data>()((set) => ({
  gender_list: Gender_List,
  season_list: Season_List,
  form_gender: _gender,
  form_season: _season,
  year_list: [],
  type_list: [],

  getGender: async () => {
    const response = await GET_Gender();
    set({ gender_list: response });
  },

  getSeason: async () => {
    const response = await GET_Season();
    set({ season_list: response });
  },

  getYear: () => {
    const currentYear = new Date().getFullYear();
    const years: YearDTO[] = [];

    for (let year = currentYear; year >= 1980; year--) {
      years.push({
        id_year: year,
        year: year,
      });
    }

    set({ year_list: years });
  },

  getType: () => {
    const type: Content_Type[] = [
      { content_type_id: 1, content_type_title: "ANIME" },
      { content_type_id: 2, content_type_title: "SERIE" },
    ];

    set({ type_list: type });
  },

  reset: () => set({ gender_list: [], season_list: [] }),
}));
