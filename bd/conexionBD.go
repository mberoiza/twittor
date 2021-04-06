package bd

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = conectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.ecdhv.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

func conectarDB() *mongo.Client {
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil{
        log.Fatal(err.Error())
        return client
    }
    err = client.Ping(context.TODO(), nil)
    if err != nil{
        log.Fatal(err.Error())
        return client
    }

    log.Println("Conexion Exitosa con la BD")
    return client
}

func ChequeoConnection() int {
    err := MongoCN.Ping(context.TODO(), nil)
    if err != nil {
        return 0
    }
    return 1
}
