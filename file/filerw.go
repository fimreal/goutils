package file

import (
	"bufio"
	"io"
	"os"
)

// os.ReadFile()
func GetFileContentByte(fileName string) ([]byte, error) {
	return os.ReadFile(fileName)
}

// os.WriteFile()
func WriteToFile(fileName string, b []byte) error {
	return os.WriteFile(fileName, b, 0644)
}

// 追加写入文件，如果不存在则创建文件
func AppendToFile(filename string, b []byte) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(b)
	return err
}

/*
	简单复制文件，如果目标文件不存在，则创建，权限 0644

返回复制字节数，和 error
*/
func CopyFile(srcFilePath string, dstFilePath string) (int64, error) {

	srcFile, _ := os.OpenFile(srcFilePath, os.O_RDONLY, 0644)
	dstFile, _ := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE, 0644)
	defer func() {
		srcFile.Close()
		dstFile.Close()
	}()
	return io.Copy(dstFile, srcFile)
}

func CopyFileBufio(srcFilePath string, dstFilePath string) error {
	srcFile, _ := os.OpenFile(srcFilePath, os.O_RDONLY, 0644)
	dstFile, _ := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE, 0644)
	defer func() {
		srcFile.Close()
		dstFile.Close()
	}()

	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(dstFile)
	buffer := make([]byte, 4096)

	for {
		_, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				// fmt.Println("文件读取完毕")
				break
			}
			return err
		}
		_, err = writer.Write(buffer)
		if err != nil {
			return err
		}
		writer.Flush()
	}
	return nil
}
