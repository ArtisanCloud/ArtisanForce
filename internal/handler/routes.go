// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	adminauth "PowerX/internal/handler/admin/auth"
	admincommon "PowerX/internal/handler/admin/common"
	admincontractway "PowerX/internal/handler/admin/contractway"
	admincustomer "PowerX/internal/handler/admin/customer"
	admindepartment "PowerX/internal/handler/admin/department"
	admindictionary "PowerX/internal/handler/admin/dictionary"
	adminemployee "PowerX/internal/handler/admin/employee"
	adminlead "PowerX/internal/handler/admin/lead"
	adminmedia "PowerX/internal/handler/admin/media"
	adminopportunity "PowerX/internal/handler/admin/opportunity"
	adminpermission "PowerX/internal/handler/admin/permission"
	adminproduct "PowerX/internal/handler/admin/product"
	adminscrmcontact "PowerX/internal/handler/admin/scrm/contact"
	adminscrmcustomer "PowerX/internal/handler/admin/scrm/customer"
	adminuserinfo "PowerX/internal/handler/admin/userinfo"
	mpcustomer "PowerX/internal/handler/mp/customer"
	"PowerX/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeNoPermJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/options/employees",
					Handler: admincommon.GetEmployeeOptionsHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/options/employee-query",
					Handler: admincommon.GetEmployeeQueryOptionsHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/options/departments",
					Handler: admincommon.GetDepartmentOptionsHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/common"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/department-tree/:depId",
					Handler: admindepartment.GetDepartmentTreeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/departments/:id",
					Handler: admindepartment.GetDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/departments",
					Handler: admindepartment.CreateDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/departments/:id",
					Handler: admindepartment.PatchDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/departments/:id",
					Handler: admindepartment.DeleteDepartmentHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/department"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/employees/actions/sync",
					Handler: adminemployee.SyncEmployeesHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/employees/:id",
					Handler: adminemployee.GetEmployeeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/employees",
					Handler: adminemployee.ListEmployeesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/employees",
					Handler: adminemployee.CreateEmployeeHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/employees/:id",
					Handler: adminemployee.UpdateEmployeeHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/employees/:id",
					Handler: adminemployee.DeleteEmployeeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/employees/actions/reset-password",
					Handler: adminemployee.ResetPasswordHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/employee"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/roles",
					Handler: adminpermission.ListRolesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/roles",
					Handler: adminpermission.CreateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/roles/:roleCode",
					Handler: adminpermission.GetRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/roles/:roleCode",
					Handler: adminpermission.PatchRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/roles/:roleCode/users",
					Handler: adminpermission.GetRoleEmployeesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/roles/:roleCode/actions/set-permissions",
					Handler: adminpermission.SetRolePermissionsHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api-list",
					Handler: adminpermission.ListAPIHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/roles/:roleCode/actions/set-employees",
					Handler: adminpermission.SetRoleEmployeesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/users/:userId/actions/set-roles",
					Handler: adminpermission.SetUserRolesHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/permission"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/access/actions/basic-login",
				Handler: adminauth.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/access/actions/exchange-token",
				Handler: adminauth.ExchangeHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/admin/auth"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/types",
					Handler: admindictionary.ListDictionaryTypesHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/types/:type",
					Handler: admindictionary.GetDictionaryTypeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/types",
					Handler: admindictionary.CreateDictionaryTypeHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/types/:type",
					Handler: admindictionary.UpdateDictionaryTypeHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/types/:type",
					Handler: admindictionary.DeleteDictionaryTypeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/items",
					Handler: admindictionary.ListDictionaryItemsHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/items/:type/:key",
					Handler: admindictionary.GetDictionaryItemHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/items",
					Handler: admindictionary.CreateDictionaryItemHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/items/:type/:key",
					Handler: admindictionary.UpdateDictionaryItemHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/items/:type/:key",
					Handler: admindictionary.DeleteDictionaryItemHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/dictionary"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user-info",
					Handler: adminuserinfo.GetUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/menu-roles",
					Handler: adminuserinfo.GetMenuRolesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/users/actions/modify-password",
					Handler: adminuserinfo.ModifyUserPasswordHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/user-center"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/leads",
					Handler: adminlead.ListLeadsHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/leads",
					Handler: adminlead.CreateLeadHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/leads/:id",
					Handler: adminlead.PatchLeadHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/leads/:id",
					Handler: adminlead.DeleteLeadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/leads/:id/actions/assign-to-employee",
					Handler: adminlead.AssignLeadToEmployeeHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/lead"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/customers/:id",
					Handler: admincustomer.GetCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/customers",
					Handler: admincustomer.ListCustomersHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/customers",
					Handler: admincustomer.CreateCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/customers/:id",
					Handler: admincustomer.PatchCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/customers/:id",
					Handler: admincustomer.DeleteCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/customers/:id/actions/employees",
					Handler: admincustomer.AssignCustomerToEmployeeHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/customer"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/medias",
					Handler: adminmedia.GetMediaListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/medias/actions/create-upload-url",
					Handler: adminmedia.CreateMediaUploadRequestHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/medias/:mediaKey",
					Handler: adminmedia.CreateOrUpdateMediaHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/medias/:key",
					Handler: adminmedia.GetMediaByKeyHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/medias/:key",
					Handler: adminmedia.DeleteMediaHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/media"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/opportunities",
					Handler: adminopportunity.GetOpportunityListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/opportunities",
					Handler: adminopportunity.CreateOpportunityHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/opportunities/:id/assign-employee",
					Handler: adminopportunity.AssignEmployeeToOpportunityHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/opportunities/:id",
					Handler: adminopportunity.UpdateOpportunityHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/opportunities/:id",
					Handler: adminopportunity.DeleteOpportunityHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/opportunity"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/products",
					Handler: adminproduct.ListProductsHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/products/:id",
					Handler: adminproduct.GetProductHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/products",
					Handler: adminproduct.CreateProductHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/products/:id",
					Handler: adminproduct.PutProductHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/products/:id",
					Handler: adminproduct.PatchProductHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/products/:id",
					Handler: adminproduct.DeleteProductHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/products/:id/actions/assign-to-product-categroy",
					Handler: adminproduct.AssignProductToProductCategoryHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/product"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/product-category-tree",
					Handler: adminproduct.ListProductCategoryTreeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/product-categories/:id",
					Handler: adminproduct.GetProductCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/product-categories",
					Handler: adminproduct.UpsertProductCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/product-categories/:id",
					Handler: adminproduct.PatchProductCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/product-categories/:id",
					Handler: adminproduct.DeleteProductCategoryHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/product"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/price-books",
					Handler: adminproduct.ListPriceBooksHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/price-books/:id",
					Handler: adminproduct.GetPriceBookHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/price-books",
					Handler: adminproduct.UpsertPriceBookHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/price-books/:id",
					Handler: adminproduct.DeletePriceBookHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/price-book-entries",
					Handler: adminproduct.ConfigPriceBookHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/product"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/stores",
					Handler: adminproduct.ListStoresHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/stores/:id",
					Handler: adminproduct.GetStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/stores",
					Handler: adminproduct.CreateStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/stores/:id",
					Handler: adminproduct.PutStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/stores/:id",
					Handler: adminproduct.PatchStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/stores/:id",
					Handler: adminproduct.DeleteStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/stores/:id/actions/assign-to-store-categroy",
					Handler: adminproduct.AssignStoreToStoreCategoryHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/product"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/artisans",
					Handler: adminproduct.ListArtisansHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/artisans/:id",
					Handler: adminproduct.GetArtisanHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/artisans",
					Handler: adminproduct.CreateArtisanHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/artisans/:id",
					Handler: adminproduct.PutArtisanHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/artisans/:id",
					Handler: adminproduct.PatchArtisanHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/artisans/:id",
					Handler: adminproduct.DeleteArtisanHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/artisans/:id/actions/assign-to-store-categroy",
					Handler: adminproduct.AssignArtisanToArtisanCategoryHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/product"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/group-tree",
					Handler: admincontractway.GetContractWayGroupTreeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/groups",
					Handler: admincontractway.GetContractWayGroupListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: admincontractway.GetContractWaysHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: admincontractway.CreateContractWayHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/:id",
					Handler: admincontractway.UpdateContractWayHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/:id",
					Handler: admincontractway.DeleteContractWayHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/contract-way"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/customers/:id",
					Handler: adminscrmcustomer.GetWeWorkCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/customers",
					Handler: adminscrmcustomer.ListWeWorkCustomersHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/customers/:id",
					Handler: adminscrmcustomer.PatchWeWorkCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/customers/actions/sync",
					Handler: adminscrmcustomer.SyncWeWorkCustomerHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/scrm/customer"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/contacts/actions/sync",
					Handler: adminscrmcontact.SyncWeWorkContactHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/employees",
					Handler: adminscrmcontact.ListWeWorkEmployeeHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/scrm/contact"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.MPCustomerJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/login",
					Handler: mpcustomer.LoginHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authByPhone",
					Handler: mpcustomer.AuthByPhoneHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authByProfile",
					Handler: mpcustomer.AuthByProfileHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/mp/customer"),
	)
}
