package controllers

import (
	"cleverbamboo.com/bee-shop-b2c/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

// TagController operations for Tag
type TagController struct {
	beego.Controller
}

// URLMapping ...
func (c *TagController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Tag
// @Param	body		body 	models.Tag	true		"body for Tag content"
// @Success 201 {int} models.Tag
// @Failure 403 body is empty
// @router / [post]
func (c *TagController) Post() {
	var v models.Tag
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddTag(&v); err == nil {
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
// @Description get Tag by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Tag
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TagController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTagById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Tag
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Tag	true		"body for Tag content"
// @Success 200 {object} models.Tag
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TagController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Tag{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateTagById(&v); err == nil {
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
// @Description delete the Tag
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TagController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteTag(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
