package controllers

import (
	"cleverbamboo.com/bee-shop-b2c/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

// ProductCategoryController operations for ProductCategory
type ProductCategoryController struct {
	beego.Controller
}

// URLMapping ...
func (c *ProductCategoryController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create ProductCategory
// @Param	body		body 	models.ProductCategory	true		"body for ProductCategory content"
// @Success 201 {int} models.ProductCategory
// @Failure 403 body is empty
// @router / [post]
func (c *ProductCategoryController) Post() {
	var v models.ProductCategory
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddProductCategory(&v); err == nil {
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

// Put ...
// @Title Put
// @Description update the ProductCategory
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ProductCategory	true		"body for ProductCategory content"
// @Success 200 {object} models.ProductCategory
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProductCategoryController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.ProductCategory{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateProductCategoryById(&v); err == nil {
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
// @Description delete the ProductCategory
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ProductCategoryController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteProductCategory(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
