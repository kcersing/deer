install_hz_latest:
	go install github.com/cloudwego/hertz/cmd/hz@latest

new_facade:
	hz new --idl=../../idl/facade/order.thrift --service=facade -f -use=deer/hertz_gen

update_facade:
	hz update --idl=../../idl/facade/member.thrift -use=deer/hertz_gen/
