package predictor

import (
	"github.com/dmitryikh/leaves"
	"log"
)

func LoadModel() (*leaves.Ensemble, error) {
	model, err := leaves.LGEnsembleFromFile(getModelPath(), true)
	if err != nil {
		log.Printf("load model error: [%v]", err)
		return nil, err
	}
	return model, nil
}

func Predict(model *leaves.Ensemble, feature []float64) (float64, error) {
	predScore := model.PredictSingle(feature, 0)
	return predScore, nil
}
