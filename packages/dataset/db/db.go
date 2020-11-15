package db

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dataset struct {
	UUID      string   "bson:'uuid,omitempty'"
	Locale    string   "bson:'locale,omitempty'"
	Multilang string   "bson:'multilang,omitempty'"
	Label     string   "bson:'label,omitempty'"
	Size      int      "bson:'size,omitempty'"
	Dimension int      "bson:'dimension,omitempty'"
	Algorithm string   "bson:'algorithm,omitempty'"
	Model     []string "bson:'model,omitempty'"
}

// Vocub ...
type Vocub struct {
	UUID string "bson:'uuid,omitempty'"
	Word string "bson:'word,omitempty'"
}

// Vectors ...
type Vectors struct {
	DatasetUUID string   "bson:'datasetuuid,omitempty'"
	WordUUID    string   "bson:'worduuid,omitempty'"
	Vectors     []string "bson:'vectors,omitempty'"
}

// NewDataset ...
func NewDataset(filename string) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()

	collection := client.Database("cs").Collection("datasets")
	vocub := client.Database("cs").Collection("vocub")
	vectors := client.Database("cs").Collection("vectors")

	genUUIDD := uuid.New()
	datasetUUID := genUUIDD.String()
	params := []string{"epoch=2", "window=5", "lr=0.001"}

	dst := dataset{datasetUUID, "ru", "yes", "tensorflow", 840, 100, "fasttext", params}

	collection.InsertOne(context.TODO(), dst)

	// Открыть файл с векторным представлением слов

	file, err := os.Open("/home/kuranov/Code/golang/service/sys/" + filename + ".vec")

	if err != nil {
		fmt.Println("Ошибка открытия файла")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	optmtext := text[1:]
	for _, line := range optmtext {
		linesplit := strings.Split(line, " ")
		lword := linesplit[:1]
		lvect := linesplit[1:]
		// Генерируем UUID для слова
		genUUIDW := uuid.New()
		wordUUID := genUUIDW.String()
		// Формируем данные для Vocub
		vcb := Vocub{wordUUID, lword[0]}
		// Формируем данные для Vectors
		vct := Vectors{datasetUUID, wordUUID, lvect}
		// Запись в БД
		vocub.InsertOne(context.TODO(), vcb)
		vectors.InsertOne(context.TODO(), vct)
	}

	//

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

}
