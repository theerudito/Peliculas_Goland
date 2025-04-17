import { create } from "zustand";

type Data = {
  gender_list: Genders[];
  season_list: Seasons[];
  type_list: TypeDTO[];
};

export const useData = create<Data>()(() => ({
  gender_list: initialGenders,
  season_list: initialSeasons,
  type_list: initialTypes,
}));
