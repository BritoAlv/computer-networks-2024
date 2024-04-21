package api

type FileTransferRequest struct {
	Source string `json:"source"`
	Destination string  `json:"destination"`
}

type FileRename struct {
	Path string `json:"path"`
	Name string  `json:"name"`
}

type ConnectRequest struct {
	IpAddress string `json:"ipAddress"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
	Port      string `json:"port"`
}

type ResponseOperation struct {
	Status    string `json:"status"`
	Succesful bool   `json:"successful"`
}

type ResponseStatus struct {
	Status string `json:"status"`
}

type PathRequest struct {
	Path string `json:"path"`
}

type ListResponse struct {
	Directories []string `json:"directories"`
	Files       []string `json:"files"`
	Successful   bool     `json:"successful"`
}