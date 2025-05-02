import axios from "axios";
import { Users } from "../models/Auth";
import { url_base } from "./Initial";

export const POST_Auth = async (obj: Users) => {
  try {
    const response = await axios.post(`${url_base}/login`, obj);
    return { success: true, data: response.data };
  } catch (error: unknown) {
    let message = "Error desconocido";
    if (axios.isAxiosError(error)) {
      message = error.response?.data?.error || message;
    }
    return { success: false, error: message };
  }
};
