import { create } from "zustand";
import { _contents, Content, ContentDTO } from "../models/Contents";

type Data = {
  form_content_type: Content;
  list_content_type: ContentDTO[];
  getContent_Types: () => void;
  postContent_Types: () => void;
  reset: () => void;
};

export const useContent = create<Data>((set) => ({
  form_content_type: _contents,
  list_content_type: [],
  getContent_Types: () => {
    const mockData: ContentDTO[] = [
      {
        content_type_id: 1,
        content_type_title: "Breaking Bad",
        content_type_cover: "breakingbad.jpg",
        content_type_year: 2008,
        content_type_url: "https://example.com/breakingbad",
        content_type_gender: "DRAMA",
      },
    ];
    set({ list_content_type: mockData });
  },
  postContent_Types: () => {},
  reset: () => set({ form_content_type: _contents }),
}));
