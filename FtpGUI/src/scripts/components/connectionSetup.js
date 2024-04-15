import { apiUrl } from "../constants.js"
import { get, post } from "../communication/httpMethods.js";
import { CreateConnectionRequest } from "../communication/requests.js";

export function connectionSetupComponent() {
    const ipAddress = document.querySelector("#ip-address");
    const userName = document.querySelector("#user-name");
    const password = document.querySelector("#password");
    const closeButton = document.querySelector("#close-button");
    const connectButton = document.querySelector("#connect-button");

    // Create connection event
    let response;
    connectButton.addEventListener("click", async () => {
        const request = new CreateConnectionRequest(
            ipAddress.value,
            userName.value,
            password.value
        );
        response = await post(apiUrl + "connect", request);

        // TODO: Store the response status into the status list
        // statusList.Push(updates)
        // Refresh status

        // TODO-PROD: Erase this logging when setting code to production. Just for debugging purposes
        console.log(response);

        cleanInputs();
    });

    // Close connection event
    closeButton.addEventListener("click", async () => {
        response = await get(apiUrl + "close");

        // TODO: Store the response status into the status list
        // statusList.Push(updates)
        // Refresh status

        // TODO-PROD: Erase this logging when setting code to production. Just for debugging purposes
        console.log(response);

        cleanInputs();
    });

    function cleanInputs() {
        ipAddress.value = "";
        userName.value = "";
        password.value = "";
    }
}