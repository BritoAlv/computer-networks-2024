import { Displayer } from "./infrastructure/displayer.js";

const apiUrl = `http://localhost:5035/`;
export const displayer = new Displayer(apiUrl);