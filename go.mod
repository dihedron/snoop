module github.com/dihedron/snoop

go 1.24.0

require (
	github.com/Masterminds/sprig/v3 v3.3.0
	github.com/dihedron/rawdata v1.0.1
	github.com/fatih/color v1.18.0
	github.com/go-playground/validator/v10 v10.25.0
	github.com/goccy/go-json v0.10.5
	github.com/jessevdk/go-flags v1.6.1
	github.com/joho/godotenv v1.5.1
	github.com/juju/rfc/v2 v2.0.0
	github.com/neilotoole/slogt v1.1.0
	github.com/rabbitmq/amqp091-go v1.10.0
	github.com/schollz/progressbar/v3 v3.18.0
	github.com/streamdal/rabbit v0.1.26
	gopkg.in/yaml.v3 v3.0.1
)

require (
	dario.cat/mergo v1.0.1 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.3.1 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/huandu/xstrings v1.5.0 // indirect
	github.com/juju/errors v1.0.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/term v0.30.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)

//replace github.com/streamdal/rabbit => github.com/dihedron/rabbit v0.1.15-0.20250108113512-27f0017f4655

replace github.com/streamdal/rabbit => ../rabbit
