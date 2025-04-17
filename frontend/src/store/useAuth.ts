import { create } from "zustand";
import { _users, Users } from "../models/Auth";

type Data = {
  user: Users;
  setLogin: (user: string, password: string) => void;
};

export const useAuth = create<Data>((set) => ({
  user: _users,
  setLogin: (user, password) =>
    set(() => ({
      user: { user, password },
    })),
  reset: () =>
    set(() => ({
      user: _users,
    })),
}));
