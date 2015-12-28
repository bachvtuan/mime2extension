package mime2extension
//build Lookup and Extension and Extensions

import (
	"path/filepath"
	//"fmt"
	"strings"
	"errors"
)

type MimeExtension struct{
	Extension string
	Mime string
}

/**
 * Given file path and it will return mime type
 * err if not found mime
 */
func Lookup( file_path string ) (error, string) {

	if file_path == ""{
		return errors.New("file_path is empty"), ""
	}

	file_path = strings.ToLower( file_path  )
	extension := filepath.Ext( file_path )

	if extension ==  "" {
		extension = file_path
	}else{
		extension = extension[1:]
	}


	for _, item := range *List_items {
		if item.Extension == extension {
			return nil, item.Mime
		}
	}

	return errors.New("Not found"), ""
}

/**
 * Given file extension when input is mimetype
 * err if not found extension
 */
func Extension( mime string ) (error, string) {

	mime = strings.ToLower( mime  )

	for _, item := range *List_items {
		if item.Mime == mime {
			return nil, item.Extension
		}
	}

	return errors.New("Not found"), ""
}

/**
 * Given list of files extension when input is mimetype
 */
func Extensions( mime string ) []string {
	
	mime = strings.ToLower( mime  )
  var return_list []string

	for _, item := range *List_items {
		if item.Mime == mime {
			return_list = append( return_list, item.Extension )
		}
	}

	return return_list
}