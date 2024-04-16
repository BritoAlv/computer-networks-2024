import { displayer } from "../globals.js"

export function localDirectoryComponent() {
    let selectedLocalFile;
    let selectedLocalDirectory;

    displayer.displayLocalDirectory();

    const uploadFileButton = document.querySelector("#upload-file-button");
    const uploadDirectoryButton = document.querySelector("#upload-directory-button");
    const fileItems = document.querySelectorAll(".file-item");
    const directoryItems = document.querySelectorAll(".directory-item");
    const refresh = document.querySelector("#local-refresh");

    fileItems.forEach(item => {
        item.addEventListener("click", () => {

            // Remove select-file class from previously selected file
            if (selectedLocalFile != undefined) {
                const previouslySelected = document.querySelector(`#${selectedLocalFile}`);
                previouslySelected.className = "file-item";
            }

            selectedLocalFile = item.id;
            item.className += " selected-file";
        });
    });

    directoryItems.forEach(item => {
        item.addEventListener("click", () => {

            // Remove select-directory class from previously selected directory
            if (selectedLocalDirectory != undefined) {
                const previouslySelected = document.querySelector(`#${selectedLocalDirectory}`);
                previouslySelected.className = "directory-item";
            }

            selectedLocalDirectory = item.id;
            item.className += " selected-directory";
        });
    });


    uploadFileButton.addEventListener("click", () => {
        // Remove the i from the start
        displayer.uploadFile(selectedLocalFile == undefined ? selectedLocalFile : selectedLocalFile.substr(1));
    });

    uploadDirectoryButton.addEventListener("click", () => {

    });
}