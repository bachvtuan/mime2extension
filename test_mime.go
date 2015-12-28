package mime2extension 

import (
  "fmt"
)


func Test() {
  //fmt.Println(mime2extension.List_items)
  fmt.Println(Lookup(".mp4") )                       // => "video/mp4" 
  fmt.Println(Lookup("/path/to/file.txt"));          // => "text/plain" 
  fmt.Println(Lookup("file.txt"));                   // => "text/plain" 
  fmt.Println(Lookup(".TXT"));                       // => "text/plain" 
  fmt.Println(Lookup("htm"));                        // => "text/html" 
  fmt.Println(Lookup("foo"));                        // => error, not found

  fmt.Println(Extension("video/mp4"));               // => "mp4" 
  fmt.Println(Extension("image/jpeg"));              // => "jpeg" 
  fmt.Println(Extension("foo"));                     // => err, not found

  fmt.Println(Extensions("video/mp4"));              // => [mp4 mp4v mpg4]
  fmt.Println(Extensions("image/jpeg"));             // [jpeg jpe jpg]
  fmt.Println(Extensions("foo"));                    // => [] empty
}