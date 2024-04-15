import { displayer } from "../constants.js";

export function connectionStatusComponent() {
    setInterval(() => displayer.update(), 2000);
}