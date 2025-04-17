import { create } from "zustand";
import { _gender, Genders } from "../models/Gender";
import { _season, Season } from "../models/Seasons";
import { _content_type, Content_Type } from "../models/Content_Type";
import { Content_Type_List, Gender_List, Season_List } from "../helpers/Data";

type Data = {
  gender: Genders;
  season: Season;
  type: Content_Type;
  gender_list: Genders[];
  season_list: Season[];
  type_list: Content_Type[];
};

export const useData = create<Data>()(() => ({
  gender_list: Gender_List,
  season_list: Season_List,
  type_list: Content_Type_List,
  gender: _gender,
  season: _season,
  type: _content_type,
}));
