package dyndns

import (
	"errors"
	"os/exec"
	"strings"
)

// GetExternalIP retrieves the external facing IP Address.
// This Variant does still rely upon others supplying the correct ip, however google
// or other larger companies might be more trustworthy.


func GetExternalIP() (string, error) {
	cmdoutput, err := exec.Command("dig", "-4", "TXT", "+short", "o-o.myaddr.l.google.com", "@ns1.google.com").CombinedOutput()
	svariable := string(cmdoutput)
	svariable = strings.Replace(svariable, "\"", "", -1)
	if err == nil {
		return svariable, err
	}




	return "", errors.New("Could not retreive external IP")
}
