import { apiUrl } from "../constants.js";
import { get } from "../communication/httpMethods.js";

export function connectionStatusComponent() {
    setInterval(updateStatus, 2000);
}

async function updateStatus() {
    const response = await get(apiUrl + "status");

    // TODO: Store the response status into the status list
    // statusList.Push(updates)
    // Refresh status

    console.log(response)
}