import { displayer } from "../globals.js";

export function connectionStatusComponent() {
    setInterval(async () => await displayer.update(), 2000);
}