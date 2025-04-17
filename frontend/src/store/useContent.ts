import { create } from "zustand";
import { _content, Content, ContentDTO } from "../models/Contents";

type Data = {
  form_content_type: Content;
  list_content_type: ContentDTO[];
  getContent_Types: () => void;
  postContent_Types: () => void;
  reset: () => void;
};

export const useContent = create<Data>((set) => ({
  form_content_type: _content,
  list_content_type: [],
  getContent_Types: () => {
    const mockData: ContentDTO[] = [
      {
        content_id: 1,
        content_title: "Breaking Bad",
        content_cover: "breakingbad.jpg",
        content_year: 2008,
        content_url: "https://example.com/breakingbad",
        content_gender: "DRAMA",
      },
    ];
    set({ list_content_type: mockData });
  },
  postContent_Types: () => {},
  reset: () => set({ form_content_type: _content }),
}));
