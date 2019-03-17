package xx

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func ExecCommand(strCommand string)(string){
	//cmd := exec.Command("/bin/bash", "-c", strCommand)
	cmd:= exec.Command(strCommand)


	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil{
		fmt.Println("Execute failed when Start:" + err.Error())
		return ""
	}

	out_bytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return ""
	}
	return string(out_bytes)
}