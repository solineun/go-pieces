package main

import "fmt"

func main() {
	f := Folder{
		Name: "root",
		Folders: []Folder{
			Folder{
				Name: "child1",
				Files: []File{
					File{Name: "1",},
					File{Name: "2"},
				},
			},
			Folder{
				Name: "child2",
				Files: []File{
					File{Name: "1"},
					File{Name: "2"},
				},
			},
		},
	}
	printFolderTree(f, " ")
}

type Folder struct {
	Name string
	Folders []Folder
	Files []File
}

type File struct {
	Name string
	Data []byte
}

func printFolderTree(root Folder, t string) {
	fmt.Println(t + root.Name)
	t += t
	for _, f := range root.Folders {
		printFolderTree(f, t)		
	}
	for _, f := range root.Files {
		fmt.Println(t + f.Name)
	}
}