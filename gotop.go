package main

import(
    "os"
    "fmt"
    "log"
    "io/ioutil"
)

func main () {
    base := "/proc"
    dirs, err := ioutil.ReadDir(base)
    if err != nil {
        log.Fatal(err)
    }
     var process_list []os.FileInfo
    for _, v := range(dirs){
        // If it is a directory and contains a file named cmdline it is a process
        if v.IsDir() {
            cmdline_name := base + "/" + v.Name() + "/cmdline"
            file, err := os.Open(cmdline_name)
            defer file.Close()
            if err == nil {
                process_list = append(process_list, v)
            }
        } else {
            fmt.Println(v.Name())
        }
    }
    for _, v := range(process_list) {
        fmt.Println(v.Name())
    }
}
