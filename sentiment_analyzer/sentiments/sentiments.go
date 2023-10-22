package sentiments

import (
	"fmt"
	"log"

	"github.com/cdipaolo/sentiment"
)

// AnalyzeSentiment ...
func AnalyzeSentiment(comment string){

//   Create sentiment analysis model
model, err := sentiment.Restore()

if err !=nil{
	log.Fatalln(err.Error())
}

var analysis *sentiment.Analysis

analysis = model.SentimentAnalysis(comment, sentiment.English)
var sentimentResult string
if analysis.Score == 1{
    sentimentResult = "Positive"
}else{
	sentimentResult = "Negative"
}

fmt.Println("SENTIMENT", sentimentResult)
}