import axios from "axios";
import { ContentDTO } from "../models/Contents";
import { AddGuiones } from "./Añadir_Guiones";
import { url_base } from "./Initial";

export const GET_Content = async () => {
  try {
    return (await axios.get(`${url_base}/content`)).data;
  } catch (error) {
    console.error("GET failed:", error);
    throw error;
  }
};

export const GET_Content_Type = async (value: number) => {
  try {
    return (await axios.get(`${url_base}/content/type_content/${value}`)).data;
  } catch (error) {
    console.error("GET failed:", error);
    throw error;
  }
};

export const GET_Find_Content = async (value: number) => {
  try {
    return (await axios.get(`${url_base}/content/find/${value}`)).data;
  } catch (error) {
    console.error("GET failed:", error);
    throw error;
  }
};

export const GET_Content_Season = async (value: string) => {
  try {
    return (await axios.get(`${url_base}/content/season/${AddGuiones(value)}`))
      .data;
  } catch (error) {
    console.error("GET failed:", error);
    throw error;
  }
};

export const GET_Content_Epidodes = async (value: string, id: number) => {
  try {
    return (
      await axios.get(
        `${url_base}/contents/episodes/${AddGuiones(value)}/${id}`
      )
    ).data;
  } catch (error) {
    console.error("GET failed:", error);
    throw error;
  }
};

export const POST_Content = async (obj: ContentDTO) => {
  try {
    return (await axios.post(`${url_base}/contents`, obj)).data;
  } catch (error) {
    console.error("POST failed:", error);
    throw error;
  }
};
