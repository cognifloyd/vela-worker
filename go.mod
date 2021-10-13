module github.com/go-vela/worker

go 1.16

require (
	github.com/Masterminds/semver/v3 v3.1.1
	github.com/gin-gonic/gin v1.7.4
	github.com/go-vela/pkg-executor v0.10.0-rc1
	github.com/go-vela/pkg-queue v0.10.0-rc1
	github.com/go-vela/pkg-runtime v0.10.0-rc1
	github.com/go-vela/sdk-go v0.10.0-rc1
	github.com/go-vela/types v0.10.0-rc1
	github.com/joho/godotenv v1.4.0
	github.com/prometheus/client_golang v1.11.0
	github.com/sirupsen/logrus v1.8.1
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
)

replace (
	// k8s-post-assemble-build branch
	github.com/go-vela/pkg-executor => github.com/cognifloyd/pkg-executor v0.10.0-rc1.0.20211012235243-eff623a928d1
	// k8s-post-assemble-build branch
	github.com/go-vela/pkg-runtime => github.com/cognifloyd/pkg-runtime v0.10.0-rc1.0.20211013005510-554c4a12bea4
	// k8s-dns-names branch
	github.com/go-vela/types => github.com/cognifloyd/vela-types v0.10.0-rc1.0.20211011222123-7045496b403e
)
