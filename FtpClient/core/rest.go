package core

func (cs *FtpSession) REST(marker string) (string, error) {
    return writeAndreadOnMemory(cs, "REST " + marker)
}