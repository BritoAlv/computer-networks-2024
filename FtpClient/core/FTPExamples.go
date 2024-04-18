package core

type FTPExample struct{
	Ip string
	Port string
	User string
	Password string
}

var DLP FTPExample = FTPExample{
	Ip: "44.241.66.173",
	Port: "21",
	User: "dlpuser",
	Password: "rNrKYTX9g7z3RgJRmxWuGHbeu",
}

var Local FTPExample = FTPExample{
	Ip : "127.0.0.1",
	Port : "21",
	User : "brito",
	Password : "password",
}

var Phone FTPExample = FTPExample{
	Ip : "192.168.43.252",
	Port : "2020",
	User : "android",
	Password : "android",
}

var Rebex FTPExample = FTPExample{
	Ip : "194.108.117.16",
	Port : "21",
	User : "demo",
	Password : "password",
}

var Scene FTPExample = FTPExample{
	Ip : "145.24.145.107",
	Port : "21",
	User : "ftp",
	Password: "email@example.com",
}

