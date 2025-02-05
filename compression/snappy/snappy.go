package snappy

import (
	"fmt"
	"log"
	"os"

	"github.com/golang/snappy"
)

func Demo() {

	// 原始数据
	data := []byte("Hello, World!")

	// 压缩数据
	compressed := snappy.Encode(nil, data)
	fmt.Printf("Compressed data size: %d bytes\n", len(compressed))

	err := os.WriteFile("compressed.snappy", compressed, 0644)
	if err != nil {
		log.Fatalf("Failed to write compressed file: %v", err)
	}
	fmt.Println("Compressed data written to 'compressed.snappy'")

	// 从文件读取压缩数据
	readCompressed, err := os.ReadFile("compressed.snappy")
	if err != nil {
		log.Fatalf("Failed to read compressed file: %v", err)
	}

	// 解压数据
	decompressed, err := snappy.Decode(nil, readCompressed)
	if err != nil {
		log.Fatalf("Failed to decompress data: %v", err)
	}
	fmt.Printf("Decompressed data: %s\n", string(decompressed))

}
