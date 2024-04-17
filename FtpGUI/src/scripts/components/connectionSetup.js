import { displayer } from "../globals.js";

export function connectionSetupComponent() {
    const ipAddress = document.querySelector("#ip-address");
    const port = document.querySelector("#port");
    const userName = document.querySelector("#user-name");
    const password = document.querySelector("#password");
    const closeButton = document.querySelector("#close-button");
    const connectButton = document.querySelector("#connect-button");

    // Create connection event
    connectButton.addEventListener("click", async () => {
        displayer.connect(ipAddress.value, port.value, userName.value, password.value);
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