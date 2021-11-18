package etherscan

// GetLogs  gets a list of log for a single address
// startBlock and endBlock cann`t be nil
func (c *Client) GetLogs(address *string, fromBlock string, toBlock string, topic0 string) (logs []Log, err error) {
	param := M{
		"fromBlock": fromBlock,
		"toBlock":   toBlock,
		"topic0":    topic0,
	}
	compose(param, "address", address)
	err = c.call("logs", "getLogs", param, &logs)
	return
}
