export class CreateConnectionRequest {
    constructor(ipAddress, userName, password) {
        this.ipAddress = ipAddress;
        this.userName = userName;
        this.password = password;
    }
}

export class ListServerRequest {
    constructor(path) {
        this.path = path
    }
}

export class TransferRequest {
    constructor(source, destination) {
        this.source = source;
        this.destination = destination;
    }
};