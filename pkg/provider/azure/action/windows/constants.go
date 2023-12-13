package windows

const (
	stackCreateWindowsDesktop = "stackCreateWindowsDesktop"
	stackSyncWindowsDesktop   = "stackSyncWindowsDesktop"

	azureWindowsDesktopID = "awd"

	cidrVN = "10.0.0.0/16"
	cidrSN = "10.0.2.0/24"

	scriptName = "setup.ps1"

	outputHost              = "awdHost"
	outputUsername          = "awdUsername"
	outputUserPassword      = "awdUserPassword"
	outputUserPrivateKey    = "awdUserPrivatekey"
	outputAdminUsername     = "awdAdminUsername"
	outputAdminUserPassword = "awdAdminUserPassword"

	prioritySpot = "Spot"
)

var defaultVMSizeByArch = map[string]string{
	"x86": "Standard_D8s_v4",
	"m1":  "mac2.metal",
	"m2":  "mac2-m2pro.metal"}

var awsArchIDbyArch = map[string]string{
	"x86": "x86_64_mac",
	"m1":  "arm64_mac",
	"m2":  "arm64_mac"}
