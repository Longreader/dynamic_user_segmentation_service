package repository

import (
	"crypto/rand"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var setPull = New()

type AuditStorage struct {
	path string
}

func NewAuditStorage() *AuditStorage {
	return &AuditStorage{path: "./storage"}
}

func (a *AuditStorage) SendAuditInformation(date string) (string, error) {

	dir := fmt.Sprintf("./storage/%s/", date)

	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		logrus.Errorf("Error: %v\n", err)
		return "", err
	}
	hashName := hex.EncodeToString(b)

	for isIn := setPull.Get(hashName); isIn; {
		b := make([]byte, 6)
		_, err := rand.Read(b)
		if err != nil {
			logrus.Errorf("Error: %v\n", err)
			return "", err
		}
		hashName = hex.EncodeToString(b)
	}
	setPull.Set(hashName)
	defer setPull.Delete(hashName)

	// Создаем новый файл CSV для записи данных
	csvFile, err := os.Create(fmt.Sprintf("./storage/csv_container/%s.csv", hashName))
	if err != nil {
		logrus.Fatal("Error occured due creating CSV file:", err)
	}

	defer func() {
		if err = csvFile.Close(); err != nil {
			logrus.Errorf("Error occured closing CSV file:", err)
		}
	}()

	// Создаем писатель для записи в CSV файл
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Получаем список файлов .csv в указанной директории
	fileList, err := filepath.Glob(filepath.Join(dir, "*.csv"))
	if err != nil {
		logrus.Fatal("Error due getting files:", err)
	}

	err = writer.Write([]string{"id;segment;operation;time;"})

	if err != nil {
		logrus.Error("Error due work with file", err)
		return "", err
	}

	// Проходимся по каждому файлу
	for _, file := range fileList {
		// Открываем каждый файл для чтения
		fileHandle, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		// Читаем строки из файла и записываем их в объединенный CSV-файл
		reader := csv.NewReader(fileHandle)
		for {
			line, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				logrus.Errorf("Error during line read:", err)

				return "", err
			}

			err = writer.Write(line)
			if err != nil {
				logrus.Error("Error due work with file", err)

				return "", err
			}
		}

		err = fileHandle.Close()
		if err != nil {
			logrus.Errorf("Error due closing file:", err)

			return "", err
		}
	}

	logrus.Debug("Audit confirmed")
	return hashName, nil
}
