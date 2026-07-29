package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	txt, err := readFile("test.txt")
	if err != nil {
		fmt.Println("failed to open file")
		return
	}
	fmt.Println("file text: ", txt)

	fmt.Println(test())
}

// функция, которая читает содержимое файла,
// если он не больше 1024 байтов,
// а так же выводит дату последнего изменения фалйа
func readFile(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", err
	}

	// вызываем file.Close() через ключевое слово defer,
	// чтобы оно вызывалось после окончания функции.
	// P.S.: закрывать файлы после их использования очень важно
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		// так как мы вызвали file.Close() через defer
		// нам не нужно вызывать его при каждом return-e
		return "", err
	}

	if len(bytes) > 1024 {
		return "", errors.New("file too big")
	}

	info, err := file.Stat()
	if err != nil {
		return "", err
	}

	fmt.Println(info.ModTime())

	return string(bytes), nil
}

func test() (s string) {
	defer func() {
		s = "123"
	}()
	s = "abc"
	return
}
