package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

/*
运行sh命令，实质就是运行/bin/sh程序
linux运行程序原理：
1.clone()先创建一个子进程
2.在子进程中，系统调用execv("/bin/sh",[])，加载可执行文件到内存，然后执行程序指令。
*/

func createProc() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// 隔离 uts,ipc,pid,mount,user,network
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUSER |
			syscall.CLONE_NEWNET,
		// 设置容器的UID和GID
		UidMappings: []syscall.SysProcIDMap{
			{
				// 容器的UID
				ContainerID: 1,
				// 宿主机的UID
				HostID: 0,
				Size:   1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				// 容器的GID
				ContainerID: 1,
				// 宿主机的GID
				HostID: 0,
				Size:   1,
			},
		},
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
