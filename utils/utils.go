package utils

import (
	"fmt"
	"os/exec"
)

func CloudScraperGet(url string) string {
	cmd := `import cloudscraper
scraper = cloudscraper.create_scraper()
print(scraper.get("%s").text)`

	output, _ := exec.Command("python", "-c", fmt.Sprintf(cmd, url)).Output()

	for {
		if output[0] != '<' {
			break
		}
		output, _ = exec.Command("python", "-c", fmt.Sprintf(cmd, url)).Output()
	}

	return string(output)
}
