package main

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"mahasiswa/models"
	"mahasiswa/pkg"
	"mahasiswa/routes"
)

func main() {

	/**
	* ========================
	*  Setup Application
	* ========================
	 */

	db := setupDatabase()
	app := setupApp()

	/**
	* ========================
	* Initialize All Route
	* ========================
	 */

	routes.InitializeRouteStudent(db, app)
	routes.InitializeRouteTeacher(db, app)

	/**
	* ========================
	*  Listening Server Port
	* ========================
	 */

	err := app.Run(":" + pkg.GodotEnv("PORT"))

	if err != nil {
		defer logrus.Error("Server is not running")
		logrus.Fatal(err)
	}
}

/**
* ========================
* Database Setup
* ========================
 */

func setupDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(pkg.GodotEnv("PG_URI")), &gorm.Config{})

	if err != nil {
		defer logrus.Error("Database connection failed")
		logrus.Fatal(err)
		return nil
	}

	err = db.AutoMigrate(&models.Student{}, &models.Student{})

	if err != nil {
		defer logrus.Error("Database migration failed")
		logrus.Fatal(err)
		return nil
	}

	return db
}

/**
* ========================
* Application Setup
* ========================
 */

func setupApp() *gin.Engine {

	app := gin.Default()

	if pkg.GodotEnv("GO_ENV") != "development" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	app.Use(helmet.Default())
	app.Use(gzip.Gzip(gzip.BestCompression))
	app.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "DELETE", "PATCH", "PUT", "OPTIONS"},
		AllowHeaders:    []string{"Content-Type", "Authorization", "Accept-Encoding"},
	}))

	return app
}
