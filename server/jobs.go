package server

import (
	"context"
	"fmt"
	"go-challenge/config"
	"go-challenge/internals/models"
	"go-challenge/internals/notification"
	"go-challenge/internals/services"
	"log"

	"github.com/robfig/cron/v3"
	"go.uber.org/fx"
)

// ImportJob will execute the imports of Open Food Fact data
// following a scheduled cron job
func ImportJob(lc fx.Lifecycle, c *config.Config, n notification.Importation, i services.Importation) {
	cr := cron.New()

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Printf("Import job scheduled to: %s \n", c.TimeExecImport)
			cr.Start()
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Printf("Import job scheduled to: %s \n", c.TimeExecImport)
			cr.Stop()
			return nil
		},
	})

	cr.AddFunc(c.TimeExecImport, func() {
		var err error
		var filenames []string

		if filenames, err = i.GetFilenames(); err != nil {
			checkErr(n.NotifyFail(err))
			return
		}

		if len(filenames) < 1 {
			message := "0 Open Food Facts files found"
			checkErr(n.NotifySuccess(message))
			return
		}

		var imports []models.Import
		if imports, err = i.ToBeImported(filenames); err != nil {
			checkErr(n.NotifyFail(err))
			return
		}

		if len(imports) < 1 {
			message := "0 files to be imported"
			checkErr(n.NotifySuccess(message))
			return
		}

		if err = i.ImportFiles(imports); err != nil {
			checkErr(n.NotifyFail(err))
			return
		}

		message := fmt.Sprintf("Files imported: %v", len(imports))
		checkErr(n.NotifySuccess(message))
	})
}

func checkErr(err error) {
	if err != nil {
		log.Printf("Notification err: %v\n", err)
	}
}
