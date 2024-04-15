import { displayer } from "../constants.js";

export function connectionSetupComponent() {
    const ipAddress = document.querySelector("#ip-address");
    const userName = document.querySelector("#user-name");
    const password = document.querySelector("#password");
    const closeButton = document.querySelector("#close-button");
    const connectButton = document.querySelector("#connect-button");

    // Create connection event
    let response;
    connectButton.addEventListener("click", async () => {
        // TODO: Display tree
        // List root directory
        displayer.connect(ipAddress.value, userName.value, password.value);
        cleanInputs();
    });

    // Close connection event
    closeButton.addEventListener("click", async () => {
        displayer.close();
        cleanInputs();
    });

    function cleanInputs() {
        ipAddress.value = "";
        userName.value = "";
        password.value = "";
    }
}