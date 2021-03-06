package job

import (
	"github.com/koshatul/satis-go/src/satis/satisphp/db"
)

// Get all repos in the repo collection
func NewFindAllJob(dbPath string) *FindAllJob {
	return &FindAllJob{
		dbPath:    dbPath,
		exitChan:  make(chan error, 1),
		ReposResp: make(chan []db.SatisRepository, 1),
	}
}

type FindAllJob struct {
	ReposResp chan []db.SatisRepository
	dbPath    string
	exitChan  chan error
}

func (j FindAllJob) ExitChan() chan error {
	return j.exitChan
}
func (j FindAllJob) Run() error {
	dbMgr := db.SatisDBManager{Path: j.dbPath}

	err := dbMgr.Load()

	j.ReposResp <- dbMgr.DB.Repositories // might be empty

	return err // might be nil
}
