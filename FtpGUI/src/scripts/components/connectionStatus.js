import { displayer } from "../globals.js";

export function connectionStatusComponent() {
    setInterval(() => displayer.update(), 1000000);
}