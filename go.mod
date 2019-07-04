module github.com/micro/examples

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f

replace github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1

require (
	github.com/99designs/gqlgen v0.9.1
	github.com/emicklei/go-restful v2.8.1+incompatible
	github.com/gin-contrib/sse v0.0.0-20190125020943-a7658810eb74 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.1
	github.com/google/uuid v1.1.1
	github.com/gorilla/websocket v1.4.0
	github.com/grpc-ecosystem/grpc-gateway v1.9.2
	github.com/hailocab/go-geoindex v0.0.0-20160127134810-64631bfe9711
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro/agent v0.1.0
	github.com/micro/go-micro/config v0.13.3
	github.com/micro/go-micro/service/grpc v0.9.0
	github.com/micro/go-micro/util/log v0.1.0
	github.com/micro/go-micro v1.7.1-0.20190627135301-d8e998ad85fe
	github.com/micro/go-plugins v1.1.1
	github.com/micro/go-web v0.6.0
	github.com/micro/micro v1.7.1-0.20190627121529-410a2eba67f1
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d
	github.com/pborman/uuid v1.2.0
	github.com/ugorji/go v1.1.5-pre // indirect
	github.com/vektah/gqlparser v1.1.2
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	google.golang.org/genproto v0.0.0-20190626174449-989357319d63
	google.golang.org/grpc v1.21.1
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
)