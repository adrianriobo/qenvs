package mac

const (
	stackDedicatedHost = "stackDedicatedHost"
	stackMacMachine    = "stackMacMachine"

	awsMacMachineID = "amm"

	urnLock = "rh:qe:aws:mac:lock"

	outputLock              = "ammLock"
	outputHost              = "ammHost"
	outputUsername          = "ammUsername"
	outputUserPassword      = "ammUserPassword"
	outputMachinePrivateKey = "ammMachinePrivatekey"
	outputUserPrivateKey    = "ammUserPrivatekey"
	outputDedicatedHostID   = "ammDedicatedHostID"
	outputDedicatedHostAZ   = "ammDedicatedHostAZ"
	outputRegion            = "ammRegion"
	// outputAdminUsername     = "ammAdminUsername"
	// outputAdminUserPassword = "ammAdminUserPassword"

	amiRegex = "amzn-ec2-macos-%s*"
	amiOwner = "628277914472"

	vncDefaultPort  int    = 5900
	diskSize        int    = 200
	defaultUsername string = "ec2-user"
	defaultSSHPort  int    = 22

	// https://www.pulumi.com/docs/intro/concepts/resources/options/customtimeouts/
	remoteTimeout string = "40m"

	backedURLTagName string = "BackedURL"
	archTagName      string = "Arch"
)

var macTypesByArch = map[string]string{
	"x86": "mac1.metal",
	"m1":  "mac2.metal",
	"m2":  "mac2-m2pro.metal"}

var awsArchIDbyArch = map[string]string{
	"x86": "x86_64_mac",
	"m1":  "arm64_mac",
	"m2":  "arm64_mac"}

// var macAMIs = map[string]string{
// 	"arm64_mac-13":  "macos-arm64-13.6.1",
// 	"arm64_mac-14":  "macos-arm64-14.1",
// 	"x86_64_mac-12": "mac12_x86",
// 	"x86_64_mac-13": "mac13_x86"}
