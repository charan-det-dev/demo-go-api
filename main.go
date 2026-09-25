package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"createdAt"`
}

func main() {

	app := fiber.New()

	fmt.Println("API Ready...")

	// GET /users endpoint
	app.Get("/users", func(c *fiber.Ctx) error {
		fmt.Println("GET /users endpoint called")

		data, err := json.MarshalIndent(users(), "", " ")
		if err != nil {

			fmt.Println("Error marshalling users data:", err)
			return err
		}

		fmt.Println(string(data))

		return c.JSON(users())
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		fmt.Println("GET /users/:id endpoint called with ID:", id)

		for _, user := range users() {
			if user.ID == id {

				fmt.Println(user)
				return c.JSON(user)
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	})

	log.Fatal(app.Listen(":8080"))
}

func users() []User {

	return []User{
		{
			CreatedAt: "2025-03-31T18:16:12.571Z",
			Name:      "Luther Streich",
			Avatar:    "https://avatars.githubusercontent.com/u/38795099",
			ID:        "1",
		},
		{
			CreatedAt: "2025-04-01T02:14:31.750Z",
			Name:      "Sonya Ondricka",
			Avatar:    "https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/73.jpg",
			ID:        "2",
		},
		{
			CreatedAt: "2025-03-31T10:49:23.624Z",
			Name:      "Travis Farrell",
			Avatar:    "https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/male/512/47.jpg",
			ID:        "3",
		},
		{
			CreatedAt: "2025-04-01T04:27:00.561Z",
			Name:      "Kristine Hirthe IV",
			Avatar:    "https://avatars.githubusercontent.com/u/50226271",
			ID:        "4",
		},
		{
			CreatedAt: "2025-03-31T23:00:44.288Z",
			Name:      "Vicky Reinger",
			Avatar:    "https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/male/512/93.jpg",
			ID:        "5",
		},
		{
			CreatedAt: "2025-03-31T20:53:06.750Z",
			Name:      "Otis Herzog",
			Avatar:    "https://avatars.githubusercontent.com/u/80682116",
			ID:        "6",
		},
		{
			CreatedAt: "2025-03-31T23:59:57.878Z",
			Name:      "Mrs. Marian Farrell",
			Avatar:    "https://avatars.githubusercontent.com/u/97060009",
			ID:        "7",
		},
		{
			CreatedAt: "2025-03-31T12:58:35.636Z",
			Name:      "Karen Stehr DVM",
			Avatar:    "https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/31.jpg",
			ID:        "8",
		},
		{
			CreatedAt: "2025-03-31T20:20:21.038Z",
			Name:      "Tonya Hauck",
			Avatar:    "https://avatars.githubusercontent.com/u/49235528",
			ID:        "9",
		},
		{
			CreatedAt: "2025-03-31T12:21:09.859Z",
			Name:      "Kerry Muller",
			Avatar:    "https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/1.jpg",
			ID:        "10",
		},
		{
			CreatedAt: "2025-03-31T13:25:41.311Z",
			Name:      "Rene Romaguera PhD",
			Avatar:    "https://avatars.githubusercontent.com/u/62373701",
			ID:        "11",
		},
		{
			CreatedAt: "2025-03-31T08:37:48.905Z",
			Name:      "April Gutmann",
			Avatar:    "https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/87.jpg",
			ID:        "12",
		},
		{
			CreatedAt: "2025-04-01T06:15:02.363Z",
			Name:      "Laverne Zulauf",
			Avatar:    "https://avatars.githubusercontent.com/u/78751297",
			ID:        "13",
		},
		{
			CreatedAt: "2025-03-31T11:33:41.249Z",
			Name:      "Mr. Edmond Feeney",
			Avatar:    "https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/male/512/32.jpg",
			ID:        "14",
		},
		{
			CreatedAt: "2025-03-31T20:50:22.718Z",
			Name:      "Taylor Donnelly",
			Avatar:    "https://avatars.githubusercontent.com/u/65592570",
			ID:        "15",
		},
		{
			CreatedAt: "2025-03-31T21:55:30.201Z",
			Name:      "Jacqueline Kuphal",
			Avatar:    "https://avatars.githubusercontent.com/u/65683142",
			ID:        "16",
		},
		{
			CreatedAt: "2025-03-31T15:39:38.318Z",
			Name:      "Yvette Feil",
			Avatar:    "https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/male/512/83.jpg",
			ID:        "17",
		},
		{
			CreatedAt: "2025-03-31T07:43:01.696Z",
			Name:      "Boyd Hermann",
			Avatar:    "https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/male/512/34.jpg",
			ID:        "18",
		},
		{
			CreatedAt: "2025-03-31T13:17:53.986Z",
			Name:      "Hattie Crist",
			Avatar:    "https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/15.jpg",
			ID:        "19",
		},
		{
			CreatedAt: "2025-03-31T09:28:47.279Z",
			Name:      "Jesse Greenfelder PhD",
			Avatar:    "https://avatars.githubusercontent.com/u/2267929",
			ID:        "20",
		},
	}
}
