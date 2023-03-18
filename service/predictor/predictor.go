package predictor

import (
	"github.com/dmitryikh/leaves"
	"log"
)

type TreePredictor struct {
	model *leaves.Ensemble
}

func (p *TreePredictor) LoadModel() error {
	modelPath := getModelPath()
	model, err := leaves.LGEnsembleFromFile(modelPath, true)
	if err != nil {
		log.Printf("load model error: [%v]", err)
		return err
	} else {
		log.Printf("get model from file: %s", modelPath)
		p.model = model
		return nil
	}
}

func (p *TreePredictor) Predict(feature []float64) (float64, error) {
	log.Printf("predict: %v", feature)
	predScore := p.model.PredictSingle(feature, 0)
	log.Printf("output: %v", predScore)
	return predScore, nil
}
