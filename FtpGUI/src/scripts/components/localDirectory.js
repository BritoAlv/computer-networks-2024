import { displayer } from "../constants.js"

export function localDirectoryComponent() {
    const uploadButton = document.querySelector("#upload-button");
    const refresh = document.querySelector("#local-refresh");

    displayer.displayLocalDirectory();

    uploadButton.addEventListener("click", () => {
        console.log("Viva Cuba!");
    });
}