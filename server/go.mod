module server

go 1.20

require (
	github.com/eclipse/paho.mqtt.golang v1.4.2
	github.com/gocql/gocql v1.4.0
	github.com/gofiber/fiber/v2 v2.46.0
	github.com/joho/godotenv v1.5.1
	github.com/scylladb/gocqlx/v2 v2.8.0
	github.com/smartystreets/assertions v1.13.1
	github.com/spf13/viper v1.16.0
	github.com/urfave/cli/v2 v2.25.5
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
)

replace github.com/gocql/gocql v1.4.0 => github.com/scylladb/gocql v1.10.0

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/klauspost/compress v1.16.3 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/savsgio/dictpool v0.0.0-20221023140959-7bf2e61cea94 // indirect
	github.com/savsgio/gotils v0.0.0-20230208104028-c358bd845dee // indirect
	github.com/scylladb/go-reflectx v1.0.1 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/tinylib/msgp v1.1.8 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.47.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
