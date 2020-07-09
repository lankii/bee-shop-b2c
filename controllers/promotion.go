package controllers

import (
	"cleverbamboo.com/bee-shop-b2c/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

// PromotionController operations for Promotion
type PromotionController struct {
	beego.Controller
}

// URLMapping ...
func (c *PromotionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Promotion
// @Param	body		body 	models.Promotion	true		"body for Promotion content"
// @Success 201 {int} models.Promotion
// @Failure 403 body is empty
// @router / [post]
func (c *PromotionController) Post() {
	var v models.Promotion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddPromotion(&v); err == nil {
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
// @Description get Promotion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Promotion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PromotionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetPromotionById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Promotion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Promotion	true		"body for Promotion content"
// @Success 200 {object} models.Promotion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PromotionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Promotion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdatePromotionById(&v); err == nil {
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
// @Description delete the Promotion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PromotionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeletePromotion(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
