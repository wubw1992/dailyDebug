package main

import (
    //"fmt"

	"io/ioutil"
	"os"
    "os/exec"
	"path/filepath"
   // "encoding/json"
    "strings"
    "github.com/linuxdeepin/go-lib/log"
)
 
var logger = log.NewLogger("test")

func ReadFile(file string) []byte {
    content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
    return content
}

func WriteFile(file string, data string) bool {
	_, err := os.Stat(filepath.Dir(file))
	if os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(file), 0755)
		if err != nil {
			logger.Warning(" >>> MkdirAll error : ", err)
			return false
		}
	}

    //path := filepath.Join(file, "123.txt")
	err = ioutil.WriteFile(file, []byte(data), 0755)
	if err != nil {
		logger.Warning(" >>> WriteFile error : ", err)
		return false
	}

	logger.Debugf("%s:\n%s", file, string(data))
	return true
}

func main() {
    data := "ondemand"
    logger.Warning(" >>> performance, powersave, userspace, ondemand, conservative. schedutil",)
    for idx, args:= range os.Args {
        logger.Warning(" >>> idx , arg : ",idx, args)
        if  args == "performance" || 
            args == "powersave" || 
            args == "userspace" || 
            args == "ondemand" || 
            args == "conservative" || 
            args == "schedutil" {
                data = args
                logger.Warning(" >>> data : ", args)
                break
        }
    }
    
    // WriteFile("/home/wubw/Desktop/1234.txt", "test1233456789")
    WriteFile("/sys/devices/system/cpu/cpu0/cpufreq/scaling_governor", data)

    out, err := exec.Command("cat", "/sys/devices/system/cpu/cpu0/cpufreq/scaling_governor").Output()
	if err != nil {
		logger.Warning(" >>> exec.Command error : ", err)
	} else {
        logger.Warningf("/sys/devices/system/cpu/cpu0/cpufreq/scaling_governor date is : %s\n", out)
    }

    logger.Warning(" ReadFile scaling_governor : ", ReadFile("/sys/devices/system/cpu/cpu0/cpufreq/scaling_governor"))
    // logger.Warning(" ReadFile scaling_available_governors : ", ReadFile("/sys/devices/system/cpu/cpu0/cpufreq/scaling_available_governors"))

    content := ReadFile("/sys/devices/system/cpu/cpu0/cpufreq/scaling_available_governors")
    value := strings.TrimSpace(string(content))
    lines := strings.Split(value, " ")

    logger.Warning(" >>> lines : ", lines, len(lines))
}