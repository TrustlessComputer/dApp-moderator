package rediskey

const (
	GetPackageStockByPackageInfoPrefix  = "GetPackageStockByPackageInfo"
	GetTrackingCodesByOrderDetailPrefix = "TrackingCode"
	GetPackageInfoByPackageData         = "GetPackageInfoByPackageData_"
	PackagesOfProduct                   = "PackagesOfProduct"
	BannerForSaleProduct                = "BannerForSaleProduct"
	TotalA384AreSold                    = "TotalA384AreSold"
	ProductInfoCart                     = "ProductInfoCart"
	BambooLimitIsSoldOut                = "BambooLimitIsSoldOut"
	AllowEditSkusInProductVendor        = "AllowEditSkusInProductVendor"
	VendorUser                          = "VendorUser"
	Customer                            = "customer"
	Product                             = "Product"
	TopThreeProductsBestSellerByDomain  = "TopThreeProductsBestSellerByDomain"
	TaxRate                             = "TaxRate"
	concatStr                           = '_'
)

var (
	GetTrackingCodesByOrderDetail = func(uuid int) string {
		return Beauty(GetTrackingCodesByOrderDetailPrefix).WithUUID(uuid).String()
	}

	// GetPackageStockByPackageInfo get redis key of dao func GetPackageStockByPackageInfo
	// This func is used in bff-service and cart-worker
	GetPackageStockByPackageInfo = func(packageInfoID, warehouseID int, country string) string {
		return Beauty(GetPackageStockByPackageInfoPrefix).WithParams(
			packageInfoID,
			warehouseID,
			country,
		).String()
	}
)
