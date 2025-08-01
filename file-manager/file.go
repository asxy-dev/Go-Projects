package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)
func main(){
	var path string
	var category string
	fmt.Println("Enter the path to folder : ")
	fmt.Scanln(&path)

	files, err := ioutil.ReadDir(path)
	if err!= nil{
		fmt.Println("Error reading folder")
	}
	for _, file := range files{
		if file.IsDir(){
			continue
		}
		ext := strings.ToLower(filepath.Ext(file.Name()))

		switch ext{
		case ".jpg",".jpeg",".png",".gif":
			category="Images"
		case ".mp3",".wav":
			category="Music"
		case ".mov",".mp4":
			category="Videos"
		case ".txt",".pdf",".rtf",".xlxs",".doc",".pptx":
			category="Documents"
		case ".zip",".jar",".rar":
			category="Zipped"
		default:
			category="Others"
		}
		folder := filepath.Join(path, category)
		os.MkdirAll(folder, 0755)

		new_folder:= filepath.Join(path, file.Name())
		new_path := filepath.Join(folder,file.Name())

		err := os.Rename(new_folder,new_path)
		if err!= nil{
			fmt.Println("Error found")
		}else {
			fmt.Println("Moved to: ",file.Name(),"->",category)
		}
	}
	
	
}