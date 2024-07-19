import axios from "axios";

const apiUrl = process.env.NEXT_PUBLIC_BACKEND_URL ?? "http://localhost:3000";

export const getTodos = async () => {
  const { data } = await axios.get(apiUrl+"/todos");
  return data;
};
