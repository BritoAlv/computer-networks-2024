import { displayer } from "../globals.js"

export async function localDirectoryComponent() {
    await displayer.displayLocalDirectory();

    const uploadFileButton = document.querySelector("#upload-file-button");
    const uploadDirectoryButton = document.querySelector("#upload-directory-button");
    const refreshButton = document.querySelector("#local-refresh");

    uploadFileButton.addEventListener("click", async () => {
        await displayer.uploadFile();
    });

    uploadDirectoryButton.addEventListener("click", async () => {
        await displayer.uploadDirectory();
    });

    refreshButton.addEventListener("click", async () => {
        await displayer.refreshLocal();
    });
}