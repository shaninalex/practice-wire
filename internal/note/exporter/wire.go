//go:build wireinject
// +build wireinject

package exporter

import (
	"github.com/google/wire"
	"github.com/shaninalex/practice-wire/internal/domain"
)

var (
	WireJSON = wire.NewSet(
		NewJSONExporter,
		wire.Bind(new(domain.IExporter), new(*JSONExporter)),
	)

	WireCSV = wire.NewSet(
		NewCSVExporter,
		wire.Bind(new(domain.IExporter), new(*CSVExporter)),
	)

	WireXML = wire.NewSet(
		NewXMLExporter,
		wire.Bind(new(domain.IExporter), new(*XMLExporter)),
	)

	WireMarkdown = wire.NewSet(
		NewMarkdownExporter,
		wire.Bind(new(domain.IExporter), new(*MarkdownExporter)),
	)
)

func ProvideJSONExporter(storage domain.IStorage) domain.IExporter {
	wire.Build(WireJSON)
	return nil
}

func ProvideCSVExporter(storage domain.IStorage) domain.IExporter {
	wire.Build(WireCSV)
	return nil
}

func ProvideXMLExporter(storage domain.IStorage) domain.IExporter {
	wire.Build(WireXML)
	return nil
}

func ProvideMarkdownExporter(storage domain.IStorage) domain.IExporter {
	wire.Build(WireMarkdown)
	return nil
}
