import axios from "axios";
import { Movies } from "../models/Movies";

const url = "http://127.0.0.1:1000/api/v1";

export const GET_Movies = async () => {
  try {
    return (await axios.get(`${url}/movie`)).data;
  } catch (error) {
    console.error("GET failed:", error);
    throw error;
  }
};

export const POST_Movie = async (obj: Movies) => {
  try {
    return (await axios.post(`${url}/movies`, obj)).data;
  } catch (error) {
    console.error("POST failed:", error);
    throw error;
  }
};
