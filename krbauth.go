package krbauth

import "os/exec"
import "strings"


func krbauth(username string, password string) bool {
    cmd := exec.Command("kinit", username)
    cmd.Stdin = strings.NewReader(password)
    err := cmd.Run()

    if err != nil {
        return false
    } else {
    	return true
    }
}
