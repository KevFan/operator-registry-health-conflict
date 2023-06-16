Example project showing `#namespace-conflict` error when importing both of the following packages from `operator-registry`
on go1.19:

```
	"github.com/operator-framework/operator-registry/pkg/client"
	"github.com/operator-framework/operator-registry/pkg/lib/bundle"
```

To show error:
```shell
 go run ./main/main.go  
```

will result in:

```shell
panic: proto: file "grpc/health/v1/health.proto" has a name conflict over grpc.health.v1.HealthCheckResponse
        previously from: "github.com/operator-framework/operator-registry/pkg/api/grpc_health_v1"
        currently from:  "google.golang.org/grpc/health/grpc_health_v1"
See https://protobuf.dev/reference/go/faq#namespace-conflict


goroutine 1 [running]:
google.golang.org/protobuf/reflect/protoregistry.glob..func1({0x1065e3b60?, 0x14000621050?}, {0x1065e3b60?, 0x14000621090})
        /Users/chfan/.gvm/pkgsets/go1.20.1/global/pkg/mod/google.golang.org/protobuf@v1.29.1/reflect/protoregistry/registry.go:56 +0x200
google.golang.org/protobuf/reflect/protoregistry.(*Files).RegisterFile.func1({0x1066037b8, 0x14000682410})
        /Users/chfan/.gvm/pkgsets/go1.20.1/global/pkg/mod/google.golang.org/protobuf@v1.29.1/reflect/protoregistry/registry.go:154 +0x23c
google.golang.org/protobuf/reflect/protoregistry.rangeTopLevelDescriptors({0x1066098e0, 0x14000550c40}, 0x140008beb40)
        /Users/chfan/.gvm/pkgsets/go1.20.1/global/pkg/mod/google.golang.org/protobuf@v1.29.1/reflect/protoregistry/registry.go:417 +0x13c
google.golang.org/protobuf/reflect/protoregistry.(*Files).RegisterFile(0x1400000e150, {0x1066098e0?, 0x14000550c40?})
        /Users/chfan/.gvm/pkgsets/go1.20.1/global/pkg/mod/google.golang.org/protobuf@v1.29.1/reflect/protoregistry/registry.go:149 +0x5f4
google.golang.org/protobuf/internal/filedesc.Builder.Build({{0x1062c6b09, 0x2c}, {0x1072d0d00, 0x22d, 0x22d}, 0x1, 0x2, 0x0, 0x1, {0x1065e86a8, ...}, ...})
        /Users/chfan/.gvm/pkgsets/go1.20.1/global/pkg/mod/google.golang.org/protobuf@v1.29.1/internal/filedesc/build.go:112 +0x1a4
google.golang.org/protobuf/internal/filetype.Builder.Build({{{0x1062c6b09, 0x2c}, {0x1072d0d00, 0x22d, 0x22d}, 0x1, 0x2, 0x0, 0x1, {0x0, ...}, ...}, ...})
        /Users/chfan/.gvm/pkgsets/go1.20.1/global/pkg/mod/google.golang.org/protobuf@v1.29.1/internal/filetype/build.go:138 +0x17c
google.golang.org/grpc/health/grpc_health_v1.file_grpc_health_v1_health_proto_init()
        /Users/chfan/.gvm/pkgsets/go1.20.1/global/pkg/mod/google.golang.org/grpc@v1.53.0/health/grpc_health_v1/health.pb.go:303 +0x1ac
google.golang.org/grpc/health/grpc_health_v1.init.0()
        /Users/chfan/.gvm/pkgsets/go1.20.1/global/pkg/mod/google.golang.org/grpc@v1.53.0/health/grpc_health_v1/health.pb.go:258 +0x1c
exit status 2

```

Using the `GOLANG_PROTOBUF_REGISTRATION_CONFLICT` flag would show the following with conflicts:
```shell
 GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn go run ./main/main.go
WARNING: proto: file "grpc/health/v1/health.proto" has a name conflict over grpc.health.v1.HealthCheckResponse
        previously from: "github.com/operator-framework/operator-registry/pkg/api/grpc_health_v1"
        currently from:  "google.golang.org/grpc/health/grpc_health_v1"
See https://protobuf.dev/reference/go/faq#namespace-conflict

WARNING: proto: file "grpc/health/v1/health.proto" has a name conflict over grpc.health.v1.HealthCheckRequest
        previously from: "github.com/operator-framework/operator-registry/pkg/api/grpc_health_v1"
        currently from:  "google.golang.org/grpc/health/grpc_health_v1"
See https://protobuf.dev/reference/go/faq#namespace-conflict

WARNING: proto: file "grpc/health/v1/health.proto" has a name conflict over grpc.health.v1.Health
        previously from: "github.com/operator-framework/operator-registry/pkg/api/grpc_health_v1"
        currently from:  "google.golang.org/grpc/health/grpc_health_v1"
See https://protobuf.dev/reference/go/faq#namespace-conflict

WARNING: proto: enum grpc.health.v1.HealthCheckResponse.ServingStatus is already registered
        previously from: "github.com/operator-framework/operator-registry/pkg/api/grpc_health_v1"
        currently from:  "google.golang.org/grpc/health/grpc_health_v1"
See https://protobuf.dev/reference/go/faq#namespace-conflict

WARNING: proto: message grpc.health.v1.HealthCheckRequest is already registered
        previously from: "github.com/operator-framework/operator-registry/pkg/api/grpc_health_v1"
        currently from:  "google.golang.org/grpc/health/grpc_health_v1"
See https://protobuf.dev/reference/go/faq#namespace-conflict

WARNING: proto: message grpc.health.v1.HealthCheckResponse is already registered
        previously from: "github.com/operator-framework/operator-registry/pkg/api/grpc_health_v1"
        currently from:  "google.golang.org/grpc/health/grpc_health_v1"
See https://protobuf.dev/reference/go/faq#namespace-conflict

```