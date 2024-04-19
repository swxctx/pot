package gopot

import (
	"encoding/json"
	"fmt"
	"net"
)

var (
	client *PotClient
)

// PotClient
type PotClient struct {
	// cmd encoder
	encoder *json.Encoder
	// response decoder
	decoder *json.Decoder
	// pot server tcp conn client
	conn net.Conn
	// config
	cfg *Config
	// module key prefix
	keyPrefix string
}

// NewClient
func NewClient(cfg *Config) (*PotClient, error) {
	// connect pot server
	conn, err := net.Dial("tcp", cfg.Address)
	if err != nil {
		return nil, fmt.Errorf("NewClient: connection to pot server err-> %v", err)
	}

	client := &PotClient{
		encoder: json.NewEncoder(conn),
		decoder: json.NewDecoder(conn),
		conn:    conn,
		cfg:     cfg,
	}
	if len(cfg.Module) > 0 {
		client.keyPrefix = fmt.Sprintf("%s:", cfg.Module)
	}

	return client, nil
}

// Close
func (c *PotClient) Close() {
	c.conn.Close()
}

// GetKey
func (c *PotClient) GetKey(key string) string {
	return fmt.Sprintf("%s%s", c.keyPrefix, key)
}
