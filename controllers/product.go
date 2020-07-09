package controllers

import (
	"cleverbamboo.com/bee-shop-b2c/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

// ProductController operations for Product
type ProductController struct {
	beego.Controller
}

// URLMapping ...
func (c *ProductController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Product
// @Param	body		body 	models.Product	true		"body for Product content"
// @Success 201 {int} models.Product
// @Failure 403 body is empty
// @router / [post]
func (c *ProductController) Post() {
	var v models.Product
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddProduct(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Product by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Product
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ProductController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetProductById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Product
// @Param	id		    path 	string	        true		"The id you want to update"
// @Param	body		body 	models.Product	true		"body for Product content"
// @Success 200 {object} models.Product
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProductController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Product{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateProductById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Product
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ProductController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteProduct(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
