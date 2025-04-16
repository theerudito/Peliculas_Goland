import { createSlice } from "@reduxjs/toolkit";
import { GenderDTO, SeasonsDTO } from "../../models/Movies";
import { Gender_List, Season_List } from "../../helpers/Data";

export const Data_Slices = createSlice({
  name: "data",
  initialState: {
    gender_list: Gender_List as GenderDTO[],
    season_list: Season_List as SeasonsDTO[],
    type_list: [
      { content_type_id: 1, name: "SERIE" },
      { content_type_id: 2, name: "ANIME" },
    ],
  },
  reducers: {},
});
