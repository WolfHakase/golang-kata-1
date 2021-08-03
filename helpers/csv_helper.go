package helpers

import (
	"encoding/csv"
	"github.com/echocat/golang-kata-1/v1/models"
	"os"
	"strings"
)

type CSVFile struct {
	Name string
	file *os.File
	reader *csv.Reader
}

func (csvFile CSVFile) openFile() error {
	openFile, err := os.Open(csvFile.Name)
	if err != nil {
		return err
	}

	csvFile.file = openFile
	return nil
}

func (csvFile CSVFile) closeFile() error {
	return csvFile.file.Close()
}

func (csvFile CSVFile) readLine() ([]string, error) {
	if csvFile.reader == nil {
		csvFile.reader = csv.NewReader(csvFile.file)
	}

	return csvFile.reader.Read()
}

func (csvFile CSVFile) GetLiteraryItem() (*models.LiteraryItem, error) {
	line, err := csvFile.readLine()
	if err != nil {
		return nil, err
	}

	item := models.LiteraryItem{}
	authors := strings.Split(line[2], ",")
	item.SetValues(line[0], line[1], authors, line[3])

	return &item, nil
}

func (csvFile CSVFile) GetLiteraryItems() (*models.LiteraryItems, error){
	err := csvFile.openFile()
	if err != nil {
		return nil, err
	}
	defer csvFile.closeFile()

	items := models.LiteraryItems{}
	for {
		item, err := csvFile.GetLiteraryItem()
		if err != nil {
			break // assume end of file, in which case no problem
		}
		items = append(items, *item)
	}

	return &items, nil
}

func (csvFile CSVFile) GetAuthor() (*models.Author, error) {
	line, err := csvFile.readLine()
	if err != nil {
		return nil, err
	}

	item := models.Author{}
	item.SetValues(line[0], line[1], line[3])

	return &item, nil
}

func (csvFile CSVFile) GetAuthors() (*models.Authors, error){
	err := csvFile.openFile()
	if err != nil {
		return nil, err
	}
	defer csvFile.closeFile()

	authors := models.Authors{}
	for {
		author, err := csvFile.GetAuthor()
		if err != nil {
			break // assume end of file, in which case no problem
		}
		authors = append(authors, *author)
	}

	return &authors, nil
}
