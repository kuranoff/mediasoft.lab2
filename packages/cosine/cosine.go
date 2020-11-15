package cosine

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"./conv"
	"./similarity"
)

type wordsuuid struct {
	UUID string "bson:'uuid,omitempty'"
	Word string "bson:'word,omitempty'"
}

type wordsvectors struct {
	Vectors []string "bson:'vectors,omitempty'"
}

type csdata struct {
	CSTitle        string
	CSWordA        string
	CSWordAUUID    string
	CSWordAVectors []string
	CSWordB        string
	CSWordBUUID    string
	CSWordBVectors []string
	CSV            float64
}

// Similarity ...
func Similarity(w http.ResponseWriter, r *http.Request) {

	wA := r.FormValue("wordA")
	wB := r.FormValue("wordB")

	// Подключение к MongoDB
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

	// Выбор коллекций Vocub и Vectors
	vocub := client.Database("cs").Collection("vocub")
	vectors := client.Database("cs").Collection("vectors")

	// Параметры поиска UUID слов wA и wB {filterf{ind}w{ords}}
	filterfw := bson.M{"word": bson.D{{"$in", bson.A{wA, wB}}}}

	var reswords []*wordsuuid

	// Поиск UUID слов wA и wB
	curw, err := vocub.Find(context.TODO(), filterfw)
	if err != nil {
		log.Fatal(err)
	}

	for curw.Next(context.TODO()) {
		var word wordsuuid
		err := curw.Decode(&word)
		if err != nil {
			log.Fatal(err)
		}
		reswords = append(reswords, &word)
	}

	if err := curw.Err(); err != nil {
		log.Fatal(err)
	}

	curw.Close(context.TODO())

	wordAuuid := reswords[0]
	wordBuuid := reswords[1]

	// Параметры поиска векторов слов wA и wB {filterf{ind}v{ectors}}
	filterfv := bson.M{"worduuid": bson.D{{"$in", bson.A{wordAuuid.UUID, wordBuuid.UUID}}}}

	var resvectors []*wordsvectors

	// Поиск векторов слов wA и wB
	curv, err := vectors.Find(context.TODO(), filterfv)
	if err != nil {
		log.Fatal(err)
	}

	for curv.Next(context.TODO()) {
		var vector wordsvectors
		err := curv.Decode(&vector)
		if err != nil {
			log.Fatal(err)
		}
		resvectors = append(resvectors, &vector)
	}

	if err := curv.Err(); err != nil {
		log.Fatal(err)
	}

	curv.Close(context.TODO())

	wordAvector := resvectors[0]
	wordBvector := resvectors[1]

	// Конвертировать векторное представление wA и wB в тип float64
	wA64 := conv.ConvertX(wordAvector.Vectors)
	wB64 := conv.ConvertX(wordBvector.Vectors)

	start := time.Now()
	// Рассчитать семантическую близость векторов
	csv := similarity.Calc(wA64, wB64)
	elapsed := time.Since(start)
	f, err := os.OpenFile("/home/kuranov/Code/golang/service/el", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Printf("%v", elapsed)

	datacs := csdata{
		CSTitle:        "CS",
		CSWordA:        wA,
		CSWordAUUID:    wordAuuid.UUID,
		CSWordAVectors: wordAvector.Vectors,
		CSWordB:        wB,
		CSWordBUUID:    wordBuuid.UUID,
		CSWordBVectors: wordBvector.Vectors,
		CSV:            csv,
	}
	cscm, _ := template.ParseFiles("templates/cosinesimilarity.html")
	cscm.Execute(w, datacs)
}
