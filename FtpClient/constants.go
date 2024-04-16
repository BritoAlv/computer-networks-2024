package main


const max_size = 32768
const timeout = 20
const no_size = -1

const append_flag = "--append"
const unique_flag = "--unique"
const binary_flag = "--binary"
const ascii_flag = "--ascii"

var FTPErrorMessages = map[string][]string{
	// GO says that errors should not be capitalized
	"110": {"Restart marker reply.", ""},
	"120": {"Service ready in nnn minutes.", ""},
	"125": {"Data connection already open; transfer starting.", ""},
	"150": {"File status okay; about to open data connection.", ""},
	"200": {"Command okay.", ""},
	"202": {"Command not implemented, superfluous at this site.", ""},
	"211": {"System status, or system help reply.", ""},
	"212": {"Directory status.", ""},
	"213": {"File status.", ""},
	"214": {"Help message.", ""},
	"215": {"NAME system type.", ""},
	"220": {"Service ready for new user.", ""},
	"221": {"Service closing control connection.", ""},
	"225": {"Data connection open; no transfer in progress.", ""},
	"226": {"Closing data connection.", ""},
	"227": {"Entering Passive Mode (h1,h2,h3,h4,p1,p2).", ""},
	"230": {"User logged in, proceed.", ""},
	"250": {"Requested file action okay, completed.", ""},
	"257": { "Path was created.", ""},
	"331": {"User name okay, need password.", ""},
	"332": {"Need account for login.", ""},
	"354": {"Start mail input; end with <CRLF>.<CRLF>", ""},
	"350": {"Requested file action pending further information.", ""},
	"421": {"", "Service not available, closing control connection."},
	"425": {"", "Can't open data connection."},
	"426": {"", "Connection closed; transfer aborted."},
	"450": {"", "Requested file action not taken: file unavailable."},
	"451": {"", "Requested action aborted: local error in processing."},
	"452": {"", "Requested action not taken: insufficient storage space in system."},
	"500": {"", "Syntax error, command unrecognized."},
	"501": {"", "Syntax error in parameters or arguments."},
	"502": {"", "Command not implemented."},
	"503": {"", "Bad sequence of commands."},
	"504": {"", "Command not implemented for that parameter."},
	"530": {"", "Not logged in."},
	"532": {"", "Need account for storing files."},
	"550": {"", "Requested action not taken: file unavailable."},
	"551": {"", "Requested action aborted: page type unknown."},
	"552": {"", "Requested file action aborted: exceeded storage allocation."},
	"553": {"", "Requested action not taken: file name not allowed."},
	"554": {"", "Transaction failed."},
}