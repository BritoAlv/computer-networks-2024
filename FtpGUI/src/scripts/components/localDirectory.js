import { displayer, selected } from "../globals.js"

export function localDirectoryComponent() {
    displayer.displayLocalDirectory();

    const uploadFileButton = document.querySelector("#upload-file-button");
    const uploadDirectoryButton = document.querySelector("#upload-directory-button");
    const refreshButton = document.querySelector("#local-refresh");

    uploadFileButton.addEventListener("click", () => {
        displayer.uploadFile();
    });

    uploadDirectoryButton.addEventListener("click", () => {
        displayer.uploadDirectory();
    });

    refreshButton.addEventListener("click", () => {
        displayer.refreshLocal();
    });
}