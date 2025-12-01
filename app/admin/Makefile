install_hz_latest:
	go install github.com/cloudwego/hertz/cmd/hz@latest

new_facade:
	hz new --idl=../../idl/api/order.thrift --service=facade -f -use=deer/hertz_gen

update_facade:
	hz update --idl=../../idl/admin/user.thrift -use=deer/hertz_gen/
	hz update --idl=../../idl/admin/dict.thrift -use=deer/hertz_gen/
	hz update --idl=../../idl/admin/menu.thrift -use=deer/hertz_gen/
	hz update --idl=../../idl/admin/role.thrift -use=deer/hertz_gen/
	hz update --idl=../../idl/admin/system.thrift -use=deer/hertz_gen/

	hz update --idl=../../idl/admin/departments.thrift -use=deer/hertz_gen/
	hz update --idl=../../idl/admin/positions.thrift -use=deer/hertz_gen/



