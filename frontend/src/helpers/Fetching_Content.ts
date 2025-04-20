import axios from "axios";
import { ContentDTO } from "../models/Contents";

const url = "http://127.0.0.1:1000/api/v1";

export const GET_Content = async (value: number) => {
  try {
    return (await axios.get(`${url}/contents/find/${value}`)).data;
  } catch (error) {
    console.error("GET failed:", error);
    throw error;
  }
};

export const POST_Content = async (obj: ContentDTO) => {
  try {
    return (await axios.post(`${url}/contents`, obj)).data;
  } catch (error) {
    console.error("POST failed:", error);
    throw error;
  }
};
