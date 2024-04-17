import { displayer, selected } from "../globals.js"

export function serverDirectoryComponent() {
    
    const downloadFileButton = document.querySelector("#download-file-button");
    const downloadDirectoryButton = document.querySelector("#download-directory-button");
    const refreshButton = document.querySelector("#server-refresh");
    const createDirectoryButton = document.querySelector("#create-directory-button");
    const createDirectoryInput = document.querySelector("#create-directory-input");
    const removeDirectoryButton = document.querySelector("#remove-directory-button");
    const removeFileButton = document.querySelector("#remove-file-button");


    downloadFileButton.addEventListener("click", () => {
        displayer.downloadFile();
    });

    downloadDirectoryButton.addEventListener("click", () => {
        displayer.downloadDirectory();
    });

    refreshButton.addEventListener("click", () => {
        displayer.refreshServer();
    });

    createDirectoryButton.addEventListener("click", () => {
        displayer.createDirectory(createDirectoryInput.value);
        createDirectoryInput.value = "";
    });

    removeDirectoryButton.addEventListener("click", () => {
        displayer.removeDirectory();
    });

    removeFileButton.addEventListener("click", () => {
        displayer.removeFile();
    });
}