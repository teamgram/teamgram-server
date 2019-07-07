package phonenumbers

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		input       string
		err         error
		expectedNum uint64
		region      string
	}{
		{
			input:       "4437990238",
			err:         nil,
			expectedNum: 4437990238,
			region:      "US",
		}, {
			input:       "(443) 799-0238",
			err:         nil,
			expectedNum: 4437990238,
			region:      "US",
		}, {
			input:       "((443) 799-023asdfghjk8",
			err:         ErrNumTooLong,
			expectedNum: 0,
			region:      "US",
		}, {
			input:       "+441932567890",
			err:         nil,
			expectedNum: 1932567890,
			region:      "GB",
		}, {
			input:       "45",
			err:         nil,
			expectedNum: 45,
			region:      "US",
		}, {
			input:       "1800AWWCUTE",
			err:         nil,
			expectedNum: 8002992883,
			region:      "US",
		}, {
			input:       "+1 1951178619",
			err:         nil,
			expectedNum: 1951178619,
			region:      "US",
		}, {
			input:       "+33 07856952",
			err:         nil,
			expectedNum: 7856952,
			region:      "",
		}, {
			input:       "190022+22222",
			err:         ErrNotANumber,
			expectedNum: 0,
			region:      "US",
		}, {
			input:       "967717105526",
			err:         nil,
			expectedNum: 717105526,
			region:      "YE",
		}, {
			input:       "+68672098006",
			err:         nil,
			expectedNum: 72098006,
			region:      "",
		},
	}

	for i, test := range tests {
		num, err := Parse(test.input, test.region)
		if err != test.err {
			t.Errorf("[test %d:err] failed: %v != %v\n", i, err, test.err)
		}
		if num.GetNationalNumber() != test.expectedNum {
			t.Errorf("[test %d:num] failed: %v != %v\n", i, num.GetNationalNumber(), test.expectedNum)
		}
	}
}

func TestConvertAlphaCharactersInNumber(t *testing.T) {
	var tests = []struct {
		input, output string
	}{
		{input: "1800AWWPOOP", output: "18002997667"},
		{input: "(800) DAW-ORLD", output: "(800) 329-6753"},
		{input: "1800-ABC-DEF", output: "1800-222-333"},
	}

	for i, test := range tests {
		out := ConvertAlphaCharactersInNumber(test.input)
		if out != test.output {
			t.Errorf("[test %d] failed, %s != %s\n", i, out, test.output)
		}
	}
}

func TestNormalizeDigits(t *testing.T) {
	var tests = []struct {
		input         string
		expected      []byte
		keepNonDigits bool
	}{
		{input: "4445556666", expected: []byte("4445556666"), keepNonDigits: false},
		{input: "(444)5556666", expected: []byte("4445556666"), keepNonDigits: false},
		{input: "(444)555a6666", expected: []byte("4445556666"), keepNonDigits: false},
		{input: "(444)555a6666", expected: []byte("(444)555a6666"), keepNonDigits: true},
	}

	for i, test := range tests {
		out := normalizeDigits(test.input, test.keepNonDigits)
		if string(out) != string(test.expected) {
			t.Errorf("[test %d] failed: %s != %s\n",
				i, string(out), string(test.expected))
		}
	}
}

func TestExtractPossibleNumber(t *testing.T) {
	var (
		input    = "(530) 583-6985 x302/x2303"
		expected = "530) 583-6985 x302" // yes, the leading '(' is missing
	)

	output := extractPossibleNumber(input)
	if output != expected {
		t.Error(output, "!=", expected)
	}
}

func TestIsViablePhoneNumer(t *testing.T) {
	var tests = []struct {
		input    string
		isViable bool
	}{
		{
			input:    "4445556666",
			isViable: true,
		}, {
			input:    "+441932123456",
			isViable: true,
		}, {
			input:    "4930123456",
			isViable: true,
		}, {
			input:    "2",
			isViable: false,
		}, {
			input:    "helloworld",
			isViable: false,
		},
	}

	for i, test := range tests {
		result := isViablePhoneNumber(test.input)
		if result != test.isViable {
			t.Errorf("[test %d] %v != %v\n", i, result, test.isViable)
		}
	}
}

func TestNormalize(t *testing.T) {
	var tests = []struct {
		in  string
		exp string
	}{
		{in: "4431234567", exp: "4431234567"},
		{in: "443 1234567", exp: "4431234567"},
		{in: "(443)123-4567", exp: "4431234567"},
		{in: "800yoloFOO", exp: "8009656366"},
		{in: "444111a2222", exp: "4441112222"},

		// from libponenumber [java] unit tests
		{in: "034-56&+#2\u00AD34", exp: "03456234"},
		{in: "034-I-am-HUNGRY", exp: "034426486479"},
		{in: "\uFF125\u0665", exp: "255"},
		{in: "\u06F52\u06F0", exp: "520"},
	}

	// TODO(ttacon): the above commented out test are because we hacked the crap
	// out of normalizeDigits, fix it

	for i, test := range tests {
		res := normalize(test.in)
		if res != test.exp {
			t.Errorf("[test %d] %s != %s\n", i, res, test.exp)
		}
	}
}

func TestIsValidNumber(t *testing.T) {
	var tests = []struct {
		input   string
		err     error
		isValid bool
		region  string
	}{
		{
			input:   "4437990238",
			err:     nil,
			isValid: true,
			region:  "US",
		}, {
			input:   "(443) 799-0238",
			err:     nil,
			isValid: true,
			region:  "US",
		}, {
			input:   "((443) 799-023asdfghjk8",
			err:     ErrNumTooLong,
			isValid: false,
			region:  "US",
		}, {
			input:   "+441932567890",
			err:     nil,
			isValid: true,
			region:  "GB",
		}, {
			input:   "45",
			err:     nil,
			isValid: false,
			region:  "US",
		}, {
			input:   "1800AWWCUTE",
			err:     nil,
			isValid: true,
			region:  "US",
		}, {
			input:   "+343511234567",
			err:     nil,
			isValid: false,
			region:  "ES",
		}, {
			input:   "+12424654321",
			err:     nil,
			isValid: true,
			region:  "BS",
		}, {
			input:   "6041234567",
			err:     nil,
			isValid: false,
			region:  "US",
		},
	}

	for i, test := range tests {
		num, err := Parse(test.input, test.region)
		if err != test.err {
			t.Errorf("[test %d:err] failed: %v != %v\n", i, err, test.err)
		}
		if test.err != nil {
			continue
		}
		if IsValidNumber(num) != test.isValid {
			t.Errorf("[test %d:validity] failed: %v != %v\n",
				i, IsValidNumber(num), test.isValid)
		}
	}
}

func TestIsValidNumberForRegion(t *testing.T) {
	var tests = []struct {
		input            string
		err              error
		isValid          bool
		validationRegion string
		region           string
	}{
		{
			input:            "4437990238",
			err:              nil,
			isValid:          true,
			validationRegion: "US",
			region:           "US",
		}, {
			input:            "(443) 799-0238",
			err:              nil,
			isValid:          true,
			region:           "US",
			validationRegion: "US",
		}, {
			input:            "((443) 799-023asdfghjk8",
			err:              ErrNumTooLong,
			isValid:          false,
			region:           "US",
			validationRegion: "US",
		}, {
			input:            "+441932567890",
			err:              nil,
			isValid:          true,
			region:           "GB",
			validationRegion: "GB",
		}, {
			input:            "45",
			err:              nil,
			isValid:          false,
			region:           "US",
			validationRegion: "US",
		}, {
			input:            "1800AWWCUTE",
			err:              nil,
			isValid:          true,
			region:           "US",
			validationRegion: "US",
		}, {
			input:            "+441932567890",
			err:              nil,
			isValid:          false,
			region:           "GB",
			validationRegion: "US",
		}, {
			input:            "1800AWWCUTE",
			err:              nil,
			isValid:          false,
			region:           "US",
			validationRegion: "GB",
		}, {
			input:            "01932 869755",
			region:           "GB",
			err:              nil,
			isValid:          true,
			validationRegion: "GB",
		}, {
			input:            "6041234567",
			region:           "US",
			err:              nil,
			isValid:          false,
			validationRegion: "US",
		},
	}

	for i, test := range tests {
		num, err := Parse(test.input, test.region)
		if err != test.err {
			t.Errorf("[test %d:err] failed: %v != %v\n", i, err, test.err)
		}
		if test.err != nil {
			continue
		}
		if IsValidNumberForRegion(num, test.validationRegion) != test.isValid {
			t.Errorf("[test %d:validity] failed: %v != %v\n",
				i, IsValidNumberForRegion(num, test.validationRegion), test.isValid)
		}
	}
}

func TestIsPossibleNumberWithReason(t *testing.T) {
	var tests = []struct {
		input  string
		region string
		err    error
		valid  ValidationResult
	}{
		{
			input:  "16502530000",
			region: "US",
			err:    nil,
			valid:  IS_POSSIBLE,
		}, {
			input:  "2530000",
			region: "US",
			err:    nil,
			valid:  IS_POSSIBLE_LOCAL_ONLY,
		}, {
			input:  "65025300001",
			region: "US",
			err:    nil,
			valid:  TOO_LONG,
		}, {
			input:  "2530000",
			region: "",
			err:    ErrInvalidCountryCode,
			valid:  IS_POSSIBLE_LOCAL_ONLY,
		}, {
			input:  "253000",
			region: "US",
			err:    nil,
			valid:  TOO_SHORT,
		}, {
			input:  "1234567890",
			region: "SG",
			err:    nil,
			valid:  IS_POSSIBLE,
		}, {
			input:  "800123456789",
			region: "US",
			err:    nil,
			valid:  TOO_LONG,
		}, {
			input:  "+1456723456",
			region: "US",
			err:    nil,
			valid:  TOO_SHORT,
		}, {
			input:  "6041234567",
			region: "US",
			err:    nil,
			valid:  IS_POSSIBLE,
		},
	}

	for i, test := range tests {
		num, err := Parse(test.input, test.region)
		if err != nil {
			if test.err == err {
				continue
			}
			t.Errorf("[test %d:err] failed: %v\n", i, err)
		}

		valid := IsPossibleNumberWithReason(num)
		if valid != test.valid {
			t.Errorf("[test %d:possible] %s failed: %v != %v\n", i, test.input, valid, test.valid)
		}
	}
}

func TestFormat(t *testing.T) {
	// useful link for validating against official lib:
	// http://libphonenumber.appspot.com/phonenumberparser?number=019+3286+9755&country=GB

	var tests = []struct {
		in     string
		exp    string
		region string
		frmt   PhoneNumberFormat
	}{
		{
			in:     "019 3286 9755",
			region: "GB",
			exp:    "01932 869755",
			frmt:   NATIONAL,
		}, {
			in:     "+44 (0) 1932 869755",
			region: "GB",
			exp:    "+44 1932 869755",
			frmt:   INTERNATIONAL,
		}, {
			in:     "4431234567",
			region: "US",
			exp:    "(443) 123-4567",
			frmt:   NATIONAL,
		}, {
			in:     "4431234567",
			region: "US",
			exp:    "+14431234567",
			frmt:   E164,
		}, {
			in:     "4431234567",
			region: "US",
			exp:    "+1 443-123-4567",
			frmt:   INTERNATIONAL,
		}, {
			in:     "4431234567",
			region: "US",
			exp:    "tel:+1-443-123-4567",
			frmt:   RFC3966,
		},
		{
			in:     "+1 100-083-0033",
			region: "US",
			exp:    "+1 1000830033",
			frmt:   INTERNATIONAL,
		},
	}

	for i, test := range tests {
		num, err := Parse(test.in, test.region)
		if err != nil {
			t.Errorf("[test %d] failed: should be able to parse, err:%v\n", i, err)
		}
		got := Format(num, test.frmt)
		if got != test.exp {
			t.Errorf("[test %d:fmt] failed %s != %s\n", i, got, test.exp)
		}
	}
}

func TestFormatForMobileDialing(t *testing.T) {
	var tests = []struct {
		in     string
		exp    string
		region string
		frmt   PhoneNumberFormat
	}{
		{
			in:     "950123456",
			region: "UZ",
			exp:    "+998950123456",
		},
	}

	for i, test := range tests {
		num, err := Parse(test.in, test.region)
		if err != nil {
			t.Errorf("[test %d] failed: should be able to parse, err:%v\n", i, err)
		}
		got := FormatNumberForMobileDialing(num, test.region, false)
		if got != test.exp {
			t.Errorf("[test %d:fmt] failed %s != %s\n", i, got, test.exp)
		}
	}
}

func TestFormatByPattern(t *testing.T) {
	var tcs = []struct {
		in          string
		region      string
		format      PhoneNumberFormat
		userFormats []*NumberFormat
		exp         string
	}{
		{
			in:     "+33122334455",
			region: "FR",
			format: E164,
			userFormats: []*NumberFormat{
				&NumberFormat{
					Pattern: s(`(\d+)`),
					Format:  s(`$1`),
				},
			},
			exp: "+33122334455",
		}, {
			in:     "+442070313000",
			region: "UK",
			format: NATIONAL,
			userFormats: []*NumberFormat{
				&NumberFormat{
					Pattern: s(`(20)(\d{4})(\d{4})`),
					Format:  s(`$1 $2 $3`),
				},
			},
			exp: "20 7031 3000",
		},
	}

	for i, tc := range tcs {
		num, err := Parse(tc.in, tc.region)
		if err != nil {
			t.Errorf("[test %d] failed: should be able to parse, err:%v\n", i, err)
		}
		got := FormatByPattern(num, tc.format, tc.userFormats)
		if got != tc.exp {
			t.Errorf("[test %d:fmt] failed %s != %s\n", i, got, tc.exp)
		}
	}
}

func TestSetItalianLeadinZerosForPhoneNumber(t *testing.T) {
	var tests = []struct {
		num          string
		numLeadZeros int32
		hasLeadZero  bool
	}{
		{
			num:          "00000",
			numLeadZeros: 4,
			hasLeadZero:  true,
		},
		{
			num:          "0123456",
			numLeadZeros: 1,
			hasLeadZero:  true,
		},
		{
			num:          "0023456",
			numLeadZeros: 2,
			hasLeadZero:  true,
		},
		{
			num:          "123456",
			numLeadZeros: 1, // it's the default value
			hasLeadZero:  false,
		},
	}

	for i, test := range tests {
		var pNum = &PhoneNumber{}
		setItalianLeadingZerosForPhoneNumber(test.num, pNum)
		if pNum.GetItalianLeadingZero() != test.hasLeadZero {
			t.Errorf("[test %d:hasLeadZero] %v != %v\n",
				i, pNum.GetItalianLeadingZero(), test.hasLeadZero)
		}
		if pNum.GetNumberOfLeadingZeros() != test.numLeadZeros {
			t.Errorf("[test %d:numLeadZeros] %v != %v\n",
				i, pNum.GetNumberOfLeadingZeros(), test.numLeadZeros)
		}

	}
}

////////// Copied from java-libphonenumber
/**
 * Unit tests for PhoneNumberUtil.java
 *
 * Note that these tests use the test metadata, not the normal metadata
 * file, so should not be used for regression test purposes - these tests
 * are illustrative only and test functionality.
 *
 * @author Shaopeng Jia
 */

// TODO(ttacon): use the test metadata and not the normal metadata

var testPhoneNumbers = map[string]*PhoneNumber{
	"ALPHA_NUMERIC_NUMBER": newPhoneNumber(1, 80074935247),
	"AE_UAN":               newPhoneNumber(971, 600123456),
	"AR_MOBILE":            newPhoneNumber(54, 91187654321),
	"AR_NUMBER":            newPhoneNumber(54, 1157774533),
	"AU_NUMBER":            newPhoneNumber(61, 236618300),
	"BS_MOBILE":            newPhoneNumber(1, 2423570000),
	"BS_NUMBER":            newPhoneNumber(1, 2423651234),
	// Note that this is the same as the example number for DE in the metadata.
	"DE_NUMBER":       newPhoneNumber(49, 30123456),
	"DE_SHORT_NUMBER": newPhoneNumber(49, 1234),
	"GB_MOBILE":       newPhoneNumber(44, 7912345678),
	"GB_NUMBER":       newPhoneNumber(44, 2070313000),
	"IT_MOBILE":       newPhoneNumber(39, 345678901),
	"IT_NUMBER": func() *PhoneNumber {
		p := newPhoneNumber(39, 236618300)
		p.ItalianLeadingZero = proto.Bool(true)
		return p
	}(),
	"JP_STAR_NUMBER": newPhoneNumber(81, 2345),
	// Numbers to test the formatting rules from Mexico.
	"MX_MOBILE1": newPhoneNumber(52, 12345678900),
	"MX_MOBILE2": newPhoneNumber(52, 15512345678),
	"MX_NUMBER1": newPhoneNumber(52, 3312345678),
	"MX_NUMBER2": newPhoneNumber(52, 8211234567),
	"NZ_NUMBER":  newPhoneNumber(64, 33316005),
	"SG_NUMBER":  newPhoneNumber(65, 65218000),
	// A too-long and hence invalid US number.
	"US_LONG_NUMBER": newPhoneNumber(1, 65025300001),
	"US_NUMBER":      newPhoneNumber(1, 6502530000),
	"US_PREMIUM":     newPhoneNumber(1, 9002530000),
	// Too short, but still possible US numbers.
	"US_LOCAL_NUMBER":        newPhoneNumber(1, 2530000),
	"US_SHORT_BY_ONE_NUMBER": newPhoneNumber(1, 650253000),
	"US_TOLLFREE":            newPhoneNumber(1, 8002530000),
	"US_SPOOF":               newPhoneNumber(1, 0),
	"US_SPOOF_WITH_RAW_INPUT": func() *PhoneNumber {
		p := newPhoneNumber(1, 0)
		p.RawInput = proto.String("000-000-0000")
		return p
	}(),
	"INTERNATIONAL_TOLL_FREE": newPhoneNumber(800, 12345678),
	// We set this to be the same length as numbers for the other
	// non-geographical country prefix that we have in our test metadata.
	// However, this is not considered valid because they differ in
	// their country calling code.
	"INTERNATIONAL_TOLL_FREE_TOO_LONG":  newPhoneNumber(800, 123456789),
	"UNIVERSAL_PREMIUM_RATE":            newPhoneNumber(979, 123456789),
	"UNKNOWN_COUNTRY_CODE_NO_RAW_INPUT": newPhoneNumber(2, 12345),
}

func newPhoneNumber(cc int, natNum uint64) *PhoneNumber {
	p := &PhoneNumber{}
	p.CountryCode = proto.Int(cc)
	p.NationalNumber = proto.Uint64(natNum)
	return p
}

func getTestNumber(alias string) *PhoneNumber {
	// there should never not be a valid number
	val := testPhoneNumbers[alias]
	return val
}

func TestGetSupportedRegions(t *testing.T) {
	if len(GetSupportedRegions()) == 0 {
		t.Error("there should be supported regions, found none")
	}
}

func TestGetMetadata(t *testing.T) {
	var tests = []struct {
		name       string
		id         string
		cc         int32
		i18nPref   string
		natPref    string
		numFmtSize int
	}{
		{
			name:       "US",
			id:         "US",
			cc:         1,
			i18nPref:   "011",
			natPref:    "1",
			numFmtSize: 2,
		}, {
			name:       "DE",
			id:         "DE",
			cc:         49,
			i18nPref:   "00",
			natPref:    "0",
			numFmtSize: 18,
		}, {
			name:       "AR",
			id:         "AR",
			cc:         54,
			i18nPref:   "00",
			natPref:    "0",
			numFmtSize: 11,
		},
	}
	for i, test := range tests {
		meta := getMetadataForRegion(test.name)
		if meta.GetId() != test.name {
			t.Errorf("[test %d:name] %s != %s\n", i, meta.GetId(), test.name)
		}
		if meta.GetCountryCode() != test.cc {
			t.Errorf("[test %d:cc] %d != %d\n", i, meta.GetCountryCode(), test.cc)
		}
		if meta.GetInternationalPrefix() != test.i18nPref {
			t.Errorf("[test %d:i18nPref] %s != %s\n",
				i, meta.GetInternationalPrefix(), test.i18nPref)
		}
		if meta.GetNationalPrefix() != test.natPref {
			t.Errorf("[test %d:natPref] %s != %s\n",
				i, meta.GetNationalPrefix(), test.natPref)
		}
		if len(meta.GetNumberFormat()) != test.numFmtSize {
			t.Errorf("[test %d:numFmt] %d != %d\n",
				i, len(meta.GetNumberFormat()), test.numFmtSize)
		}
	}
}

func TestIsNumberGeographical(t *testing.T) {
	if !isNumberGeographical(getTestNumber("AU_NUMBER")) {
		t.Error("Australia should be a geographical number")
	}
	if isNumberGeographical(getTestNumber("INTERNATIONAL_TOLL_FREE")) {
		t.Error("An international toll free number should not be geographical")
	}
}

func TestGetLengthOfGeographicalAreaCode(t *testing.T) {
	var tests = []struct {
		numName string
		length  int
	}{
		{numName: "US_NUMBER", length: 3},
		{numName: "US_TOLLFREE", length: 0},
		{numName: "GB_NUMBER", length: 2},
		{numName: "GB_MOBILE", length: 0},
		{numName: "AR_NUMBER", length: 2},
		{numName: "AU_NUMBER", length: 1},
		{numName: "IT_NUMBER", length: 2},
		{numName: "SG_NUMBER", length: 0},
		{numName: "US_SHORT_BY_ONE_NUMBER", length: 0},
		{numName: "INTERNATIONAL_TOLL_FREE", length: 0},
	}
	for i, test := range tests {
		l := GetLengthOfGeographicalAreaCode(getTestNumber(test.numName))
		if l != test.length {
			t.Errorf("[test %d:length] %d != %d for %s\n", i, l, test.length, test.numName)
		}
	}
}

func TestGetCountryMobileToken(t *testing.T) {
	if GetCountryMobileToken(GetCountryCodeForRegion("MX")) != "1" {
		t.Error("Mexico should have a mobile token == \"1\"")
	}
	if GetCountryMobileToken(GetCountryCodeForRegion("SE")) != "" {
		t.Error("Sweden should have a mobile token")
	}
}

func TestGetNationalSignificantNumber(t *testing.T) {
	var tests = []struct {
		name, exp string
	}{
		{"US_NUMBER", "6502530000"}, {"IT_MOBILE", "345678901"},
		{"IT_NUMBER", "0236618300"}, {"INTERNATIONAL_TOLL_FREE", "12345678"},
	}
	for i, test := range tests {
		natsig := GetNationalSignificantNumber(getTestNumber(test.name))
		if natsig != test.exp {
			t.Errorf("[test %d] %s != %s\n", i, natsig, test.exp)
		}
	}
}

func TestGetExampleNumberForType(t *testing.T) {
	if !reflect.DeepEqual(getTestNumber("DE_NUMBER"), GetExampleNumber("DE")) {
		t.Error("the example number for Germany should have been the " +
			"same as the test number we're using")
	}
	if !reflect.DeepEqual(
		getTestNumber("DE_NUMBER"), GetExampleNumberForType("DE", FIXED_LINE)) {
		t.Error("the example number for Germany should have been the " +
			"same as the test number we're using [FIXED_LINE]")
	}
	// For the US, the example number is placed under general description,
	// and hence should be used for both fixed line and mobile, so neither
	// of these should return null.
	if GetExampleNumberForType("US", FIXED_LINE) == nil {
		t.Error("FIXED_LINE example for US should not be nil")
	}
	if GetExampleNumberForType("US", MOBILE) == nil {
		t.Error("FIXED_LINE example for US should not be nil")
	}
	// CS is an invalid region, so we have no data for it.
	if GetExampleNumberForType("CS", MOBILE) != nil {
		t.Error("there should not be an example MOBILE number for the " +
			"invalid region \"CS\"")
	}
	// RegionCode 001 is reserved for supporting non-geographical country
	// calling code. We don't support getting an example number for it
	// with this method.
	if GetExampleNumber("UN001") != nil {
		t.Error("there should not be an example number for UN001 " +
			"that is retrievable by this method")
	}
}

func TestGetExampleNumberForNonGeoEntity(t *testing.T) {
	if !reflect.DeepEqual(
		getTestNumber("INTERNATIONAL_TOLL_FREE"),
		GetExampleNumberForNonGeoEntity(800)) {
		t.Error("there should be an example 800 number")
	}
	if !reflect.DeepEqual(
		getTestNumber("UNIVERSAL_PREMIUM_RATE"),
		GetExampleNumberForNonGeoEntity(979)) {
		t.Error("there should be an example number for 979")
	}
}

func TestNormalizeDigitsOnly(t *testing.T) {
	if NormalizeDigitsOnly("034-56&+a#234") != "03456234" {
		t.Errorf("didn't fully normalize digits only")
	}
}

func TestNormalizeDiallableCharsOnly(t *testing.T) {
	if normalizeDiallableCharsOnly("03*4-56&+a#234") != "03*456+234" {
		t.Error("did not correctly remove non-diallable characters")
	}
}

type testCase struct {
	num           string
	parseRegion   string
	expectedE164  string
	validRegion   string
	isValid       bool
	isValidRegion bool
}

type timeZonesTestCases struct {
	num              string
	expectedTimeZone string
}

type prefixMapTestCases struct {
	num      string
	lang     string
	expected string
}

func runTestBatch(t *testing.T, tests []testCase) {
	for _, test := range tests {
		n, err := Parse(test.num, test.parseRegion)
		if err != nil {
			t.Errorf("Failed to parse number %s: %s", test.num, err)
		}

		if IsValidNumber(n) != test.isValid {
			t.Errorf("Number %s: validity mismatch: expected %t got %t.", test.num, test.isValid, !test.isValid)
		}

		if IsValidNumberForRegion(n, test.validRegion) != test.isValidRegion {
			t.Errorf("Number %s: region validity mismatch: expected %t got %t.", test.num, test.isValidRegion, !test.isValidRegion)
		}

		s := Format(n, E164)
		if s != test.expectedE164 {
			t.Errorf("Expected '%s', got '%s'", test.expectedE164, s)
		}
	}
}

func TestItalianLeadingZeroes(t *testing.T) {

	tests := []testCase{
		{
			num:           "0491 570 156",
			parseRegion:   "AU",
			expectedE164:  "+61491570156",
			validRegion:   "AU",
			isValid:       true,
			isValidRegion: true,
		},
		{
			num:           "02 5550 1234",
			parseRegion:   "AU",
			expectedE164:  "+61255501234",
			validRegion:   "AU",
			isValid:       true,
			isValidRegion: true,
		},
		{
			num:           "+39.0399123456",
			parseRegion:   "IT",
			expectedE164:  "+390399123456",
			validRegion:   "IT",
			isValid:       true,
			isValidRegion: true,
		},
	}

	runTestBatch(t, tests)
}

func TestARNumberTransformRule(t *testing.T) {
	tests := []testCase{
		{
			num:           "+541151123456",
			parseRegion:   "AR",
			expectedE164:  "+541151123456",
			validRegion:   "AR",
			isValid:       true,
			isValidRegion: true,
		},
		{
			num:           "+540111561234567",
			parseRegion:   "AR",
			expectedE164:  "+5491161234567",
			validRegion:   "AR",
			isValid:       true,
			isValidRegion: true,
		},
	}

	runTestBatch(t, tests)
}

func TestLeadingOne(t *testing.T) {
	tests := []testCase{
		{
			num:           "15167706076",
			parseRegion:   "US",
			expectedE164:  "+15167706076",
			validRegion:   "US",
			isValid:       true,
			isValidRegion: true,
		},
	}

	runTestBatch(t, tests)
}

func TestNewIndianPhones(t *testing.T) {
	tests := []testCase{
		{
			num:           "7999999543",
			parseRegion:   "IN",
			expectedE164:  "+917999999543",
			validRegion:   "IN",
			isValid:       true,
			isValidRegion: true,
		},
	}

	runTestBatch(t, tests)
}

func TestBurkinaFaso(t *testing.T) {
	tests := []testCase{
		{
			num:           "+22658125926",
			parseRegion:   "",
			expectedE164:  "+22658125926",
			validRegion:   "BF",
			isValid:       true,
			isValidRegion: true,
		},
	}

	runTestBatch(t, tests)
}

func TestParsing(t *testing.T) {
	tests := []struct {
		number   string
		country  string
		expected string
	}{
		{"0788383383", "RW", "+250788383383"},
		{"+250788383383 ", "KE", "+250788383383"},
		{"+250788383383", "", "+250788383383"},
		{"(917)992-5253", "US", "+19179925253"},
		{"+62877747666", "", "+62877747666"},
		{"0877747666", "ID", "+62877747666"},
		{"07531669965", "GB", "+447531669965"},
		{"+22658125926", "", "+22658125926"},
	}

	for _, tc := range tests {
		num, err := Parse(tc.number, tc.country)
		if err != nil {
			t.Errorf("Error parsing number: %s: %s", tc.number, err)
		} else {
			formatted := Format(num, E164)
			if formatted != tc.expected {
				t.Errorf("Error parsing number '%s:%s', got %s when expecting %s", tc.number, tc.country, formatted, tc.expected)
			}
		}
	}
}

func TestGetTimeZonesForPrefix(t *testing.T) {
	tests := []timeZonesTestCases{
		{
			num:              "+442073238299",
			expectedTimeZone: "Europe/London",
		},
		{
			num:              "+61491570156",
			expectedTimeZone: "Australia/Adelaide",
		},
		{
			num:              "+61255501234",
			expectedTimeZone: "Australia/Sydney",
		},
		{
			num:              "+390399123456",
			expectedTimeZone: "Europe/Rome",
		},
		{
			num:              "+541151123456",
			expectedTimeZone: "America/Buenos_Aires",
		},
		{
			num:              "+15167706076",
			expectedTimeZone: "America/New_York",
		},
		{
			num:              "+917999999543",
			expectedTimeZone: "Asia/Calcutta",
		},
		{
			num:              "+540111561234567",
			expectedTimeZone: "America/Buenos_Aires",
		},
		{
			num:              "+18504320800",
			expectedTimeZone: "America/Chicago",
		},
		{
			num:              "+14079395277",
			expectedTimeZone: "America/New_York",
		},
		{
			num:              "+18508632167",
			expectedTimeZone: "America/Chicago",
		},
		{
			num:              "+40213158207",
			expectedTimeZone: "Europe/Bucharest",
		},
		// UTC +5:45
		{
			num:              "+97714240520",
			expectedTimeZone: "Asia/Katmandu",
		},
		// UTC -3:30
		{
			num:              "+17097264534",
			expectedTimeZone: "America/St_Johns",
		},
		{
			num:              "0000000000",
			expectedTimeZone: "Etc/Unknown",
		},
	}

	for _, test := range tests {
		timeZones, err := GetTimezonesForPrefix(test.num)
		if err != nil {
			t.Errorf("Failed to getTimezone for the number %s: %s", test.num, err)
		}

		if len(timeZones) == 0 {
			t.Errorf("Expected at least 1 timezone.")
		}

		if timeZones[0] != test.expectedTimeZone {
			t.Errorf("Expected '%s', got '%s' for '%s'", test.expectedTimeZone, timeZones[0], test.num)
		}

		num, err := Parse(test.num, "")
		if err != nil {
			continue
		}

		timeZones, err = GetTimezonesForNumber(num)
		if err != nil {
			t.Errorf("Failed to getTimezone for the number %s: %s", num, err)
		}

		if len(timeZones) == 0 {
			t.Errorf("Expected at least 1 timezone.")
		}

		if timeZones[0] != test.expectedTimeZone {
			t.Errorf("Expected '%s', got '%s' for '%s'", test.expectedTimeZone, timeZones[0], num)
		}
	}
}

func TestGetCarrierForNumber(t *testing.T) {
	tests := []prefixMapTestCases{
		{num: "+8613702032331", lang: "en", expected: "China Mobile"},
		{num: "+8613702032331", lang: "zh", expected: "中国移动"},
		{num: "+6281377468527", lang: "en", expected: "Telkomsel"},
		{num: "+8613323241342", lang: "en", expected: "China Telecom"},
		{num: "+61491570156", lang: "en", expected: "Telstra"},
		{num: "+917999999543", lang: "en", expected: "Reliance Jio"},
		{num: "+593992218722", lang: "en", expected: "Claro"},
	}
	for _, test := range tests {
		number, err := Parse(test.num, "ZZ")
		if err != nil {
			t.Errorf("Failed to parse number %s: %s", test.num, err)
		}
		carrier, err := GetCarrierForNumber(number, test.lang)
		if err != nil {
			t.Errorf("Failed to getCarrier for the number %s: %s", test.num, err)
		}
		if test.expected != carrier {
			t.Errorf("Expected '%s', got '%s' for '%s'", test.expected, carrier, test.num)
		}
	}
}

func TestGetGeocodingForNumber(t *testing.T) {
	tests := []prefixMapTestCases{
		{num: "+8613702032331", lang: "en", expected: "Tianjin"},
		{num: "+8613702032331", lang: "zh", expected: "天津市"},
		{num: "+863197785050", lang: "zh", expected: "河北省邢台市"},
		{num: "+8613323241342", lang: "en", expected: "Baoding, Hebei"},
		{num: "+917999999543", lang: "en", expected: "Ahmedabad Local, Gujarat"},
		{num: "+17047181840", lang: "en", expected: "North Carolina"},
		{num: "+12542462158", lang: "en", expected: "Texas"},
		{num: "+16193165996", lang: "en", expected: "California"},
		{num: "+12067799191", lang: "en", expected: "Washington State"},
	}
	for _, test := range tests {
		number, err := Parse(test.num, "ZZ")
		if err != nil {
			t.Errorf("Failed to parse number %s: %s", test.num, err)
		}
		geocoding, err := GetGeocodingForNumber(number, test.lang)
		if err != nil {
			t.Errorf("Failed to getGeocoding for the number %s: %s", test.num, err)
		}
		if test.expected != geocoding {
			t.Errorf("Expected '%s', got '%s' for '%s'", test.expected, geocoding, test.num)
		}
	}
}
func TestMaybeStripExtension(t *testing.T) {
	var tests = []struct {
		input     string
		number    uint64
		extension string
		region    string
	}{
		{
			input:     "1234576 ext. 1234",
			number:    1234576,
			extension: "1234",
			region:    "US",
		}, {
			input:     "1234-576",
			number:    1234576,
			extension: "",
			region:    "US",
		}, {
			input:     "1234576-123#",
			number:    1234576,
			extension: "123",
			region:    "US",
		}, {
			input:     "1234576 ext.123#",
			number:    1234576,
			extension: "123",
			region:    "US",
		},
	}

	for i, test := range tests {
		num, _ := Parse(test.input, test.region)
		if num.GetNationalNumber() != test.number {
			t.Errorf("[test %d:num] failed: %v != %v\n", i, num.GetNationalNumber(), test.number)
		}

		if num.GetExtension() != test.extension {
			t.Errorf("[test %d:num] failed: %v != %v\n", i, num.GetExtension(), test.extension)
		}
	}
}

func TestGetSupportedCallingCodes(t *testing.T) {
	var tests = []struct {
		code    int
		present bool
	}{
		{
			1,
			true,
		}, {
			800,
			true,
		}, {
			593,
			true,
		}, {
			44,
			true,
		}, {
			999,
			false,
		},
	}

	supported := GetSupportedCallingCodes()
	for i, tc := range tests {
		if supported[tc.code] != tc.present {
			t.Errorf("[test %d:num] failed for code %d: %v != %v\n", i, tc.code, tc.present, supported[tc.code])
		}
	}
}

func TestMergeLengths(t *testing.T) {
	var tests = []struct {
		l1     []int32
		l2     []int32
		merged []int32
	}{
		{
			[]int32{1, 5},
			[]int32{2, 3, 4},
			[]int32{1, 2, 3, 4, 5},
		},
		{
			[]int32{1},
			[]int32{3, 4},
			[]int32{1, 3, 4},
		},
		{
			[]int32{1, 2, 5},
			[]int32{4},
			[]int32{1, 2, 4, 5},
		},
	}

	for i, tc := range tests {
		merged := mergeLengths(tc.l1, tc.l2)
		if !reflect.DeepEqual(merged, tc.merged) {
			t.Errorf("[test %d:num] failed for l1: %v and l2: %v: %v != %v\n", i, tc.l1, tc.l2, tc.merged, merged)
		}
	}
}

func TestRegexCacheWrite(t *testing.T) {
	pattern1 := "TestRegexCacheWrite"
	if _, found1 := readFromRegexCache(pattern1); found1 {
		t.Errorf("pattern |%v| is in the cache", pattern1)
	}
	regex1 := regexFor(pattern1)
	cachedRegex1, found1 := readFromRegexCache(pattern1)
	if !found1 {
		t.Errorf("pattern |%v| is not in the cache", pattern1)
	}
	if regex1 != cachedRegex1 {
		t.Error("expected the same instance, but got a different one")
	}
	pattern2 := pattern1 + "."
	if _, found2 := readFromRegexCache(pattern2); found2 {
		t.Errorf("pattern |%v| is in the cache", pattern2)
	}
}

func TestRegexCacheRead(t *testing.T) {
	pattern1 := "TestRegexCacheRead"
	if _, found1 := readFromRegexCache(pattern1); found1 {
		t.Errorf("pattern |%v| is in the cache", pattern1)
	}
	regex1 := regexp.MustCompile(pattern1)
	writeToRegexCache(pattern1, regex1)
	if cachedRegex1 := regexFor(pattern1); cachedRegex1 != regex1 {
		t.Error("expected the same instance, but got a different one")
	}
	cachedRegex1, found1 := readFromRegexCache(pattern1)
	if !found1 {
		t.Errorf("pattern |%v| is not in the cache", pattern1)
	}
	if cachedRegex1 != regex1 {
		t.Error("expected the same instance, but got a different one")
	}
	pattern2 := pattern1 + "."
	if _, found2 := readFromRegexCache(pattern2); found2 {
		t.Errorf("pattern |%v| is in the cache", pattern2)
	}
}

func TestRegexCacheStrict(t *testing.T) {
	const expectedResult = "(41) 3020-3445"
	phoneToTest := &PhoneNumber{
		CountryCode:    proto.Int32(55),
		NationalNumber: proto.Uint64(4130203445),
	}
	firstRunResult := Format(phoneToTest, NATIONAL)
	if expectedResult != firstRunResult {
		t.Errorf("phone number formatting not as expected")
	}
	// This adds value to the regex cache that would break the following lookup if the regex-s
	// in cache were not strict.
	Format(&PhoneNumber{
		CountryCode:    proto.Int32(973),
		NationalNumber: proto.Uint64(17112724),
	}, NATIONAL)
	secondRunResult := Format(phoneToTest, NATIONAL)

	if expectedResult != secondRunResult {
		t.Errorf("phone number formatting not as expected")
	}
}

func s(str string) *string {
	return &str
}
