package datainfo

import (
	"time"
)

type (
	// SequenceGeneratorInfo - sequence generator query
	SequenceGeneratorInfo struct {
		UpsertQuery     string
		ResultQuery     string
		NamePlaceHolder string
	}
	DataOption func(do *DataInfo) error
	DataInfo   struct {
		Schema                 *string                // Schema to use
		ReferenceMode          *bool                  // Indicates that the data is in reference mode
		ReferenceModePrefix    *string                // Reference mode prefix. The default is 'ref'.
		InterpolateTables      *bool                  // Interpolate tables that has been enclosed by {}
		ConnectionString       *string                // Connection string of data
		DriverName             *string                // Driver name to use
		HelperID               *string                // Helper ID to use
		ParameterInSequence    *bool                  // Parameter is in sequence
		ParameterPlaceHolder   *string                // Parameter place holder
		StringEnclosingChar    *string                // Gets or sets the character that encloses a string in the query
		StringEscapeChar       *string                // Gets or Sets the character that escapes a reserved character such as the character that encloses a s string
		ReservedWordEscapeChar *string                // Reserved word escape chars. For escaping with different opening and closing characters, just set to both. Example. `[]` for SQL server
		MaxOpenConnection      *int                   // Maximum open connection
		MaxIdleConnection      *int                   // Maximum idle connection
		MaxConnectionLifetime  *int                   // Max connection lifetime
		MaxConnectionIdleTime  *int                   // Max idle connection lifetime
		Ping                   *bool                  // Ping connection
		UserName               *string                // User name to indicate who initiated the transaction call
		SequenceGenerator      *SequenceGeneratorInfo // Sequence generator
	}
)

var (
	refModePfx, paramPh,
	strEncChar, strEscChar,
	resWrdEscChar string
	intTbls, paramInSeq bool
	maxOpnConn, maxIdlConn,
	maxConnLt, maxConnIdlLt int
)

func init() {
	refModePfx = `ref`
	paramPh = `?`
	strEncChar = `'`
	strEscChar = `\`
	resWrdEscChar = `[]`
	intTbls = true
	paramInSeq = true
	maxOpnConn = 25
	maxIdlConn = 25
	maxConnLt = int(5 * time.Minute)
	maxConnIdlLt = int(3 * time.Minute)
}

// New initializes a common data info and accepts further option where the user could change
func New(options ...DataOption) *DataInfo {
	n := DataInfo{
		Schema:                 new(string),
		ReferenceMode:          new(bool),
		ReferenceModePrefix:    &refModePfx,
		InterpolateTables:      &intTbls,
		ParameterInSequence:    &paramInSeq,
		ParameterPlaceHolder:   &paramPh,
		StringEnclosingChar:    &strEncChar,
		StringEscapeChar:       &strEscChar,
		ReservedWordEscapeChar: &resWrdEscChar,
		MaxOpenConnection:      &maxOpnConn,
		MaxIdleConnection:      &maxIdlConn,
		MaxConnectionLifetime:  &maxConnLt,
		MaxConnectionIdleTime:  &maxConnIdlLt,
		Ping:                   new(bool),
	}
	for _, o := range options {
		if o == nil {
			continue
		}
		o(&n)
	}
	return &n
}

// NewMinimal initializes a common data info by using the minimal set. Further options are accepted via options parameter.
func NewMinimal(connStr, schema, driver, user string, options ...DataOption) *DataInfo {
	opts := []DataOption{
		ConnectionString(connStr),
		Schema(schema),
		DriverName(driver),
		UserName(user),
	}
	if len(options) > 0 {
		opts = append(opts, options...)
	}
	return New(
		opts...,
	)
}

// Schema sets the schema of the data
func Schema(sch string) DataOption {
	return func(d *DataInfo) error {
		d.Schema = new(string)
		*d.Schema = sch
		return nil
	}
}

// ReferenceMode indicates that the data is in reference mode
//
// By default, table or views will have a `ref` prefix unless changed by ReferenceModePrefix
func ReferenceMode(indeed bool) DataOption {
	return func(d *DataInfo) error {
		d.ReferenceMode = new(bool)
		*d.ReferenceMode = indeed
		return nil
	}
}

// ReferenceModePrefix is reference mode prefix. The default is 'ref'.
func ReferenceModePrefix(pfx string) DataOption {
	return func(d *DataInfo) error {
		d.ReferenceModePrefix = new(string)
		*d.ReferenceModePrefix = pfx
		return nil
	}
}

// InterpolateTables interpolates tables that has been enclosed by {}
func InterpolateTables(indeed bool) DataOption {
	return func(d *DataInfo) error {
		d.InterpolateTables = new(bool)
		*d.InterpolateTables = indeed
		return nil
	}
}

// ConnectionString sets the connection string of the data access library
func ConnectionString(conn string) DataOption {
	return func(d *DataInfo) error {
		d.ConnectionString = new(string)
		*d.ConnectionString = conn
		return nil
	}
}

// Driver name sets the driver name of the data access library
func DriverName(name string) DataOption {
	return func(d *DataInfo) error {
		d.DriverName = new(string)
		*d.DriverName = name
		return nil
	}
}

// HelperID sets the helper ID to use for datahelperlite implementation
func HelperID(id string) DataOption {
	return func(d *DataInfo) error {
		d.HelperID = new(string)
		*d.HelperID = id
		return nil
	}
}

// ParameterInSequence sets that the parameter placeholder is in sequence
func ParameterInSequence(indeed bool) DataOption {
	return func(d *DataInfo) error {
		d.ParameterInSequence = new(bool)
		*d.ParameterInSequence = indeed
		return nil
	}
}

// ParameterPlaceHolder sets the parameter place holder. The default is `?`.
func ParameterPlaceHolder(holder string) DataOption {
	return func(d *DataInfo) error {
		d.ParameterPlaceHolder = new(string)
		*d.ParameterPlaceHolder = holder
		return nil
	}
}

// StringEnclosingChar sets the character that encloses a string in the query
func StringEnclosingChar(char string) DataOption {
	return func(d *DataInfo) error {
		d.StringEnclosingChar = new(string)
		*d.StringEnclosingChar = char
		return nil
	}
}

// StringEscapeChar sets the character that escapes a reserved character such as the character that encloses a s string
func StringEscapeChar(char string) DataOption {
	return func(d *DataInfo) error {
		d.StringEscapeChar = new(string)
		*d.StringEscapeChar = char
		return nil
	}
}

// ReservedWordEscapeChar sets the reserved word escape character(s). For escaping with different opening and closing characters, just set to both. Example. `[]` for SQL server
func ReservedWordEscapeChar(char string) DataOption {
	return func(d *DataInfo) error {
		d.ReservedWordEscapeChar = new(string)
		*d.ReservedWordEscapeChar = char
		return nil
	}
}

// MaxOpenConnection sets maximum open connection
func MaxOpenConnection(max int) DataOption {
	return func(d *DataInfo) error {
		d.MaxOpenConnection = new(int)
		*d.MaxOpenConnection = max
		return nil
	}
}

// MaxIdleConnection sets the maximum idle connection
func MaxIdleConnection(max int) DataOption {
	return func(d *DataInfo) error {
		d.MaxIdleConnection = new(int)
		*d.MaxIdleConnection = max
		return nil
	}
}

// MaxConnectionLifetime sets the max connection lifetime of each open connection
func MaxConnectionLifetime(max int) DataOption {
	return func(d *DataInfo) error {
		d.MaxConnectionLifetime = new(int)
		*d.MaxConnectionLifetime = max
		return nil
	}
}

// MaxConnectionIdleTime sets the max idle connection lifetime
func MaxConnectionIdleTime(max int) DataOption {
	return func(d *DataInfo) error {
		d.MaxConnectionIdleTime = new(int)
		*d.MaxConnectionIdleTime = max
		return nil
	}
}

// Ping turns on or off whether a connection needs to ping before executing a query
func Ping(indeed bool) DataOption {
	return func(d *DataInfo) error {
		d.Ping = new(bool)
		*d.Ping = indeed
		return nil
	}
}

// UserName sets the user who initiated the transaction call
func UserName(name string) DataOption {
	return func(d *DataInfo) error {
		d.UserName = new(string)
		*d.UserName = name
		return nil
	}
}

// SequenceGenerator sets the sequence generator
func SequenceGenerator(seqGen SequenceGeneratorInfo) DataOption {
	return func(d *DataInfo) error {
		d.SequenceGenerator = &seqGen
		return nil
	}
}
