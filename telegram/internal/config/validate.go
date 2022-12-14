// Code generated by go-validator; DO NOT EDIT.
// Package config contains models and autogenerated validation code
package config

// Validate validates struct accordingly to fields tags
func (c Config) Validate() []string {
	var errs []string
	if e := c.Delivery.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}
	if e := c.Logger.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}

	return errs
}

// Validate validates struct accordingly to fields tags
func (d Delivery) Validate() []string {
	var errs []string
	if e := d.HTTPServer.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}
	if e := d.TelegramBot.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}

	return errs
}

// Validate validates struct accordingly to fields tags
func (h HTTPServer) Validate() []string {
	var errs []string
	if h.ListenAddress == "" {
		errs = append(errs, "listen_address::is_required")
	}
	if h.ReadTimeout == 0 {
		errs = append(errs, "read_timeout::is_required")
	}
	if h.WriteTimeout == 0 {
		errs = append(errs, "write_timeout::is_required")
	}
	if h.BodySizeLimitBytes == 0 {
		errs = append(errs, "body_size_limit_bytes::is_required")
	}
	if h.GracefulTimeout == 0 {
		errs = append(errs, "graceful_timeout::is_required")
	}

	return errs
}

// Validate validates struct accordingly to fields tags
func (t TelegramBot) Validate() []string {
	var errs []string
	if t.Token == "" {
		errs = append(errs, "token::is_required")
	}

	return errs
}
