package factory

import (
	"os"
	"os/exec"

	"github.com/Sirupsen/logrus"
)

type Runc struct {
	name string
}

func (this *Runc) SetRT(runtime string) {
	this.name = "runc"
}

func (this *Runc) GetRT() string {
	return "runc"
}

func (this *Runc) StartRT(specDir string) error {
	logrus.Debugf("Launcing runtime")

	cmd := exec.Command("runc", "start")
	cmd.Dir = specDir
	cmd.Stdin = os.Stdin
	out, err := cmd.CombinedOutput()
	logrus.Debugf("Command done")
	if string(out) != "" {
		logrus.Printf("container output=%s\n", out)
	} else {
		logrus.Debugf("container output= nil\n")
	}
	if err != nil {
		return err
	}
	return nil
}

func (this *Runc) StopRT() error {
	return nil
}
