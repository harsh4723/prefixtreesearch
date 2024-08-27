// package main

// import (
// 	"fmt"
// 	"strings"
// )

// // ARTNode represents a node in the Adaptive Radix Tree.
// type ARTNode struct {
// 	children map[string]*ARTNode
// 	isLeaf   bool
// 	objects  []string // Store objects if this node is a leaf
// }

// // NewARTNode creates a new ARTNode.
// func NewARTNode() *ARTNode {
// 	return &ARTNode{
// 		children: make(map[string]*ARTNode),
// 		isLeaf:   false,
// 	}
// }

// // Insert inserts an object into the ART using a folder-based path.
// func (n *ARTNode) Insert(path string, object string) {
// 	folders := strings.Split(path, "/")
// 	current := n
// 	for _, folder := range folders {
// 		if folder == "" {
// 			continue
// 		}
// 		if current.children[folder] == nil {
// 			current.children[folder] = NewARTNode()
// 		}
// 		current = current.children[folder]
// 	}
// 	current.isLeaf = true
// 	current.objects = append(current.objects, object)
// }

// // ListContents lists all folders and objects under a specific folder prefix.
// func (n *ARTNode) ListContents(prefix string) []string {
// 	folders := strings.Split(prefix, "/")
// 	fmt.Println("folders ...", folders)
// 	fmt.Println("len..", len(folders))
// 	current := n
// 	for _, folder := range folders {
// 		if folder == "" {
// 			continue
// 		}
// 		if current.children[folder] == nil {
// 			return []string{}
// 		}
// 		current = current.children[folder]
// 	}
// 	fmt.Println("Harh current", current.objects)
// 	return n.collectContents(current)
// }

// // collectContents collects all folders and objects in the subtree rooted at the current node.
// func (n *ARTNode) collectContents(node *ARTNode) []string {
// 	var result []string
// 	// Collect folders
// 	for folderName, _ := range node.children {
// 		result = append(result, folderName+"/")
// 		fmt.Println("foldernamee..", folderName)
// 		// if childNode.isLeaf {
// 		// 	fmt.Println("Is leafnodeee", childNode.objects)
// 		// 	result = append(result, childNode.objects...)
// 		// }
// 	}
// 	// Collect objects at this node
// 	if node.isLeaf {
// 		result = append(result, node.objects...)
// 	}
// 	return result
// }

// func main() {
// 	root := NewARTNode()

// 	// Insert some objects
// 	root.Insert("dir1/file1.txt", "file1.txt")
// 	root.Insert("dir1/file2.txt", "file2.txt")
// 	root.Insert("dir1/subdir1/file3.txt", "file3.txt")
// 	root.Insert("dir2/file4.txt", "file4.txt")
// 	root.Insert("dir1/subdir2/", "subdir2") // Inserting a folder

// 	// List all contents (folders and objects) under "dir1/"
// 	fmt.Println("Contents in dir1/:", root.ListContents("dir1/subdir2/"))

// 	// List all contents (folders and objects) under "dir1/subdir1/"
// 	// fmt.Println("Contents in dir1/subdir1/:", root.ListContents("dir1/subdir1/"))

// 	// // List all contents (folders and objects) under "dir2/"
// 	// fmt.Println("Contents in dir2/:", root.ListContents("dir2/"))
// }

package main

import (
	"fmt"

	"github.com/arriqaaq/art"
)

func main() {
	tree := art.NewTree()

	// Insert
	// tree.Insert([]byte("dir1/sub1/text1.txt"), "file")
	// tree.Insert([]byte("dir1/sub1/text2.txt"), "file")
	// tree.Insert([]byte("dir1/sub1/text3.txt"), "file")
	// tree.Insert([]byte("dir1/sub1/sub2/text4.txt"), "file")
	// tree.Insert([]byte("dir1/sub2/text5.txt"), "file")
	// value := tree.Search([]byte("hello"))
	// fmt.Println("value=", value)

	// // Delete
	// tree.Insert([]byte("wonderful"), "life")
	// tree.Insert([]byte("foo"), "bar")
	// deleted := tree.Delete([]byte("foo"))
	// fmt.Println("deleted=", deleted)

	// // Search
	// value = tree.Search([]byte("hello"))
	// fmt.Println("value=", value)

	// // Traverse (with callback function)
	// tree.Each(func(node *art.Node) {
	// 	if node.IsLeaf() {
	// 		fmt.Println("value=", node.Value())
	// 	}
	// })

	// // Iterator
	// for it := tree.Iterator(); it.HasNext(); {
	// 	value := it.Next()
	// 	if value.IsLeaf() {
	// 		fmt.Println("value=", value.Value())
	// 	}
	// }

	// Prefix Scan
	// tree.Insert([]byte("api"), "bar")
	// tree.Insert([]byte("api.com"), "bar")
	// tree.Insert([]byte("api.com.xyz"), "bar")
	tree.Insert([]byte("dir1/sub1/text2.txt"), "file")
	tree.Insert([]byte("dir1/sub1/text1.txt"), "file")
	tree.Insert([]byte("dir1/sub1/"), "folder")
	tree.Insert([]byte("dir1/sub1/text3.txt"), "file")
	tree.Insert([]byte("dir1/sub1/sub2/text4.txt"), "file")
	tree.Insert([]byte("dir2/sub2/text5.txt"), "file")
	leafFilter := func(n *art.Node) {
		if n.IsLeaf() {
			fmt.Println("value=", string(n.Key()), n.Value())
		}
	}
	tree.Scan([]byte("dir1/"), leafFilter)
	deleted := tree.Delete([]byte("dir1/sub1/sub2/text4.txt"))
	fmt.Println("deleted=", deleted)

	tree.Scan([]byte("dir1/"), leafFilter)

}
