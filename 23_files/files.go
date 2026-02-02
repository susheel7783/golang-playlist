package main

import (
	"bufio" // Used for buffered reading & writing (efficient IO)
	"fmt"   // Used for printing output
	"os"    // Used for file & directory operations
)

func main() {

	// ==================================================
	// 1️⃣ OPEN A FILE AND READ FILE INFORMATION
	// ==================================================

	/*
	// Open a file in read-only mode
	f, err := os.Open("example.txt")
	if err != nil {
		// panic stops the program if error occurs
		panic(err)
	}

	// Ensure file is closed when main() exits
	defer f.Close()

	// Get file metadata (name, size, permissions, etc.)
	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	// Print file information
	fmt.Println("file name:", fileInfo.Name())       // Name of the file
	fmt.Println("is directory:", fileInfo.IsDir())  // true if folder
	fmt.Println("file size:", fileInfo.Size())       // Size in bytes
	fmt.Println("permissions:", fileInfo.Mode())    // File permissions
	fmt.Println("modified at:", fileInfo.ModTime()) // Last modified time
	*/

	// ==================================================
	// 2️⃣ READ FILE USING Read() (LOW-LEVEL)
	// ==================================================

	/*
	// Open file
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Create a byte slice (buffer) of size 12
	buf := make([]byte, 12)

	// Read data into buffer
	n, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	// n = number of bytes actually read
	fmt.Println("bytes read:", n)

	// Print each byte as character
	for i := 0; i < n; i++ {
		fmt.Println("data:", string(buf[i]))
	}
	*/

	// ==================================================
	// 3️⃣ READ FILE USING os.ReadFile (EASIEST WAY)
	// ==================================================

	/*
	// Reads entire file into memory at once
	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}

	// Convert bytes to string and print
	fmt.Println(string(data))
	*/

	// ==================================================
	// 4️⃣ READ DIRECTORY CONTENTS
	// ==================================================

	/*
	// Open current directory (../ means parent directory)
	dir, err := os.Open("../")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	// Read all files and folders in directory
	entries, err := dir.ReadDir(-1)
	if err != nil {
		panic(err)
	}

	// Loop through directory entries
	for _, entry := range entries {
		fmt.Println(entry.Name(), "is directory:", entry.IsDir())
	}
	*/

	// ==================================================
	// 5️⃣ CREATE AND WRITE TO A FILE
	// ==================================================

	/*
	// Create a new file (overwrites if exists)
	f, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Write string data
	f.WriteString("hi go\n")
	f.WriteString("nice language\n")

	// Write byte slice
	bytes := []byte("Hello Golang\n")
	f.Write(bytes)

	fmt.Println("file written successfully")
	*/

	// ==================================================
	// 6️⃣ COPY FILE USING BUFFERED IO (BEST PRACTICE)
	// ==================================================

	/*
	// Open source file
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	// Create destination file
	destFile, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	// Create buffered reader & writer
	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	// Read and write byte by byte
	for {
		b, err := reader.ReadByte()

		// If end of file reached, stop loop
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		// Write byte to destination
		err = writer.WriteByte(b)
		if err != nil {
			panic(err)
		}
	}

	// Flush buffered data to disk
	writer.Flush()

	fmt.Println("file copied successfully")
	*/

	// ==================================================
	// 7️⃣ DELETE A FILE
	// ==================================================

	/*
	// Delete file permanently
	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("file deleted successfully")
	*/
}
-----------------------------------------------
os.Open → open file

defer file.Close() → avoid memory leaks

os.ReadFile → easiest way to read file

bufio → efficient for large files

os.Create → create / overwrite file

os.Remove → delete file
