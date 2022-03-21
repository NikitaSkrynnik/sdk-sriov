module github.com/networkservicemesh/sdk-sriov

go 1.16

require (
	github.com/ghodss/yaml v1.0.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/networkservicemesh/api v1.2.1-0.20220315001249-f33f8c3f2feb
	github.com/networkservicemesh/sdk v0.5.1-0.20220316101237-288caa7bbc1c
	github.com/networkservicemesh/sdk-kernel v0.0.0-20220316101641-0103343013f0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	go.uber.org/goleak v1.1.12
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/sys v0.0.0-20220307203707-22a9840ba4d7
	google.golang.org/grpc v1.42.0
)

replace github.com/networkservicemesh/sdk-kernel => github.com/NikitaSkrynnik/sdk-kernel v0.0.0-20220321082030-704de772333a
