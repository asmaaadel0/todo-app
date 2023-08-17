package internal

// Task is a struct representing a task
type Task struct {
	// Id uniqe for each task
	Id int `json:"id"`
	// Title descripe each task
	Title string `json:"title"`
	// Completed if it's complete or not
	Completed bool `json:"completed"`
}
