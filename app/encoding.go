package app

import (
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	"github.com/xpladev/xpla/app/params"

	"github.com/xpladev/ethermint/encoding/codec"
)

// MakeTestEncodingConfig creates an EncodingConfig for testing. This function
// should be used only in tests or when creating a new app instance (NewApp*()).
// App user shouldn't create new codecs - use the app.AppCodec instead.
// [DEPRECATED]
func MakeTestEncodingConfig() params.EncodingConfig {
	encodingConfig := params.MakeTestEncodingConfig()
	codec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	codec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	// for removing vesting module
	vestingtypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	return encodingConfig
}
