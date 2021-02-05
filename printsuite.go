package termtools

import (
	"errors"
	"fmt"
)

var (
	ErrUnknownPrinter               = errors.New("error: no such printer")
	ErrFailedToAdd                  = errors.New("error: failed to add printer: name already taken or nil pointer passed")
	ErrFailedToProcessPrinterConfig = errors.New("error: failed to process PrinterConfig: field Name must not be empty when adding to PrintSuite")
	ErrEmptyName                    = errors.New("error: printer name may not be empty string")
)

// PrintSuite holds pointers to one ore more instances of Printer and
// allows to switch available printers on the fly to use differrent
// output styles. An instance of printer is embedded into PrintSuite so
// you can call printer methods on it directly.
type PrintSuite struct {
	Printer
	available map[string]*Printer
}

// Configure accepts zero or more PrinterConfig and adds printers to
// PrintSuite. If one or more configs fail to process the method will
// return an error listing names that failed to add.
func (suite *PrintSuite) Configure(configs ...PrinterConfig) error {
	suite.ensureMapExists()
	failing := ""
	for _, conf := range configs {
		if conf.Name == "" {
			return ErrFailedToProcessPrinterConfig
		}
		if p, err := NewPrinter(conf); err == nil {
			suite.available[conf.Name] = p
		} else {
			failing += conf.Name + " "
		}
	}
	if failing != "" {
		return fmt.Errorf("error: failed to add the following names: %v", failing)
	}
	return nil
}

// Use returns instance of printer with requested printername. It returns no error
// even if printer name is incorrect. In this case a default Printer instance is returned.
func (suite *PrintSuite) Use(printername string) *Printer {
	suite.ensureMapExists()
	if printer, ok := suite.available[printername]; ok {
		return printer
	}
	return &Printer{}
}

// UseDefault acts in the same manner as Use and returns printer with no style options set
// which will output the same as fmt module.
func (suite *PrintSuite) UseDefault() *Printer {
	return &Printer{}
}

// SwitchTo sets the default PrintSuite printer to printer with requested name.
// If printername is not known the method will return an error without changes to
// current configuration.
func (suite *PrintSuite) SwitchTo(printername string) error {
	suite.ensureMapExists()
	if printer, ok := suite.available[printername]; ok && printer != nil {
		suite.Printer = *printer
		// note that we're dereferencing here, so changes to currently
		// used printer configuration will not affect stored configurations
		return nil
	}
	return ErrUnknownPrinter
}

// SwitchToDefault switches active printer of PrintSuite to default Printer{} (with no settings)
func (suite *PrintSuite) SwitchToDefault() {
	suite.Printer = Printer{}
}

// AddPrinter accepts name of printer and pointer to printer. If nil pointer or empty string
// is passed or if printername is already added the method will fail with an error.
func (suite *PrintSuite) AddPrinter(printername string, p *Printer) error {
	if printername == "" {
		return ErrEmptyName
	}
	suite.ensureMapExists()
	if _, ok := suite.available[printername]; !ok && p != nil && printername != "" {
		suite.available[printername] = p
		return nil
	}
	return ErrFailedToAdd
}

// ensureMapExists is triggered on each call to PrintSuite methods which
// operate on suite.available. This makes sure that we never derefence
// nil map. This is needed because this way we can intialize PrintSuite
// without calling any functions, i.e. just by declaring var ps PrintSuite
// or saying ps := PRintSuite{}
func (suite *PrintSuite) ensureMapExists() {
	if suite.available == nil {
		suite.available = make(map[string]*Printer)
	}
}
