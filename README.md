# mime2extension

The golang library handle convertion between mime and extension.


# Usage

### Get library

```
 go get github.com/bachvtuan/mime2extension
```

### Import library to your code

```
import "github.com/bachvtuan/mime2extension"

```


### Lookup( file_path ) [ error, string ]


Return mimetype for file_path

```
fmt.Println(mime2extension.Lookup(".mp4") )                       // => "video/mp4" 
fmt.Println(mime2extension.Lookup("/path/to/file.txt"));          // => "text/plain" 
fmt.Println(mime2extension.Lookup("file.txt"));                   // => "text/plain" 
fmt.Println(mime2extension.Lookup(".TXT"));                       // => "text/plain" 
fmt.Println(mime2extension.Lookup("htm"));                        // => "text/html" 
fmt.Println(mime2extension.Lookup("foo"));                        // => error, not found
```

### Extension( mimetype ) [ error, string ] 

Return first match file extension for mimetype is given

```
fmt.Println(mime2extension.Extension("video/mp4"));               // => "mp4" 
fmt.Println(mime2extension.Extension("image/jpeg"));              // => "jpeg" 
fmt.Println(mime2extension.Extension("foo"));                     // => err, not found
```

### Extensions( mimetype ) []string

Return the array of all file extensions associate with mimetype
One mimetype can be belong in multiple extension, so this is may useful for some purposes

```
fmt.Println(mime2extension.Extensions("video/mp4"));              // => [mp4 mp4v mpg4]
fmt.Println(mime2extension.Extensions("image/jpeg"));             // [jpeg jpe jpg]
fmt.Println(mime2extension.Extensions("foo"));                    // => [] empty
```

## TEST AND FULL EXAMPLES

Check file test_mime.go 


## LICENSE

MIT