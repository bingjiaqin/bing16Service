package git

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/spf13/viper"
)

func CommitAndPush(commitInfo string) error {
	workspace := viper.GetString("project.blog_root")
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = workspace

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("exec git add error: ", err)
		fmt.Println(string(output))
		return errors.New("exec git add error")
	}

	fmt.Println(string(output))

	cmd1 := exec.Command("git", "commit", "-a", "-m", commitInfo)
	cmd1.Dir = workspace
	output1, err1 := cmd1.CombinedOutput()

	if err1 != nil {
		fmt.Println("exec git commit error: ", err1)
		fmt.Println(string(output1))
		return errors.New("exec git commit error")
	}

	fmt.Println(string(output1))

	cmd2 := exec.Command("git", "push", "origin", "main")
	cmd2.Dir = workspace
	output2, err2 := cmd2.CombinedOutput()

	if err2 != nil {
		fmt.Println("exec git push error: ", err2)
		fmt.Println(string(output2))
		return errors.New("exec git push error")
	}

	return nil
}
