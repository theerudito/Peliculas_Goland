import axios from "axios";
import { ContentDTO } from "../models/Contents";

const url = "http://127.0.0.1:1000/api/v1";

export const GET_Find_Content = async (value: number) => {
  try {
    return (await axios.get(`${url}/contents/find/${value}`)).data;
  } catch (error) {
    console.error("GET failed:", error);
    throw error;
  }
};

export const GET_Content_Type = async (value: number) => {
  try {
    return (await axios.get(`${url}/contents/content-type/${value}`)).data;
  } catch (error) {
    console.error("GET failed:", error);
    throw error;
  }
};

export const GET_Content_Season = async (value: string) => {
  try {
    // const data = await axios.get(`${url}/contents/season/${value}`);

    // console.log("GET Content Season", data.data);

    return (await axios.get(`${url}/contents/season/${value}`)).data;
  } catch (error) {
    console.error("GET failed:", error);
    throw error;
  }
};

export const GET_Content_Epidodes = async (value: number) => {
  try {
    //const data = await axios.get(`${url}/contents/episodes/${value}`);

    //console.log("GET Content Episodes", data.data);

    return (await axios.get(`${url}/contents/episodes/${value}`)).data;
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
