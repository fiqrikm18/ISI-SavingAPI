package repositories

import "github.com/fiqrikm18/ISITransaction/internal/configs"

type IMutationRepository interface {
	CreateMutation(accountId int, amount float64, mutationType string) error
}

type MutationRepository struct {
	dbConf *configs.DatabaseConnection
}

func NewMutationRepository() *MutationRepository {
	dbConfig := configs.NewDatabaseConnection()

	return &MutationRepository{
		dbConf: dbConfig,
	}
}

func (repository *MutationRepository) CreateMutation(accountId int, amount float64, mutationType string) error {
	qs := `INSERT INTO mutations (account_id, amount, type, created_at) VALUES ($1, $2, $3, NOW())`
	_, err := repository.dbConf.DB.Exec(qs, accountId, amount, mutationType)
	if err != nil {
		return err
	}

	return nil
}
