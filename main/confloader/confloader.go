package confloader

import (
	"io"
	"os"
)

type configFileLoader func(string) (io.ReadCloser, error)

var (
	EffectiveConfigFileLoader configFileLoader
)

func LoadConfig(file string) (io.ReadCloser, error) {
	if EffectiveConfigFileLoader == nil {
		return os.Stdin, nil
	}
	return EffectiveConfigFileLoader(file)
}

// export CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 CC=/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/clang CXX=/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/clang++ CGO_CFLAGS="-isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/iPhoneOS13.2.sdk -miphoneos-version-min=7.0 -fembed-bitcode -arch arm64" CGO_CXXFLAGS="-isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/iPhoneOS13.2.sdk -miphoneos-version-min=7.0 -fembed-bitcode -arch arm64" CGO_LDFLAGS="-isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/iPhoneOS13.2.sdk -miphoneos-version-min=7.0 -fembed-bitcode -arch arm64" CGO_ENABLED=1 GOPATH=$GOPATH GO111MODULE=off go build -tags ios -prefix V2Ray -o $HOME/v2ray -buildmode=c-archive -ldflags "-s -w" -x -v
