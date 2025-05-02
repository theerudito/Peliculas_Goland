import { url_base } from "./Initial";
import { Genders } from "../models/Gender";
import { Season } from "../models/Seasons";
import axios from "axios";

export const GET_Gender = async () => {
  try {
    const response = await axios.get<Genders[]>(`${url_base}/gender`);
    return { success: true, data: response.data };
  } catch (error: unknown) {
    let message = "Error desconocido";
    if (axios.isAxiosError(error)) {
      message = error.response?.data?.error || message;
    }
    return { success: false, error: message };
  }
};

export const POST_Gender = async (obj: Genders) => {
  try {
    const response = await axios.post(`${url_base}/gender`, obj);
    return { success: true, data: response.data };
  } catch (error: unknown) {
    let message = "Error desconocido";
    if (axios.isAxiosError(error)) {
      message = error.response?.data?.error || message;
    }
    return { success: false, error: message };
  }
};

export const GET_Season = async () => {
  try {
    const response = await axios.get<Season[]>(`${url_base}/season`);
    return { success: true, data: response.data };
  } catch (error: unknown) {
    let message = "Error desconocido";
    if (axios.isAxiosError(error)) {
      message = error.response?.data?.error || message;
    }
    return { success: false, error: message };
  }
};

export const POST_Season = async (obj: Season) => {
  try {
    const response = await axios.post(`${url_base}/season`, obj);
    return { success: true, data: response.data };
  } catch (error: unknown) {
    let message = "Error desconocido";
    if (axios.isAxiosError(error)) {
      message = error.response?.data?.error || message;
    }
    return { success: false, error: message };
  }
};
