module github.com/teamgram/teamgram-server

go 1.16

require (
	github.com/bwmarrin/snowflake v0.3.0
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/go-ini/ini v1.63.2 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/minio/minio-go v6.0.14+incompatible // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/nyaruka/phonenumbers v1.0.74 // indirect
	github.com/oschwald/geoip2-golang v1.5.0
	github.com/panjf2000/gnet v1.5.3
	github.com/teamgram/marmota v0.0.0-20220224022037-4b5ab312525f
	github.com/teamgram/proto v0.0.0-20220226015348-e671d4bec270
	github.com/zeromicro/go-zero v1.3.0
	google.golang.org/grpc v1.43.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/mail.v2 v2.3.1 // indirect
	mvdan.cc/xurls/v2 v2.3.0 // indirect
)

replace (
	github.com/panjf2000/gnet => github.com/teamgram/gnet v1.6.5-0.20220203114726-06bfacbd8548
	github.com/zeromicro/go-zero v1.3.0 => github.com/teamgramio/go-zero v1.3.1-0.20220221143929-c2897dcaf14c
)
