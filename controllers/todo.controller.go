package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"trandung/server/configs"
	"trandung/server/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var todoCollection *mongo.Collection = configs.GetCollection(configs.DB, "todos")

func GetTodos(c *gin.Context) {
	cur, err := todoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	var results []models.Todo
	if err = cur.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"todos":  results,
	})
}
func GetTodoById(c *gin.Context) {
	var id = c.Param("id")
	todoIdObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "failure",
			"message": "Id invalid",
		})
		return
	}
	var result models.Todo
	err = todoCollection.FindOne(context.TODO(), bson.D{{"_id", todoIdObject}}).Decode(&result)
	if result.Title == "" && result.Des == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"message": "Todo not found!",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"item":   result,
		})
		return
	}
}
func AddTodo(c *gin.Context) {
	var newTodo models.Todo
	newTodo.Id = primitive.NewObjectID()
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	res, err := todoCollection.InsertOne(context.Background(), newTodo)
	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID
	fmt.Println(res)
	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"_id":    id,
	})
}

func DeleteTodoById(c *gin.Context) {
	var id = c.Param("id")
	todoIdObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "failure",
			"message": "Id invalid",
		})
		return
	}
	var result models.Todo
	err = todoCollection.FindOneAndDelete(context.TODO(), bson.D{{"_id", todoIdObject}}).Decode(&result)
	if result.Title == "" && result.Des == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"message": "Todo not found!",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Deleted todo successfully!",
		})
		return
	}
}
func UpdateById(c *gin.Context) {
	var id = c.Param("id")
	todoIdObject, err := primitive.ObjectIDFromHex(id)
	var updateTodo models.Todo
	if err := c.BindJSON(&updateTodo); err != nil {
		return
	}
	fmt.Println(updateTodo.Title)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": "Id invalid",
		})
		return
	}
	var result models.Todo
	err = todoCollection.FindOne(context.TODO(), bson.D{{"_id", todoIdObject}}).Decode(&result)
	if result.Title == "" && result.Des == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"message": "Todo not found!",
		})
		return
	} else {
		opts := options.Update().SetUpsert(true)
		result, err := todoCollection.UpdateOne(context.TODO(), bson.D{{"_id", todoIdObject}}, bson.D{{"$set", bson.D{{"title", updateTodo.Title}, {"des", updateTodo.Des}, {"completed", updateTodo.Completed}}}}, opts)
		if err != nil {
			log.Fatal(err)
		}
		if result.MatchedCount != 0 {
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"item":   "Updated todo successfully!",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "failure",
				"message": "Todo update failure!",
			})
			return
		}
	}
}
