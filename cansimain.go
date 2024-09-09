package Cansi

import "fmt"

func SuperT(x string, option int, rgb ...[3]int) string {
	reset := "\x1B[0m"
	switch option {
	case 0:
		colorTx := fmt.Sprintf("\x1B[38;2;%d;%d;%dm", rgb[0][0], rgb[0][1], rgb[0][2])
		return fmt.Sprintf("%s%s%s", colorTx, x, reset)
	case 1:
		colorBg := fmt.Sprintf("\x1B[48;2;%d;%d;%dm", rgb[0][0], rgb[0][1], rgb[0][2])
		return fmt.Sprintf("%s%s%s", colorBg, x, reset)
	case 2:
		colorTx := fmt.Sprintf("\x1B[38;2;%d;%d;%dm", rgb[0][0], rgb[0][1], rgb[0][2])
		colorBg := fmt.Sprintf("\x1B[48;2;%d;%d;%dm", rgb[1][0], rgb[1][1], rgb[1][2])
		return fmt.Sprintf("%s%s%s%s", colorTx, colorBg, x, reset)
	default:
		return ""
	}
}

func SuperP(x *string, option int, rgb ...[3]int) {
	reset := "\x1B[0m"
	switch option {
	case 0:
		colorTx := fmt.Sprintf("\x1B[38;2;%d;%d;%dm", rgb[0][0], rgb[0][1], rgb[0][2])
		*x = fmt.Sprintf("%s%s%s", colorTx, *x, reset)
	case 1:
		colorBg := fmt.Sprintf("\x1B[48;2;%d;%d;%dm", rgb[0][0], rgb[0][1], rgb[0][2])
		*x = fmt.Sprintf("%s%s%s", colorBg, *x, reset)
	case 2:
		colorTx := fmt.Sprintf("\x1B[38;2;%d;%d;%dm", rgb[0][0], rgb[0][1], rgb[0][2])
		colorBg := fmt.Sprintf("\x1B[48;2;%d;%d;%dm", rgb[1][0], rgb[1][1], rgb[1][2])
		*x = fmt.Sprintf("%s%s%s%s", colorTx, colorBg, *x, reset)
	}
}
