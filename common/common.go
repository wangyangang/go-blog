package common

import (
	"go-blog/config"
	"go-blog/models"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/templates/")
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}
