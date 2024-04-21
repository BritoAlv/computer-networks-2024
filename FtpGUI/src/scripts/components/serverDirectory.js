import { displayer, selected } from "../globals.js"

export function serverDirectoryComponent() {
    
    const downloadFileButton = document.querySelector("#download-file-button");
    const downloadDirectoryButton = document.querySelector("#download-directory-button");
    const refreshButton = document.querySelector("#server-refresh");
    const createDirectoryButton = document.querySelector("#create-directory-button");
    const createDirectoryInput = document.querySelector("#create-directory-input");
    const removeDirectoryButton = document.querySelector("#remove-directory-button");
    const removeFileButton = document.querySelector("#remove-file-button");
    const renameFileButton = document.querySelector("#rename-file-button");
    const renameFileInput = document.querySelector("#rename-file-input");


    downloadFileButton.addEventListener("click", async () => {
        await displayer.downloadFile();
    });

    downloadDirectoryButton.addEventListener("click", async () => {
        await displayer.downloadDirectory();
    });

    refreshButton.addEventListener("click", async () => {
        await displayer.refreshServer();
    });

    createDirectoryButton.addEventListener("click", async () => {
        await displayer.createDirectory(createDirectoryInput.value);
        createDirectoryInput.value = "";
    });

    removeDirectoryButton.addEventListener("click", async () => {
        await displayer.removeDirectory();
    });

    removeFileButton.addEventListener("click", async () => {
        await displayer.removeFile();
    });

    renameFileButton.addEventListener("click", async () => {
        await displayer.renameFile(renameFileInput.value);
        renameFileInput.value = "";
    });
}