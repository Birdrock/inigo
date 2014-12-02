package helpers

import (
	"os/exec"

	"github.com/onsi/gomega"
)

func Copy(sourcePath, destinationPath string) {
	err := exec.Command("cp", "-a", sourcePath, destinationPath).Run()
	gomega.Ω(err).ShouldNot(gomega.HaveOccurred())
}
