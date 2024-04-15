export class CreateConnectionRequest {
    constructor(ipAddress, userName, password) {
        this.ipAddress = ipAddress;
        this.userName = userName;
        this.password = password;
    }
}