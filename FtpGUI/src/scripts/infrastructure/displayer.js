import { selected } from "../globals.js";
import { DirectoryTree } from "./directory.js";
import { Requester } from "./requester.js";
import { ConnectRequest, PathRequest, TransferRequest } from "./requests.js";

export class Displayer {
    #apiUrl;
    #requester;
    #serverDirectoryTree;
    #localDirectoryTree;
    #statusList;

    constructor(apiUrl) {
        this.#apiUrl = apiUrl;
        this.#requester = new Requester();
        this.#serverDirectoryTree = new DirectoryTree("root");
        this.#localDirectoryTree = new DirectoryTree("root");
        this.#statusList = [];
    }

    async connect(ipAddress, port, userName, password) {
        const connectionRequest = new ConnectRequest(
            ipAddress,
            port,
            userName,
            password
        );
        const connectionResponse = await this.#requester.post(this.#apiUrl + "connect", connectionRequest);

        this.#displayStatus(connectionResponse.status);

        if (!connectionResponse.successful)
            return;

        await this.displayServerDirectory()
    }

    async close() {
        const response = await this.#requester.get(this.#apiUrl + "close");

        this.#displayStatus(response.status);

        if (!response.successful)
            return

        const serverDirectory = document.querySelector("#server-directory");

        // Reset server directory tree
        this.#serverDirectoryTree = new DirectoryTree("root");
        // Reset server directory html
        serverDirectory.innerHTML = "";

        // Reset selected items
        selected.localFile = undefined;
        selected.localDirectory = undefined;
        selected.serverFile = undefined;
        selected.serverDirectory = undefined;
    }

    async update() {
        const response = await this.#requester.get(this.#apiUrl + "status");

        if (response.status == "")
            return;

        this.#displayStatus(response.status);
    }

    async displayServerDirectory(directoryId = undefined) {
        if (directoryId == undefined)
            directoryId = this.#serverDirectoryTree.root.id;

        const directory = this.#serverDirectoryTree.findDirectory(directoryId);
        const path = directory.path;

        const request = new PathRequest(path);
        const response = await this.#requester.post(this.#apiUrl + "list/server", request);

        if (!response.successful) {
            this.#displayStatus("Error while listing directory");
            return;
        }

        // Insert directories into directory tree
        response.directories.forEach(dir => {
            if (!directory.directories.map(d => d.name).includes(dir))
                this.#serverDirectoryTree.insertDirectory(directoryId, dir);
        });

        // Remove nonexisting directories
        const directoriesToRemove = directory.directories.filter(dir => !response.directories.includes(dir.name)).map(dir => dir.id);

        directoriesToRemove.forEach(dir => {
            this.#serverDirectoryTree.removeDirectory(dir);
        });

        // Insert files into directory tree
        response.files.forEach(f => {
            if (!directory.files.map(fp => fp.name).includes(f))
                this.#serverDirectoryTree.insertFile(directoryId, f);
        });

        // Remove nonexisting files
        const filesToRemove = directory.files.filter(f => !response.files.includes(f.name)).map(f => f.id);

        filesToRemove.forEach(f => {
            this.#serverDirectoryTree.removeFile(f);
        });

        this.#setServerDirectoryHtml();

        this.#displayStatus("Directory successfully listed");
    }

    async displayLocalDirectory(directoryId = undefined) {
        if (directoryId == undefined)
            directoryId = this.#localDirectoryTree.root.id;

        const directory = this.#localDirectoryTree.findDirectory(directoryId);
        const path = directory.path;

        const request = new PathRequest(path);
        const response = await this.#requester.post(this.#apiUrl + "list/local", request);

        if (!response.successful) {
            this.#displayStatus("Error while listing directory");
            return;
        }

        // Insert directories into directory tree
        response.directories.forEach(dir => {
            if (!directory.directories.map(d => d.name).includes(dir))
                this.#localDirectoryTree.insertDirectory(directoryId, dir);
        });

        // Remove nonexisting directories
        const directoriesToRemove = directory.directories.filter(dir => !response.directories.includes(dir.name)).map(dir => dir.id);

        directoriesToRemove.forEach(dir => {
            this.#localDirectoryTree.removeDirectory(dir);
        });

        // Insert files into directory tree
        response.files.forEach(f => {
            if (!directory.files.map(fp => fp.name).includes(f))
                this.#localDirectoryTree.insertFile(directoryId, f);
        });

        // Remove nonexisting files
        const filesToRemove = directory.files.filter(f => !response.files.includes(f.name)).map(f => f.id);

        filesToRemove.forEach(f => {
            this.#localDirectoryTree.removeFile(f);
        });

        this.#setLocalDirectoryHtml();
    }

    async uploadFile() {
        if (selected.localFile == undefined || selected.serverDirectory == undefined) {
            this.#displayStatus("Error while uploading file. Must select a file and a destination directory in order to upload");
            return;
        }

        const source = this.#localDirectoryTree.findFile(selected.localFile.substring(1)).path();
        const destination = this.#serverDirectoryTree.findDirectory(selected.serverDirectory.substring(1)).path;

        const request = new TransferRequest(source, destination);
        const response = await this.#requester.post(this.#apiUrl + "files/upload", request);

        this.#displayStatus(response.status);
    }

    async uploadDirectory() {
        if (selected.localDirectory == undefined || selected.serverDirectory == undefined) {
            this.#displayStatus("Error while uploading directory. Must select a source and destination directory in order to upload");
            return;
        }

        const source = this.#localDirectoryTree.findDirectory(selected.localDirectory.substring(1)).path;
        const destination = this.#serverDirectoryTree.findDirectory(selected.serverDirectory.substring(1)).path;

        const request = new TransferRequest(source, destination);
        const response = await this.#requester.post(this.#apiUrl + "directories/upload", request);

        this.#displayStatus(response.status);
    }

    async downloadFile() {
        if (selected.localDirectory == undefined || selected.serverFile == undefined) {
            this.#displayStatus("Error while downloading file. Must select a file and a destination directory in order to download");
            return;
        }

        const source = this.#serverDirectoryTree.findFile(selected.serverFile.substring(1)).path();
        const destination = this.#localDirectoryTree.findDirectory(selected.localDirectory.substring(1)).path;

        const request = new TransferRequest(source, destination);
        const response = await this.#requester.post(this.#apiUrl + "files/download", request);

        this.#displayStatus(response.status);
    }

    async downloadDirectory() {
        if (selected.localDirectory == undefined || selected.serverDirectory == undefined) {
            this.#displayStatus("Error while downloading directory. Must select a source and destination directory in order to download");
            return;
        }

        const source = this.#serverDirectoryTree.findDirectory(selected.serverDirectory.substring(1)).path;
        const destination = this.#localDirectoryTree.findDirectory(selected.localDirectory.substring(1)).path;

        const request = new TransferRequest(source, destination);
        const response = await this.#requester.post(this.#apiUrl + "directories/download", request);

        this.#displayStatus(response.status);
    }

    async refreshLocal() {
        if (selected.localDirectory == undefined) {
            this.#displayStatus("Error while refreshing local directory. Must select a directory in order to refresh it");
            return;
        }

        await this.displayLocalDirectory(selected.localDirectory.substring(1));

        // Reset selected items
        selected.localFile = undefined;
        selected.localDirectory = undefined;

    }

    async refreshServer() {
        if (selected.serverDirectory == undefined) {
            this.#displayStatus("Error while refreshing server directory. Must select a directory in order to refresh it");
            return;
        }

        await this.displayServerDirectory(selected.serverDirectory.substring(1));

        // Reset selected items
        selected.serverFile = undefined;
        selected.serverDirectory = undefined;
    }

    async createDirectory(directoryName) {
        if (selected.serverDirectory == undefined) {
            this.#displayStatus("Error while creating server directory. Must select a directory in order to create one");
            return;
        }

        const directory = this.#serverDirectoryTree.findDirectory(selected.serverDirectory.substring(1));
        const path = `${directory.path}/${directoryName}`;
        const request = new PathRequest(path);
        const response = await this.#requester.post(this.#apiUrl + "directories/create", request);

        this.#displayStatus(response.status);

        if (!response.successful)
            return;

        this.#serverDirectoryTree.insertDirectory(directory.id, directoryName);

        this.#setServerDirectoryHtml();
    }

    async removeDirectory() {
        if (selected.serverDirectory == undefined) {
            this.#displayStatus("Error while removing server directory. Must select a directory in order to remove it");
            return;
        }

        const directory = this.#serverDirectoryTree.findDirectory(selected.serverDirectory.substring(1));
        const path = directory.path;
        const request = new PathRequest(path);
        const response = await this.#requester.post(this.#apiUrl + "directories/remove", request);

        this.#displayStatus(response.status);

        if (!response.successful)
            return;

        selected.serverDirectory = undefined;
        this.#serverDirectoryTree.removeDirectory(directory.id);

        this.#setServerDirectoryHtml();
    }

    async removeFile() {
        if (selected.serverFile == undefined) {
            this.#displayStatus("Error while removing server file. Must select a file in order to remove it");
            return;
        }

        const file = this.#serverDirectoryTree.findFile(selected.serverFile.substring(1));
        const path = file.path();
        const request = new PathRequest(path);
        const response = await this.#requester.post(this.#apiUrl + "files/remove", request);

        this.#displayStatus(response.status);

        if (!response.successful)
            return;

        selected.serverFile = undefined;
        this.#serverDirectoryTree.removeFile(file.id);

        this.#setServerDirectoryHtml();
    }

    #displayStatus(status) {
        const statusContainer = document.querySelector("#status");

        this.#statusList.unshift(status);
        let data = "";
        this.#statusList.forEach(s => {
            data += `<li>${s}</li>`
        });
        statusContainer.innerHTML = `<ul>${data}</ul>`;
    }

    #setLocalDirectoryHtml() {
        const localDirectory = document.querySelector("#local-directory");
        localDirectory.innerHTML = this.#localDirectoryTree.toHtml();
        this.#setLocalDirectoryEvents();
    }

    #setServerDirectoryHtml() {
        const serverDirectory = document.querySelector("#server-directory");
        serverDirectory.innerHTML = this.#serverDirectoryTree.toHtml();
        this.#setServerDirectoryEvents(false);
    }

    #setLocalDirectoryEvents() {
        const fileItems = document.querySelectorAll(`#local-directory .file-item`);
        const directoryItems = document.querySelectorAll("#local-directory .directory-item");
        fileItems.forEach(item => {
            item.addEventListener("click", () => {
                if (selected.localFile == item.id) {
                    item.className = `file-item`;
                    selected.localFile = undefined;
                    return;
                }

                // Remove select-file class from previously selected file
                if (selected.localFile != undefined) {
                    const previouslySelected = document.querySelector(`#${selected.localFile}`);
                    previouslySelected.className = `file-item`;
                }

                selected.localFile = item.id;
                item.className += " selected-file";
            });
        });

        directoryItems.forEach(item => {
            item.addEventListener("click", () => {
                if (selected.localDirectory == item.id) {
                    item.className = `directory-item`;
                    selected.localDirectory = undefined;
                    return;
                }

                // Remove select-directory class from previously selected directory
                if (selected.localDirectory != undefined) {
                    const previouslySelected = document.querySelector(`#${selected.localDirectory}`);
                    previouslySelected.className = "directory-item";
                }

                selected.localDirectory = item.id;
                item.className += " selected-directory";
            });
        });

        directoryItems.forEach(item => {
            item.addEventListener("dblclick", () => {
                const directory = this.#localDirectoryTree.findDirectory(item.id.substring(1));

                directory.display = !directory.display;

                this.displayLocalDirectory(directory.id);
            });
        })
    }

    #setServerDirectoryEvents() {
        const fileItems = document.querySelectorAll(`#server-directory .file-item`);
        const directoryItems = document.querySelectorAll("#server-directory .directory-item");
        fileItems.forEach(item => {
            item.addEventListener("click", () => {
                if (selected.serverFile == item.id) {
                    item.className = `file-item`;
                    selected.serverFile = undefined;
                    return;
                }

                // Remove select-file class from previously selected file
                if (selected.serverFile != undefined) {
                    const previouslySelected = document.querySelector(`#${selected.serverFile}`);
                    previouslySelected.className = `file-item`;
                }

                selected.serverFile = item.id;
                item.className += " selected-file";
            });
        });

        directoryItems.forEach(item => {
            item.addEventListener("click", () => {
                if (selected.serverDirectory == item.id) {
                    item.className = `directory-item`;
                    selected.serverDirectory = undefined;
                    return;
                }

                // Remove select-directory class from previously selected directory
                if (selected.serverDirectory != undefined) {
                    const previouslySelected = document.querySelector(`#${selected.serverDirectory}`);
                    previouslySelected.className = "directory-item";
                }

                selected.serverDirectory = item.id;
                item.className += " selected-directory";
            });
        });

        directoryItems.forEach(item => {
            item.addEventListener("dblclick", () => {
                const directory = this.#serverDirectoryTree.findDirectory(item.id.substring(1));

                directory.display = !directory.display;

                this.displayServerDirectory(directory.id);
            });
        })
    }
}