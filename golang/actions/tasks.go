package actions

import (
	"management/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Task)
// DB Table: Plural (tasks)
// Resource: Plural (Tasks)
// Path: Plural (/tasks)
// View Template Folder: Plural (/templates/tasks/)

// TasksResource is the resource for the Task model
type TasksResource struct {
	buffalo.Resource
}

// List gets all Tasks. This function is mapped to the path
// GET /tasks
func (v TasksResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return InternalError(c)
	}

	tasks := &models.Tasks{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Tasks from the DB
	if err := q.All(tasks); err != nil {
		return InternalError(c)
	}

	// Add the paginator to the context so it can be used in the template.

	return c.Render(http.StatusOK, r.JSON(Response{
		Data:       tasks,
		Pagination: q.Paginator,
	}))
}

// Show gets the data for one Task. This function is mapped to
// the path GET /tasks/{task_id}
func (v TasksResource) Show(c buffalo.Context) error {

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return InternalError(c)
	}

	// Allocate an empty User
	task := &models.Task{}

	// To find the User the parameter user_id is used.
	if err := tx.Find(task, c.Param("task_id")); err != nil {
		return Error(c, http.StatusForbidden, "task.not_found")
	}

	return c.Render(http.StatusOK, r.JSON(Response{Data: task}))
}

// New renders the form for creating a new Task.
// This function is mapped to the path GET /tasks/new
func (v TasksResource) New(c buffalo.Context) error {
	return c.Render(200, r.JSON(&models.Task{}))
}

// Create adds a Task to the DB. This function is mapped to the
// path POST /tasks
func (v TasksResource) Create(c buffalo.Context) error {
	// Allocate an empty Task
	task := &models.Task{}

	// Bind task to the html form elements
	if err := c.Bind(task); err != nil {
		return Error(c, http.StatusForbidden, "task.create.failed")
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return InternalError(c)
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(task)
	if err != nil {
		return InternalError(c)
	}

	if verrs.HasAny() {

		// Render again the new.html template that the user can
		// correct the input.
		return Error(c, http.StatusForbidden, "task.create.failed", verrs)
	}

	// and redirect to the tasks index page
	return Success(c, "task.create.success", task)
}

// Edit renders a edit form for a Task. This function is
// mapped to the path GET /tasks/{task_id}/edit
func (v TasksResource) Edit(c buffalo.Context) error {

	return c.Redirect(http.StatusMovedPermanently, "/")
}

// Update changes a Task in the DB. This function is mapped to
// the path PUT /tasks/{task_id}
func (v TasksResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errNoTransaction
	}

	// Allocate an empty Task
	task := &models.Task{}

	if err := tx.Find(task, c.Param("task_id")); err != nil {
		return Error(c, http.StatusNotFound, "task.update.failed")
	}

	// Bind Task to the html form elements
	if err := c.Bind(task); err != nil {
		return Error(c, http.StatusForbidden, "task.update.failed")
	}

	verrs, err := tx.ValidateAndUpdate(task)

	if err != nil {
		return InternalError(c)
	}

	if verrs.HasAny() {

		// Render again the edit.html template that the user can
		// correct the input.
		return Error(c, http.StatusForbidden, "task.update.failed", verrs)
	}

	return Success(c, "task.update.success", task)
}

// Destroy deletes a Task from the DB. This function is mapped
// to the path DELETE /tasks/{task_id}
func (v TasksResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return InternalError(c)
	}

	// Allocate an empty Task
	task := &models.Task{}

	// To find the Task the parameter task_id is used.
	if err := tx.Find(task, c.Param("task_id")); err != nil {
		return Error(c, http.StatusNotFound, "task.destroy.failed")
	}

	if err := tx.Destroy(task); err != nil {
		return Error(c, http.StatusForbidden, "task.destroy.failed")
	}

	return Success(c, "task.destroyed.success")
}
