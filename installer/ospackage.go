package installer
type Ospackage struct {
	Name            string
	Install_method  string
func (p Ospackage) Install() error {
	if IsInstalled(p.Name) {
		return nil
	}

	switch p.Install_method {
	case "winget":
		return exec.Command("winget", "install", "--id", p.Name, "-e", "--accept-package-agreements", "--accept-source-agreements").Run()
	case "apt":
		return exec.Command("sh", "-c", "sudo apt-get install "+p.Name+" -y").Run()
	case "pacman":
		return exec.Command("sh", "-c", "sudo pacman -S "+p.Name+" --noconfirm").Run()
	}

	return nil
}
}
