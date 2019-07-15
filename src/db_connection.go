package main

import (
	"net/http"
	// "strconv"
	"time"

	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//-------------------------------------------------------- Structs for database tables -----------------------------------------------------------------------
type Url struct {
	gorm.Model
	UrlName           string `gorm:"unique;not null" json:"url_name"`
	Crawl_timeout     int    `json:"crawl_timeout`
	Frequency         int    `json:frequency`
	Failure_threshold int    `json:failure_threshold`
	Health            int    `gorm:"default:2"`
}

type UrlHits struct {
	gorm.Model
	Hit_number int
	Status int
}

//-------------------------------------------------------------- Migrating tables ---------------------------------------------------------------------

var db *gorm.DB

func init() {

	//open a db connection

	var err error

	db, err = gorm.Open("mysql", "root:nikitesh@/url_health?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	//Migrate the schema

	fmt.Println("Creating URL table ")
	db.AutoMigrate(&Url{})

	// fmt.Println("Creating UrlHits table")
	// db.AutoMigrate(&UrlHits{})

}

// ------------------------------------------------------------ setting up routes ----------------------------------------------------------------------
func main() {
	defer db.Close()

	router := gin.Default()

	v1 := router.Group("/CRUD")
	{
		v1.POST("/", createUrl)
		v1.GET("/", fetchAllUrl)

		// 	v1.GET("/:id", fetchSingleUrl)
		// 	v1.PUT("/:id", updateUrl)
		// 	v1.DELETE("/:id", deleteUrl)

	}
	router.Run()

	fmt.Println("Connected to database")
}


// createUrl add new row in url_health table

func createUrl(c *gin.Context) {
	var x []Url
	c.Bind(&x)

	fmt.Println(x)
	fmt.Println("\n\n")
	for i := 0; i < len(x); i++ {
		db.Save(&x[i])
		fmt.Println("inserting data into table Url ")
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Url added successfully!"})
	}

}



// fetchAllTodo fetch all todos

func fetchAllUrl(c *gin.Context) {
	var urls []Url

	db.Find(&urls)

	if len(urls) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No urls found!, kindly insert some urls."})
		return
	}
	for i := 0 ; i < len(urls) ; i++ {
		fmt.Println("Checking for url :", urls[i].UrlName)

		timeout := time.Duration(100 * time.Second)
		client := http.Client {
			Timeout: timeout,
		}
		resp, err := client.Get(urls[i].UrlName)

		if err != nil {
			fmt.Println(err)
			fmt.Println("url is unhealthy")
		} else {
			fmt.Println(resp.StatusCode, resp.Status)
			if resp.StatusCode == 200 {
				fmt.Println("chalrela hai")
				fmt.Println(resp.Status)
			} else {
				fmt.Println("nahi chalrela hai")
			}

		}

	}

}





















// 	//transforms the todos for building a good response
// 	for _, item := range todos {
// 		completed := false
// 		if item.Completed == 1 {
// 			completed = true
// 		} else {
// 			completed = false
// 		}
// 		_todos = append(_todos, transformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
// 	}
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
// }

// // fetchSingleTodo fetch a single todo
// func fetchSingleTodo(c *gin.Context) {
// 	var todo todoModel
// 	todoID := c.Param("id")

// 	db.First(&todo, todoID)

// 	if todo.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
// 		return
// 	}

// 	completed := false
// 	if todo.Completed == 1 {
// 		completed = true
// 	} else {
// 		completed = false
// 	}

// 	_todo := transformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
// }

// // updateTodo update a todo
// func updateTodo(c *gin.Context) {
// 	var todo todoModel
// 	todoID := c.Param("id")

// 	db.First(&todo, todoID)

// 	if todo.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
// 		return
// 	}

// 	db.Model(&todo).Update("title", c.PostForm("title"))
// 	completed, _ := strconv.Atoi(c.PostForm("completed"))
// 	db.Model(&todo).Update("completed", completed)
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
// }

// // deleteTodo remove a todo
// func deleteTodo(c *gin.Context) {
// 	var todo todoModel
// 	todoID := c.Param("id")

// 	db.First(&todo, todoID)

// 	if todo.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
// 		return
// 	}

// 	db.Delete(&todo)
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
// }