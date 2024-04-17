import { displayer, selected } from "../globals.js"

export function serverDirectoryComponent() {
    
    const downloadFileButton = document.querySelector("#download-file-button");
    const downloadDirectoryButton = document.querySelector("#download-directory-button");
    const refreshButton = document.querySelector("#server-refresh");


    downloadFileButton.addEventListener("click", () => {
        displayer.downloadFile();
    });

    downloadDirectoryButton.addEventListener("click", () => {
        displayer.downloadDirectory();
    });

    refreshButton.addEventListener("click", () => {
        displayer.refreshServer();
    });
}