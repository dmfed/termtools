package termtools

import "errors"

var (
	ErrUnknownPrinter = errors.New("error: no such printer")
	ErrFailedToAdd    = errors.New("error: failed to add printer: name already taken or nil pointer passed")
)

// PrintSuite holds pointers to one ore more instances of Printer and
// allows to switch available printers on the fly to use differrent
// output styles. An instance of printer is embedded into PrintSuite so
// you can call printer methods on it directly.
type PrintSuite struct {
	available map[string]*Printer
	Printer
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

// Default switches active printer of PRintSuite to default Printer{} (with no settings)
func (suite *PrintSuite) Default() {
	suite.Printer = Printer{}
}

// SwitchTo sets the default PrintSuite printer to printer with requested name
func (suite *PrintSuite) SwitchTo(printername string) error {
	suite.ensureMapExists()
	if printer, ok := suite.available[printername]; ok && printer != nil {
		suite.Printer = *printer
		// note that we're dereferencing here, so changes to currently
		// used printer will not affect stored configurations
		return nil
	}
	return ErrUnknownPrinter
}

// AddPrinter accepts
func (suite *PrintSuite) AddPrinter(printername string, p *Printer) error {
	suite.ensureMapExists()
	if _, ok := suite.available[printername]; !ok && p != nil {
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
