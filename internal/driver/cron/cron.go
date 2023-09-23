package cron

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/config"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/factory"
	"github.com/rochmanramadhani/go-lazisnu-api/pkg/util/gracefull"
)

var stopper = gracefull.NilStopper

// !TODO: add stopper
func Init(cfg *config.Configuration, f factory.Factory) gracefull.ProcessStopper {
	if !cfg.Driver.Cron.Enabled {
		return stopper
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	s := gocron.NewScheduler(loc)
	s.SingletonModeAll()
	s.WaitForScheduleAll()

	// daily
	//_, _ = s.Every(1).Day().Do(func() {
	//	logrus.Info("Cron reminder stock daily, running ...")
	//	err := f.Usecase.ReminderStockHistory.GenerateRecurring(context.Background(), constant.REMINDER_STOCK_DAILY)
	//	if err != nil {
	//		logrus.Error("Cron reminder stock daily error", err)
	//	}
	//})

	// monthly
	//_, _ = s.Every(1).MonthLastDay().Do(func() {
	//	logrus.Info("Cron reminder stock monthly, running ...")
	//	err := f.Usecase.ReminderStockHistory.GenerateRecurring(context.Background(), constant.REMINDER_STOCK_MONTHLY)
	//	if err != nil {
	//		logrus.Error("Cron reminder stock monthly error", err)
	//	}
	//})

	s.StartAsync()

	return stopper
}
