package db

import {

    "context"
    "log"
    "api/configs"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
}

var mongoCN= connectDB()
var clientOptions = options.Client().ApplyURI(config.MongoDB.URI) // use env variables

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
    log.Printfln("Conexion exitosa con la DB")
    return client
}