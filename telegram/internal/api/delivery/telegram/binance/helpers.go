/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package binance

import "fmt"

// toBinanceCourseURL - convert to binance course url.
func toBinanceCourseURL(url, message, currency string) string {
	return fmt.Sprintf("%s%s%s", url, message, currency)
}
