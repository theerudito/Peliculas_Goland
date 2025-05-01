import axios from "axios";
import { url_base } from "./Initial";

export const GET_Content = async () => {
  try {
    return (await axios.get(`${url_base}/content`)).data;
  } catch (error) {
    console.error("GET failed:", error);
    throw error;
  }
};

export const GET_Gender = async () => {
  try {
    return (await axios.get(`${url_base}/gender`)).data;
  } catch (error) {
    console.error("GET failed:", error);
    throw error;
  }
};

export const GET_Season = async () => {
  try {
    return (await axios.get(`${url_base}/season`)).data;
  } catch (error) {
    console.error("GET failed:", error);
    throw error;
  }
};
