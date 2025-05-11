import { create } from "zustand";
import { _users, Users } from "../models/Auth";
import { POST_Auth } from "../helpers/Fetching_Auth";

type Data = {
  form_auth: Users;

  postLogin: (obj: Users) => void;
  isLogin: boolean;
  reset: () => void;
};

export const useAuth = create<Data>()((set, get) => ({
  form_auth: _users,
  isLogin: false,

  postLogin: async (obj: Users) => {
    const result = await POST_Auth(obj);
    if (result.success === true) {
      set({ isLogin: true });
      get().reset();
      return result.data;
    }

    set({ isLogin: false });
    return result.error;
  },

  reset: () => set({ form_auth: _users, isLogin: false }),
}));
