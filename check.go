package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"runtime"
)

var is_64 bool
var pwd string
var ddsn string

func check() bool {

	fmt.Println("your OS is", runtime.GOARCH, runtime.GOOS)

	switch runtime.GOARCH {
	case "amd64":
		is_64 = true
		ddsn = "/Server/bin64/dontstarve_dedicated_server_nullrenderer_x64"
	case "386":
		is_64 = false
		ddsn = "/Server/bin/dontstarve_dedicated_server_nullrenderer"
	default:
		fmt.Println("not support your arch")
		os.Exit(1)
	}

	pwd, _ = os.Getwd()
	// 这里不要 err 不然会出问题
	// fmt.Println(pwd)

	user_now, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return false
	}
	if user_now.Username == "root" {
		fmt.Println("please don't exec it in root user")
		return false
	}

	if _, err := exec.LookPath("wget"); err != nil {
		fmt.Println("需要依赖：wget")
		return false
	}
	if _, err := exec.LookPath("lua"); err != nil {
		fmt.Println("需要依赖：lua")
		return false
	}

	if !has_steamcmd() {
		fmt.Println("需要前半本体，开始下载：")
		if err := download_steamcmd(); err != nil {
			fmt.Println(err)
			return false
		}

		fmt.Println("解压：")
		tar := exec.Command(
			"tar", "-xvzf",
			"Steam/steamcmd_linux.tar.gz",
			"--directory", "Steam",
		)
		bash(tar)
	}

	if !has_ldd(pwd + "/Steam/linux32/steamcmd") {
		return false
	}

	if !has_server_bin() {
		fmt.Println("安装下载后半本体：")
		if err := install_server(); err != nil {
			fmt.Println(err)
			return false
		}
		if err := install_server(); err != nil {
			fmt.Println(err)
			return false
		}
	}

	// 检查 bin 的依赖
	if !has_ldd(pwd + ddsn) {
		return false
	}

	return true

}

func has_steamcmd() bool {
	_, err1 := os.Stat("/Steam/steamcmd.sh")
	check1 := errors.Is(err1, os.ErrNotExist)
	_, err2 := os.Stat("/Steam/steamcmd_linux.tar.gz")
	check2 := errors.Is(err2, os.ErrNotExist)
	return check1 && check2
}

func download_steamcmd() error {
	if err := os.MkdirAll("Steam", os.ModePerm); err != nil {
		return err
	}
	download := exec.Command(
		"wget", "-q", "--show-progress", "--progress=bar:force",
		"--output-document", "Steam/steamcmd_linux.tar.gz",
		"https://steamcdn-a.akamaihd.net/client/installer/steamcmd_linux.tar.gz",
	)
	if err := bash(download); err != nil {
		return err
	}
	return nil
}
func has_server_bin() bool {
	_, err := os.Stat(ddsn)
	return errors.Is(err, os.ErrNotExist)
}

func install_server() error {
	if err := os.MkdirAll("Server", os.ModePerm); err != nil {
		return err
	}
	install := exec.Command(
		// ~/Steam/steamcmd.sh +force_install_dir $DST_DIR +login anonymous +app_update 343050 validate +quit
		"bash", "Steam/steamcmd.sh",
		"+force_install_dir", pwd+"/Server", // 绝对路径
		"login", "anonymous", // 匿名登录
		"+app_update", "343050", // 343050 是饥荒在steam 中的 id
		"validate", "+quit",
	)
	if err := bash2(install); err != nil {
		return err
	}
	return nil
}

func has_ldd(exe string) bool {
	// fmt.Println(pwd + ddsn)
	bin_ldd1 := exec.Command("ldd", exe)
	bin_ldd2 := exec.Command("grep", "not")
	// bash3(bin_ldd1, bin_ldd2)
	if err := bash3(bin_ldd1, bin_ldd2); err != nil {
		fmt.Println(err)
		return false
	}
	// 这里出现 not found 了吗？增加一个用户判断
	return true
}
