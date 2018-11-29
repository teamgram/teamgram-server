package libphonenumber

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
			expectedNum: 951178619,
			region:      "US",
		}, {
			input:       "+33 07856952",
			err:         nil,
			expectedNum: 7856952,
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

func Test_normalizeDigits(t *testing.T) {
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

func Test_extractPossibleNumber(t *testing.T) {
	var (
		input    = "(530) 583-6985 x302/x2303"
		expected = "530) 583-6985 x302" // yes, the leading '(' is missing
	)

	output := extractPossibleNumber(input)
	if output != expected {
		t.Error(output, "!=", expected)
	}
}

func Test_isViablePhoneNumer(t *testing.T) {
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

func Test_normalize(t *testing.T) {
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

func Test_IsValidNumber(t *testing.T) {
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

func Test_IsValidNumberForRegion(t *testing.T) {
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

func TestFormat(t *testing.T) {
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
			exp:    "+1 000830033",
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

func Test_setItalianLeadinZerosForPhoneNumber(t *testing.T) {
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

func Test_testNumberLengthAgainstPattern(t *testing.T) {
	var tests = []struct {
		pattern  string
		num      string
		expected ValidationResult
	}{
		{
			"\\d{7}(?:\\d{3})?",
			"1234567",
			IS_POSSIBLE,
		},
		{
			"\\d{7}(?:\\d{3})?",
			"1234567890",
			IS_POSSIBLE,
		},
		{
			"\\d{7}(?:\\d{3})?",
			"12345678",
			TOO_LONG,
		},
		{
			"\\d{7}(?:\\d{3})?",
			"123456",
			TOO_SHORT,
		},
		{
			"\\d{7}(?:\\d{3})?",
			"abc1234567",
			TOO_SHORT,
		},
	}

	for i, test := range tests {
		pat := regexp.MustCompile(test.pattern)
		res := testNumberLengthAgainstPattern(pat, test.num)
		if res != test.expected {
			t.Errorf("[test %d] failed: should be %v, got %v", i, test.expected, res)
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
	"AR_NUMBER":            newPhoneNumber(54, 1187654321),
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

func Test_GetSupportedRegions(t *testing.T) {
	if len(GetSupportedRegions()) == 0 {
		t.Error("there should be supported regions, found none")
	}
}

func Test_getMetadata(t *testing.T) {
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

func Test_isLeadingZeroPossible(t *testing.T) {
	if !isLeadingZeroPossible(39) {
		t.Error("Leading 0 should be possible in Italy")
	}
	if isLeadingZeroPossible(1) {
		t.Error("Leading 0 should not be possible in the USA")
	}
	if !isLeadingZeroPossible(800) {
		t.Error("Leading 0 should be possible for International toll free")
	}
	if isLeadingZeroPossible(889) {
		t.Error("Leading 0 should not be possible in non-existent region")
	}
}

func Test_isNumberGeographical(t *testing.T) {
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
			t.Errorf("[test %d:length] %d != %d\n", i, l, test.length)
		}
	}
}

func TestGetCountryMobileToken(t *testing.T) {
	if "1" != GetCountryMobileToken(GetCountryCodeForRegion("MX")) {
		t.Error("Mexico should have a mobile token == \"1\"")
	}
	if "" != GetCountryMobileToken(GetCountryCodeForRegion("SE")) {
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

func Test_GetExampleNumberForType(t *testing.T) {
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
	if "03456234" != NormalizeDigitsOnly("034-56&+a#234") {
		t.Errorf("didn't fully normalize digits only")
	}
}

func Test_normalizeDiallableCharsOnly(t *testing.T) {
	if "03*456+234" != normalizeDiallableCharsOnly("03*4-56&+a#234") {
		t.Error("did not correctly remove non-diallable characters")
	}
}

type testCase struct {
	num          string
	region       string
	expectedE164 string
	valid        bool
}

type timeZonesTestCases struct {
	num              string
	expectedTimeZone string
}

func runTestBatch(t *testing.T, tests []testCase) {
	for _, test := range tests {
		n, err := Parse(test.num, test.region)
		if err != nil {
			t.Errorf("Failed to parse number %s: %s", test.num, err)
		}

		if IsValidNumberForRegion(n, test.region) != test.valid {
			t.Errorf("Number %s: validity mismatch: expected %t got %t.", test.num, test.valid, !test.valid)
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
			num:          "0491 570 156",
			region:       "AU",
			expectedE164: "+61491570156",
			valid:        true,
		},
		{
			num:          "02 5550 1234",
			region:       "AU",
			expectedE164: "+61255501234",
			valid:        true,
		},
		{
			num:          "+39.0399123456",
			region:       "IT",
			expectedE164: "+390399123456",
			valid:        true,
		},
	}

	runTestBatch(t, tests)
}

func TestARNumberTransformRule(t *testing.T) {
	tests := []testCase{
		{
			num:          "+541151123456",
			region:       "AR",
			expectedE164: "+541151123456",
			valid:        true,
		},
		{
			num:          "+540111561234567",
			region:       "AR",
			expectedE164: "+5491161234567",
			valid:        true,
		},
	}

	runTestBatch(t, tests)
}

func TestLeadingOne(t *testing.T) {
	tests := []testCase{
		{
			num:          "15167706076",
			region:       "US",
			expectedE164: "+15167706076",
			valid:        true,
		},
	}

	runTestBatch(t, tests)
}

func TestNewIndianPhones(t *testing.T) {
	tests := []testCase{
		{
			num:          "7999999543",
			region:       "IN",
			expectedE164: "+917999999543",
			valid:        true,
		},
	}

	runTestBatch(t, tests)
}

func TestGetTimeZonesForRegion(t *testing.T) {
	tests := []timeZonesTestCases{
		{
			num:              "+442073238299",
			expectedTimeZone: "Europe/London",
		},
		{
			num:              "+61491570156",
			expectedTimeZone: "Australia/Sydney",
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
		timeZones, err := GetTimeZonesForRegion(test.num)
		if err != nil {
			t.Errorf("Failed to getTimezone for the number %s: %s", test.num, err)
		}

		if len(timeZones) == 0 {
			t.Errorf("Expected at least 1 timezone.")
		}

		if timeZones[0] != test.expectedTimeZone {
			t.Errorf("Expected '%s', got '%s'", test.expectedTimeZone, timeZones[0])
		}
	}
}
