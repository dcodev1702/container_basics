// Author: Liz Rice w/ Aqua Security
// Site: https://www.lizrice.com/
// go run main.go run         <command> <args>
// docker         run <image> <command> <args>
package main

import (
  "fmt"
  "os"
  "os/exec"
  "syscall"
)

func main() {
   switch os.Args[1] {
     case "run":
       run()
     case "child":
       child()
     default:
       panic("I'm confused!")
   }
}

func run() {
   fmt.Printf("Running %v as UID %d as PID %d\n", os.Args[2:], os.Geteuid(), os.Getpid())

   cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
   cmd.Stdin  = os.Stdin
   cmd.Stdout = os.Stdout
   cmd.Stderr = os.Stderr
   cmd.SysProcAttr = &syscall.SysProcAttr {
       Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNS | syscall.CLONE_NEWPID,
       Credential: &syscall.Credential {Uid: 0, Gid: 0},
       UidMappings: []syscall.SysProcIDMap {
	  { ContainerID: 0, HostID: os.Getuid(), Size: 1, },
       },
       GidMappings: []syscall.SysProcIDMap {
	  { ContainerID: 0, HostID: os.Getgid(), Size: 1, },
       },
   }
   must(cmd.Run())
}

func child() {

   pwd, err := os.Getwd()
   if err != nil {
     panic(err)
   }
   fmt.Printf("Running %v as UID %d as PID %d\n", os.Args[2:], os.Geteuid(), os.Getpid())

   must(syscall.Sethostname([]byte("container-00")))
   must(syscall.Chroot(pwd + "/alpinefs"))
   must(os.Chdir("/"))
   must(syscall.Mount("proc", "proc", "proc", 0, ""))
	
   cmd := exec.Command(os.Args[2], os.Args[3:]...)
   cmd.Stdin  = os.Stdin
   cmd.Stdout = os.Stdout
   cmd.Stderr = os.Stderr
	
   must(cmd.Run())
   must(syscall.Unmount("proc", 0))
}

func must(err error) {
   if err != nil {
      panic(err)
   }
}
