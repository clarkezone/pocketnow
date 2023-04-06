package geocacheservice

type CosmosDBWriter struct {
}

func (c *CosmosDBWriter) Process(msg Message) {
	// TODO foreach, call convert, call COSMOSWrite via an interface
}
