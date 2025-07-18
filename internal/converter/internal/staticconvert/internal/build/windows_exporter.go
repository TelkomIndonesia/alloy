package build

import (
	"strings"

	"github.com/grafana/alloy/internal/component/discovery"
	"github.com/grafana/alloy/internal/component/prometheus/exporter/windows"
	"github.com/grafana/alloy/internal/static/integrations/windows_exporter"
)

func (b *ConfigBuilder) appendWindowsExporter(config *windows_exporter.Config, instanceKey *string) discovery.Exports {
	args := toWindowsExporter(config)
	return b.appendExporterBlock(args, config.Name(), instanceKey, "windows")
}

// Splits a string such as "example1,example2"
// into a list such as []string{"example1", "example2"}.
func split(collectorList string) []string {
	if collectorList == "" {
		return []string{}
	}
	return strings.Split(collectorList, ",")
}

func toWindowsExporter(config *windows_exporter.Config) *windows.Arguments {
	return &windows.Arguments{
		EnabledCollectors: split(config.EnabledCollectors),
		Dfsr: windows.DfsrConfig{
			SourcesEnabled: split(config.Dfsr.SourcesEnabled),
		},
		Exchange: windows.ExchangeConfig{
			EnabledList: split(config.Exchange.EnabledList),
		},
		Filetime: windows.FiletimeConfig{
			FilePatterns: config.Filetime.FilePatterns,
		},
		IIS: windows.IISConfig{
			AppBlackList:  config.IIS.AppBlackList,
			AppWhiteList:  config.IIS.AppWhiteList,
			SiteBlackList: config.IIS.SiteBlackList,
			SiteWhiteList: config.IIS.SiteWhiteList,
			AppExclude:    config.IIS.AppExclude,
			AppInclude:    config.IIS.AppInclude,
			SiteExclude:   config.IIS.SiteExclude,
			SiteInclude:   config.IIS.SiteInclude,
		},
		LogicalDisk: windows.LogicalDiskConfig{
			BlackList: config.LogicalDisk.BlackList,
			WhiteList: config.LogicalDisk.WhiteList,
			Include:   config.LogicalDisk.Include,
			Exclude:   config.LogicalDisk.Exclude,
		},
		MSSQL: windows.MSSQLConfig{
			EnabledClasses: split(config.MSSQL.EnabledClasses),
		},
		MSCluster: windows.MSClusterConfig{
			EnabledList: split(config.MSCluster.EnabledList),
		},
		Network: windows.NetworkConfig{
			BlackList: config.Network.BlackList,
			WhiteList: config.Network.WhiteList,
			Exclude:   config.Network.Exclude,
			Include:   config.Network.Include,
		},
		NetFramework: windows.NetFrameworkConfig{
			EnabledList: split(config.NetFramework.EnabledList),
		},
		PerformanceCounter: windows.PerformanceCounterConfig{
			Objects: config.PerformanceCounter.Objects,
		},
		PhysicalDisk: windows.PhysicalDiskConfig{
			Exclude: config.PhysicalDisk.Exclude,
			Include: config.PhysicalDisk.Include,
		},
		Printer: windows.PrinterConfig{
			Exclude: config.Printer.Exclude,
			Include: config.Printer.Include,
		},
		Process: windows.ProcessConfig{
			BlackList: config.Process.BlackList,
			WhiteList: config.Process.WhiteList,
			Exclude:   config.Process.Exclude,
			Include:   config.Process.Include,
		},
		ScheduledTask: windows.ScheduledTaskConfig{
			Exclude: config.ScheduledTask.Exclude,
			Include: config.ScheduledTask.Include,
		},
		Service: windows.ServiceConfig{
			Include: config.Service.Include,
			Exclude: config.Service.Exclude,
		},
		SMB: windows.SMBConfig{
			EnabledList: split(config.SMB.EnabledList),
		},
		SMBClient: windows.SMBClientConfig{
			EnabledList: split(config.SMBClient.EnabledList),
		},
		SMTP: windows.SMTPConfig{
			BlackList: config.SMTP.BlackList,
			WhiteList: config.SMTP.WhiteList,
			Exclude:   config.SMTP.Exclude,
			Include:   config.SMTP.Include,
		},
		TextFile: windows.TextFileConfig{
			TextFileDirectory: config.TextFile.TextFileDirectory,
		},
		TCP: windows.TCPConfig{
			EnabledList: split(config.TCP.EnabledList),
		},
		DNS: windows.DNSConfig{
			EnabledList: split(config.DNS.EnabledList),
		},
		Update: windows.UpdateConfig{
			Online:         config.Update.Online,
			ScrapeInterval: config.Update.ScrapeInterval,
		},
	}
}
