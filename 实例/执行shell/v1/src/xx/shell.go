package xx

import (
	"fmt"
	"os/exec"
)

func Exe_shell()  {
	cmd := exec.Command("ipconfig")  // 执行shell
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
