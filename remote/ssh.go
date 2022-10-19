package remote

import (
	"errors"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"time"
)

type SShInfo struct {
	Host        string
	User        string
	Password    string
	LoginMethod string
	KeyPath     string
	Port        string
	Icommand    string
	Icommands   []string
}

func (s *SShInfo) One() ([]string, error) {
	var result []string

	config := &ssh.ClientConfig{
		Timeout:         time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            s.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if s.LoginMethod == "passwd" {
		config.Auth = []ssh.AuthMethod{ssh.Password(s.Password)}
	} else {
		config.Auth = []ssh.AuthMethod{StringKey(s.KeyPath)}
	}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)

	fmt.Println(addr, config)

	remote, err := ssh.Dial("tcp", addr, config)

	defer remote.Close()

	if err != nil {
		fmt.Println("创建ssh client 失败", err)
		return result, err
	} else {
		//创建ssh-session
		session, err := remote.NewSession()
		if err != nil {
			fmt.Println("创建ssh session 失败", err)
			return result, err
		}

		defer session.Close()

		//执行远程命令
		combo, err := session.CombinedOutput(s.Icommand)
		if err != nil {
			fmt.Println("远程执行cmd 失败", err)
			return result, err
		}

		result = append(result, string(combo))
		return result, nil
	}
	return result, errors.New("无数据")
}

func ShhRun(user, password, addr, keyPath, loginMethod, command string) (string, error) {
	config := &ssh.ClientConfig{
		Timeout:         time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}

	if loginMethod == "passwd" {
		config.Auth = []ssh.AuthMethod{ssh.Password(password)}
	} else {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(keyPath)}
	}

	remote, err := ssh.Dial("tcp", addr, config)

	defer remote.Close()

	if err != nil {
		fmt.Println("创建ssh client 失败", err)
		return "", err
	} else {
		//创建ssh-session
		session, err := remote.NewSession()
		if err != nil {
			fmt.Println("创建ssh session 失败", err)
			return "", err
		}

		defer session.Close()

		//执行远程命令
		combo, err := session.CombinedOutput(command)
		if err != nil {
			fmt.Println("远程执行cmd 失败", err)
			return "", err
		}
		return string(combo), nil
	}
	return "", errors.New("无数据")
}

func (s *SShInfo) Commands() ([]string, error) {

	var result []string

	config := &ssh.ClientConfig{
		Timeout:         time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            s.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}

	if s.LoginMethod == "passwd" {
		config.Auth = []ssh.AuthMethod{ssh.Password(s.Password)}
	} else {
		config.Auth = []ssh.AuthMethod{StringKey(s.KeyPath)}
	}

	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)

	remote, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		fmt.Println("创建ssh client 失败", err)
		return result, err
	}

	defer remote.Close()

	if len(s.Icommands) > 0 {
		for _, v := range s.Icommands {
			//创建ssh-session
			session, err := remote.NewSession()
			if err != nil {
				fmt.Println("创建ssh session 失败", err)
			}

			defer session.Close()
			//执行远程命令
			combo, err := session.CombinedOutput(v)
			if err != nil {
				fmt.Println("远程执行cmd 失败", err)
			}
			result = append(result, string(combo))
		}
	}
	return result, nil
}


//路径传入
func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		fmt.Println("find key's home dir failed", err)
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		fmt.Println("ssh key file read failed", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		fmt.Println("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}

//字符串传入
func StringKey(key string) ssh.AuthMethod {
	keyByte := []byte(key)
	signer, err := ssh.ParsePrivateKey(keyByte)
	if err != nil {
		fmt.Println("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}