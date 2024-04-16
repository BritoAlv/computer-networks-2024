import { selected } from "../globals.js";
import { DirectoryTree } from "./directory.js";
import { Requester } from "./requester.js";
import { CreateConnectionRequest, ListServerRequest, TransferRequest } from "./requests.js";

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

    async connect(ipAddress, userName, password) {
        const connectionRequest = new CreateConnectionRequest(
            ipAddress,
            userName,
            password
        );
        const connectionResponse = await this.#requester.post(this.#apiUrl + "connect", connectionRequest);

        this.#displayStatus(connectionResponse.status);

        if (!connectionResponse.successful)
            return;

        this.displayServerDirectory()
    }

    async close() {
        const response = await this.#requester.get(this.#apiUrl + "close");

        this.#displayStatus(response.status);

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

        // Wait until server is functioning
        // const listRequest = new ListServerRequest(path);
        // const listResponse = await this.#requester.post(this.#apiUrl+"list/server", listRequest);
        // if(!listResponse.successful) {
        //     this.#displayStatus("Error while listing directory");
        //     return;
        // }

        // Mock response
        const directories = ["Pictures", "Music", "Videos", "Books"];
        const files = ["main.c", "lib.c", "script.py"];

        // Insert directories into directory tree
        directories.forEach(dir => {
            if (!directory.directories.map(d => d.name).includes(dir))
                this.#serverDirectoryTree.insertDirectory(directoryId, dir);
        });

        // Insert files into directory tree
        files.forEach(f => {
            if (!directory.files.map(fp => fp.name).includes(f))
                this.#serverDirectoryTree.insertFile(directoryId, f);
        });

        const serverDirectory = document.querySelector("#server-directory");

        // Display html
        serverDirectory.innerHTML = this.#serverDirectoryTree.toHtml();

        this.#setServerDirectoryEvents(false);
    }

    async displayLocalDirectory(directoryId = undefined) {
        if (directoryId == undefined)
            directoryId = this.#localDirectoryTree.root.id;

        const directory = this.#localDirectoryTree.findDirectory(directoryId);
        const path = directory.path;

        // Wait until server is functioning
        // const listRequest = new ListServerRequest(path);
        // const listResponse = await this.#requester.post(this.#apiUrl+"list/local", listRequest);

        // if(!listResponse.successful) {
        //     this.#displayStatus("Error while listing directory");
        //     return;
        // }

        // Mock response
        const directories = ["Movies", "Lectures", "Projects"];
        const files = ["pic.jpeg", "music.mp3"];

        // Insert directories into directory tree
        directories.forEach(dir => {
            if (!directory.directories.map(d => d.name).includes(dir))
                this.#localDirectoryTree.insertDirectory(directoryId, dir);
        });

        // Insert files into directory tree
        files.forEach(f => {
            if (!directory.files.map(fp => fp.name).includes(f))
                this.#localDirectoryTree.insertFile(directoryId, f);
        });

        const localDirectory = document.querySelector("#local-directory");

        // Display html
        localDirectory.innerHTML = this.#localDirectoryTree.toHtml();

        this.#setLocalDirectoryEvents();
    }

    async uploadFile() {
        if (selected.localFile == undefined || selected.serverDirectory == undefined) {
            this.#displayStatus("Error while uploading file. Must select a file and a destination directory in order to upload");
            return;
        }

        const source = this.#localDirectoryTree.findFile(selected.localFile.substr(1)).path();
        const destination = this.#serverDirectoryTree.findDirectory(selected.serverDirectory.substr(1)).path;

        const request = new TransferRequest(source, destination);
        const response = await this.#requester.post(this.#apiUrl + "files/upload", request);

        this.#displayStatus(response.status);
    }

    async uploadDirectory() {
        if (selected.localDirectory == undefined || selected.serverDirectory == undefined) {
            this.#displayStatus("Error while uploading file. Must select a source and destination directory in order to upload");
            return;
        }

        const source = this.#localDirectoryTree.findDirectory(selected.localDirectory.substr(1)).path;
        const destination = this.#serverDirectoryTree.findDirectory(selected.serverDirectory.substr(1)).path;

        const request = new TransferRequest(source, destination);
        const response = await this.#requester.post(this.#apiUrl + "directories/upload", request);

        this.#displayStatus(response.status);
    }

    async downloadFile() {
        if (selected.localDirectory == undefined || selected.serverFile == undefined) {
            this.#displayStatus("Error while uploading file. Must select a file and a destination directory in order to download");
            return;
        }

        const source = this.#serverDirectoryTree.findFile(selected.serverFile.substr(1)).path();
        const destination = this.#localDirectoryTree.findDirectory(selected.localDirectory.substr(1)).path;

        const request = new TransferRequest(source, destination);
        const response = await this.#requester.post(this.#apiUrl + "files/download", request);

        this.#displayStatus(response.status);
    }

    async downloadDirectory() {
        if (selected.localDirectory == undefined || selected.serverDirectory == undefined) {
            this.#displayStatus("Error while uploading file. Must select a source and destination directory in order to download");
            return;
        }

        const source = this.#serverDirectoryTree.findDirectory(selected.serverDirectory.substr(1)).path;
        const destination = this.#localDirectoryTree.findDirectory(selected.localDirectory.substr(1)).path;

        const request = new TransferRequest(source, destination);
        const response = await this.#requester.post(this.#apiUrl + "directories/download", request);

        this.#displayStatus(response.status);
    }

    #displayStatus(status) {
        const statusContainer = document.querySelector("#status");

        this.#statusList.push(status);
        let data = "";
        this.#statusList.forEach(s => {
            data += `<li>${s}</li>`
        });
        statusContainer.innerHTML = data;
    }

    #setLocalDirectoryEvents() {
        const fileItems = document.querySelectorAll(`#local-directory .file-item`);
        const directoryItems = document.querySelectorAll("#local-directory .directory-item");
        fileItems.forEach(item => {
            item.addEventListener("click", () => {

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

                // Remove select-directory class from previously selected directory
                if (selected.localDirectory != undefined) {
                    const previouslySelected = document.querySelector(`#${selected.localDirectory}`);
                    previouslySelected.className = "directory-item";
                }

                selected.localDirectory = item.id;
                item.className += " selected-directory";
            });
        });
    }

    #setServerDirectoryEvents() {
        const fileItems = document.querySelectorAll(`#server-directory .file-item`);
        const directoryItems = document.querySelectorAll("#server-directory .directory-item");
        fileItems.forEach(item => {
            item.addEventListener("click", () => {

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

                // Remove select-directory class from previously selected directory
                if (selected.serverDirectory != undefined) {
                    const previouslySelected = document.querySelector(`#${selected.serverDirectory}`);
                    previouslySelected.className = "directory-item";
                }

                selected.serverDirectory = item.id;
                item.className += " selected-directory";
            });
        });
    }
}