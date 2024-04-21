export class ConnectRequest {
    constructor(ipAddress, port, userName, password) {
        this.IpAddress = ipAddress;
        this.Port = port;
        this.UserName = userName;
        this.Password = password;
    }
}

export class PathRequest {
    constructor(path) {
        this.Path = path
    }
}

export class TransferRequest {
    constructor(source, destination) {
        this.Source = source;
        this.Destination = destination;
    }
};

export class RenameRequest {
    constructor(path, name) {
        this.Path = path;
        this.Name = name;
    }
}