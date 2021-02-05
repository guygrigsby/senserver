package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type Test struct {
	Text string
}

func TestParse(t *testing.T) {

	tests := []Test{
		{Text: `
coretemp-isa-0000
Adapter: ISA adapter
Package id 0:  +60.0°C  (high = +80.0°C, crit = +98.0°C)
Core 0:        +60.0°C  (high = +80.0°C, crit = +98.0°C)
Core 1:        +56.0°C  (high = +80.0°C, crit = +98.0°C)
Core 2:        +46.0°C  (high = +80.0°C, crit = +98.0°C)
Core 3:        +58.0°C  (high = +80.0°C, crit = +98.0°C)

nct6776-isa-0290
Adapter: ISA adapter
Vcore:           1.42 V  (min =  +0.00 V, max =  +1.74 V)
in1:             1.81 V  (min =  +0.00 V, max =  +0.00 V)  ALARM
AVCC:            3.34 V  (min =  +2.98 V, max =  +3.63 V)
+3.3V:           3.33 V  (min =  +2.98 V, max =  +3.63 V)
in4:           984.00 mV (min =  +0.00 V, max =  +0.00 V)  ALARM
in5:             1.68 V  (min =  +0.00 V, max =  +0.00 V)  ALARM
in6:           880.00 mV (min =  +0.00 V, max =  +0.00 V)  ALARM
3VSB:            3.42 V  (min =  +2.98 V, max =  +3.63 V)
Vbat:            3.26 V  (min =  +2.70 V, max =  +3.63 V)
fan1:          1117 RPM  (min =    0 RPM)
fan2:          1259 RPM  (min =    0 RPM)
fan3:             0 RPM  (min =    0 RPM)
fan4:          1101 RPM  (min =    0 RPM)
fan5:           725 RPM  (min =    0 RPM)
SYSTIN:         +30.0°C  (high =  +0.0°C, hyst =  +0.0°C)  ALARM  sensor = thermistor
CPUTIN:         +42.0°C  (high = +80.0°C, hyst = +75.0°C)  sensor = thermistor
AUXTIN:         +37.5°C  (high = +80.0°C, hyst = +75.0°C)  sensor = thermistor
PECI Agent 0:   +60.0°C  (high = +80.0°C, hyst = +75.0°C)
			 (crit = +98.0°C)
PCH_CHIP_TEMP:   +0.0°C
PCH_CPU_TEMP:    +0.0°C
PCH_MCH_TEMP:    +0.0°C
intrusion0:    ALARM
intrusion1:    OK
beep_enable:   disabled
		`},
	}

	for _, tc := range tests {
		res, err := Parse([]byte(tc.Text))
		require.NoError(t, err)
		fmt.Printf("%+v", res)
	}
}
