package handler

import (
	// "net/http"
	"context"

	"github.com/go-faster/jx"
	api "github.com/jamesw201/go-starter/example/api"
	"github.com/jamesw201/go-starter/example/ent"
	"github.com/jamesw201/go-starter/example/ent/keyboard"
	"github.com/jamesw201/go-starter/example/ent/keycapmodel"
	"github.com/jamesw201/go-starter/example/ent/switchmodel"
)

// OgentHandler implements the ogen generated Handler interface and uses Ent as data layer.
type OgentHandler struct {
	client *ent.Client
}

// NewOgentHandler returns a new OgentHandler.
func NewOgentHandler(c *ent.Client) *OgentHandler { 
  return &OgentHandler{c} 
}

// rawError renders err as json string.
func rawError(err error) jx.Raw {
	var e jx.Encoder
	e.Str(err.Error())
	return e.Bytes()
}

func (h *OgentHandler) GetKeyboard(ctx context.Context, params api.GetKeyboardParams) (*api.Keyboard, error) {
  e, err := h.client.Keyboard.Get(ctx, params.ID)
	// q := h.client.Keyboard.Query().Where(keyboard.IDEQ(params.ID))
	// e, err := q.Only(ctx)
	if err != nil {
		switch {
		// case ent.IsNotFound(err):
		// 	return &R404{
		// 		Code:   http.StatusNotFound,
		// 		Status: http.StatusText(http.StatusNotFound),
		// 		Errors: rawError(err),
		// 	}, nil
		// case ent.IsNotSingular(err):
		// 	return &R409{
		// 		Code:   http.StatusConflict,
		// 		Status: http.StatusText(http.StatusConflict),
		// 		Errors: rawError(err),
		// 	}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}

  // Reload the entity to attach all eager-loaded edges.
	// q := h.client.Keyboard.Query().Where(keyboard.ID(e.ID))
	// e, err = q.Only(ctx)
	// if err != nil {
	// 	// This should never happen.
	// 	return nil, err
	// }
  // sw, swerr := h.client.SwitchModel.Query().Where(switchmodel.IDEQ(e.))
  sw, swerr := h.client.Keyboard.QuerySwitches(e).Only(ctx)
  if swerr != nil {
		return nil, err
	}
  kc, kcerr := h.client.Keyboard.QueryKeycaps(e).Only(ctx)
  if kcerr != nil {
		return nil, err
	}

  keyboard, keyerr := NewKeyboardRead(e, sw, kc)
  return keyboard, keyerr
  // return NewKeyboardRead(e), nil
}

func (h *OgentHandler) KeyboardGet(ctx context.Context) ([]api.Keyboard, error) {
  return nil, nil
}

func saveSwitchModel(h *OgentHandler, ctx context.Context, req *api.Keyboard) (*ent.SwitchModel, error) {
  ksm := h.client.SwitchModel.Create()
  ksm.SetID(req.GetSwitches().ID)
  ksm.SetName(req.Switches.Name)
  ksm.SetSwitchType(switchmodel.SwitchType(req.GetSwitches().SwitchType))
  e, err := ksm.Save(ctx)
	if err != nil {
		switch {
		// case ent.IsNotSingular(err):
		// 	return req, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
  // Reload the entity to attach all eager-loaded edges.
	q := h.client.SwitchModel.Query().Where(switchmodel.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
  return e, nil
}

func saveKeyCaps(h *OgentHandler, ctx context.Context, req *api.Keyboard) (*ent.KeycapModel, error) {
  kkc := h.client.KeycapModel.Create()
  kkc.SetID(req.Keycaps.ID)
  kkc.SetName(req.Keycaps.Name)
  kkc.SetProfile(req.Keycaps.Profile)
  kkc.SetMaterial(keycapmodel.Material(req.Keycaps.Material))
  
  e, err := kkc.Save(ctx)
	if err != nil {
		switch {
		// case ent.IsNotSingular(err):
		// 	return req, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
  // Reload the entity to attach all eager-loaded edges.
	q := h.client.KeycapModel.Query().Where(keycapmodel.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
  return e, nil
}

func (h *OgentHandler) AddKeyboard(ctx context.Context, req *api.Keyboard) (*api.Keyboard, error) {
  k := h.client.Keyboard.Create()
  k.SetName(req.Name)
  k.SetPrice(req.Price)

  sm, smerr := saveSwitchModel(h, ctx, req)
  if smerr != nil {
    return nil, smerr
  }

  kc, kcerr := saveKeyCaps(h, ctx, req)
  if kcerr != nil {
    return nil, kcerr
  }
  
  k.SetSwitches(sm)
  k.SetKeycaps(kc)
  // k.SetDiscount(req.Discount)
  
	// Persist to storage.
	e, err := k.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return req, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Keyboard.Query().Where(keyboard.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
  return NewKeyboardCreate(e), nil
}
// // CreateCustomer handles POST /customers requests.
// func (h *OgentHandler) CreateCustomer(ctx context.Context, req *CreateCustomerReq) (CreateCustomerRes, error) {
// 	b := h.client.Customer.Create()
// 	// Add all fields.
// 	b.SetName(req.Name)
// 	b.SetLocation(req.Location)
// 	// Add all edges.
// 	b.AddProductIDs(req.Products...)
// 	// Persist to storage.
// 	e, err := b.Save(ctx)
// 	if err != nil {
// 		switch {
// 		case ent.IsNotSingular(err):
// 			return &R409{
// 				Code:   http.StatusConflict,
// 				Status: http.StatusText(http.StatusConflict),
// 				Errors: rawError(err),
// 			}, nil
// 		case ent.IsConstraintError(err):
// 			return &R409{
// 				Code:   http.StatusConflict,
// 				Status: http.StatusText(http.StatusConflict),
// 				Errors: rawError(err),
// 			}, nil
// 		default:
// 			// Let the server handle the error.
// 			return nil, err
// 		}
// 	}
// 	// Reload the entity to attach all eager-loaded edges.
// 	q := h.client.Customer.Query().Where(customer.ID(e.ID))
// 	e, err = q.Only(ctx)
// 	if err != nil {
// 		// This should never happen.
// 		return nil, err
// 	}
// 	return NewCustomerCreate(e), nil
// }

// // ReadCustomer handles GET /customers/{id} requests.
// func (h *OgentHandler) ReadCustomer(ctx context.Context, params ReadCustomerParams) (ReadCustomerRes, error) {
// 	q := h.client.Customer.Query().Where(customer.IDEQ(params.ID))
// 	e, err := q.Only(ctx)
// 	if err != nil {
// 		switch {
// 		case ent.IsNotFound(err):
// 			return &R404{
// 				Code:   http.StatusNotFound,
// 				Status: http.StatusText(http.StatusNotFound),
// 				Errors: rawError(err),
// 			}, nil
// 		case ent.IsNotSingular(err):
// 			return &R409{
// 				Code:   http.StatusConflict,
// 				Status: http.StatusText(http.StatusConflict),
// 				Errors: rawError(err),
// 			}, nil
// 		default:
// 			// Let the server handle the error.
// 			return nil, err
// 		}
// 	}
// 	return NewCustomerRead(e), nil
// }

// // UpdateCustomer handles PATCH /customers/{id} requests.
// func (h *OgentHandler) UpdateCustomer(ctx context.Context, req *UpdateCustomerReq, params UpdateCustomerParams) (UpdateCustomerRes, error) {
// 	b := h.client.Customer.UpdateOneID(params.ID)
// 	// Add all fields.
// 	if v, ok := req.Name.Get(); ok {
// 		b.SetName(v)
// 	}
// 	if v, ok := req.Location.Get(); ok {
// 		b.SetLocation(v)
// 	}
// 	// Add all edges.
// 	if req.Products != nil {
// 		b.ClearProducts().AddProductIDs(req.Products...)
// 	}
// 	// Persist to storage.
// 	e, err := b.Save(ctx)
// 	if err != nil {
// 		switch {
// 		case ent.IsNotFound(err):
// 			return &R404{
// 				Code:   http.StatusNotFound,
// 				Status: http.StatusText(http.StatusNotFound),
// 				Errors: rawError(err),
// 			}, nil
// 		case ent.IsConstraintError(err):
// 			return &R409{
// 				Code:   http.StatusConflict,
// 				Status: http.StatusText(http.StatusConflict),
// 				Errors: rawError(err),
// 			}, nil
// 		default:
// 			// Let the server handle the error.
// 			return nil, err
// 		}
// 	}
// 	// Reload the entity to attach all eager-loaded edges.
// 	q := h.client.Customer.Query().Where(customer.ID(e.ID))
// 	e, err = q.Only(ctx)
// 	if err != nil {
// 		// This should never happen.
// 		return nil, err
// 	}
// 	return NewCustomerUpdate(e), nil
// }

// // DeleteCustomer handles DELETE /customers/{id} requests.
// func (h *OgentHandler) DeleteCustomer(ctx context.Context, params DeleteCustomerParams) (DeleteCustomerRes, error) {
// 	err := h.client.Customer.DeleteOneID(params.ID).Exec(ctx)
// 	if err != nil {
// 		switch {
// 		case ent.IsNotFound(err):
// 			return &R404{
// 				Code:   http.StatusNotFound,
// 				Status: http.StatusText(http.StatusNotFound),
// 				Errors: rawError(err),
// 			}, nil
// 		case ent.IsConstraintError(err):
// 			return &R409{
// 				Code:   http.StatusConflict,
// 				Status: http.StatusText(http.StatusConflict),
// 				Errors: rawError(err),
// 			}, nil
// 		default:
// 			// Let the server handle the error.
// 			return nil, err
// 		}
// 	}
// 	return new(DeleteCustomerNoContent), nil

// }

// // ListCustomer handles GET /customers requests.
// func (h *OgentHandler) ListCustomer(ctx context.Context, params ListCustomerParams) (ListCustomerRes, error) {
// 	q := h.client.Customer.Query()
// 	page := 1
// 	if v, ok := params.Page.Get(); ok {
// 		page = v
// 	}
// 	itemsPerPage := 30
// 	if v, ok := params.ItemsPerPage.Get(); ok {
// 		itemsPerPage = v
// 	}
// 	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

// 	es, err := q.All(ctx)
// 	if err != nil {
// 		switch {
// 		case ent.IsNotFound(err):
// 			return &R404{
// 				Code:   http.StatusNotFound,
// 				Status: http.StatusText(http.StatusNotFound),
// 				Errors: rawError(err),
// 			}, nil
// 		case ent.IsNotSingular(err):
// 			return &R409{
// 				Code:   http.StatusConflict,
// 				Status: http.StatusText(http.StatusConflict),
// 				Errors: rawError(err),
// 			}, nil
// 		default:
// 			// Let the server handle the error.
// 			return nil, err
// 		}
// 	}
// 	r := NewCustomerLists(es)
// 	return (*ListCustomerOKApplicationJSON)(&r), nil
// }

// // ListCustomerProducts handles GET /customers/{id}/products requests.
// func (h *OgentHandler) ListCustomerProducts(ctx context.Context, params ListCustomerProductsParams) (ListCustomerProductsRes, error) {
// 	q := h.client.Customer.Query().Where(customer.IDEQ(params.ID)).QueryProducts()
// 	page := 1
// 	if v, ok := params.Page.Get(); ok {
// 		page = v
// 	}
// 	itemsPerPage := 30
// 	if v, ok := params.ItemsPerPage.Get(); ok {
// 		itemsPerPage = v
// 	}
// 	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)
// 	es, err := q.All(ctx)
// 	if err != nil {
// 		switch {
// 		case ent.IsNotFound(err):
// 			return &R404{
// 				Code:   http.StatusNotFound,
// 				Status: http.StatusText(http.StatusNotFound),
// 				Errors: rawError(err),
// 			}, nil
// 		case ent.IsNotSingular(err):
// 			return &R409{
// 				Code:   http.StatusConflict,
// 				Status: http.StatusText(http.StatusConflict),
// 				Errors: rawError(err),
// 			}, nil
// 		default:
// 			// Let the server handle the error.
// 			return nil, err
// 		}
// 	}
// 	r := NewCustomerProductsLists(es)
// 	return (*ListCustomerProductsOKApplicationJSON)(&r), nil
// }

// // CreateProduct handles POST /products requests.
// func (h *OgentHandler) CreateProduct(ctx context.Context, req *CreateProductReq) (CreateProductRes, error) {
// 	b := h.client.Product.Create()
// 	// Add all fields.
// 	b.SetName(req.Name)
// 	b.SetPrice(req.Price)
// 	b.SetCurrency(req.Currency)
// 	// Add all edges.
// 	// Persist to storage.
// 	e, err := b.Save(ctx)
// 	if err != nil {
// 		switch {
// 		case ent.IsNotSingular(err):
// 			return &R409{
// 				Code:   http.StatusConflict,
// 				Status: http.StatusText(http.StatusConflict),
// 				Errors: rawError(err),
// 			}, nil
// 		case ent.IsConstraintError(err):
// 			return &R409{
// 				Code:   http.StatusConflict,
// 				Status: http.StatusText(http.StatusConflict),
// 				Errors: rawError(err),
// 			}, nil
// 		default:
// 			// Let the server handle the error.
// 			return nil, err
// 		}
// 	}
// 	// Reload the entity to attach all eager-loaded edges.
// 	q := h.client.Product.Query().Where(product.ID(e.ID))
// 	e, err = q.Only(ctx)
// 	if err != nil {
// 		// This should never happen.
// 		return nil, err
// 	}
// 	return NewProductCreate(e), nil
// }

// // ReadProduct handles GET /products/{id} requests.
// func (h *OgentHandler) ReadProduct(ctx context.Context, params ReadProductParams) (ReadProductRes, error) {
// 	q := h.client.Product.Query().Where(product.IDEQ(params.ID))
// 	e, err := q.Only(ctx)
// 	if err != nil {
// 		switch {
// 		case ent.IsNotFound(err):
// 			return &R404{
// 				Code:   http.StatusNotFound,
// 				Status: http.StatusText(http.StatusNotFound),
// 				Errors: rawError(err),
// 			}, nil
// 		case ent.IsNotSingular(err):
// 			return &R409{
// 				Code:   http.StatusConflict,
// 				Status: http.StatusText(http.StatusConflict),
// 				Errors: rawError(err),
// 			}, nil
// 		default:
// 			// Let the server handle the error.
// 			return nil, err
// 		}
// 	}
// 	return NewProductRead(e), nil
// }

// // UpdateProduct handles PATCH /products/{id} requests.
// func (h *OgentHandler) UpdateProduct(ctx context.Context, req *UpdateProductReq, params UpdateProductParams) (UpdateProductRes, error) {
// 	b := h.client.Product.UpdateOneID(params.ID)
// 	// Add all fields.
// 	if v, ok := req.Name.Get(); ok {
// 		b.SetName(v)
// 	}
// 	if v, ok := req.Price.Get(); ok {
// 		b.SetPrice(v)
// 	}
// 	if v, ok := req.Currency.Get(); ok {
// 		b.SetCurrency(v)
// 	}
// 	// Add all edges.
// 	// Persist to storage.
// 	e, err := b.Save(ctx)
// 	if err != nil {
// 		switch {
// 		case ent.IsNotFound(err):
// 			return &R404{
// 				Code:   http.StatusNotFound,
// 				Status: http.StatusText(http.StatusNotFound),
// 				Errors: rawError(err),
// 			}, nil
// 		case ent.IsConstraintError(err):
// 			return &R409{
// 				Code:   http.StatusConflict,
// 				Status: http.StatusText(http.StatusConflict),
// 				Errors: rawError(err),
// 			}, nil
// 		default:
// 			// Let the server handle the error.
// 			return nil, err
// 		}
// 	}
// 	// Reload the entity to attach all eager-loaded edges.
// 	q := h.client.Product.Query().Where(product.ID(e.ID))
// 	e, err = q.Only(ctx)
// 	if err != nil {
// 		// This should never happen.
// 		return nil, err
// 	}
// 	return NewProductUpdate(e), nil
// }

// // DeleteProduct handles DELETE /products/{id} requests.
// func (h *OgentHandler) DeleteProduct(ctx context.Context, params DeleteProductParams) (DeleteProductRes, error) {
// 	err := h.client.Product.DeleteOneID(params.ID).Exec(ctx)
// 	if err != nil {
// 		switch {
// 		case ent.IsNotFound(err):
// 			return &R404{
// 				Code:   http.StatusNotFound,
// 				Status: http.StatusText(http.StatusNotFound),
// 				Errors: rawError(err),
// 			}, nil
// 		case ent.IsConstraintError(err):
// 			return &R409{
// 				Code:   http.StatusConflict,
// 				Status: http.StatusText(http.StatusConflict),
// 				Errors: rawError(err),
// 			}, nil
// 		default:
// 			// Let the server handle the error.
// 			return nil, err
// 		}
// 	}
// 	return new(DeleteProductNoContent), nil

// }

// // ListProduct handles GET /products requests.
// func (h *OgentHandler) ListProduct(ctx context.Context, params ListProductParams) (ListProductRes, error) {
// 	q := h.client.Product.Query()
// 	page := 1
// 	if v, ok := params.Page.Get(); ok {
// 		page = v
// 	}
// 	itemsPerPage := 30
// 	if v, ok := params.ItemsPerPage.Get(); ok {
// 		itemsPerPage = v
// 	}
// 	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

// 	es, err := q.All(ctx)
// 	if err != nil {
// 		switch {
// 		case ent.IsNotFound(err):
// 			return &R404{
// 				Code:   http.StatusNotFound,
// 				Status: http.StatusText(http.StatusNotFound),
// 				Errors: rawError(err),
// 			}, nil
// 		case ent.IsNotSingular(err):
// 			return &R409{
// 				Code:   http.StatusConflict,
// 				Status: http.StatusText(http.StatusConflict),
// 				Errors: rawError(err),
// 			}, nil
// 		default:
// 			// Let the server handle the error.
// 			return nil, err
// 		}
// 	}
// 	r := NewProductLists(es)
// 	return (*ListProductOKApplicationJSON)(&r), nil
// }
