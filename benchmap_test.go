package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkSwitch10(b *testing.B) {
	s := fmt.Sprintf("string-number_%v", rand.Intn(len(map10)+1))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchSwitch10(s)
	}
}

func BenchmarkSwitch100(b *testing.B) {
	s := fmt.Sprintf("string-number_%v", rand.Intn(len(map100))+1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchSwitch100(s)
	}
}

func benchSwitch10(s string) int {
	r := 0
	switch s {
	case "string-number_1":
		r = 1
	case "string-number_2":
		r = 2
	case "string-number_3":
		r = 3
	case "string-number_4":
		r = 4
	case "string-number_5":
		r = 5
	case "string-number_6":
		r = 6
	case "string-number_7":
		r = 7
	case "string-number_8":
		r = 8
	case "string-number_9":
		r = 9
	case "string-number_10":
		r = 10
	default:
		r = -1
	}
	return r
}

func benchSwitch100(s string) int {
	r := 0
	switch s {
	case "string-number_1":
		r = 1
	case "string-number_2":
		r = 2
	case "string-number_3":
		r = 3
	case "string-number_4":
		r = 4
	case "string-number_5":
		r = 5
	case "string-number_6":
		r = 6
	case "string-number_7":
		r = 7
	case "string-number_8":
		r = 8
	case "string-number_9":
		r = 9
	case "string-number_10":
		r = 10
	case "string-number_11":
		r = 11
	case "string-number_12":
		r = 12
	case "string-number_13":
		r = 13
	case "string-number_14":
		r = 14
	case "string-number_15":
		r = 15
	case "string-number_16":
		r = 16
	case "string-number_17":
		r = 17
	case "string-number_18":
		r = 18
	case "string-number_19":
		r = 19
	case "string-number_20":
		r = 20
	case "string-number_21":
		r = 21
	case "string-number_22":
		r = 22
	case "string-number_23":
		r = 23
	case "string-number_24":
		r = 24
	case "string-number_25":
		r = 25
	case "string-number_26":
		r = 26
	case "string-number_27":
		r = 27
	case "string-number_28":
		r = 28
	case "string-number_29":
		r = 29
	case "string-number_30":
		r = 30
	case "string-number_31":
		r = 31
	case "string-number_32":
		r = 32
	case "string-number_33":
		r = 33
	case "string-number_34":
		r = 34
	case "string-number_35":
		r = 35
	case "string-number_36":
		r = 36
	case "string-number_37":
		r = 37
	case "string-number_38":
		r = 38
	case "string-number_39":
		r = 39
	case "string-number_40":
		r = 40
	case "string-number_41":
		r = 41
	case "string-number_42":
		r = 42
	case "string-number_43":
		r = 43
	case "string-number_44":
		r = 44
	case "string-number_45":
		r = 45
	case "string-number_46":
		r = 46
	case "string-number_47":
		r = 47
	case "string-number_48":
		r = 48
	case "string-number_49":
		r = 49
	case "string-number_50":
		r = 50
	case "string-number_51":
		r = 51
	case "string-number_52":
		r = 52
	case "string-number_53":
		r = 53
	case "string-number_54":
		r = 54
	case "string-number_55":
		r = 55
	case "string-number_56":
		r = 56
	case "string-number_57":
		r = 57
	case "string-number_58":
		r = 58
	case "string-number_59":
		r = 59
	case "string-number_60":
		r = 60
	case "string-number_61":
		r = 61
	case "string-number_62":
		r = 62
	case "string-number_63":
		r = 63
	case "string-number_64":
		r = 64
	case "string-number_65":
		r = 65
	case "string-number_66":
		r = 66
	case "string-number_67":
		r = 67
	case "string-number_68":
		r = 68
	case "string-number_69":
		r = 69
	case "string-number_70":
		r = 70
	case "string-number_71":
		r = 71
	case "string-number_72":
		r = 72
	case "string-number_73":
		r = 73
	case "string-number_74":
		r = 74
	case "string-number_75":
		r = 75
	case "string-number_76":
		r = 76
	case "string-number_77":
		r = 77
	case "string-number_78":
		r = 78
	case "string-number_79":
		r = 79
	case "string-number_80":
		r = 80
	case "string-number_81":
		r = 81
	case "string-number_82":
		r = 82
	case "string-number_83":
		r = 83
	case "string-number_84":
		r = 84
	case "string-number_85":
		r = 85
	case "string-number_86":
		r = 86
	case "string-number_87":
		r = 87
	case "string-number_88":
		r = 88
	case "string-number_89":
		r = 89
	case "string-number_90":
		r = 90
	case "string-number_91":
		r = 91
	case "string-number_92":
		r = 92
	case "string-number_93":
		r = 93
	case "string-number_94":
		r = 94
	case "string-number_95":
		r = 95
	case "string-number_96":
		r = 96
	case "string-number_97":
		r = 97
	case "string-number_98":
		r = 98
	case "string-number_99":
		r = 99
	case "string-number_100":
		r = 100
	default:
		r = -1
	}
	return r
}
func BenchmarkMap10(b *testing.B) {
	s := fmt.Sprintf("string-number_%v", rand.Intn(len(map10)+1))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchMap10(s)
	}
}
func BenchmarkMap100(b *testing.B) {
	s := fmt.Sprintf("string-number_%v", rand.Intn(len(map100)+1))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchMap100(s)
	}
}

var map10 = map[string]int{
	"string-number_1":  1,
	"string-number_2":  2,
	"string-number_3":  3,
	"string-number_4":  4,
	"string-number_5":  5,
	"string-number_6":  6,
	"string-number_7":  7,
	"string-number_8":  8,
	"string-number_9":  9,
	"string-number_10": 10,
}

var map100 = map[string]int{
	"string-number_1":   1,
	"string-number_2":   2,
	"string-number_3":   3,
	"string-number_4":   4,
	"string-number_5":   5,
	"string-number_6":   6,
	"string-number_7":   7,
	"string-number_8":   8,
	"string-number_9":   9,
	"string-number_10":  10,
	"string-number_11":  11,
	"string-number_12":  12,
	"string-number_13":  13,
	"string-number_14":  14,
	"string-number_15":  15,
	"string-number_16":  16,
	"string-number_17":  17,
	"string-number_18":  18,
	"string-number_19":  19,
	"string-number_20":  20,
	"string-number_21":  21,
	"string-number_22":  22,
	"string-number_23":  23,
	"string-number_24":  24,
	"string-number_25":  25,
	"string-number_26":  26,
	"string-number_27":  27,
	"string-number_28":  28,
	"string-number_29":  29,
	"string-number_30":  30,
	"string-number_31":  31,
	"string-number_32":  32,
	"string-number_33":  33,
	"string-number_34":  34,
	"string-number_35":  35,
	"string-number_36":  36,
	"string-number_37":  37,
	"string-number_38":  38,
	"string-number_39":  39,
	"string-number_40":  40,
	"string-number_41":  41,
	"string-number_42":  42,
	"string-number_43":  43,
	"string-number_44":  44,
	"string-number_45":  45,
	"string-number_46":  46,
	"string-number_47":  47,
	"string-number_48":  48,
	"string-number_49":  49,
	"string-number_50":  50,
	"string-number_51":  51,
	"string-number_52":  52,
	"string-number_53":  53,
	"string-number_54":  54,
	"string-number_55":  55,
	"string-number_56":  56,
	"string-number_57":  57,
	"string-number_58":  58,
	"string-number_59":  59,
	"string-number_60":  60,
	"string-number_61":  61,
	"string-number_62":  62,
	"string-number_63":  63,
	"string-number_64":  64,
	"string-number_65":  65,
	"string-number_66":  66,
	"string-number_67":  67,
	"string-number_68":  68,
	"string-number_69":  69,
	"string-number_70":  70,
	"string-number_71":  71,
	"string-number_72":  72,
	"string-number_73":  73,
	"string-number_74":  74,
	"string-number_75":  75,
	"string-number_76":  76,
	"string-number_77":  77,
	"string-number_78":  78,
	"string-number_79":  79,
	"string-number_80":  80,
	"string-number_81":  81,
	"string-number_82":  82,
	"string-number_83":  83,
	"string-number_84":  84,
	"string-number_85":  85,
	"string-number_86":  86,
	"string-number_87":  87,
	"string-number_88":  88,
	"string-number_89":  89,
	"string-number_90":  90,
	"string-number_91":  91,
	"string-number_92":  92,
	"string-number_93":  93,
	"string-number_94":  94,
	"string-number_95":  95,
	"string-number_96":  96,
	"string-number_97":  97,
	"string-number_98":  98,
	"string-number_99":  99,
	"string-number_100": 100,
}

func benchMap10(s string) int {
	return map10[s]
}
func benchMap100(s string) int {
	return map100[s]
}
