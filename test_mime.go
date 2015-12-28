package main 

import (
  "github.com/bachvtuan/mime2extension"
  "fmt"
)


func main() {
  //fmt.Println(mime2extension.List_items)
  fmt.Println(mime2extension.Lookup(".mp4") )                       // => "video/mp4" 
  fmt.Println(mime2extension.Lookup("/path/to/file.txt"));          // => "text/plain" 
  fmt.Println(mime2extension.Lookup("file.txt"));                   // => "text/plain" 
  fmt.Println(mime2extension.Lookup(".TXT"));                       // => "text/plain" 
  fmt.Println(mime2extension.Lookup("htm"));                        // => "text/html" 
  fmt.Println(mime2extension.Lookup("foo"));                        // => error, not found

  fmt.Println(mime2extension.Extension("video/mp4"));               // => "mp4" 
  fmt.Println(mime2extension.Extension("image/jpeg"));              // => "jpeg" 
  fmt.Println(mime2extension.Extension("foo"));                     // => err, not found

  fmt.Println(mime2extension.Extensions("video/mp4"));              // => [mp4 mp4v mpg4]
  fmt.Println(mime2extension.Extensions("image/jpeg"));             // [jpeg jpe jpg]
  fmt.Println(mime2extension.Extensions("foo"));                    // => [] empty
}