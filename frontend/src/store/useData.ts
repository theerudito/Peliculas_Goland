import { create } from "zustand";
import { _gender, Genders } from "../models/Gender";
import { _season, Season } from "../models/Seasons";
import { Content_Type } from "../models/Content_Type";
import { Gender_List, Season_List } from "../helpers/Data";
import {
  GET_Gender,
  GET_Season,
  POST_Gender,
  POST_Season,
} from "../helpers/Fetching_Data";
import { YearDTO } from "../models/Years";

type Data = {
  // LISTADO
  gender_list: Genders[];
  season_list: Season[];
  year_list: YearDTO[];
  type_list: Content_Type[];

  // FUNCIONES
  getSeason: () => void;
  postSeason: (obj: Season) => void;
  postGender: (obj: Genders) => void;
  getGender: () => void;
  getYear: () => void;
  getType: () => void;
  reset: () => void;

  // FORMULARIOS
  form_gender: Genders;
  form_season: Season;
};

export const useData = create<Data>()((set, get) => ({
  // LISTADO
  gender_list: Gender_List,
  season_list: Season_List,
  year_list: [],
  type_list: [],

  // FUNCIONES
  getGender: async () => {
    const response = await GET_Gender();
    if (response.data && Array.isArray(response.data)) {
      set({
        gender_list: response.data.length === 0 ? [] : response.data,
      });
    } else {
      set({ gender_list: [] });
    }
  },

  postGender: async (obj: Genders) => {
    const result = await POST_Gender(obj);

    if (result.success === true) {
      get().reset();
      get().getGender();
      return result.data;
    }

    return result.error;
  },

  getSeason: async () => {
    const response = await GET_Season();
    if (response.data && Array.isArray(response.data)) {
      set({
        season_list: response.data.length === 0 ? [] : response.data,
      });
    } else {
      set({ season_list: [] });
    }
  },

  postSeason: async (obj: Season) => {
    const result = await POST_Season(obj);

    if (result.success === true) {
      get().reset();
      get().getSeason();
      return result.data;
    }

    return result.error;
  },

  getYear: () => {
    const currentYear = new Date().getFullYear();
    const years: YearDTO[] = [];

    for (let year = currentYear; year >= 1980; year--) {
      years.push({
        year_id: year,
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

  reset: () => set({ form_gender: _gender, form_season: _season }),

  // FORMULARIOS
  form_gender: _gender,
  form_season: _season,
}));
