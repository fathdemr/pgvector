package Config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	appDBHost         string
	appDBPort         string
	appDBUserName     string
	appDBUserPassword string
	appDBName         string
	Db                *gorm.DB
	Viper             = viper.New()
	OpenAIKey         string
)

func InitDb() error {

	err := ReadConfig()
	if err != nil {
		return err
	}

	cnnString, _ := GetCnnString()

	db, err := gorm.Open(postgres.Open(cnnString), &gorm.Config{})
	if err != nil {
		return err
	}

	fmt.Println("DB Connected")

	Db = db

	return nil
}

func GetCnnString() (string, error) {

	cnnString := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=disable",
		appDBHost,
		appDBUserName,
		appDBUserPassword,
		appDBName,
		appDBPort)

	return cnnString, nil
}

func ReadConfig() error {
	Viper.SetConfigFile(".env")
	err := Viper.ReadInConfig()
	if err != nil {
		return err
	}

	appDBHost = Viper.GetString("DB_HOST_NAME")
	appDBUserName = Viper.GetString("DB_USER")
	appDBUserPassword = Viper.GetString("DB_PASSWORD")
	appDBName = Viper.GetString("DB_NAME")
	appDBPort = Viper.GetString("DB_PORT")
	OpenAIKey = Viper.GetString("OPENAI_API_KEY")

	return nil
}
