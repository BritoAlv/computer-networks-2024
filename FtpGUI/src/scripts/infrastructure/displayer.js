import { DirectoryTree } from "./directory.js";
import { Requester } from "./requester.js";
import { CreateConnectionRequest, ListServerRequest } from "./requests.js";

export class Displayer {
    #apiUrl;
    #requester;
    #serverDirectoryTree;
    #statusList;

    constructor(apiUrl) {
        this.#apiUrl = apiUrl;
        this.#requester = new Requester();
        this.#serverDirectoryTree = new DirectoryTree("root");
        this.#statusList = [];
    }

    async connect(ipAddress, userName, password) {
        const connectionRequest = new CreateConnectionRequest(
            ipAddress.value,
            userName.value,
            password.value
        );
        const connectionResponse = await this.#requester.post(this.#apiUrl + "connect", connectionRequest);

        this.#displayStatus(connectionResponse.status);

        if(!connectionResponse.successful)
            return;

        // Wait until server is functioning
        // const listRequest = new ListServerRequest("/");
        // const listResponse = await this.#requester.post(this.#apiUrl+"list", listRequest);

        // if(!listResponse.successful) {
        //     this.#displayStatus("Error while listing directory");
        //     return;
        // }

        // Mock response
        const directories = ["Pictures", "Music", "Videos", "Books"];
        const files = ["main.c", "lib.c", "script.py"];

        const rootId = this.#serverDirectoryTree.root.id;

        // Insert directories into directory tree
        directories.forEach(dir => {
            this.#serverDirectoryTree.insertDirectory(rootId, dir);
        }); 
    
        // Insert files into directory tree
        files.forEach(f => {
            this.#serverDirectoryTree.insertFile(rootId, f);
        });

        const serverDirectory = document.querySelector("#server-directory");

        // Display html
        serverDirectory.innerHTML = this.#serverDirectoryTree.toHtml();
    }

    async close() {
        const response = await this.#requester.get(this.#apiUrl + "close");

        this.#displayStatus(response.status);

        const serverDirectory = document.querySelector("#server-directory");

        // Reset server directory tree
        this.#serverDirectoryTree = new DirectoryTree("root");
        // Reset server directory html
        serverDirectory.innerHTML = "";
    }

    async update() {
        const response = await this.#requester.get(this.#apiUrl + "status");

        if (response.status == "")
            return;

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
}