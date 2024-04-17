import { Displayer } from "./infrastructure/displayer.js";

const apiUrl = `http://localhost:5035/`;
export const displayer = new Displayer(apiUrl);
export const selected = {
    localFile: undefined,
    localDirectory: undefined,
    serverFile: undefined,
    serverDirectory: undefined 
}