import axios from "axios";
import { Content, ContentDTO } from "../models/Contents";
import { url_base } from "./Initial";

export const GET_Content = async () => {
  try {
    const response = await axios.get<ContentDTO>(`${url_base}/content`);
    return { success: true, data: response.data };
  } catch (error: unknown) {
    let message = "Error desconocido";
    if (axios.isAxiosError(error)) {
      message = error.response?.data?.error || message;
    }
    return { success: false, error: message };
  }
};

export const GET_Content_Type = async (value: number) => {
  try {
    const response = await axios.get<ContentDTO>(
      `${url_base}/content/type_content/${value}`
    );
    return { success: true, data: response.data };
  } catch (error: unknown) {
    let message = "Error desconocido";
    if (axios.isAxiosError(error)) {
      message = error.response?.data?.error || message;
    }
    return { success: false, error: message };
  }
};

export const GET_Find_Content = async (value: number) => {
  try {
    const response = await axios.get<ContentDTO>(
      `${url_base}/content/find/${value}`
    );
    return { success: true, data: response.data };
  } catch (error: unknown) {
    let message = "Error desconocido";
    if (axios.isAxiosError(error)) {
      message = error.response?.data?.error || message;
    }
    return { success: false, error: message };
  }
};

export const POST_Content = async (obj: Content) => {
  try {
    const response = await axios.post(`${url_base}/content`, obj);
    return { success: true, data: response.data };
  } catch (error: unknown) {
    let message = "Error desconocido";
    if (axios.isAxiosError(error)) {
      message = error.response?.data?.error || message;
    }
    return { success: false, error: message };
  }
};
