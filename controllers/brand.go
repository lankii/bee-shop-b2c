package controllers

import (
	"cleverbamboo.com/bee-shop-b2c/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

// BrandController operations for Brand
type BrandController struct {
	beego.Controller
}

// URLMapping ...
func (c *BrandController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Brand
// @Param	body		body 	models.Brand	true		"body for Brand content"
// @Success 201 {int} models.Brand
// @Failure 403 body is empty
// @router / [post]
func (c *BrandController) Post() {
	var v models.Brand
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddBrand(&v); err == nil {
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
// @Description get Brand by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Brand
// @Failure 403 :id is empty
// @router /:id [get]
func (c *BrandController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetBrandById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Brand
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Brand	true		"body for Brand content"
// @Success 200 {object} models.Brand
// @Failure 403 :id is not int
// @router /:id [put]
func (c *BrandController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Brand{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateBrandById(&v); err == nil {
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
// @Description delete the Brand
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *BrandController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteBrand(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
