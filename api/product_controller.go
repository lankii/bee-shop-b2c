package api

import (
	"cleverbamboo.com/bee-shop-b2c/common"
	"cleverbamboo.com/bee-shop-b2c/helpers"
	"cleverbamboo.com/bee-shop-b2c/model_views"
	"cleverbamboo.com/bee-shop-b2c/models"
	"encoding/json"
	"strconv"
	"strings"
)

type ProductController struct {
	BaseController
}

// URLMapping ...
func (c *ProductController) URLMapping() {
	c.Mapping("AddProduct", c.AddProduct)
	c.Mapping("GetAllProduct", c.GetAllProduct)
	c.Mapping("UpdateProduct", c.UpdateProduct)
}

// @Title Add Product
// @Description create Product
// @Success 201
// @Failure 500
// @router /add [post]
func (c *ProductController) AddProduct() {
	var productCategoryId int
	var productType int
	var name string
	var caption string
	var price float64
	var cost float64
	var marketPrice float64
	var image string
	var unit string
	var point int64
	var stock int
	var brand int
	var tag []string
	var keyword string
	var isTop int
	var isMarketable int
	var isList int

	// productCategoryId
	if v := c.GetString("productCategoryId"); v != "" {
		productCategoryId, _ = strconv.Atoi(v)
	}
	// productType
	if v := c.GetString("productType"); v != "" {
		productType, _ = strconv.Atoi(v)
	}
	// caption
	if v := c.GetString("caption"); v != "" {
		caption = v
	}
	// name
	if v := c.GetString("name"); v != "" {
		name = v
	}
	// price
	if v := c.GetString("price"); v != "" {
		price, _ = strconv.ParseFloat(v, 64)
	}
	// cost
	if v := c.GetString("cost"); v != "" {
		cost, _ = strconv.ParseFloat(v, 64)
	}
	// marketPrice
	if v := c.GetString("marketPrice"); v != "" {
		marketPrice, _ = strconv.ParseFloat(v, 64)
	}
	// image
	if v := c.GetString("image"); v != "" {
		image = v
	}
	// unit
	if v := c.GetString("unit"); v != "" {
		unit = v
	}
	// point
	if v := c.GetString("point"); v != "" {
		point, _ = strconv.ParseInt(v, 10, 64)
	}
	// stock
	if v := c.GetString("stock"); v != "" {
		stock, _ = strconv.Atoi(v)
	}
	// brand
	if v := c.GetString("brand"); v != "" {
		brand, _ = strconv.Atoi(v)
	}
	// tag
	if v := c.GetString("tag"); v != "" {
		tag = strings.Split(v, ",")
	}
	// keyword
	if v := c.GetString("keyword"); v != "" {
		keyword = v
	}
	// is_top
	if v := c.GetString("isTop"); v != "" {
		isTop, _ = strconv.Atoi(v)
	}
	// is_marketable
	if v := c.GetString("isMarketable"); v != "" {
		isMarketable, _ = strconv.Atoi(v)
	}
	// is_list
	if v := c.GetString("isList"); v != "" {
		isList, _ = strconv.Atoi(v)
	}

	var product models.Product

	/**
	 * 商品名称不能相同
	 */

	/**
	 * 雪花算法
	 */
	node, err := helpers.NewWorker(int64(productCategoryId))
	if err != nil {
		c.ServerError(err)
		return
	}

	product.Sn = strconv.FormatInt(node.GetId(), 10)
	product.ProductCategoryId = &models.ProductCategory{Id: productCategoryId}
	product.Type = productType
	product.Name = name
	product.Caption = caption
	product.Price = price
	product.Cost = cost
	product.MarketPrice = marketPrice
	product.Image = image
	product.Unit = unit
	product.Point = point
	product.Stock = stock
	product.BrandId = &models.Brand{Id: brand}
	product.Keyword = keyword
	product.IsTop = int8(isTop)
	product.IsMarketable = int8(isMarketable)
	product.IsList = int8(isList)

	id, err := models.AddProduct(&product)
	if err != nil {
		c.ServerError(err)
		return
	}
	/**
	 * tag
	 */
	if len(tag) > 0 {
		for _, tag := range tag {
			tag, _ := strconv.Atoi(tag)
			if err = models.AddProductTag(int(id), tag); err != nil {
				c.ServerError(err)
				return
			}
		}
	}

	c.JsonResult(common.GetHttpStatus("created"), common.ErrOK, common.Success, nil)
}

// @Title GetAllProduct
// @Description Get all product list by some filed
// @Param	query	    query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	sortby	    query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	    query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ...
// @Param	pageNumber	query	string	false	"Start position of result set. Must be an integer"
// @Param	pageSize	query	int	    false	"Limit the size of result set. Must be an integer"
// @router / [get]
// @Success 200 {object} model_views.Product
// @Failure 500
func (c *ProductController) GetAllProduct() {
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var pageSize int64 = 10
	var pageNumber int64 = 1

	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.JsonResult(common.GetHttpStatus("ok"), common.ErrError, "fail", "Error: invalid query key/value pair")
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}
	// pageNumber: 1 (default is 1)
	if v := c.GetString("pageNumber"); v != "" {
		pageNumber, _ = strconv.ParseInt(v, 10, 64)
	}
	// pageSize: 10 (default is 10)
	if v := c.GetString("pageSize"); v != "" {
		pageSize, _ = strconv.ParseInt(v, 10, 64)
	}
	// start position of result set
	offset := (pageNumber - 1) * pageSize
	l, err := models.GetAllProduct(query, nil, sortby, order, offset, pageSize)
	if err != nil {
		c.ServerError(err)
		return
	}

	/**
	 * 分页并只返回前端需要的字段
	 */
	var pageList []interface{}
	for _, v := range l {
		productView := model_views.Product{}
		product := v.(models.Product)

		productView.Id = product.Id
		productView.Sn = product.Sn
		productView.Name = product.Name
		productView.ProductCategoryId = product.ProductCategoryId.Id
		productView.ProductCategoryName = product.ProductCategoryId.Name
		productView.Price = product.Price
		productView.Stock = product.Stock
		productView.IsMarketable = product.IsMarketable
		productView.IsList = product.IsList
		productView.IsTop = product.IsTop
		productView.CreationDate = product.CreationDate
		pageList = append(pageList, productView)
	}

	/**
	 * 查询 Product Count
	 */
	cnt, err := models.GetProductCount(query)
	if err != nil {
		c.ServerError(err)
		return
	}
	pages, err := helpers.NewPagination(pageList, int(cnt), int(pageSize), int(pageNumber))

	if err != nil {
		c.ServerError(err)
		return
	}

	c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, common.Success, *pages)
}

// @Title UpdateProduct
// @Description Update product by some field
// @Param   id			path    string          true    "The id you want to update"
// @Param	body		body 	models.Product	true	"body for Product content"
// @router  /update/:id [put]
// @Success 200 nil
// @Fail    500
func (c *ProductController) UpdateProduct() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	// GetProductById
	currentProduct, err := models.GetProductById(id)
	if err != nil {
		c.ServerError(err)
		return
	}

	// Unmarshal RequestBody
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &currentProduct)
	if err != nil {
		c.ServerError(err)
		return
	}

	// UpdateProductById
	err = models.UpdateProductById(currentProduct)
	if err != nil {
		c.ServerError(err)
		return
	}

	c.JsonResult(common.GetHttpStatus("ok"), common.ErrOK, "success", nil)
}
