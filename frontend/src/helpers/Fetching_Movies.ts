import axios from "axios";
import { Movies } from "../models/Movies";
import { url_base } from "./Initial";

export const GET_Movies = async () => {
  try {
    return (await axios.get(`${url_base}/movie`)).data;
  } catch (error) {
    console.error("GET failed:", error);
    throw error;
  }
};

export const POST_Movie = async (obj: Movies) => {
  try {
    return (await axios.post(`${url_base}/movies`, obj)).data;
  } catch (error) {
    console.error("POST failed:", error);
    throw error;
  }
};
