package main

import (
	"errors"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const max_size = 1024

func starts_with(command, prefix string) bool {
	// Function to check if a string starts with a prefix
	return len(command) >= len(prefix) && command[:len(prefix)] == prefix
}

func is_valid_ip(ip string) bool {
	// Function to check if a string is a valid IP
	// This is a very simple check, it only checks if the string has 3 dots and 4 numbers
	numbers := strings.Split(ip, ".")
	if len(numbers) != 4 {
		return false
	}
	for _, number := range numbers {
		value, err := strconv.Atoi(number)
		if err != nil {
			return false
		}
		if value < 0 || value > 255 {
			return false
		}
	}
	return true
}

func open_conection(connDataConfig string) (net.Conn, error) {
	connDataIP, connDataPort := parse_get_connection_ftp(connDataConfig)
	connData, err := net.Dial("tcp", connDataIP+":"+connDataPort)
	if err != nil {
		return nil, err
	}
	return connData, nil
}

func parse_get_connection_ftp(input string) (string, string) {
	// This method is crying for refactoring.
	// Parse the input string to get the IP and Port
	// Example: "227 Entering Passive Mode (127,0,0,1,195,149)"
	// The port is calculated as (195*256+149)
	// Output: "127.0.0.1", "195*256+149".
	split1 := strings.Split(input, "(")
	split2 := strings.Split(split1[1], ",")
	split3 := strings.Split(split2[5], ")")
	ip := split2[0] + "." + split2[1] + "." + split2[2] + "." + split2[3]
	first_part_port, _ := strconv.ParseInt(split2[4], 10, 32)
	second_part_port, _ := strconv.ParseInt(split3[0], 10, 32)

	port := strconv.FormatInt((first_part_port*256 + second_part_port), 10)
	return ip, port
}

func writeonMemory(connData *net.Conn, data []byte) (int, error) {
	// I wll read from server in chunks of size 1024 KBs.
	pr := 0
	for {
		(*connData).SetWriteDeadline(time.Now().Add(60 * time.Second))
		size_to_write := min(max_size, len(data)-pr)
		n, err := (*connData).Write(data[pr : pr+size_to_write])
		if err != nil {
			return pr, nil
		}
		if n != size_to_write {
			return n, errors.New("not everything was written")
		}
		pr += n
		if pr == len(data) {
			break
		}
	}
	return pr, nil
}

func readOnMemory(connData *net.Conn) ([]byte, error) {
	// I wll read from server in chunks of size 1024 KBs.
	tmp := make([]byte, max_size)
	data := make([]byte, 0)
	for {
		(*connData).SetReadDeadline(time.Now().Add(60 * time.Second))
		n, err := (*connData).Read(tmp)
		if err != nil {
			if err != io.EOF {
				r := make([]byte, 0)
				return r, err
			}
			break
		}
		data = append(data, tmp[:n]...)
		if n < max_size {
			break
		}
	}
	return data, nil
}

func writeAndreadOnMemory(connData *net.Conn, data []byte) ([]byte, error) {
	_, err := writeonMemory(connData, data)
	if err != nil {
		return nil, err
	}
	return readOnMemory(connData)
}

func readOnFile(connData *net.Conn, file *os.File, size int64) error {
	// I wll read from server in chunks of size 1024 KBs.
	offset := int64(0)
	tmp := make([]byte, max_size)
	for ; offset != size ; {
		(*connData).SetReadDeadline(time.Now().Add(60 * time.Second))
		n, err := (*connData).Read(tmp)
		if err != nil {
			continue
		}
		file.WriteAt(tmp[:n], offset)
		offset += int64(n)
	}
	return nil
}

var FTPErrorMessages = map[string][]string{
	// GO says that errors should not be capitalized 
	"200": {"Command okay.", ""},
	"500": {"", "Syntax error, command unrecognized."},
	"501": {"", "Syntax error in parameters or arguments."},
	"202": {"Command not implemented, superfluous at this site.", ""},
	"502": {"", "Command not implemented."},
	"503": {"", "Bad sequence of commands."},
	"504": {"", "Command not implemented for that parameter."},
	"110": {"Restart marker reply.", ""},
	"211": {"System status, or system help reply.", ""},
	"212": {"Directory status.", ""},
	"213": {"File status.", ""},
	"214": {"Help message.", ""},
	"215": {"NAME system type.", ""},
	"120": {"Service ready in nnn minutes.", ""},
	"220": {"Service ready for new user.", ""},
	"221": {"Service closing control connection.", ""},
	"421": {"", "Service not available, closing control connection."},
	"125": {"Data connection already open; transfer starting.", ""},
	"225": {"Data connection open; no transfer in progress.", ""},
	"425": {"", "Can't open data connection."},
	"226": {"Closing data connection.", ""},
	"426": {"", "Connection closed; transfer aborted."},
	"227": {"Entering Passive Mode (h1,h2,h3,h4,p1,p2).", ""},
	"230": {"User logged in, proceed.", ""},
	"530": {"", "Not logged in."},
	"331": {"User name okay, need password.", ""},
	"332": {"Need account for login.", ""},
	"532": {"", "Need account for storing files."},
	"150": {"File status okay; about to open data connection.", ""},
	"250": {"Requested file action okay, completed.", ""},
	"257": {"\"PATHNAME\" created.", ""},
	"350": {"Requested file action pending further information.", ""},
	"450": {"", "Requested file action not taken: file unavailable."},
	"550": {"", "Requested action not taken: file unavailable."},
	"451": {"", "Requested action aborted: local error in processing."},
	"551": {"", "Requested action aborted: page type unknown."},
	"452": {"", "Requested action not taken: insufficient storage space in system."},
	"552": {"", "Requested file action aborted: exceeded storage allocation."},
	"553": {"", "Requested action not taken: file name not allowed."},
	"354": {"Start mail input; end with <CRLF>.<CRLF>", ""},
	"554": {"", "Transaction failed."},
}

func ParseFTPCode(code string) (string, error) {
	message, exists := FTPErrorMessages[code]
	if !exists {
		return "",errors.New("codigo de error desconocido")
	}
	msg := message[0]
	if( message[1] == ""){
		return msg, nil
	}
	return msg, errors.New(message[1])
}