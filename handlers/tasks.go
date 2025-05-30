package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m4d4r4666/fiber-crud/db"
	"github.com/m4d4r4666/fiber-crud/models"
)

func GetTasks(c *fiber.Ctx) error {
	rows, err := db.DB.Query("SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil {
		return c.Status(500).SendString("Error in Query")
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return c.Status(500).SendString("Scan error")
		}
		tasks = append(tasks, task)
	}
	return c.JSON(tasks)
}

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).SendString("Uncompatible query type")
	}

	_, err := db.DB.Exec("INSERT INTO tasks (id,title, description, status) values($1,$2,$3,$4)",
		task.ID, task.Title, task.Description, task.Status)
	if err != nil {
		return c.Status(500).SendString("Error while inserting values to DB")
	}
	return c.Status(201).SendString("Task  was created sucessfully")
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(400).SendString("Uncompatible query type")
	}

	_, err := db.DB.Exec("UPDATE tasks SET status=$1 where id=$2",
		task.Status, id)
	if err != nil {
		return c.Status(500).SendString("Error while updating data")
	}
	return c.SendString("Task updated successfully")
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := db.DB.Exec("DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		return c.Status(500).SendString("Error while deleting task")
	}
	return c.SendString("Task was deleted successfully")
}
