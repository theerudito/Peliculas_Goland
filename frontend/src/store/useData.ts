import { create } from "zustand";
import { Genders } from "../models/Gender";
import { Season } from "../models/Seasons";
import { Content_Type } from "../models/Content_Type";
import { Content_Type_List, Gender_List, Season_List } from "../helpers/Data";

type Data = {
  gender_list: Genders[];
  season_list: Season[];
  type_list: Content_Type[];
};

export const useData = create<Data>()(() => ({
  gender_list: Gender_List,
  season_list: Season_List,
  type_list: Content_Type_List,
}));
