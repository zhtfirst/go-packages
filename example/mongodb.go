package example

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/zhtfirst/go-packages/config"
	"github.com/zhtfirst/go-packages/mongodb"
)

func MongoDbConnect() {
	config.Setup("") // 初始化配置

	mongodb.InitMongoDB(config.GetString("mongo_conf", "hosts"))
	fmt.Println("MongoDB 连接数据库:", mongodb.MongoDB)

	// 常用操作
	//// 添加一条
	//s := Student{Name: "wang fan", Age: 26}
	//insertOne(s)

	//// 添加多条
	//s1 := Student{Name: "wang fan1", Age: 26}
	//s2 := Student{Name: "wang fan2", Age: 26}
	//s3 := Student{Name: "wang fan3", Age: 26}
	//students := []interface{}{s1, s2, s3}
	//insertMore(students)

	//// 查询
	//findData()

	//// 更新
	//update()

	// 删除
	delete()

}

type Student struct {
	Name string
	Age  int
}

func insertOne(s Student) {
	c := mongodb.MongoDB.Database("go_db").Collection("student")
	ior, err := c.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ior.InsertedID: %v\n", ior.InsertedID)
}

func insertMore(students []interface{}) {
	c := mongodb.MongoDB.Database("go_db").Collection("student")
	imr, err := c.InsertMany(context.TODO(), students)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("imr.InsertedIDs: %v\n", imr.InsertedIDs)
}

func findData() {
	c := mongodb.MongoDB.Database("go_db").Collection("student")
	// c2, _ := c.Find(context.Background(), bson.D{{Key: "name", Value: "wang fan"}})
	// c2, _ := c.Find(context.Background(), bson.D{{"name", "wang fan"}})
	//c2, _ := c.Find(context.Background(), bson.D{})
	c2, _ := c.Find(context.Background(), bson.M{"name": "wang fan3"}) // ObjectId("632a76cbb09b7afe1e1bc385")

	defer c2.Close(context.Background())
	for c2.Next(context.Background()) {
		var result bson.D
		err := c2.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("result: %v\n", result)
		fmt.Printf("result.Map(): %v\n", result.Map())
		fmt.Printf("result.Map()[\"name\"]: %v\n", result.Map()["name"])
	}
	if err := c2.Err(); err != nil {
		log.Fatal(err)
	}
}

func update() {
	c := mongodb.MongoDB.Database("go_db").Collection("student")
	update := bson.D{{"$set", bson.D{{Key: "name", Value: "王帆"}, {Key: "age", Value: 19}}}}
	// 更新条件 name wang fan 更新问 王帆 年龄 18
	ur, err := c.UpdateMany(context.TODO(), bson.D{{"name", "wang fan"}}, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ur.ModifiedCount: %v\n", ur.ModifiedCount)
}

func delete() {
	c := mongodb.MongoDB.Database("go_db").Collection("student")
	dr, err := c.DeleteMany(context.TODO(), bson.D{{Key: "name", Value: "wang fan2"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dr.DeletedCount: %v\n", dr.DeletedCount)
}
