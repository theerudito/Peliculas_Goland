import axios from "axios";

const url = "http://127.0.0.1:1000/api/v1/movies";

export const GET_Movies = async () => {
  const res = await axios.get(url);

  return res.data;
};
