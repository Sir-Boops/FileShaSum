package main

import "os"
import "io"
//import "fmt"
import "strings"
import "io/ioutil"
import "encoding/hex"
import "encoding/json"
import "crypto/sha256"
import "path/filepath"

type FileSum struct {
  FileName string `json:"name"`
  FileSum string `json:"sha256sum"`
}

func main() {

  var files []FileSum

  filepath.Walk(os.Args[1], func(path string, f os.FileInfo, err error) error {
    if !f.Mode().IsDir() {
      //fmt.Fprintf(os.Stdout, "%s", file)
      files = append(files, FileSum{
        FileName: strings.ReplaceAll(strings.TrimPrefix(path, os.Args[1] + "\\"), "\\", "/"),
        FileSum: Sha256Sum(path)})
    }
    return nil
  })

  encjson, _ := json.Marshal(files)
  loc, _ := os.Getwd()
  WriteFile(encjson, loc + "/json.json")
}


func Sha256Sum(Path string) (string) {
  f, _ := os.Open(Path)
  defer f.Close()
  h := sha256.New()
  io.Copy(h, f)
  return hex.EncodeToString(h.Sum(nil))
}

func WriteFile(Data []byte, Path string) {
  // Deletes the file before writing to it
  if CheckForFile(Path) {
    os.Remove(Path)
  }
  ioutil.WriteFile(Path, Data, os.ModePerm)
}

func CheckForFile(Path string) (bool) {
  // True if found
  // False if not
  ans := false
  if _, err := os.Stat(Path); err == nil {
    // File/Folder found!
    ans = true
  }
  return ans
}
