module github.com/dihedron/snoop

go 1.23.4

require (
	github.com/dihedron/rawdata v1.0.1
	github.com/go-playground/validator/v10 v10.23.0
	github.com/jessevdk/go-flags v1.6.1
	github.com/joho/godotenv v1.5.1
	github.com/juju/rfc/v2 v2.0.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/rabbitmq/amqp091-go v1.10.0
	github.com/streamdal/rabbit v0.1.26
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/gabriel-vasile/mimetype v1.4.7 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/juju/errors v1.0.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)

replace github.com/streamdal/rabbit => github.com/dihedron/rabbit v0.1.15-0.20250108113512-27f0017f4655
