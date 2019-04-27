protoc -I jarviscrawlercore/ jarviscrawlercore/result.proto --go_out=plugins=grpc:jarviscrawlercore
protoc -I proto/ proto/dtdata.proto --go_out=plugins=grpc:proto