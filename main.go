package main

import (
	"be-authen/config"
	"be-authen/database"
	"be-authen/di"
	"be-authen/server"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {

	conf := config.GetConfig()

	db := database.NewPostgresDatabase(conf)
	// rabbitMQURL := os.Getenv("AMQP_URL")
	// if rabbitMQURL == "" {
	// 	log.Fatal("AMQP_URL is not set in the environment")
	// }
	// rabbitMQ, err := commons.InitRabbitMQ(rabbitMQURL)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	// }
	// defer rabbitMQ.Close()
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(conf.Aws.Region),
		Credentials: credentials.NewStaticCredentials(
			conf.Aws.AccessKey,
			conf.Aws.SecretKey,
			"",
		),
	}))

	// สร้าง SNS Client
	snsClient := sns.New(sess)

	gormDB := db.GetDb()
	container := di.NewContainer(gormDB, snsClient)

	ginServer := server.NewGinServer(conf, db, container)

	ginServer.Start()
}
