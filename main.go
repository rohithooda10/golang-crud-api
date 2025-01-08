package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)
type Item struct {
	ID int `json:"id"`
	Name string`json:"name"`
	Price int64`json:"price"`
}

type Response struct {
	Status int
	ErrorMessage string
	Data interface{}
}

type BaseHandler struct {}

type ItemHandler struct {
	BaseHandler
	Items []Item
	mutex sync.RWMutex
}

func (b *BaseHandler) sendResponse(ctx *fiber.Ctx, status int, err error, data interface{}){
	var response Response
	response.Status = status
	if err != nil {
		response.ErrorMessage = err.Error()
	}
	response.Data = data
	ctx.Status(status).JSON(response)
}

// Add new item
func (handler *ItemHandler) CreateItem(ctx *fiber.Ctx) error {
	var newItem Item
	err := ctx.BodyParser(&newItem)
	if err != nil {
		handler.sendResponse(ctx, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError, "")
		return fiber.ErrInternalServerError
	}
	handler.mutex.Lock()
	defer handler.mutex.Unlock()
	handler.Items = append(handler.Items, newItem)
	handler.sendResponse(ctx, 200, nil, "item added")
	return nil
}

// Get all items
func (handler *ItemHandler) GetItems(ctx *fiber.Ctx) error {
	handler.mutex.RLock()
	defer handler.mutex.RUnlock()
	handler.sendResponse(ctx, 200, nil, handler.Items)
	return nil
}

// Get item specified by id
func (handler *ItemHandler) GetItem(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		handler.sendResponse(ctx, fiber.ErrBadRequest.Code, fmt.Errorf("missing id"), "")
		return fiber.ErrBadRequest
	}
	handler.mutex.RLock()
	defer handler.mutex.RUnlock()
	for _, item := range handler.Items{
		if item.ID == id {
			handler.sendResponse(ctx, 200, nil, item)
			return nil
		}
	}
	return fiber.ErrNotFound
}

// Update item
func (handler *ItemHandler) UpdateItem(ctx *fiber.Ctx) error {
	var updatedItem Item
	id, err1 := strconv.Atoi(ctx.Params("id"))
	if err2 := ctx.BodyParser(&updatedItem); err1 != nil || err2 != nil {
		handler.sendResponse(ctx, fiber.ErrBadRequest.Code, fmt.Errorf("missing id or new item"), "")
		return fiber.ErrBadRequest
	}
	handler.mutex.Lock()
	defer handler.mutex.Unlock()
	for index, item := range handler.Items{
		if item.ID == id {
			handler.Items[index] = updatedItem
			handler.sendResponse(ctx, 200, nil, updatedItem)
			return nil
		}
	}
	return fiber.ErrNotFound
}

// Delete item
func (handler *ItemHandler) DeleteItem(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		handler.sendResponse(ctx, fiber.ErrBadRequest.Code, fmt.Errorf("missing id"), "")
		return fiber.ErrBadRequest
	}
	handler.mutex.Lock()
	defer handler.mutex.Unlock()
	for index, item := range handler.Items{
		if item.ID == id {
			handler.Items = append(handler.Items[:index], handler.Items[index + 1:]...)
			handler.sendResponse(ctx, 200, nil, "item deleted")
			return nil
		}
	}
	return fiber.ErrNotFound
}

func main(){
	app := fiber.New()

	itemHandler := &ItemHandler{
		Items: []Item{},
		mutex: sync.RWMutex{},
	}

	// Handlers
	app.Post("/api/v1/item", itemHandler.CreateItem)
	app.Get("/api/v1/items", itemHandler.GetItems)
	app.Get("/api/v1/item/:id", itemHandler.GetItem)
	app.Patch("/api/v1/item/:id", itemHandler.UpdateItem)
	app.Delete("/api/v1/item/:id", itemHandler.DeleteItem)

	// Start the server
	app.Listen(":8080")
}