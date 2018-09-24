package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

type Config map[string]string

func ReadConfig(filename string) (Config, error) {
    // init with some bogus data
    config := Config{
        "nwsHexKey":     "padraoa8773564ebc8f7abdcaac6bd2137dd07",
        "appHexKey":     "padrao1e745697990164f22531bf11e4614ad1",
        "devHexAddr":    "padrao018a6355",
	"broker":        "padraotcp://localhost:1884",
	"username":       "padrao",
	"password":       "padrao",	

    }
    if len(filename) == 0 {
        return config, nil
    }
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := bufio.NewReader(file)

    for {
        line, err := reader.ReadString('\n')

        // check if the line has = sign
        // and process the line. Ignore the rest.
        if equal := strings.Index(line, "="); equal >= 0 {
            if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
                value := ""
                if len(line) > equal {
                    value = strings.TrimSpace(line[equal+1:])
                }
                // assign the config map
                config[key] = value
            }
        }
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }
    }
    return config, nil
}

func main() {
    // for this tutorial, we will hard code it to config.txt
    config, err := ReadConfig(`/home/docker/go/src/github.com/willianxz/golang/data.txt`)

    if err != nil {
        fmt.Println(err)
    }
   
   
    
    // assign values from config file to variables
    nwsHexKey := config["nwsHexKey"]
    appHexKey := config["appHexKey"]
    devHexAddr := config["devHexAddr"]
    broker := config["broker"]
    username := config["username"]
    password := config["password"]

    fmt.Println("NwsHexKey :", nwsHexKey)
    fmt.Println("AppHexKey :", appHexKey)
    fmt.Println("DevHexAddr :", devHexAddr)
    fmt.Println("Broker :", broker)
    fmt.Println("Username :", username)
    fmt.Println("Password :", password)
}
