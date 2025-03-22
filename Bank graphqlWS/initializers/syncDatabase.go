package initializers

import "Bank_graphqlWS/graph/model"

func SyncDatabase() {

	DB.AutoMigrate(&model.Bank{})
	DB.AutoMigrate(&model.Purchase{})
	DB.AutoMigrate(&model.Insurance{})
}
