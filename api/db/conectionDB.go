package db

import (

    "context"
    "log"
    "api/configs"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
/*mongoCN database object */
var mongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI(configs.GetConfig().MongoDB.URI) // use env variables

/* ConnectDB connects to database */
func ConnectDB() *mongo.Client {
    client, err := mongo.Connect(context.TODO(), clientOptions)
    //check connection
    if err != nil {
        log.Fatal(err)
        return client
    }
    //ping to db
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
        return client
    }
    log.Println("Conexion exitosa con la DB")
    return client
}

/* CheckConnection Ping to check the database connection */
func CheckConnection() int {
    err := mongoCN.Ping(context.TODO(), nil)
    if err != nil {
        return 0
    }
    return 1
}